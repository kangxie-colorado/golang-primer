/*
The series, 11 + 22 + 33 + ... + 1010 = 10405071317.

Find the last ten digits of the series, 11 + 22 + 33 + ... + 10001000.
*/

package main

import (
	"fmt"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

func main() {
	sum := []int{}
	for i := 1; i <= 1000; i++ {
		sum = libs.AddTwoIntSlices(sum, libs.XPowY(i, i))
	}

	fmt.Println(sum)
}
