// https://leetcode.com/problems/maximum-ice-cream-bars/

package main

import "sort"

/*
is this not greedy?

*/

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	cnt := 0
	for _, c := range costs {
		if c <= coins {
			cnt++
			coins -= c
		}
	}

	return cnt

}

/*
Runtime: 218 ms, faster than 62.50% of Go online submissions for Maximum Ice Cream Bars.
Memory Usage: 9.4 MB, less than 75.00% of Go online submissions for Maximum Ice Cream Bars.

ugh.. how is this a medium problem?
*/
