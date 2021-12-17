package lib

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/kangxie-colorado/golang-primer/messaging/lib"
	log "github.com/sirupsen/logrus"
)

var mu = sync.Mutex{}
var kvstore = make(map[string]string)

var raftserver *lib.RaftServer

func appendEntry(raft *lib.RaftServer, msg string, commited chan bool) {
	raft.AppendNewEntry(msg, commited)
}

func Set(msg string, cl net.Conn) {
	key, value := parseKeyValue(msg)
	log.Infoln("Calling SET key:", key, "to value:", value)

	log.Infof("Send SET %v:%v to raft\n", key, value)
	commited := make(chan bool)
	go appendEntry(raftserver, msg, commited)
	<-commited
	log.Infoln("Raft has committed")

	// this will need some lock to protect against race conditions
	mu.Lock()
	defer mu.Unlock()
	kvstore[key] = value
	lib.SendMessageStr(cl, "OK")
}

func Get(msg string, cl net.Conn) {
	key := parseKey(msg)
	log.Infoln("Calling GET key", key)

	mu.Lock()
	defer mu.Unlock()
	if val, ok := kvstore[key]; ok {
		lib.SendMessageStr(cl, val)
	} else {
		lib.SendMessageStr(cl, "Key not exsitent")
	}
}

func Del(msg string, cl net.Conn) {
	key := parseKey(msg)
	log.Infoln("Calling DEL key", key)

	log.Infof("Send DEL %v to raft\n", key)
	commited := make(chan bool)
	go appendEntry(raftserver, msg, commited)
	<-commited

	mu.Lock()
	defer mu.Unlock()
	if _, ok := kvstore[key]; ok {
		delete(kvstore, key)
		lib.SendMessageStr(cl, "Deleted")
	} else {
		lib.SendMessageStr(cl, "Key not exsitent")
	}
}

func HandleKVClient(cl net.Conn) {
	for {
		msgBytes, err := lib.RecvMessage(cl)
		if err != nil {
			log.Errorln("Error when receiving messages", err.Error())
			break
		}

		// this kind of waiting is not ideal
		// but net.Conn.Read() won't block so I don't know a better way yet
		// actually, EOF might got ahead of this block: EOF will cause the err!=nil to fire
		if msgBytes == nil {
			log.Infoln("Received nothing, sleep for 1s")
			time.Sleep(1 * time.Second)
			continue
		}

		msg := string(msgBytes)
		handleKVMsg(msg, cl)

	}
}

func handleKVMsg(msg string, cl net.Conn) {
	switch msg[:3] {
	case "SET":
		Set(msg, cl)

	case "GET":

		Get(msg, cl)

	case "DEL":

		Del(msg, cl)

	default:
		log.Errorln("Methond Unknown!")
	}
}

func playLogSet(msg string) {
	key, value := parseKeyValue(msg)
	log.Infoln("Playing log, SET key:", key, "to value:", value)

	kvstore[key] = value
}

func playLogDel(msg string) {
	key := parseKey(msg)
	log.Infoln("Playing log, DEL key", key)

	if _, ok := kvstore[key]; ok {
		delete(kvstore, key)
	} else {
		log.Infoln("Key not existent!")
	}
}

func handleRaftLog(msg string) {
	log.Debugf("Reading off raft log %v", msg)
	switch msg[:3] {
	case "SET":
		playLogSet(msg)

	case "DEL":
		playLogDel(msg)

	default:
		log.Errorln("Methond Unknown!")
	}
}

func RaftCallback(entires []lib.RaftLogEntry) {
	for _, e := range entires {
		msg := fmt.Sprintf("%v", e.Item)
		handleRaftLog(msg)
	}
}

func KVServer(sock lib.SocketDescriptor, raftserverID int) {
	// this part hooks up with raft
	// arbitrarily appoint raftserver 0 as the leader
	// how do I do that? actually at this point, there is no leader, every raftnode just append to itself and then sending to others
	// but only one server got the input, so purpose is served
	raftserver = lib.CreateARaftServer(raftserverID, RaftCallback)
	log.Infoln("Starting Raft Server")
	raftserver.Start()
	// don't forgot this gating method
	raftserver.Net().BeginSending()

	// below was without raft
	log.Infoln("Staring KV server")
	ln, err := net.Listen(sock.ConnType, sock.ConnHost+":"+sock.ConnPort)
	if err != nil {
		log.Errorf("Cannot listen on %+v\n", sock)
		os.Exit(1)
	}

	defer ln.Close()

	for {
		cl, err := ln.Accept()
		if err != nil {
			log.Errorln("Error connecting:", err.Error())
			return
		}
		log.Infoln("Client Connected: ", cl.RemoteAddr().String())

		go HandleKVClient(cl)
	}

}
