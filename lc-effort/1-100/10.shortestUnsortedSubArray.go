// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

package main

import (
	"sort"
)

/** analysis
	no matter it looks difficult or easy, analyze the problem thoroughly

	observations:
		sort a sub-array so the whole array is sorted
		*sub-array is an abstract -- whole array is a sub-array

		naturally, I can sort the whole array then compare from both end and easily find out the sub-array that needs to be sorted

	regarding O(n) solution, think about it later

subArrayToSort(nums):
	start = 0
	end = len(nums)

	sorted = sort(nums)
	for i:=0;i<len(nums);i++ :
		if nums[i] != sorted[i]
			start = i
			for j=len(nums)-1;j>i;j--:
				if nums[j] != sorted[j]
				end = j
				return j-i+1

	return -1

petty easy?

**/

func copySlice(s []int) []int {
	ret := make([]int, len(s))
	for i, e := range s {
		ret[i] = e
	}

	return ret
}

func _1_findUnsortedSubarray(nums []int) int {
	sorted := copySlice(nums)
	sort.Ints(sorted)

	for i := 0; i < len(nums); i++ {
		if nums[i] != sorted[i] {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] != sorted[j] {
					return j - i + 1
				}
			}
		}
	}

	return 0

}

/**

Success
Details
Runtime: 24 ms, faster than 84.85% of Go online submissions for Shortest Unsorted Continuous Subarray.
Memory Usage: 6.6 MB, less than 42.42% of Go online submissions for Shortest Unsorted Continuous Subarray.

^^^ this is even in the better tier of solution
	let me think the O(n) solution

	so actually this is just simply to look for the leftest out-of-order element and the right one
findUnsortedSubarray(nums)
	left:=0
	right:=len(nums)-1

	for i:=left+1;i<len(nums);i++, left++
		if nums[i] < nums[left]
			for j:= len(nums)-2;j>left;j--, right--
				if nums[j]>nums[right]
					return right-left+1

	return 0
**/
func wrong_findUnsortedSubarray(nums []int) int {
	left := 0
	right := len(nums) - 1

	for i := left + 1; i < len(nums); {
		if nums[i] == nums[left] {
			i++
			continue
		}

		if nums[i] < nums[left] {
			for j := right - 1; j >= left; {
				if nums[j] == nums[right] {
					j--
					continue
				}

				if nums[j] > nums[right] {
					return right - left + 1
				} else {
					right = j + 1
				}

				j--
				right--
			}

		} else {
			left = i - 1
		}

		i++
		left++
	}

	return 0
}

// issue: [1,3,2,2,2] when there is equal elements
// when there is the left/right should not move

// hm... so O(n) is indeed much more difficult than it looks
// e.g. [1,2,2,2,4,1,5,6]

/* rethink
maybe the constraint is a clue

Constraints:
	1 <= nums.length <= 10^4
	-10^5 <= nums[i] <= 10^5

just allocate an array of size 2*10^5 + 1 then read thru the nums and fill the array
for example
[2,6,2,8,10,-9,15] will fill
index-base + 2: 2 (two times)
index-base + 6: 1 (one time)
index-base - 9: 1
...
index-base: 10^5 (this is the middle, which is 0)

then read thru the array and write the sorted array then compare..
this is indeed O(n) but kind of weird let me code it up and see
*/

func _2_findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}

	tenPowFive := 100000
	toFill := make([]int, 2*tenPowFive+1)
	indexBase := tenPowFive

	for _, n := range nums {
		toFill[indexBase+n]++
	}

	sorted := make([]int, len(nums))
	idxInSorted := 0
	for i, n := range toFill {
		for ; n != 0; n-- {
			sorted[idxInSorted] = i - indexBase
			idxInSorted++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != sorted[i] {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] != sorted[j] {
					return j - i + 1
				}
			}
		}
	}

	return 0

}

func findMinMax(nums []int) (int, int) {
	min := 100000
	max := -100000

	for _, n := range nums {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return min, max

}

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}

	min, max := findMinMax(nums)
	toFill := make([]int, max-min+1)
	indexBase := -min

	for _, n := range nums {
		toFill[indexBase+n]++
	}

	sorted := make([]int, len(nums))
	idxInSorted := 0
	for i, n := range toFill {
		for ; n != 0; n-- {
			sorted[idxInSorted] = i - indexBase
			idxInSorted++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != sorted[i] {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] != sorted[j] {
					return j - i + 1
				}
			}
		}
	}

	return 0

}

/**
	Success
	Details
	Runtime: 242 ms, faster than 6.06% of Go online submissions for Shortest Unsorted Continuous Subarray.
	Memory Usage: 17.2 MB, less than 6.06% of Go online submissions for Shortest Unsorted Continuous Subarray.

	it is O(n) but of course the big array is a bit too wasteful
	so instead let me find out the min/max of the array and setup the auxliary from Min/Max
	min could be > 0, so be careful of calculation of the index-base: basically (0 - Min)

	Success
	Details
	Runtime: 43 ms, faster than 28.79% of Go online submissions for Shortest Unsorted Continuous Subarray.
	Memory Usage: 6.8 MB, less than 37.88% of Go online submissions for Shortest Unsorted Continuous Subarray.

remove this
	if sort.IntsAreSorted(nums) {
		return 0
	}

	Success
	Details
	Runtime: 24 ms, faster than 84.85% of Go online submissions for Shortest Unsorted Continuous Subarray.
	Memory Usage: 7.3 MB, less than 9.09% of Go online submissions for Shortest Unsorted Continuous Subarray.


**/
