package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kangxie-colorado/golang-primer/messaging/lib"
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
	log.SetReportCaller(true)
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

/**
// test with disable/enable network links
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
	raftnets[2].Enable(0)

	time.Sleep(3 * time.Second)

	for i := 0; i < 5; i++ {
		for time := 0; time < 5; time++ {
			fmt.Println(raftnets[i].Receive())
		}
	}
}
**/

/***
// test with multi-clients
func main() {

	prog := "server"

	if len(os.Args) > 1 {
		prog = os.Args[1]
	}

	if prog == "server" {
		initLog("server.log", log.InfoLevel)
		log.Debugln("********************************************************************************************")
		lib_test.KVServer(transport.SocketDescriptor{"tcp", "localhost", "15000"})
	} else if prog == "client" {
		clientID := 1
		if len(os.Args) > 2 {
			clientID, _ = strconv.Atoi(os.Args[2])
		}

		initLog("client"+strconv.Itoa(clientID)+".log", log.InfoLevel)
		log.Debugln("********************************************************************************************")

		kvclient := lib_test.CreateKVClient(transport.SocketDescriptor{"tcp", "localhost", "15000"})
		if clientID == 1 {
			kvclient.Get("foo")

			kvclient.Set("foo", "bar")
			time.Sleep(3 * time.Second)
			kvclient.Get("foo")
		} else {
			kvclient.Get("foo")

			kvclient.Set("foo", "bar24")
			kvclient.Get("foo2")
			kvclient.Set("foo2", "bar245")
			kvclient.Get("foo")
			kvclient.Del("foo")

		}

	} else {
		fmt.Println("Wrong program type!")
	}
}
**/

