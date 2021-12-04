// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.
package main

import (
	"fmt"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

func sumOfPrimsUnder(n int) int {
	sum := 0

	for i := 2; i < n; i++ {
		if libs.IsPrime(i) {
			sum += i
		}
	}

	return sum
}

func main() {
	println(sumOfPrimsUnder(10))
	fmt.Println(123)

}
