// https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/submissions/

/*
	it is not hard to see than it can at most be len(n) or the biggest element
	but that is not strict enough

	it can be much smaller than biggest element
	so it needs a transition

	max:=1 , the start point
	then

	for i:=1;i<len(arr);i++ {
        max = min(arr[i], max+1)
    }

	might be not easy to come to this at first
	but if you think max as an array
	max[i] = min(max[i-1]+1, arr[i])

	then it is so obvious you don't even need the array


*/

package main

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)

	max := 1

	for i := 1; i < len(arr); i++ {
		max = min(arr[i], max+1)
	}

	return max
}

/*
Runtime: 76 ms, faster than 100.00% of Go online submissions for Maximum Element After Decreasing and Rearranging.
Memory Usage: 9.8 MB, less than 33.33% of Go online submissions for Maximum Element After Decreasing and Rearranging.
Next challenges:
*/
