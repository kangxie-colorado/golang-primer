// https://leetcode.com/problems/2-keys-keyboard/

package main

import (
	"fmt"
	"math"
)

/*
	analysis:
	of course, the most steps you need is n-1; only copy once, then paste

	but the ask is min step
	first observation is each time you copy, you can not go back

	A: copy (buffer:A)
	paste -> AA (buffer: A)
	copy|paste?

	Also notice when you copy, the number of A doesn't change until you paste; but the buffer changes.
	so there is a buffer state

	if design this like a N steps
	N[1] N[2] .. N[n]

	N[k] will be either
	- N[k] = 2*N[k-1] <-- copy then paste
	- N[k] = N[k-1] + N[k-2] <-- paste N[k-2] twice...(buffer start with zero)
		makes the assumption, if you keeps copying but no pasting, that cannot be the min step so don't consider that
^^^ not correct maybe (the 2nd one)

	okay, with this formula, I can actually search back
	or DP, fill forward
	A
	next step
	1: AA

	next step
	2: AAA
	2: AAAA

	next step
	3: AAAA
	3: AAAAAA
	3: AAAAAA
	3: AAAAAAAAA

	what states to maintain
	1. previous numbers - what is in the buffer
	2. current numbers - refresh the buffer
	at any certian step, you can either paste the buffer or refresh the buffer then copy ( not sure if copy counts an operation.. it does look like it should)

	DP.. when there is a better way emerging we can substitue
	so this suddenly becaomes another distra game?

	but you notice that you cannot go to previous.. that when you buffer changed into 3, you cannot go back to using 2
	maintain the link...(state: count->{buffer, now, steps} => count->{buffer, now, steps})

	yeah, maybe with these extra states tracking, this is solvable
	let me try this then backwards

	then let me think over the contrains
		Constraints:
			1 <= n <= 1000

min_steps(n)
	type State struct {
		buffer int
		steps int
	}

	tracking = map[int] State{
		1 : State{
			buffer: 0
			steps: 0
		}
	}


^^^ paper evolution shows this is very messy,
	hard to code it up
	of course everytime, I could replace the frontline using a new one... it could work and I could even try
	but there seems to be no priority queue place so the performance may suffer... could try still

	but now lets go back to the top-down method

	N[k] will be either
	- N[k] = 2*N[k-1] <-- copy then paste (2 ops in this one: copy+paste)
	- N[k] = N[k-1] + buffer...(buffer start with zero)  (1 op in this one: paste)
		buffer can be from N[k-2] time or N[k-2] itself
		buffer can be any possible N[0..k-2], hmmm....

	is this correct?
	yes, it looks correct
	but the buffer can be so many values makes things complicated


^^^ another mind trap
	okay I must be missing something, it is okay to let go ego and return/repeat
	even I am kind of stuck I see the formula is actually a joke it is actually

	N[k] = N[k-1] + buffer
		buffer being N[0..k-1]
		change buffer need an extra op

	that, buffer being N[0..k-1] is also wrong
	there are paths that will not cross...

	oh...

still going back to frontline-expand solution, even it is a brute-force one
just code it up.. let me see

*/

type State struct {
	number int
	steps  int
	buffer int
}

func _minSteps(n int) int {
	start := State{
		1,
		0,
		0,
	}

	tracking := []State{start}
	for true {
		nextTracking := []State{}
		for _, s := range tracking {
			if s.number == n {
				return s.steps
			}

			// copy myself, paste myself and steps + 2
			next1 := State{
				s.number * 2,
				s.steps + 2,
				s.number,
			}
			nextTracking = append(nextTracking, next1)

			// continue to paste the previous buffer
			// empty paste is a waste of steps (only copy no paste is too... so ignore that)
			if s.buffer == 0 {
				continue
			}

			next2 := State{
				s.number + s.buffer,
				s.steps + 1,
				s.buffer,
			}
			nextTracking = append(nextTracking, next2)

			// I already see an optimization, this overlap
			// but can I, that will increase the complexicty of algorithem a lot because it then become a priority queue like issue

		}
		tracking = nextTracking
	}
	return -1
}

/*
	Damn myself. the above is the BFS...
	and yes, this is actually the BFS... to search for shortest path?
	and above might be logically right but resource forbidden

	28 / 126 test cases passed.
		Runtime Error Message:
		runtime: out of memory: cannot allocate 109051904-byte block (440041472 in use)
		fatal error: out of memory
		runtime.throw({0x4c9756, 0x3289})
*/

/**
	That naive BFS search gives me some clue
	lead to case research

	29 -> 29, it blows out the memory... hmm.. prime number
	9 -> 6 (1->2 2ops, 2->3, 1 ops, 3->6 2 ops, 6->9 1 ops, total: 6 ops)
	8 -> 6 (1->2 2ops, 2->4 2ops, 4->8 2ops, total: 6 ops)

	so this problem is kind of math invovled...
	so when you do case study, you can not forget the target, especially when it is a single input... it can lead to good insights

	so I believe
	n=1, 0
	n=2, 2 (copy+paste)
	is a prime.. then n steps

	otherwise, find its biggest fatore (which can be found by finding the smallest factor), then multipled from there
	f(bf) + n/bf  (this will include the step for the making the copy)


**/

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := int(2); i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func biggestFactor(num int) int {

	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return num / i
		}
	}

	return 1
}

func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	if IsPrime(n) {
		return n
	}

	bf := biggestFactor(n)
	return minSteps(bf) + n/bf

}

/**

	Success
	Details
	Runtime: 0 ms, faster than 100.00% of Go online submissions for 2 Keys Keyboard.
	Memory Usage: 1.9 MB, less than 56.41% of Go online submissions for 2 Keys Keyboard.

	I think math proof may lay in the fact that the last buffer cannot be bigger than the biggest factor
	otherwise it will only surpass...

	however, the BFS solution gave me a hint.. and once again I learned building up the state tracking
	obviously this is a DFS or recursive nature problem

	okay the example solution is very tricky but same thing
		func minSteps(n int) int {
		for i := n-1; i > 0; i-- {
			if n%i == 0 {									// this is searching for biggest factore
					return n/i + minSteps(i)
				}
			}
			return 0
		}

this also lets me know, at least research a few simple cases
covering some numbers.. up to 20
if it is a number game...

**/

func testTwoKeysKeyboard() {
	for i := 1; i <= 30; i++ {
		fmt.Println(_minSteps(i))
	}

}
