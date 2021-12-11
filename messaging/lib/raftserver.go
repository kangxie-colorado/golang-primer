package lib

import "time"

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

func (raftserver *RaftServer) prevTerm() int {
	if len(raftserver.raftlog.items) == 0 {
		// doesn't matter, append at index 0 always succeeds
		return -1
	}

	return raftserver.raftlog.items[len(raftserver.raftlog.items)-1].term
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
	return true
}

func (raftserver *RaftServer) Start() {
	raftserver.raftnet.Start()

	// start the net as above ^
	// maybe start some book keeping?
}
