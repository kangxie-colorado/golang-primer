// https://leetcode.com/problems/sliding-window-maximum/

package main

/*
worked a few simple examples.. looks like, use a deque keeping the number and idx
lets it left maintain the max, and when window slids out of the range, it pops left

also on the right side, use it like a non-increasing stack.. if top is smaller, then pop from right..
this way, it seems giving the max on left each time and maintain the possibility of queuing next max
*/

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	dq := [][]int{} // 0: idx, 1: val
	for i := 0; i < len(nums); i++ {
		if len(dq) > 0 && i-k >= dq[0][0] {
			dq = dq[1:]
		}

		for len(dq) > 0 && dq[len(dq)-1][1] < nums[i] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, []int{i, nums[i]})

		// some confusion of when to take the dq[0]
		// should be after all the modifications
		if i >= k-1 {
			res = append(res, dq[0][1])
		}
	}

	return res
}

/*
Runtime: 644 ms, faster than 18.48% of Go online submissions for Sliding Window Maximum.
Memory Usage: 16.8 MB, less than 5.22% of Go online submissions for Sliding Window Maximum.

Runtime: 257 ms, faster than 83.33% of Go online submissions for Sliding Window Maximum.
Memory Usage: 11.5 MB, less than 42.97% of Go online submissions for Sliding Window Maximum.
*/
