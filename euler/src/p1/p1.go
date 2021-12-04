package main

import (
	"fmt"
	"os"
	"runtime"
)

func assert(b bool, s string) {
	if !b {
		println(s)
		os.Exit(1)
	}
}

func current_function_name() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}

func getSum(num int, step int) int {
	sum := 0
	incre := 0
	for n := num; n > step; n -= step {
		incre += step
		sum += incre
	}

	return sum
}

func getSum35(num int) int {
	sum := getSum(num, 3)
	sum += getSum(num, 5)
	sum -= getSum(num, 15)

	return sum
}

func main() {
	test_given_sumto(3, 0)
	test_given_sumto(6, 8)
	test_given_sumto(7, 14)
	test_given_sumto(10, 23)
	test_given_sumto(11, 33)
	test_given_sumto(16, 33+12+15)

	test_given_sumto(1000, 233168)
}

func test_given_sumto(given int, sumto int) {
	sum := getSum35(given)
	str := fmt.Sprintf("given: %d, expected: %d, got: %d", given, sumto, sum)
	assert(sum == sumto, str)
}

func test_give3_sum0() {
	test_given_sumto(3, 0)
}

func test_give6_sum8() {
	test_given_sumto(6, 8)
}
