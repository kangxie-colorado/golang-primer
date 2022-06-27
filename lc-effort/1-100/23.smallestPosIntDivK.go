// https://leetcode.com/problems/smallest-integer-divisible-by-k/

/*
analysis
	apparently you cannot hold the number in any numeric form
	10003 -> 1423 or something, so not possible hold it by numeric

	so array

	[1] -> 1
	[1,1] -> 11
	[1,1,1] -> 111
	...

	which one is the high bit? make it easier...
	maybe not a difference.

	calculate div by arrray, I have done this before


*/

package main

import (
	"fmt"
	"math"
)

// need special case for num < 10
// let try a more generic one
func _1_divSliceByInt(num []int, d int) []int {
	res := []int{}

	buf := num[0]
	rem := 0
	nextIdx := 1

	// can only borrow n-1 times
	// but need to walk all the num, so the first seed need to happen first
	res = append(res, buf/d)
	buf = buf%d*10 + num[nextIdx]

	for nextIdx < len(num) {
		if buf < d {
			res = append(res, 0)
			rem = buf % d

			buf = rem*10 + num[nextIdx]
			nextIdx++
		} else {
			res = append(res, buf/d)
			rem = buf % d

			buf = rem*10 + num[nextIdx]
			nextIdx++
		}
	}

	if rem != 0 {
		return nil
	}

	leadingZeros := 0
	for _, r := range res {
		if r == 0 {
			leadingZeros++
		} else {
			break
		}
	}

	return res[leadingZeros:]

}

func divSliceByInt(num []int, d int) []int {
	res := []int{}

	buf := num[0]
	rem := num[0] % d
	nextIdx := 1

	for i := 0; nextIdx < len(num); i++ {

		if buf < d {
			res = append(res, 0)
		} else {
			res = append(res, buf/d)
		}
		rem = buf % d
		buf = rem*10 + num[nextIdx]
		nextIdx++

		// no more values to borrow
		// just do it..
		// or do it out side loop... no diff
		if nextIdx == len(num) {

			res = append(res, buf/d)
			rem = buf % d
		}
	}

	if rem != 0 {
		return nil
	}

	leadingZeros := 0
	for _, r := range res {
		if r == 0 {
			leadingZeros++
		} else {
			break
		}
	}

	return res[leadingZeros:]

}

func _1_smallestRepunitDivByK(k int) int {
	res := -1

	if k%2 == 0 || k%5 == 0 {
		return res
	}

	oneCount := int(math.Log10(float64(k)))
	ones := make([]int, oneCount+1)

	for i := range ones {
		ones[i] = 1
	}

	for true {
		if divSliceByInt(ones, k) != nil {
			res = len(ones)
			break
		}

		ones = append(ones, 1)
	}

	return res
}

/*

63 / 70 test cases passed.
Status: Time Limit Exceeded
Submitted: 0 minutes ago
Last executed input:
19927

apparently you did a lot of repeated calculations
for [1,1,1] to [1,1,1,1]

you actually need to remainder part
if rem=2
then 2*10+1

if at this point it divides.. then it is over

*/

func smallestRepunitDivByK(k int) int {
	res := -1

	if k%2 == 0 || k%5 == 0 {
		return res
	}

	startLen := 0
	startOne := 0
	ones := []int{1, 11, 111, 1111, 11111, 111111}

	for i, one := range ones {
		if one == k {
			return i + 1
		}
		startLen = i + 1
		startOne = one

		if k < one {
			break
		}
	}

	rem := startOne
	for true {
		if rem%k == 0 {
			return startLen
		}

		rem = rem%k*10 + 1
		startLen++
	}

	return res
}

/*
Success
Details
Runtime: 4 ms, faster than 46.15% of Go online submissions for Smallest Integer Divisible by K.
Memory Usage: 2 MB, less than 30.77% of Go online submissions for Smallest Integer Divisible by K.
Next challenges:
*/

func testSmallestRepunitDivByK() {
	fmt.Println(divSliceByInt([]int{1, 1, 1}, 3))

	fmt.Println(smallestRepunitDivByK(1))
	fmt.Println(smallestRepunitDivByK(2))
	fmt.Println(smallestRepunitDivByK(3))
}
