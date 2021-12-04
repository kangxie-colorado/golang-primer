/**

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9 = 7 + 2×12
15 = 7 + 2×22
21 = 3 + 2×32
25 = 7 + 2×32
27 = 19 + 2×22
33 = 31 + 2×12

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?

**/

/***
analysis:
  first thinking is to end the search of prime that is <= the number and see if the rest is a square
  but calculate the prime is hard enough, and you have to calculate so many times

  then 2nd thinking is to end the search of the square
  notice
  num = p + 2 *n^2
  then (num-p)/2 = n^2

  the search space is [1:sqrt((num-p)/2)], which is included in [1:sqrt(num-2/2)], essential [1:sqrt(num/2)]
  this way, we only at most test sqrt(num) times prime

  this way, don't even need to worry about filtering non-prime number, since it won't be satisfied with this formula anyway

***/

package main

import (
	"fmt"
	"math"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

func main() {
OUTER:
	for num := 35; ; num += 2 {
		if libs.IsPrime(num) {
			continue
		}
		for sqrt := 1; sqrt <= int(math.Sqrt(float64(num/2))); sqrt++ {
			p := num - sqrt*sqrt*2
			if libs.IsPrime(p) {
				// this number could be written in such way
				continue OUTER
			}
		}

		fmt.Println(num)
		break
	}

}
