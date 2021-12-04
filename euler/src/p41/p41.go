/**

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?

*/

package main

import (
	"fmt"
	"math"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

func main() {
	for num := 87654321; ; num-- {
		n := int(math.Log10(float64(num))) + 1
		if libs.IsNumPandigitToN(num, n) && libs.IsPrime(num) {
			fmt.Println(n, num)
			break
		}
	}
}
