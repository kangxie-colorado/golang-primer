/***
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
***/

package main

import (
	"fmt"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

var f1or2 = []int{1}

func fibonacciNth(N int) []int {
	// assume N>=3
	if N < 3 {
		return f1or2
	}

	twoBefore := f1or2
	oneBefore := f1or2
	nthFib := []int{}
	for i := 3; i <= N; i++ {
		nthFib = libs.AddTwoIntSlices(twoBefore, oneBefore)
		twoBefore = oneBefore
		oneBefore = nthFib
	}

	return nthFib
}

func main() {
	twoBefore := f1or2
	oneBefore := f1or2

	for n := 3; ; n++ {
		nthFib := libs.AddTwoIntSlices(twoBefore, oneBefore)
		if len(nthFib) == 1000 {
			fmt.Println(n, nthFib)
			break
		}

		twoBefore = oneBefore
		oneBefore = nthFib
	}
}
