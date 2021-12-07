package lib

import (
	"net"
	"os"
	"time"

	"github.com/kangxie-colorado/golang-primer/messaging/lib"
	log "github.com/sirupsen/logrus"
)

var kvstore = make(map[string]string)

func Set(key, value string) {
	kvstore[key] = value
}

func Get(key string) string {
	if val, ok := kvstore[key]; ok {
		return val
	} else {
		return "Key not exsitent"
	}
}

func Del(key string) string {
	if _, ok := kvstore[key]; ok {
		delete(kvstore, key)
		return "Deleted"
	} else {
		return "Key not exsitent"
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
			log.Infoln("Received nothing sleep for 1s")
			time.Sleep(1 * time.Second)
			continue
		}

		msg := string(msgBytes)
		switch msg[:3] {
		case "SET":
			key, value := parseKeyValue(msg)
			log.Infoln("Calling SET key:", key, "to value:", value)
			Set(key, value)
			lib.SendMessageStr(cl, "OK")

		case "GET":
			key := parseKey(msg)
			log.Infoln("Calling GET key", key)

			lib.SendMessageStr(cl, Get(key))

		case "DEL":
			key := parseKey(msg)
			log.Infoln("Calling DEL key", key)

			lib.SendMessageStr(cl, Del(key))

		default:
			log.Errorln("Methond Unknown!")
		}
	}
}

func KVServer(sock lib.SocketDescriptor) {
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
