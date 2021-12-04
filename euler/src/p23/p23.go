/**
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
**/

package main

import "github.com/kangxie-colorado/golang-primer/euler/libs"

func divisors(num int) []int {
	divs := []int{}
	for n := 1; n < num; n++ {
		if num%n == 0 {
			divs = append(divs, n)
		}
	}

	return divs
}

func d(n int) int {
	sum := 0
	for _, n := range divisors(n) {
		sum += n
	}

	return sum
}

func getAllAbundantNumsBelow(upper int) []int {
	nums := []int{}
	for n := 1; n < upper; n++ {
		if n < d(n) {
			nums = append(nums, n)
		}
	}

	return nums
}

var abundantNums = getAllAbundantNumsBelow(28123 + 1)

func numCanBeSumOf2Abundant(num int) bool {
	for _, a_n := range abundantNums {
		if libs.IsNumInSlice(num-a_n, abundantNums) {
			return true
		}
	}

	return false
}

func main() {
	sum := 0
	for n := 1; n < 28123+1; n++ {
		if !numCanBeSumOf2Abundant(n) {
			sum += n
		}
	}

	println(sum)
}
