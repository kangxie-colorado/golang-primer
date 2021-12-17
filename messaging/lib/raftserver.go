package lib

import (
	"sort"
	"sync"

	log "github.com/sirupsen/logrus"
)

type RaftServer struct {
	raftlog *RaftLog
	raftnet *RaftNet

	myID        int
	followerIDs []int

	leaderTerm int

	// book keeping

	lastCommitIdx int
	commitIdx     int
	replicatedIdx [NetWorkSize]int

	// gurad around commitIdx read/write by
	// this way, I can use raftserver.lock() like https://kaviraj.me/understanding-condition-variable-in-go/
	// entering wait() will release the lock iirc
	lock sync.Mutex
	cond *sync.Cond

	// callback to application
	raftcallback func([]RaftLogEntry)
}

func CreateARaftServer(id int, callback func([]RaftLogEntry)) *RaftServer {
	var raftserver = RaftServer{}
	raftserver.myID = id
	raftserver.raftlog = &RaftLog{}
	raftserver.raftnet = CreateARaftNet(id)

	raftserver.lock = sync.Mutex{}
	raftserver.cond = sync.NewCond(&raftserver.lock)

	var followerIDs = []int{}
	for k := range RaftNetConfig {
		if k != id {
			followerIDs = append(followerIDs, k)
		}

	}

	// to save sending commitIdx more than one times, or actually yeah we can switch to heartbeat
	// but that also seems uncessary, when there is no commit, why boterh sending commitIdx
	raftserver.lastCommitIdx = -1
	raftserver.commitIdx = -1
	raftserver.replicatedIdx = [NetWorkSize]int{
		-1, -1, -1, -1, -1,
	}

	raftserver.followerIDs = followerIDs

	raftserver.raftcallback = callback

	// maybe currentTerm should start with something, not 0
	// but here, the server starts from beginning, as in the sense of the raftlog is empty...
	// so it being 0
	// in other cases, if there is a already a raftlog, it maybe different
	// keep it simple for now

	// ^ applies to commitIdx as well

	return &raftserver
}

func (raftserver *RaftServer) Net() *RaftNet {
	return raftserver.raftnet
}

func (raftserver *RaftServer) currentTermFromLog() int {
	if len(raftserver.raftlog.items) == 0 {
		// doesn't matter, append at index 0 always succeeds
		return -1
	}

	return raftserver.raftlog.items[len(raftserver.raftlog.items)-1].Term
}

// set term?
// leader election happened, new leader will update its own term
// follower will update on messages
// that is much later
func (raftserver *RaftServer) currentLeaderTerm() int {
	return raftserver.leaderTerm
}

// prototype of watiForCommit()
func (raftserver *RaftServer) watiForCommit(writtenIdx int, commited chan bool) {
	raftserver.lock.Lock()
	log.Debugf("lock: %v is locked, waiting on cond var: %v", raftserver.lock, raftserver.cond)

	//log.Debugf("commitIdx now %v, and wirttenIdx: %v", raftserver.commitIdx, writtenIdx)

	for raftserver.commitIdx < writtenIdx {
		raftserver.cond.Wait()
	}

	log.Debugf("commitIdx now %v, and wirttenIdx: %v", raftserver.commitIdx, writtenIdx)
	commited <- true

	raftserver.lock.Unlock()
	log.Debugf("lock: %v is unlocked", raftserver.lock)

}

func (raftserver *RaftServer) AppendNewEntry(msg string, commited chan bool) {
	// append to leader
	// from the respective of leader, the currentTermFromLog before appending, is the prevTerm parameter to the function
	// currentLeaderTerm() >= currentTermFromLog()
	writeIdx := len(raftserver.raftlog.items)
	success := raftserver.LeaderAppendEntries(writeIdx, raftserver.currentTermFromLog(), []RaftLogEntry{RaftLogEntry{raftserver.currentLeaderTerm(), msg}})

	if success {
		// for now, not blocked waiting for commited
		// need to think how to wait for commit ... better in a go routine actuall
		log.Infoln("Waiting Raft to commit")
		go raftserver.watiForCommit(writeIdx, commited)
		// commited <- true
	}
}

func (raftserver *RaftServer) LeaderAppendEntries(index, prevTerm int, entries []RaftLogEntry) bool {
	success := raftserver.raftlog.AppendEntries(index, prevTerm, entries)

	if success {
		// book keeping leader itself
		raftserver.replicatedIdx[raftserver.myID] = index

		// send AppendEntries Msg to followers
		for _, f := range raftserver.followerIDs {
			m := CreateAppendEntriesMsg(raftserver.myID, index, prevTerm, entries)
			// simple fixed length type field for now, of course can wrap this with even one more layer.
			// not the major focus
			log.Debugf("Sending raftnode:%v", f)
			raftserver.raftnet.Send(f, m.Encoding())
		}

	}
	return success
}

func (raftserver *RaftServer) Start() {
	go raftserver.Receive()
	go raftserver.raftnet.Start()
	// start the net as above ^
	// maybe start some book keeping?
	// like when a server starts, it assumes all followers log are same as its own
	/**
	 for f in followers:
	 	f.commitIdx = raftserver.commitIdx
		f.nextIdx = raftserver.nextIdx()
		// of such

	but for tonight, let me only make the log replication happen and that would be a victory
	 **/
}

