package lib

import (
	"net"

	"github.com/enriquebris/goconcurrentqueue"
)

type SocketDescriptor struct {
	ConnType string
	ConnHost string
	ConnPort string
}

var RaftNetConfig = map[int]SocketDescriptor{
	0: SocketDescriptor{"tcp", "localhost", "15000"},
	1: SocketDescriptor{"tcp", "localhost", "15001"},
	2: SocketDescriptor{"tcp", "localhost", "15002"},
	3: SocketDescriptor{"tcp", "localhost", "15003"},
	4: SocketDescriptor{"tcp", "localhost", "15004"},
}

const NetWorkSize int = 5

type RaftNet struct {
	Id          int
	inbox       *goconcurrentqueue.FIFO
	outboxs     [NetWorkSize]*goconcurrentqueue.FIFO // sent to myself? just discard in the sender
	outlinks    [NetWorkSize]net.Conn
	enabled     [NetWorkSize]bool
	linkEnabled [NetWorkSize]chan bool
	activated   [NetWorkSize]chan bool
}

const MSGTYPEFIELDLEN int = 11
const APPENDENTRYMSG string = "APPENDENTRY"
const APPENDENTRYRSP string = "APPENDRESPS"
const COMMITUPDATE string = "COMMITUPDAT"
