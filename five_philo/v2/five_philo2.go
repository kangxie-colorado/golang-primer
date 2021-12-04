package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PhiloAction struct {
	id     int // philo's id
	action int // pickup:1 or putdown:0
}

var fork_chans = []chan int{
	make(chan int, 1),
	make(chan int, 1), // default is not 1? it is 0?
	make(chan int, 1),
	make(chan int, 1),
	make(chan int, 1),
}

func indentedPrint(id int, a ...interface{}) {
	indents := ""
	for i := 0; i < id; i++ {
		indents += "    "
	}

	fmt.Println(indents, a)
}

func do(id int, what string, ith int) {
	d := rand.Int31n(3000)
	time.Sleep(time.Duration(d) * time.Millisecond)
	status := fmt.Sprintf("Philo %d finishes to  %s for %d-th time", id, what, ith)
	indentedPrint(id, status)
}

func pickup(id int, forkid int) {
	status := fmt.Sprintf("Philo %d tried to pickup fork %d", id, forkid)
	indentedPrint(id, status)
}

var done = make(chan int)

func phil2(id int) {
	for i := 1; i < 100; i++ {
		do(id, "think", i)

		// pick up left fork
		pickup(id, id)
		fork_chans[id] <- id

		// pick up right fork
		pickup(id, (id+1)%5)
		fork_chans[(id+1)%5] <- id

		indentedPrint(id, "Has two forks now")

		do(id, "eat", i)

		// put it down
		<-fork_chans[(id+1)%5]

		<-fork_chans[id]

	}

	done <- id
}

func main() {
	for id := 0; id < 5; id++ {
		go phil2(id)
	}

	for id := 0; id < 5; id++ {
		fmt.Println(<-done, "DONE")
	}

}
