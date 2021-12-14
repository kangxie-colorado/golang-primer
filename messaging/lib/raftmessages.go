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
	SenderId int
	Index    int
	PrevTerm int
	Entries  []RaftLogEntry
}

func (m *AppendEntriesMsg) Repr() string {
	return fmt.Sprintf("AppendEntriesMsg{SenderId=%v, Index=%v, PrevTerm=%v, Entries=[]RaftLogEntry{%v}}", m.SenderId, m.Index, m.PrevTerm, m.Entries)
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

func CreateAppendEntriesMsg(sender, index, prevTerm int, entries []RaftLogEntry) AppendEntriesMsg {
	return AppendEntriesMsg{sender, index, prevTerm, entries}
}

// this needs some delicate serialize/deserialize scheme
// if using string here? that would really not so easy
// learn or build?

func (m *AppendEntriesMsg) Encoding() string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		log.Errorln("failed gob Encode", err)
	}
	return APPENDENTRYMSG + base64.StdEncoding.EncodeToString(b.Bytes())
}

// caller allocates the memory
/** calling pattern?
	str[:6] will be the type
	append = AppendEntriesMsg{}
	append.Decoding(str[6:])
**/
func (m *AppendEntriesMsg) Decoding(str string) {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Errorln("failed base64 Decode", err)
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil {
		log.Errorln("failed gob Decode", err)
	}
}

// what is the better way to do this code sharing
// this polymorphism in golang?
//
func (m *AppendEntriesResp) Encoding() string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		log.Errorln("failed gob Encode", err)
	}
	return APPENDENTRYRSP + base64.StdEncoding.EncodeToString(b.Bytes())
}

// caller allocates the memory
func (m *AppendEntriesResp) Decoding(str string) {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Errorln("failed base64 Decode", err)
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil {
		log.Errorln("failed gob Decode", err)
	}
}
