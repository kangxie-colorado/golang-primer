// https://leetcode.com/problems/partition-labels/

/*
so just merge intervals..

*/

package main

import "sort"

func partitionLabels(s string) []int {
	intervals := make([]pair, 26)

	// build intervals
	for i, c := range s {
		tup := intervals[c-'a']
		if tup.n1 == 0 && tup.n2 == 0 {
			intervals[c-'a'] = pair{i + 1, i + 1}
		} else {
			intervals[c-'a'] = pair{tup.n1, i + 1}
		}
	}

	// merge intervals.. actually no need to real merge.. just calculate the length of it
	// but need it be sorted
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].n1 < intervals[j].n1 ||
			(intervals[i].n1 == intervals[j].n1 && intervals[i].n2 < intervals[j].n2)
	})

	overlap := func(p1, p2 pair) bool {
		return p1.n2 > p2.n1
	}

	res := []int{}
	for i := 0; i < len(intervals); i++ {
		I := intervals[i]
		if I.n1 == 0 {
			continue
		}

		for i+1 < len(intervals) && overlap(I, intervals[i+1]) {
			I = pair{min(intervals[i+1].n1, I.n1), max(intervals[i+1].n2, I.n2)}
			i++
		}
		l := I.n2 - I.n1 + 1
		res = append(res, l)
	}

	return res
}

/*
Runtime: 4 ms, faster than 51.69% of Go online submissions for Partition Labels.
Memory Usage: 2.1 MB, less than 69.10% of Go online submissions for Partition Labels.

Runtime: 0 ms, faster than 100.00% of Go online submissions for Partition Labels.
Memory Usage: 2.2 MB, less than 69.10% of Go online submissions for Partition Labels.
*/
