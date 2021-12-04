/***
Take the number 192 and multiply it by each of 1, 2, and 3:

192 × 1 = 192
192 × 2 = 384
192 × 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
***/

package main

import (
	"fmt"
	"math"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

func isConcatProd(num int) bool {

	logTen := int(math.Log10(float64(num)))

	// concat means more than 1 numbers, so a single digit should be counted out
	// for args{3} should be returning false
	// thus logTen > 0;
	for logTen > 0 {
		d1 := num / int(math.Pow10(int(logTen)))
		numCal := d1
		n := 2

		for numCal < num {
			dNext := n * d1
			logTenDNext := int(math.Log10(float64(dNext)))
			numCal = numCal*int(math.Pow10(int(logTenDNext)+1)) + dNext
			n += 1
		}

		if numCal == num {
			return true
		}

		logTen -= 1
	}

	return false

}

func main() {

	for num := 987654321; ; num-- {
		if libs.HasDuplicatedDigit(num) || libs.HasZeroDigit(num) {
			continue
		}

		if isConcatProd(num) {
			fmt.Println(num)
			break
		}
	}
}
