/****

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17
Find the sum of all 0 to 9 pandigital numbers with this property.

***/

package main

import (
	"fmt"
	"math"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

// return digits [s:d] inclusive
func sliceNumber(num int, s int, e int) int {
	powOf10 := int(math.Log10(float64(num)))

	if s < 0 || e < 0 || s > powOf10+1 || e > powOf10+1 || (e-s) > powOf10 {
		return -1
	}

	return (num / int(math.Pow10(powOf10-e+1))) % int(math.Pow10(e-s+1))

}

func main() {
	sumOfRangeDriver(1_023_456_789, 9876543211)
}

func sumOfRangeRecursive(start, end int, ch chan int) {
	primes := []int{2, 3, 5, 7, 11, 13, 17}
	sum := 0

	if end-start < 1000000 {
	OUTER:
		for n := start; n < end; n++ {
			if !libs.IsNumPandigitAtoB(n, 0, 9) {
				continue

			}
			for d := 2; d < 9; d++ {
				if sliceNumber(n, d, d+2)%primes[d-2] != 0 {
					continue OUTER
				}
			}

			fmt.Println(n)
			sum += n
		}

		ch <- sum

	} else {
		go sumOfRange(start, start+(end-start)/2, ch)
		go sumOfRange(start+(end-start)/2, end, ch)

		sum1, sum2 := <-ch, <-ch
		ch <- sum1 + sum2
	}
}

func sumOfRange(start, end int, ch chan int) {
	primes := []int{2, 3, 5, 7, 11, 13, 17}
	sum := 0

OUTER:
	for n := start; n < end; n++ {
		if !libs.IsNumPandigitAtoB(n, 0, 9) {
			continue

		}
		for d := 2; d < 9; d++ {
			if sliceNumber(n, d, d+2)%primes[d-2] != 0 {
				continue OUTER
			}
		}

		fmt.Println(n)
		sum += n
	}
	ch <- sum
}

func sumOfRangeDriver(start, end int) {
	sum := 0
	len := end - start
	batchSize := 1_000_000_000

	ch := make(chan int, len/batchSize)
	for n := start; n < end; n += batchSize {
		intervalEnd := n + batchSize
		if intervalEnd > end {
			intervalEnd = end
		}
		go sumOfRange(n, intervalEnd, ch)
	}

	for i := 0; i < (end-start)/batchSize; i++ {
		aNumber := <-ch
		fmt.Println(aNumber)
		sum += aNumber
	}

	fmt.Println(sum)
}
