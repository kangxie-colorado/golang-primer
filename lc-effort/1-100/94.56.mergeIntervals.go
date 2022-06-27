// https://leetcode.com/problems/merge-intervals/

package main

import (
	"fmt"
	"sort"
)

/*
this is just the sub-problem of insert-interval
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	res := [][]int{}

	for i := 0; i < len(intervals); i++ {
		if i+1 < len(intervals) && intervals[i][1] < intervals[i+1][0] {
			// non-overlapping
			res = append(res, intervals[i])

		} else {
			// because of sort, i cannot be to the right of
			newInterval := intervals[i]
			for i+1 < len(intervals) && newInterval[1] >= intervals[i+1][0] {
				newInterval[1] = max(newInterval[1], intervals[i+1][1])
				i++
			}
			res = append(res, newInterval)

		}
	}

	return res
}

/*
Runtime: 37 ms, faster than 37.01% of Go online submissions for Merge Intervals.
Memory Usage: 7.1 MB, less than 40.83% of Go online submissions for Merge Intervals.
*/

func testMergeInterval() {

	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}
