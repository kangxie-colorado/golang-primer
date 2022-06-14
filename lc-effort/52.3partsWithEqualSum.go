// https://leetcode.com/problems/partition-array-into-three-parts-with-equal-sum/

package main

import "fmt"

/*
	this is an easy problem
	but I don't feel it being so easy and the accept rate is actually pretty low  < 50%
	Accepted 65,144
	Submissions 147,087

	let me see; so
	1. it has to have a whole sum that can be divided by 3
	2. it cannot change the order to mix/match

	so yeah I see a possible solution
	I just calcute the partial sum

	at some point, it must be 1/3 of total
	at some point, it must be 2/2 of total
*/

func _1_canThreePartsEqualSum(arr []int) bool {
	// calculate the partial sum
	prevSum := 0
	for i := range arr {
		arr[i] += prevSum
		prevSum = arr[i]
	}

	total := arr[len(arr)-1]
	if total%3 == 0 {
		oneThirdFound := false
		twoThirdFound := false
		for i, s := range arr {
			if s == total/3 {
				oneThirdFound = true

			} else if oneThirdFound && s == total/3*2 {
				// the checking must be in order..
				// you cannot find 2/3 first then 1/3
				twoThirdFound = true
			} else if oneThirdFound && twoThirdFound && s == total && i == len(arr)-1 {
				return true
			}
		}
	}

	return false
}

/*
Input:
[1,-1,1,-1]
Output:
true
Expected:
false

yeah, I see why its acceptance rate is so low
so many edge cases..

so actually you must find the third as well

Input:
[0,0,0,0]
Output:
false
Expected:
true

ok.. ok...
I must get each idx... (at least the final idx)

oh... actually [0,0,0,0] should still pass
well...  I do need to actually partition the array so for sure


*/

func _2_canThreePartsEqualSum(arr []int) bool {
	// calculate the partial sum
	prevSum := 0
	for i := range arr {
		arr[i] += prevSum
		prevSum = arr[i]
	}

	total := arr[len(arr)-1]
	if total%3 == 0 {
		oneThirdIdx := -1
		twoThirdIdx := -1
		threeThirdidx := len(arr) - 1
		for i, s := range arr {
			if oneThirdIdx == -1 && s == total/3 {
				oneThirdIdx = i
			} else if oneThirdIdx != -1 && s == total/3*2 {
				twoThirdIdx = i
				break
			}

		}

		if oneThirdIdx != twoThirdIdx && twoThirdIdx != threeThirdidx && oneThirdIdx != threeThirdidx &&
			oneThirdIdx != -1 && twoThirdIdx != -1 && arr[threeThirdidx]-arr[twoThirdIdx] == total/3 {
			return true
		}

	}

	return false
}

/*
Runtime: 78 ms, faster than 12.50% of Go online submissions for Partition Array Into Three Parts With Equal Sum.
Memory Usage: 7.7 MB, less than 40.63% of Go online submissions for Partition Array Into Three Parts With Equal Sum.

damn.. even an easy problem
still not so easy to get it right...

I think it should be this way.. better flow

*/

func canThreePartsEqualSum(arr []int) bool {
	// calculate the partial sum
	prevSum := 0
	for i := range arr {
		arr[i] += prevSum
		prevSum = arr[i]
	}

	total := arr[len(arr)-1]
	if total%3 == 0 {
		oneThirdIdx := -1
		for i := 0; i < len(arr); i++ {
			if arr[i] == total/3 {
				oneThirdIdx = i
				for j := oneThirdIdx + 1; j < len(arr); j++ {
					if j != len(arr)-1 && arr[j] == total/3*2 {
						return arr[len(arr)-1]-arr[j] == total/3
					}
				}
			}
		}

	}

	return false
}

func testCanThreePartsEqualSum() {
	fmt.Println(canThreePartsEqualSum([]int{6, 1, 1, 13, -1, 0, -10, 20}))
	fmt.Println(canThreePartsEqualSum([]int{1, -1, 1, -1}))
	fmt.Println(canThreePartsEqualSum([]int{0, 0, 0, 0}))
}
