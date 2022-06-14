// https://leetcode.com/problems/next-greater-element-ii/

package main

import (
	"fmt"
	"math"
)

/*
I get the idea of using a stack
just pop data out when there is a bigger number coming up

and when it comes, keep popping the stack until it is empty or hit a even bigger number
if an even bigger number appears, then push the current number onto stack
*/

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	filled := 0

	maxN := math.MinInt
	for i, n := range nums {
		maxN = max(maxN, n)
		res[i] = 1_000_000_001
	}

	stack := [][]int{}
	for i := 0; filled < len(nums); i = (i + 1) % len(nums) {
		num := nums[i]

		if len(stack) != 0 {
			j := len(stack) - 1
			for ; j >= 0 && stack[j][0] < num; j-- {
				res[stack[j][1]] = num
				filled++
			}
			stack = stack[:j+1]
		}

		if num == maxN {
			res[i] = -1
			filled++
		} else {
			if res[i] == 1_000_000_001 {
				stack = append(stack, []int{num, i})
			}

		}

	}

	return res

}

/*
Runtime: 30 ms, faster than 71.43% of Go online submissions for Next Greater Element II.
Memory Usage: 7.1 MB, less than 47.37% of Go online submissions for Next Greater Element II.


although it passes, there are some rough patches.

*/

func testNextGreaterElements() {

	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
