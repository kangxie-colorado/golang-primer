// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

package main

/*
sliding window
	track state: sum and len(natually)
	len-sum>1, then it violate..


*/

func _1_longestSubarray(nums []int) int {
	maxLen := 0
	i, j := 0, 0
	sum := 0
	for j < len(nums) {
		sum += nums[j]
		// now we need i<=j; why could it go over?
		// stupid, because windLen is locked in
		// but when i changes.. windLen should change as well...
		// how stupid.. (we don't need i<=j)

		for j-i+1-sum > 1 {
			sum -= nums[i]
			i++
		}

		j++
		// cannot use windLen, it could have changed
		// -1, because we need to delete one element
		maxLen = max(maxLen, j-i-1)
	}

	return maxLen
}

/*
Runtime: 36 ms, faster than 96.43% of Go online submissions for Longest Subarray of 1's After Deleting One Element.
Memory Usage: 8.5 MB, less than 14.29% of Go online submissions for Longest Subarray of 1's After Deleting One Element.

yea... in the code you watch out when you tie the change varirable to a previously declared/initiated varible...
and in for loop, the variable should be the changng one...
*/

func longestSubarray(nums []int) int {

	i, j := 0, 0
	sum := 0
	for j < len(nums) {
		sum += nums[j]

		if j-i+1-sum > 1 {
			sum -= nums[i]
			i++
		}

		j++

	}
	// -1 because we need to delete one
	return j - i - 1
}

/*
Runtime: 44 ms, faster than 60.71% of Go online submissions for Longest Subarray of 1's After Deleting One Element.
Memory Usage: 8.5 MB, less than 10.71% of Go online submissions for Longest Subarray of 1's After Deleting One Element.
*/
