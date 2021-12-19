package lib

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type RaftMessage interface {
	Encoding() string
	Decoding(string)
	Repr() string
}

type AppendEntriesMsg struct {
	SenderId        int
	LeaderTerm      int
	LeaderCommitIdx int

	Index    int
	PrevTerm int
	Entries  []RaftLogEntry
}

func (m *AppendEntriesMsg) Repr() string {
	return fmt.Sprintf("AppendEntriesMsg{SenderId=%v, LeaderTerm=%v, LeaderCommitIdx=%v, Index=%v, PrevTerm=%v, Entries=[]RaftLogEntry{%v}}",
		m.SenderId, m.LeaderTerm, m.LeaderCommitIdx, m.Index, m.PrevTerm, m.Entries)
}

type AppendEntriesResp struct {
	SenderId     int
	Success      bool
	Index        int
	NumOfEntries int // index + NumOfEntries is the followers' latest index pos
	Term         int
}

func (m *AppendEntriesResp) Repr() string {
	return fmt.Sprintf("AppendEntriesResp{SenderId=%v, Success=%v, Index=%v, NumOfEntries=%v, Term=%v}", m.SenderId, m.Success, m.Index, m.NumOfEntries, m.Term)
}

// for test convienience ony?
func CreateAppendEntriesMsg(sender, leaderTerm, leaderCommitIdx, index, prevTerm int, entries []RaftLogEntry) AppendEntriesMsg {
	return AppendEntriesMsg{sender, leaderTerm, leaderCommitIdx, index, prevTerm, entries}
}

type CommitUpdate struct {
	SenderId  int
	CommitIdx int
}

func (m *CommitUpdate) Repr() string {
	return fmt.Sprintf("CommitUpdate{SenderId=%v, CommitIdx=%v}", m.SenderId, m.CommitIdx)

}

func encoding(m interface{}) bytes.Buffer {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		log.Errorln("failed gob Encode", err)
	}

	return b
}

func getDecoder(str string) *gob.Decoder {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Errorln("failed base64 Decode", err)
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)

	return d
}

func (m *AppendEntriesMsg) Encoding() string {
	b := encoding(m)
	return APPENDENTRYMSG + base64.StdEncoding.EncodeToString(b.Bytes())
}

// caller allocates the memory
/** calling pattern?
	str[:6] will be the type
	append = AppendEntriesMsg{}
	append.Decoding(str[6:])
**/
func (m *AppendEntriesMsg) Decoding(str string) {
	d := getDecoder(str)
	err := d.Decode(&m)
	if err != nil {
		log.Errorln("failed gob Decode", err)
	}
}

// what is the better way to do this code sharing
// this polymorphism in golang?
// I haven't studied this yet - now, keep it duplicae and simple, but only in this file
func (m *AppendEntriesResp) Encoding() string {
	b := encoding(m)
	return APPENDENTRYRSP + base64.StdEncoding.EncodeToString(b.Bytes())
}

// caller allocates the memory
func (m *AppendEntriesResp) Decoding(str string) {
	d := getDecoder(str)
	err := d.Decode(&m)
	if err != nil {
		log.Errorln("failed gob Decode", err)
	}
}

func (m *CommitUpdate) Encoding() string {
	b := encoding(m)
	return COMMITUPDATE + base64.StdEncoding.EncodeToString(b.Bytes())
}

// caller allocates the memory
func (m *CommitUpdate) Decoding(str string) {
	d := getDecoder(str)
	err := d.Decode(&m)
	if err != nil {
		log.Errorln("failed gob Decode", err)
	}
}
