package lib

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type RaftServer struct {
	raftlog *RaftLog
	raftnet *RaftNet

	myID        int
	followerIDs []int

	term int

	// book keeping
	commitIdx int
}

func CreateARaftServer(id int) *RaftServer {
	var raftserver = RaftServer{}
	raftserver.myID = id
	raftserver.raftlog = &RaftLog{}
	raftserver.raftnet = CreateARaftNet(id)

	var followerIDs = []int{}
	for k := range RaftNetConfig {
		if k != id {
			followerIDs = append(followerIDs, k)
		}
	}
	raftserver.followerIDs = followerIDs

	// maybe currentTerm should start with something, not 0
	// but here, the server starts from beginning, as in the sense of the raftlog is empty...
	// so it being 0
	// in other cases, if there is a already a raftlog, it maybe different
	// keep it simple for now

	// ^ applies to commitIdx as well

	return &raftserver
}

func (raftserver *RaftServer) nextIndex() int {
	return len(raftserver.raftlog.items)
}

func (raftserver *RaftServer) Net() *RaftNet {
	return raftserver.raftnet
}

func (raftserver *RaftServer) prevTerm() int {
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

func (raftserver *RaftServer) currentTerm() int {
	return raftserver.term
}

// prototype of watiForCommit()
func (raftserver *RaftServer) watiForCommit(writtenIdx int, commited chan bool) {
	for {
		if raftserver.commitIdx >= writtenIdx {
			commited <- true
			return
		}
		// <- checkForCommit // some chanel to check for commit, can be signalled when updating commit index of bookkeeping
		time.Sleep(1 * time.Second)
	}
}

func (raftserver *RaftServer) AppendNewEntry(msg string, commited chan bool) {
	// append to leader
	success := raftserver.LeaderAppendEntries(raftserver.nextIndex(), raftserver.prevTerm(), []RaftLogEntry{RaftLogEntry{raftserver.currentTerm(), msg}})

	if success {
		// for now, not blocked waiting for commited
		// need to think how to wait for commit ... better in a go routine actuall
		// go raftserver.watiForCommit(raftserver.nextIndex(), commited)
		commited <- true
	}
}

func (raftserver *RaftServer) LeaderAppendEntries(index, prevTerm int, entries []RaftLogEntry) bool {
	success := raftserver.raftlog.AppendEntries(index, prevTerm, entries)

	if success {
		// send AppendEntries Msg to followers
		for _, f := range raftserver.followerIDs {
			m := CreateAppendEntriesMsg(index, prevTerm, entries)
			// simple fixed length type field for now, of course can wrap this with even one more layer.
			// not the major focus
			log.Debugf("Sending raftnode:%v", f)
			raftserver.raftnet.Send(f, "APPEND"+ToGOB64(&m))
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
		log.Infoln("Waiting for a message")
		msgB64Encoding, _ := raftserver.raftnet.Receive()
		switch msgB64Encoding[:6] {
		case "APPEND":
			msg := FromGOB64(msgB64Encoding[6:])
			log.Infof("AppendEntriesMsg received %v\n", msg)
			raftserver.FollowerAppendEntries(msg.Index, msg.PrevTerm, msg.Entries)
		default:
			log.Errorln("Unknown Message Type!")
		}

	}
}

func (raftserver *RaftServer) FollowerAppendEntries(index, prevTerm int, entries []RaftLogEntry) bool {
	success := raftserver.raftlog.AppendEntries(index, prevTerm, entries)
	return success
}
