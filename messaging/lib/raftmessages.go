package lib

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	log "github.com/sirupsen/logrus"
)

type RaftMessage interface {
	Encoding() string
	Decoding(string)
}

type AppendEntriesMsg struct {
	SenderId int
	Index    int
	PrevTerm int
	Entries  []RaftLogEntry
}

type AppendEntriesResp struct {
	SenderId     int
	Success      bool
	Index        int
	NumOfEntries int // index + NumOfEntries is the followers' latest index pos
	Term         int
}

func CreateAppendEntriesMsg(sender, index, prevTerm int, entries []RaftLogEntry) AppendEntriesMsg {
	return AppendEntriesMsg{sender, index, prevTerm, entries}
}

// this needs some delicate serialize/deserialize scheme
// if using string here? that would really not so easy
// learn or build?

func (append *AppendEntriesMsg) Encoding() string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(append)
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
func (append *AppendEntriesMsg) Decoding(str string) {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Errorln("failed base64 Decode", err)
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&append)
	if err != nil {
		log.Errorln("failed gob Decode", err)
	}
}

// what is the better way to do this code sharing
// this polymorphism in golang?
//
func (append *AppendEntriesResp) Encoding() string {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(append)
	if err != nil {
		log.Errorln("failed gob Encode", err)
	}
	return APPENDENTRYRSP + base64.StdEncoding.EncodeToString(b.Bytes())
}

// caller allocates the memory
func (append *AppendEntriesResp) Decoding(str string) {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Errorln("failed base64 Decode", err)
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&append)
	if err != nil {
		log.Errorln("failed gob Decode", err)
	}
}
