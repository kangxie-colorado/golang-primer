package lib

type AppendEntriesMsg struct {
	Index, PrevTerm int
	Entries         []RaftLogEntry
}

func CreateAppendEntriesMsg(index, prevTerm int, entries []RaftLogEntry) AppendEntriesMsg {
	return AppendEntriesMsg{index, prevTerm, entries}
}

// this needs some delicate serialize/deserialize scheme
// if using string here? that would really not so easy
// learn or build?
