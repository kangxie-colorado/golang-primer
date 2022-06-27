// https://leetcode.com/problems/jump-game/
package main

import "fmt"

/*
I only see a backtracking...
but start from the max jump

not really backtracking..

just dfs actually
*/

func _dfs_canJump(nums []int) bool {
	var jump func(start int) bool
	m := make([]int, len(nums))
	jump = func(start int) bool {
		if start >= len(nums)-1 {
			// actually if you can jump over, you can jump onto the finish line
			return true
		}
		if nums[start] == 0 {
			return false
		}
		if m[start] != 0 {
			return m[start] == 1
		}

		for i := nums[start]; i >= 1; i-- {
			if jump(start + i) {
				m[start] = 1
				return true
			}
		}
		m[start] = -1
		return false
	}

	return jump(0)
}

/*
missing the scenario of jumping out of bounds
		if start >= len(nums)-1 {
			return start == len(nums)-1
		}

		and of course.. TLE
		can do memorization

as expected
Runtime: 239 ms, faster than 22.38% of Go online submissions for Jump Game.
Memory Usage: 9.6 MB, less than 5.54% of Go online submissions for Jump Game.

pass but not too fast..
DP obviously not very applicable here..

might be a DP but O(n^2)
*/

func _dp_canJump(nums []int) bool {
	nums[len(nums)-1] = len(nums) - 1
	for j := len(nums) - 2; j >= 0; j-- {
		d := nums[j]

		for ; d > 0; d-- {
			if j+d >= len(nums) {
				continue
			}
			if nums[j+d] != 0 {
				// if nums[j+d] is not a trap, then j can reach idx j+d
				// nums[j+d] stores the furthest index from there it can reach
				nums[j] = max(nums[j], nums[j+d])
				if nums[j] >= len(nums)-1 {
					break
				}
			}

		}
	}

	return nums[0] >= len(nums)-1
}

/*
Runtime: 1833 ms, faster than 7.26% of Go online submissions for Jump Game.
Memory Usage: 6.9 MB, less than 85.99% of Go online submissions for Jump Game.

yeah I know this should be slower...
*/

func _use_map_dfs_canJump(nums []int) bool {
	var jump func(start int) bool
	m := make(map[int]bool)
	jump = func(start int) bool {
		if start >= len(nums)-1 {
			// actually if you can jump over, you can jump onto the finish line
			return true
		}
		if nums[start] == 0 {
			return false
		}
		if v, found := m[start]; found {
			return v
		}

		for i := nums[start]; i >= 1; i-- {
			if jump(start + i) {
				m[start] = true
				return true
			}
		}
		m[start] = false
		return false
	}

	return jump(0)
}

/*
^^ using a map, slower
Runtime: 1687 ms, faster than 7.76% of Go online submissions for Jump Game.
Memory Usage: 9.5 MB, less than 5.54% of Go online submissions for Jump Game.

I failed to see this O(n) solution  -- just keep tracking the last index that can reach the end
my previous dp is kind of like this... but not quite the same...

but then I can make it work...
nah..

this tracking a different state and makes O(d) to O(1) is interesting..
*/

func _tracking_last_reachable_idx_canJump(nums []int) bool {
	lastReachableIdx := len(nums) - 1

	for j := len(nums) - 2; j >= 0; j-- {
		if j+nums[j] >= lastReachableIdx {
			lastReachableIdx = j
		}
	}

	return lastReachableIdx == 0
}

func canJump(nums []int) bool {
	furthest := len(nums) - 1
	nums[furthest] = furthest

	for j := len(nums) - 2; j >= 0; j-- {
		if j+nums[j] >= furthest || nums[j+nums[j]] >= furthest {
			nums[j] = max(j+nums[j], furthest)
		}
	}

	return nums[0] >= furthest
}

func testCanJump() {
	fmt.Println(canJump([]int{2, 5, 0, 0}))
	fmt.Println(canJump([]int{2, 0}))
	fmt.Println(canJump([]int{0, 2, 3}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
