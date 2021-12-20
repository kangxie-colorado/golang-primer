package lib

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	Undefined nodeRole = iota
	Follower
	Candidate
	Leader
)

func (noderole nodeRole) asString() string {
	switch noderole {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	default:
		return "Undefined"
	}
}

type RaftState struct {
	raftlog *RaftLog // share the reference with raftserver

	myId         int
	myRole       nodeRole
	currentTerm  int // maybe not the same as the term from log
	votedFor     int
	whoVoteForMe []int

	// book keeping
	commitIdx     int
	replicatedIdx [NetWorkSize]int

	nextIdx [NetWorkSize]int // not in use yet
}

func (raftstate *RaftState) Repr() string {
	return fmt.Sprintf("RaftState{myId=%v, myRole=%v, currentTerm=%v, votedFor=%v}",
		raftstate.myId, raftstate.myRole.asString(), raftstate.currentTerm, raftstate.votedFor)
}

func (raftstate *RaftState) initRaftState(id int) {
	raftstate.myId = id

	raftstate.myRole = Follower
	// simple hack here to do some early test
	if id == 0 {
		raftstate.myRole = Candidate
	}

	raftstate.currentTerm = 0
	raftstate.votedFor = -1
	raftstate.whoVoteForMe = []int{}

	raftstate.raftlog = &RaftLog{}

	raftstate.commitIdx = -1
	raftstate.replicatedIdx = [NetWorkSize]int{
		-1, -1, -1, -1, -1,
	}

	raftlogLen := 0
	raftstate.nextIdx = [NetWorkSize]int{
		raftlogLen, raftlogLen, raftlogLen, raftlogLen, raftlogLen,
	}

}

func (raftstate *RaftState) handleTimer() {

}

func (raftstate *RaftState) handleAppendEntriesMsg(msg *AppendEntriesMsg, raftserver *RaftServer) {
	// if msg.term > my term, convert to follower
	// if msg.term < my term, just ignore it... no need to reply
	log.Debugf("Current I am %s, received %s", raftstate.Repr(), msg.Repr())
	if raftstate.myRole == Follower {
		if msg.LeaderTerm >= raftstate.currentTerm {
			raftstate.currentTerm = msg.LeaderTerm
			raftstate.FollowerAppendEntries(msg, raftserver)
		}

	}

	if raftstate.myRole == Candidate {
		if msg.LeaderTerm > raftstate.currentTerm {
			raftstate.myRole = Follower
			raftstate.currentTerm = msg.LeaderTerm
			raftstate.FollowerAppendEntries(msg, raftserver)
		}

	}

	if raftstate.myRole == Leader {
		// network partition
		if msg.LeaderTerm > raftstate.currentTerm {
			raftstate.myRole = Follower
			raftstate.currentTerm = msg.LeaderTerm
			raftstate.FollowerAppendEntries(msg, raftserver)
		}

	}

	// apparently above can refactor into one block; notice the > vs >=
	// keep it this way until later, easier for debugging

}

func (raftstate *RaftState) handleAppendEntriesResp(msg *AppendEntriesResp, raftserver *RaftServer) {
	log.Debugf("Current I am %s, received %s", raftstate.Repr(), msg.Repr())

	// only learder is possible to receive this message
	// in network partiotion and lesser side, the fake leader will continue to receieve the together partioned follower
	// but in that sense, it is still leader but won't be able to commit at all
	// because we ignore the msg term < my term, so message with a bigger term won't come back
	// this simplifies things; this fake leader can become follower when it receives the new leader's heartbeat
	if raftstate.myRole != Leader {
		panic("Not A Leader! But Received AppendEntriesResp " + msg.Repr())
	}

	// deal as normal
	raftstate.ProcessAppendEntriesResp(msg, raftserver)

}

func (raftstate *RaftState) handleRequestVoteMsg(msg *RequestVoteMsg, raftserver *RaftServer) {
	log.Debugf("Current I am %s, received %s", raftstate.Repr(), msg.Repr())

	// appears follower should deal with this message, for sure
	// also appears candidate should deal with this message
	// If votedFor is null or candidateId, and candidate’s log is at least as up-to-date as receiver’s log, grant vote (§5.2, §5.4)
	// maybe all server should just deal with it... and candidate is special case
	voteGranted := false
	if msg.Term > raftstate.currentTerm {
		raftstate.myRole = Follower
		raftstate.votedFor = msg.SenderId
		voteGranted = true
		raftstate.currentTerm = msg.Term
	}

	if msg.Term == raftstate.currentTerm && msg.LastLogTerm >= raftstate.prevTermFromLog() &&
		msg.LastLogIdx >= len(raftstate.raftlog.items)-1 &&
		(raftstate.votedFor == -1 || raftstate.votedFor == msg.SenderId) {
		raftstate.votedFor = msg.SenderId
		voteGranted = true

	}

	resp := RequestVoteResp{raftstate.myId, raftstate.currentTerm, voteGranted}
	raftserver.SendRequestVoteResp(msg.SenderId, &resp)

}

