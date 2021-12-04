/***

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?

**/

package main

import (
	"fmt"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

var s = make(map[int]struct{})
var exists = struct{}{}

// the minimum number with 4 prime factors are 2*3*4*5 = 120
// so lets start with 120

func main() {
	num := 120
	for {
		uniqPrimeFactorsCount, _ := libs.UniqNumbersInSlice(libs.PrimeFactors(num))

		uniqPrimeFactorsCount1, _ := libs.UniqNumbersInSlice(libs.PrimeFactors(num + 1))
		uniqPrimeFactorsCount2, _ := libs.UniqNumbersInSlice(libs.PrimeFactors(num + 2))
		uniqPrimeFactorsCount3, _ := libs.UniqNumbersInSlice(libs.PrimeFactors(num + 3))

		if uniqPrimeFactorsCount == 4 && uniqPrimeFactorsCount1 == 4 && uniqPrimeFactorsCount2 == 4 && uniqPrimeFactorsCount3 == 4 {
			fmt.Println(num)
			break
		}

		num += 1
	}
}
