package main

// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

import (
	"fmt"
	"os"
)

func assert(b bool, s string) {
	if !b {
		println(s)
		os.Exit(1)
	}
}

func getEvenFibSumUnderTerm(upper int) int {
	sum := 0
	if upper < 3 {
		return 0
	}

	fib1 := 2
	fib2 := 3
	fib3 := 5

	for fib1 < upper {
		sum += fib1
		fib1 = fib2 + fib3
		fib2 = fib1 + fib3
		fib3 = fib1 + fib2

	}

	return sum
}

func main() {
	test_given_evenFibSum(2, 0)
	test_given_evenFibSum(3, 2)
	test_given_evenFibSum(5, 2)
	test_given_evenFibSum(13, 10)
	test_given_evenFibSum(22, 10)

	test_given_evenFibSum(55, 44)
	test_given_evenFibSum(89, 44)
	test_given_evenFibSum(144, 44)
	test_given_evenFibSum(145, 188)
	test_given_evenFibSum(146, 188)
	test_given_evenFibSum(4000000, 4613732)

}

func test_given_evenFibSum(given int, evenFibSumExp int) {
	evenFibSum := getEvenFibSumUnderTerm(given)
	str := fmt.Sprintf("given: %d, expected: %d, got: %d", given, evenFibSumExp, evenFibSum)
	assert(evenFibSumExp == evenFibSum, str)
}
