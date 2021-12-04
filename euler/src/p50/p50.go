/***

The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?

*/

/**

analysis:
	this number has following properties
		1. it is a prime
		2. consecutive prime sum, then if sum - max(ele) will still be a prime and a consecutive prime sum

		there is some recursive going on here

	and the searching space is not super, so lets just populate all the primes under 1Million first, actually this can be a libs function

	okay, got that done in libs.PrimesUnderN
	then this becomes a reduce function

	I just need to reduce myself to 0 with 6 consecutive primes,
*/

package main

import (
	"fmt"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

var primesUnder1M = libs.PrimesUnderN(1000000)

func findConsecutivePrimes(sum int, primeSpace []int) []int {
	primes := []int{}

	for start := len(primeSpace) - 1; start >= 0; start-- {
		if primeSpace[start] >= sum {
			continue
		}
		runningSum := primeSpace[start]
		next := start - 1
		for ; next >= 0 && runningSum < sum; next-- {
			runningSum += primeSpace[next]
			if runningSum == sum {
				if start+1-next > len(primes) {
					primes = primeSpace[next : start+1]
				}
			}
		}

	}

	return primes

}

func consecutiveLength(sum int, primeSpace []int) int {
	return len(findConsecutivePrimes(sum, primeSpace))
}

func main() {
	maxLen := 0
	finalP := -1
	for _, p := range primesUnder1M {
		l := consecutiveLength(p, primesUnder1M)
		if l > maxLen {
			maxLen = l
			finalP = p
		}
	}

	fmt.Println(finalP, maxLen)
}
