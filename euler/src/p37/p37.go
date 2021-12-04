/***
The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
***/
package main

import (
	"fmt"
	"math"
)

func elem_in_array(val int, vals []int) bool {
	for i := range vals {
		if val == vals[i] {
			return true
		}
	}

	return false
}

var primes = []int{2, 3}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	if elem_in_array(num, primes) {
		return true
	} else {
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				return false
			}
		}
	}
	primes = append(primes, num)
	return true
}

func leftTruncate(num int) int {
	return num % int(math.Pow10(int(math.Log10(float64(num)))))
}

func rightTruncate(num int) int {
	num /= 10
	return num
}

func main() {
	sum := 0
	howManyFound := 0

OUTER:
	for num := 10; howManyFound < 11; num++ {

		rightCopy := num
		for rightCopy > 0 {
			if !isPrime(rightCopy) {
				continue OUTER
			}

			rightCopy = rightTruncate(rightCopy)
		}

		leftCopy := num
		for leftCopy > 0 {
			if !isPrime(leftCopy) {
				continue OUTER
			}

			leftCopy = leftTruncate(leftCopy)
		}

		sum += num
		howManyFound += 1
	}

	fmt.Println(sum)
}
