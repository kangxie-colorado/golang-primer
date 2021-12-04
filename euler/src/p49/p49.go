/**

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?

**/

package main

import (
	"fmt"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

func main() {

	for num := 1001; num <= 3340; num += 2 {
		if libs.IsPrime(num) && libs.IsPrime(num+3330) && libs.IsPrime(num+6660) {
			digits := libs.DigitSet(num)
			digits2 := libs.DigitSet(num + 3330)
			digits3 := libs.DigitSet(num + 6660)

			if libs.SameSet(digits, digits2) && libs.SameSet(digits, digits3) {
				fmt.Println(num*10000*10000 + 10000*(num+3330) + num + 6660)

			}

		}
	}
}
