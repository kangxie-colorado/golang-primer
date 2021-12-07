package lib

import (
	"fmt"
	"net"
	"os"

	"github.com/kangxie-colorado/golang-primer/messaging/lib"
	log "github.com/sirupsen/logrus"
)

func CreateKVClient(sock lib.SocketDescriptor) KVClient {
	log.Infoln("Staring KV client")

	log.Infoln("Connecting to " + sock.ConnType + sock.ConnHost + ":" + sock.ConnPort)

	conn, err := net.Dial(sock.ConnType, sock.ConnHost+":"+sock.ConnPort)
	if err != nil {
		log.Errorln("Error Conecting:", err.Error())
		os.Exit(1)
	}

	log.Infoln("Connected: ", conn.RemoteAddr(), conn.LocalAddr())
	return KVClient{conn}
}

type KVClient struct {
	conn net.Conn
}

func (cl KVClient) Set(key, value string) {
	payload := "SET" + encodeKeyValue(key, value)
	lib.SendMessage(cl.conn, []byte(payload))

	resp, _ := lib.RecvMessageStr(cl.conn)
	fmt.Println(resp)
}

func (cl KVClient) Get(key string) {
	payload := "GET" + encodeKey(key)
	lib.SendMessage(cl.conn, []byte(payload))

	resp, _ := lib.RecvMessageStr(cl.conn)
	fmt.Println(resp)
}

func (cl KVClient) Del(key string) {
	payload := "DEL" + encodeKey(key)
	lib.SendMessage(cl.conn, []byte(payload))

	resp, _ := lib.RecvMessageStr(cl.conn)
	fmt.Println(resp)
}
