// https://leetcode.com/problems/non-overlapping-intervals/

package main

import "sort"

/*
not much idea
thinking maybe turn this into a histgram.. then just count each overlapping sections
but that is very complicated

then thinking brute force.. not much idea
then thinking sort

then use the merge interval technique, compare each two
and if overlapping, get rid of the bigger one..

seems like it could work
let me see
*/

func _wrong_removing_longer_eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	intvLen := func(intv []int) int {
		return intv[1] - intv[0]
	}

	overlap := func(int1, int2 []int) bool {
		l := min(int1[0], int2[0])
		r := max(int1[1], int2[1])

		return r-l < intvLen(int1)+intvLen(int2)
	}

	count := 0
	for i := 0; i < len(intervals); i++ {
		temp := intervals[i]
		for i+1 < len(intervals) && overlap(temp, intervals[i+1]) {
			// overlapping, get rid of the bigger one
			if intvLen(temp) > intvLen(intervals[i+1]) {
				temp = intervals[i+1]
			}
			count++
			i++
		}
	}

	return count
}

/*
failed here
[[-52,31],[-73,-26],[82,97],[-65,-11],[-62,-49],[95,99],[58,95],[-31,49],[66,98],[-63,2],[30,47],[-40,-26]]
hmm...

is the logical alright?

yeah thining three intervals and they just cross by 1
int1 : 10000 (0,10000)
int2 : 4  (9999, 10003)
int3 : 1  (10002, 10003)

removing int2 is good
but the previous logic will need to remove 2

okay.. although it is wrong.. still good to excercise turning thought into code

maybe the removing critirea is to remove the right most one...
becaue what happens to the left, won't be interleaving with rest

think, cut int1 to 9999-10000, int2 is 9999-10003; now int2 becomes the longer one


*/

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	intvLen := func(intv []int) int {
		return intv[1] - intv[0]
	}

	overlap := func(int1, int2 []int) bool {
		l := min(int1[0], int2[0])
		r := max(int1[1], int2[1])

		return r-l < intvLen(int1)+intvLen(int2)
	}

	count := 0
	for i := 0; i < len(intervals); i++ {
		temp := intervals[i]
		for i+1 < len(intervals) && overlap(temp, intervals[i+1]) {
			// overlapping, get rid of the rightmost one
			if temp[1] > intervals[i+1][1] {
				temp = intervals[i+1]
			}
			count++
			i++
		}
	}

	return count
}

/*
Runtime: 353 ms, faster than 39.69% of Go online submissions for Non-overlapping Intervals.
Memory Usage: 18 MB, less than 31.68% of Go online submissions for Non-overlapping Intervals.

and it works..

Runtime: 255 ms, faster than 83.59% of Go online submissions for Non-overlapping Intervals.
Memory Usage: 16.3 MB, less than 88.17% of Go online submissions for Non-overlapping Intervals.

cannot do any much better than this...

also pay a little more attention to this loop
	for i := 0; i < len(intervals); i++ {
		temp := intervals[i]
		for i+1 < len(intervals) && overlap(temp, intervals[i+1]) {
			// overlapping, get rid of the rightmost one
			if temp[1] > intervals[i+1][1] {
				temp = intervals[i+1]
			}
			count++
			i++
		}
	}

	i++ happens in both places
	it seems confusing but it is write

thinking of a few scenarios
1. i+1 to the len(), then i will be len()-1, give the outer loop a chance to check
2. when overlap(temp, intervals[i+1]) fails, i+1 will point to the slot that breaks it;
	and i will be point to one element before it..
	the next element to check is i+1, it start a new cycle.. so outer loop will move i to it...
	this could be written in other way.. but this works
*/