func (raftstate *RaftState) handleRequestVoteResp(msg *RequestVoteResp, raftserver *RaftServer) {
	log.Debugf("Current I am %s, received %s", raftstate.Repr(), msg.Repr())

	if msg.Term > raftstate.currentTerm {
		raftstate.myRole = Follower
	}

	if msg.VoteGranted {
		raftstate.whoVoteForMe = append(raftstate.whoVoteForMe, msg.SenderId)
		// need to deal with myself vote
		if len(raftstate.whoVoteForMe) >= 3 {
			raftstate.myRole = Leader
			go raftserver.LeaderNoop()
		}
	}

}

func (raftstate *RaftState) prevTermFromLog() int {
	if len(raftstate.raftlog.items) == 0 {
		// doesn't matter, append at index 0 always succeeds
		return -1
	}

	return raftstate.raftlog.items[len(raftstate.raftlog.items)-1].Term
}

func (raftstate *RaftState) FollowerAppendEntries(msg *AppendEntriesMsg, raftserver *RaftServer) {
	success := raftstate.raftlog.AppendEntries(msg.Index, msg.PrevTerm, msg.Entries)
	resp := AppendEntriesResp{raftstate.myId, success, msg.Index, len(msg.Entries), raftstate.prevTermFromLog()}

	raftserver.raftnet.Send(msg.SenderId, resp.Encoding())

	// if the commitIdx is newer than myself, I might want to look into update my commit
	// especially when it is a heartbeat message
	if len(msg.Entries) == 0 {
		raftserver.ProcessCommitUpdate(msg.LeaderCommitIdx)
	}
}

func (raftstate *RaftState) ProcessAppendEntriesResp(msg *AppendEntriesResp, raftserver *RaftServer) {
	if msg.Success {
		log.Debugf("Follower %v able to append to its own %v\n", msg.SenderId, msg.Repr())
		// establish the consensus here
		raftstate.replicatedIdx[msg.SenderId] = msg.Index + msg.NumOfEntries - 1

		// determine the highest replicatedIDX which has 3 or more shows, including leader itself
		// bookkeeping of leader itself happens in LeaderAppendEntries()
		newCommitIdx := DetermineCommitIdx(raftstate.replicatedIdx)
		raftserver.lock.Lock()
		log.Debugf("lock: %v is locked, raftserver.commitIdx is %v, newCommitIdx is %v", raftserver.lock, raftserver.raftstate.commitIdx, newCommitIdx)
		if raftstate.commitIdx < newCommitIdx && raftstate.raftlog.items[newCommitIdx].Term == raftserver.raftstate.currentTerm {

			raftstate.commitIdx = newCommitIdx

			// we then should tell followers of this update
			// using heartbeat message can be an option
			// now just use a CommitUpdate msg, can incorporate with heartbeat if that is how paper describes
			// actually with heartbeat, this is unnecessary, wait until I have heartbeat setup,
			// evaluate if I can remove this
			// we'd better remove this actually, two sources of commitIdx cause unnecessary contention
			/**
			commitUpdateMsg := CommitUpdate{raftserver.myID, raftserver.commitIdx}
			for _, f := range raftserver.followerIDs {
				raftserver.raftnet.Send(f, commitUpdateMsg.Encoding())
			}
			**/

		}
		raftserver.lock.Unlock()
		log.Debugf("lock: %v is unlocked", raftserver.lock)

		// we should now signla checkForCommit channel
		// any client/goroutine waiting for its entry to be committed can now move on
		raftserver.cond.Broadcast()

	} else {
		log.Errorf("Follower %v not able to append to its own %v, back tracking\n", msg.SenderId, msg.Repr())
		prevTerm := 0
		newIndex := msg.Index - 1
		if newIndex > 0 {
			prevTerm = raftstate.raftlog.items[newIndex-1].Term
		}
		backoffMsg := AppendEntriesMsg{raftserver.myID, raftstate.currentTerm, raftstate.commitIdx, newIndex, prevTerm, raftstate.raftlog.items[newIndex:]}
		// two places sending to follower
		// would there be some refactoring
		// conditional send to follower: condition being if index<follower.commitIdx, no need to send any more...
		raftserver.raftnet.Send(msg.SenderId, backoffMsg.Encoding())
	}

}

func (raftstate *RaftState) LeaderAppendEntries(index, prevTerm int, entries []RaftLogEntry, raftserver *RaftServer) bool {
	success := raftstate.raftlog.AppendEntries(index, prevTerm, entries)

	if success {
		// book keeping leader itself
		raftstate.replicatedIdx[raftserver.myID] = index

		// send AppendEntries Msg to followers
		for _, f := range raftserver.followerIDs {
			m := CreateAppendEntriesMsg(raftstate.myId, raftstate.currentTerm, raftstate.commitIdx, index, prevTerm, entries)
			// simple fixed length type field for now, of course can wrap this with even one more layer.
			// not the major focus
			log.Debugf("Sending raftnode:%v, msg: %v", f, m.Repr())
			raftserver.raftnet.Send(f, m.Encoding())
		}

	}
	return success
}