func (raftserver *RaftServer) Receive() {
	for {
		log.Debugln("Waiting for a message")
		msgB64Encoding, _ := raftserver.raftnet.Receive()

		// route based on the message type
		// can refactor to a function if there is better solution
		switch msgB64Encoding[:MSGTYPEFIELDLEN] {
		case APPENDENTRYMSG:
			msg := AppendEntriesMsg{}
			msg.Decoding(msgB64Encoding[MSGTYPEFIELDLEN:])
			log.Infof("AppendEntriesMsg received: %v\n", msg.Repr())
			raftserver.FollowerAppendEntries(&msg)
		case APPENDENTRYRSP:
			msg := AppendEntriesResp{}
			msg.Decoding(msgB64Encoding[MSGTYPEFIELDLEN:])
			log.Infof("AppendEntriesResp received: %v\n", msg.Repr())
			raftserver.ProcessAppendEntriesResp(&msg)
		case COMMITUPDATE:
			msg := CommitUpdate{}
			msg.Decoding(msgB64Encoding[MSGTYPEFIELDLEN:])
			log.Infof("CommitUpdate received: %v\n", msg.Repr())
			raftserver.ProcessCommitUpdate(&msg)
		default:
			log.Errorln("Unknown Message Type!")
		}

	}
}

func (raftserver *RaftServer) FollowerAppendEntries(msg *AppendEntriesMsg) {
	success := raftserver.raftlog.AppendEntries(msg.Index, msg.PrevTerm, msg.Entries)
	resp := AppendEntriesResp{raftserver.myID, success, msg.Index, len(msg.Entries), raftserver.currentTermFromLog()}

	raftserver.raftnet.Send(msg.SenderId, resp.Encoding())
}

func (raftserver *RaftServer) ProcessAppendEntriesResp(msg *AppendEntriesResp) {
	if msg.Success {
		log.Debugf("Follower %v able to append to its own %v\n", msg.SenderId, msg.Repr())
		// establish the consensus here
		raftserver.replicatedIdx[msg.SenderId] = msg.Index + msg.NumOfEntries - 1

		// determine the highest replicatedIDX which has 3 or more shows, including leader itself
		// bookkeeping of leader itself happens in LeaderAppendEntries()
		raftserver.lock.Lock()
		raftserver.commitIdx = raftserver.DetermineCommitIdx()
		log.Debugf("lock: %v is locked, commitIdx is %v", raftserver.lock, raftserver.commitIdx)

		raftserver.lock.Unlock()
		log.Debugf("lock: %v is unlocked, commitIdx is %v", raftserver.lock, raftserver.commitIdx)

		// we should now signla checkForCommit channel
		// any client/goroutine waiting for its entry to be committed can now move on
		raftserver.cond.Broadcast()

		// we then should tell followers of this update
		// using heartbeat message can be an option
		// now just use a CommitUpdate msg, can incorporate with heartbeat if that is how paper describes
		if raftserver.lastCommitIdx < raftserver.commitIdx {
			commitUpdateMsg := CommitUpdate{raftserver.myID, raftserver.commitIdx}
			for _, f := range raftserver.followerIDs {
				raftserver.raftnet.Send(f, commitUpdateMsg.Encoding())
			}

			raftserver.lastCommitIdx = raftserver.commitIdx
		}

	} else {
		log.Errorf("Follower %v not able to append to its own %v, back tracking\n", msg.SenderId, msg.Repr())
		prevTerm := 0
		newIndex := msg.Index - 1
		if newIndex > 0 {
			prevTerm = raftserver.raftlog.items[newIndex-1].Term
		}
		backoffMsg := AppendEntriesMsg{raftserver.myID, newIndex, prevTerm, raftserver.raftlog.items[newIndex:]}
		// two places sending to follower
		// would there be some refactoring
		// conditional send to follower: condition being if index<follower.commitIdx, no need to send any more...
		raftserver.raftnet.Send(msg.SenderId, backoffMsg.Encoding())
	}

}

func (raftserver *RaftServer) DetermineCommitIdx() int {
	// [5,4,4,3,3] => 4
	// [5,4,3,3,3] => 3
	// looks like it is just to get the middle number?
	cpy := make([]int, NetWorkSize)
	copy(cpy, raftserver.replicatedIdx[:])
	sort.Ints(cpy)

	log.Debugf("the replication indexes are %v\n", cpy)

	return cpy[NetWorkSize/2]
}

func (raftserver *RaftServer) ProcessCommitUpdate(msg *CommitUpdate) {
	// in case the replication has not caught up on the follower, just wait until next nudge
	// it could be a heartbeat or next message
	if len(raftserver.raftlog.items) > msg.CommitIdx {
		// msg.commitIdx is already written in leader and should be written in follower, so should plus 1 to form the right bound
		// raftserver.commitIdx, 1)starting with -1, 2)also should be written already: to avoid rewriting it, plus 1 generalize both
		go raftserver.raftcallback(raftserver.raftlog.items[raftserver.commitIdx+1 : msg.CommitIdx+1])
		raftserver.commitIdx = msg.CommitIdx
	}

}
