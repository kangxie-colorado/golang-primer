// https://leetcode.com/problems/container-with-most-water/

package main

/*
I can only see the O(n^2) solution
is there O(n) or O(nlgn)???

after thinking monotonical stack, dp, which doesn't work out really
this may be a two pointers issue

l,r : r-l * min(R,L)
	now moving the smaller between R/L, because if you move bigger one, it is destined to becomes smaller

but I aint sure if this is complete or not
have to try
*/

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxC := 0
	for l < r {
		maxC = max(maxC, (r-l)*min(height[r], height[l]))

		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return maxC
}

/*
Runtime: 129 ms, faster than 38.65% of Go online submissions for Container With Most Water.
Memory Usage: 8.9 MB, less than 55.06% of Go online submissions for Container With Most Water.

hahahaha
so sliding window actually could take care of some scenarios to cover all windows..
esp, when some windows are for sure to be skipped because it just slip away from the target further
like when L<R, you don't want to move R.. because those windows will just be smaller.. dominated by left bar
so those can be skipped...

very good finding...

but why so slow? this is O(N), can there be a O(lgn)??
*/
