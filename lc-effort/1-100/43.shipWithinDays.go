// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
// see analysis in binarySearch.go

package main

import (
	"fmt"
	"math"
)

func getDays(cap int, weights []int) int {
	// at least need one day
	days := 1
	weight := 0
	for _, w := range weights {
		if w > cap {
			return math.MaxInt
		}

		weight += w

		if weight > cap {
			days++
			weight = w
		}
	}

	return days
}

func shipWithinDays(weights []int, days int) int {
	l, r := 1, 500*5*10000

	for l < r {
		mid := l + (r-l)/2

		if getDays(mid, weights) <= days {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

/*
Runtime: 57 ms, faster than 28.75% of Go online submissions for Capacity To Ship Packages Within D Days.
Memory Usage: 7 MB, less than 38.75% of Go online submissions for Capacity To Ship Packages Within D Days.


the monotonicity is
We dig out the monotonicity of this problem: if we can successfully ship all packages within D days with capacity m, then we can definitely ship them all with any capacity larger than m.


*/

func testGetDays() {
	fmt.Println(getDays(2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(getDays(9, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(getDays(10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(getDays(15, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(getDays(16, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(getDays(120, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