/***
// gobtest

****************************************************************************************
// [tw-mbp-kxie test (master)]$ go run *go server
// Qv+BAwEBEEFwcGVuZEVudHJpZXNNc2cB/4IAAQMBBUluZGV4AQQAAQhQcmV2VGVybQEEAAEHRW50cmllcwH/hgAAACH/hQIBARJbXWxpYi5SYWZ0TG9nRW50cnkB/4YAAf+EAAAs/4MDAQEMUmFmdExvZ0VudHJ5Af+EAAECAQRUZXJtAQQAAQRJdGVtARAAAAA4/4IDAgIGc3RyaW5nDA0AC1NFVCBGT08gQkFSAAECAQZzdHJpbmcMDgAMU0VUIEZPTyBCQVIyAAA=
// {0 0 [{0 SET FOO BAR}]}
// {0 0 [{0 SET FOO BAR} {1 SET FOO BAR2}]}
// Qv+BAwEBEEFwcGVuZEVudHJpZXNNc2cB/4IAAQMBBUluZGV4AQQAAQhQcmV2VGVybQEEAAEHRW50cmllcwH/hgAAACH/hQIBARJbXWxpYi5SYWZ0TG9nRW50cnkB/4YAAf+EAAAs/4MDAQEMUmFmdExvZ0VudHJ5Af+EAAECAQRUZXJtAQQAAQRJdGVtARAAAABV/4IBAgIDAQQBBnN0cmluZwwOAAxTRVQgRk9PIEJBUjMAAQYBBnN0cmluZwwQAA5TRVQgRk9PMiBCQVIyMwABCAEGc3RyaW5nDAkAB0RFTCBGT08AAA==
// {1 0 [{2 SET FOO BAR3} {3 SET FOO2 BAR23} {4 DEL FOO}]}
****************************************************************************************

func main() {
	gob.Register(lib.AppendEntriesMsg{})

	m := lib.CreateAppendEntriesMsg(0, 0, []lib.RaftLogEntry{lib.CreateRaftLogEntry(0, "SET FOO BAR"), lib.CreateRaftLogEntry(1, "SET FOO BAR2")})
	fmt.Println(lib.ToGOB64(&m))

	// upto SET FOO BAR
	m1 := lib.FromGOB64("Qv+BAwEBEEFwcGVuZEVudHJpZXNNc2cB/4IAAQMBBUluZGV4AQQAAQhQcmV2VGVybQEEAAEHRW50cmllcwH/hgAAACH/hQIBARJbXWxpYi5SYWZ0TG9nRW50cnkB/4YAAf+EAAAs/4MDAQEMUmFmdExvZ0VudHJ5Af+EAAECAQRUZXJtAQQAAQRJdGVtARAAAAAd/4IDAQIGc3RyaW5nDA0AC1NFVCBGT08gQkFSAAA=")
	fmt.Printf("%v\n", m1)

	// upto SET FOO BAR2
	m2 := lib.FromGOB64("Qv+BAwEBEEFwcGVuZEVudHJpZXNNc2cB/4IAAQMBBUluZGV4AQQAAQhQcmV2VGVybQEEAAEHRW50cmllcwH/hgAAACH/hQIBARJbXWxpYi5SYWZ0TG9nRW50cnkB/4YAAf+EAAAs/4MDAQEMUmFmdExvZ0VudHJ5Af+EAAECAQRUZXJtAQQAAQRJdGVtARAAAAA4/4IDAgIGc3RyaW5nDA0AC1NFVCBGT08gQkFSAAECAQZzdHJpbmcMDgAMU0VUIEZPTyBCQVIyAAA=")
	fmt.Printf("%v\n", m2)

	m3 := lib.CreateAppendEntriesMsg(1, 0, []lib.RaftLogEntry{lib.CreateRaftLogEntry(2, "SET FOO BAR3"), lib.CreateRaftLogEntry(3, "SET FOO2 BAR23"), lib.CreateRaftLogEntry(4, "DEL FOO")})
	fmt.Println(lib.ToGOB64(&m3))

	m4 := lib.FromGOB64("Qv+BAwEBEEFwcGVuZEVudHJpZXNNc2cB/4IAAQMBBUluZGV4AQQAAQhQcmV2VGVybQEEAAEHRW50cmllcwH/hgAAACH/hQIBARJbXWxpYi5SYWZ0TG9nRW50cnkB/4YAAf+EAAAs/4MDAQEMUmFmdExvZ0VudHJ5Af+EAAECAQRUZXJtAQQAAQRJdGVtARAAAABV/4IBAgIDAQQBBnN0cmluZwwOAAxTRVQgRk9PIEJBUjMAAQYBBnN0cmluZwwQAA5TRVQgRk9PMiBCQVIyMwABCAEGc3RyaW5nDAkAB0RFTCBGT08AAA==")
	fmt.Printf("%v\n", m4)

}
***/

func main() {

	prog := "server"

	if len(os.Args) > 1 {
		prog = os.Args[1]
	}

	if prog == "server" {
		raftID := 0
		if len(os.Args) > 2 {
			raftID, _ = strconv.Atoi(os.Args[2])
		}
		initLog("server"+strconv.Itoa(raftID)+".log", log.DebugLevel)
		log.Debugln("********************************************************************************************")

		port := 25000 + raftID
		lib_test.KVServer(lib.SocketDescriptor{"tcp", "localhost", strconv.Itoa(port)}, raftID)

	} else if prog == "client" {
		clientID := 1
		if len(os.Args) > 2 {
			clientID, _ = strconv.Atoi(os.Args[2])
		}

		initLog("client"+strconv.Itoa(clientID)+".log", log.InfoLevel)
		log.Debugln("********************************************************************************************")

		kvclient := lib_test.CreateKVClient(lib.SocketDescriptor{"tcp", "localhost", "25000"})
		if clientID == 1 {
			kvclient.Get("foo")

			kvclient.Set("foo", "bar")
			time.Sleep(3 * time.Second)
			kvclient.Get("foo")
		} else {
			kvclient.Get("foo")

			kvclient.Set("foo", "bar24")
			kvclient.Get("foo2")
			kvclient.Set("foo2", "bar245")
			kvclient.Get("foo")
			kvclient.Del("foo")

		}

	} else {
		fmt.Println("Wrong program type!")
	}
}
