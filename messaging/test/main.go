package main

import (
	"fmt"
	"os"
	"strconv"

	transport "github.com/kangxie-colorado/golang-primer/messaging/lib"
	lib_test "github.com/kangxie-colorado/golang-primer/messaging/test/lib"
	log "github.com/sirupsen/logrus"
)

func initLog(filename string, logLevel log.Level) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(file)
	log.SetLevel(logLevel)
	//log.SetReportCaller(true)
}

/**
func main() {
	prog := os.Args[1]

	if prog == "server" {
		SimpleServer(transport.SocketDescriptor{"tcp", "localhost", "15000"})
	} else if prog == "client" {
		sz, _ := strconv.Atoi(os.Args[2])
		SimpleClient(transport.SocketDescriptor{"tcp", "localhost", "15000"}, sz)
	} else {
		fmt.Println("Wrong program type!")
	}
}
**/

/**
func main() {
	prog := os.Args[1]

	if prog == "server" {
		initLog("server.log", log.DebugLevel)
		lib_test.EchoServer(transport.SocketDescriptor{"tcp", "localhost", "15000"})
	} else if prog == "client" {
		initLog("client.log", log.InfoLevel)
		lib_test.EchoClient(transport.SocketDescriptor{"tcp", "localhost", "15000"})
	} else {
		fmt.Println("Wrong program type!")
	}
}
**/

/**
func main() {
	prog := os.Args[1]

	if prog == "server" {
		initLog("server.log", log.DebugLevel)
		lib_test.KVServer(transport.SocketDescriptor{"tcp", "localhost", "15000"})
	} else if prog == "client" {
		initLog("client.log", log.InfoLevel)
		kvclient := lib_test.CreateKVClient(transport.SocketDescriptor{"tcp", "localhost", "15000"})
		kvclient.Get("foo")

		kvclient.Set("foo", "bar")
		kvclient.Get("foo")

	} else {
		fmt.Println("Wrong program type!")
	}
}
**/

func main() {
	prog := os.Args[1]

	if prog == "server" {
		lib_test.SimpleServer(transport.SocketDescriptor{"tcp", "localhost", "15000"})
	} else if prog == "client" {
		sz, _ := strconv.Atoi(os.Args[2])
		lib_test.SimpleClient(transport.SocketDescriptor{"tcp", "localhost", "15000"}, sz)
	} else {
		fmt.Println("Wrong program type!")
	}
}
