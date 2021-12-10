package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kangxie-colorado/golang-primer/messaging/lib"
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
// quick test a simple server , clien->conn
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
// echo server test
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
// kv server test
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

/***
func main() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	var raftnets = make([]lib.RaftNet, 5)

	for i := 0; i < 5; i++ {
		raftnets[i] = *lib.CreateARaftNet(i)
		go raftnets[i].Start()
	}

	// disable 0->2 but and 2->0
	//
	raftnets[0].Disable(2)
	raftnets[2].Disable(0)
	time.Sleep(3 * time.Second)

	for i := 0; i < 5; i++ {
		raftnets[i].BeginSending()
	}

	for i := 0; i < 5; i++ {
		for to := 0; to < 5; to++ {
			raftnets[i].Send(to, fmt.Sprintf("msg %v to %v", i, to))
			time.Sleep(3 * time.Millisecond)
		}
	}

	for i := 0; i < 5; i++ {
		for time := 0; time < 5; time++ {
			fmt.Println(raftnets[i].Receive())
		}
	}
}
***/

/**
// start net per process
func main() {
	id, _ := strconv.Atoi(os.Args[1])
	lib.CreateARaftNet(id).Start()

}
**/

func main() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	var raftnets = make([]lib.RaftNet, 5)

	for i := 0; i < 5; i++ {
		raftnets[i] = *lib.CreateARaftNet(i)
		go raftnets[i].Start()
	}

	// disable 0->2 but and 2->0
	//
	raftnets[0].Disable(2)
	raftnets[2].Disable(0)
	time.Sleep(3 * time.Second)

	for i := 0; i < 5; i++ {
		raftnets[i].BeginSending()
	}

	for i := 0; i < 5; i++ {
		for to := 0; to < 5; to++ {
			raftnets[i].Send(to, fmt.Sprintf("msg %v to %v", i, to))
			time.Sleep(3 * time.Millisecond)
		}
	}

	raftnets[0].Enable(2)
	time.Sleep(3 * time.Second)

	for i := 0; i < 5; i++ {
		for time := 0; time < 5; time++ {
			fmt.Println(raftnets[i].Receive())
		}
	}
}
