/***
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: As 1! = 1 and 2! = 2 are not sums they are not included.
***/
package main

import "fmt"

var facotrials = make(map[int]int)

func facotrial(n int) int {
	if val, ok := facotrials[n]; ok {
		return val
	}

	if n <= 1 {
		return 1
	}

	return n * facotrial(n-1)
}

func main() {

	for i := 0; i <= 9; i++ {
		facotrials[i] = facotrial(i)
	}

	for num := 10; num < 10_000_000; num++ {
		digitFactorial := 0
		copy := num
		for copy > 0 {
			d := copy % 10
			copy /= 10
			digitFactorial += facotrials[d]
		}

		if digitFactorial == num {
			fmt.Println(num)
		}
	}
}
