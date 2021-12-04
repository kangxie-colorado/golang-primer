package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

var room_enter = make(chan int)
var room_exit = make(chan int)
var done = make(chan int)

func waitForRoom(id int, chForward chan int, chBack chan int) {
	result := <-chBack
	for result != 1 {
		chForward <- id
		result = <-chBack
	}
}

func waitForFork(id int, chForward chan PhiloAction, chBack chan int) {
	result := <-chBack
	for result != 1 {
		chForward <- PhiloAction{id, 1}
		result = <-chBack
	}
}

func phil(id int) {
	for i := 1; i < 11; i++ {
		do(id, "think", i)

		// enter the room
		room_enter <- id
		waitForRoom(id, room_enter, room_confirm_chans[id])

		// pick up left fork
		fork_chans[id] <- PhiloAction{id, 1}
		waitForFork(id, fork_chans[id], philo_chans[id])

		// pick up right fork
		fork_chans[(id+1)%5] <- PhiloAction{id, 1}
		waitForFork(id, fork_chans[(id+1)%5], philo_chans[id])

		do(id, "eat", i)

		// put it down
		fork_chans[id] <- PhiloAction{id, 0}
		fork_chans[(id+1)%5] <- PhiloAction{id, 0}

		// exit the room
		room_exit <- id
	}

	done <- id
}

type PhiloAction struct {
	id     int // philo's id
	action int // pickup:1 or putdown:0
}

var fork_chans = []chan PhiloAction{
	make(chan PhiloAction),
	make(chan PhiloAction),
	make(chan PhiloAction),
	make(chan PhiloAction),
	make(chan PhiloAction),
}

var philo_chans = []chan int{
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
}

var room_confirm_chans = []chan int{
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
	make(chan int),
}

func fork(id int) {
	// forks are not picked up initially
	my_state := 0
	my_user := -1
	for {
		philoAction := <-fork_chans[id]
		if my_state == 0 && philoAction.action == 1 {
			// pickup, send ack back
			my_state = 1
			my_user = philoAction.id
			philo_chans[philoAction.id] <- 1
		} else if my_state == 1 && philoAction.action == 0 && my_user == philoAction.id {
			// putdown
			my_state = 0
			my_user = -1
		} else if my_state == 1 && philoAction.action == 1 && my_user != philoAction.id {
			// the contention, someone else tries to pick it up
			// they should block on their side
			if rand.Int31n(300000) > 299998 {
				fmt.Printf("Phlio %d tries to pick up fork %d, but it is being used by %d, he has to wait\n", philoAction.id, id, my_user)
			}
			philo_chans[philoAction.id] <- 0
		} else {
			errorString := fmt.Sprintf("Not right, my_state: %d, my users: %d, attempting Philo: %d, attempting action: %d\n", my_state, my_user, philoAction.id, philoAction.action)
			panic(errorString)
		}

		//fmt.Printf("my_id: %d, my_state: %d, my users: %d, attempting Philo: %d, attempting action: %d\n", id, my_state, my_user, philoAction.id, philoAction.action)
	}
}

func room() {
	occupancy := 0
	for {
		select {
		case someone_enters := <-room_enter:
			if occupancy >= 4 {
				// block until someone left
				if rand.Int31n(30000) > 29998 {
					// reduce printing this message
					indentedPrint(someone_enters, "Philo", someone_enters, "want to enter, but it will be full, waiting...")
				}

				room_confirm_chans[someone_enters] <- 0
			} else {
				room_confirm_chans[someone_enters] <- 1
				occupancy += 1
				//indentedPrint(someone_enters, "Philo", someone_enters, "enters the room.", "occupancy is", occupancy)
			}
		case <-room_exit:
			occupancy -= 1
			//indentedPrint(someone_exits, "Phlio", someone_exits, "exits the room.", "occupancy is", occupancy)
		}

	}
}

/*
func main() {
	for id := 0; id < 5; id++ {
		go phil(id)
		go fork(id)
	}

	go room()

	for id := 0; id < 5; id++ {
		fmt.Println(<-done, "DONE")
	}

}
*/
