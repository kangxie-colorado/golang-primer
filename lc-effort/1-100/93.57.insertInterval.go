// https://leetcode.com/problems/insert-interval/

package main

import "fmt"

/*
so there should be this brute force way
just go check one by one and merge if necessary

keep an prefix, merging, suffix for easy contactation

then there should be a binary search method
let me try brute force first to verify
*/

func _brute_force_insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{} // the prefix so far

	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			// new interval is to the left
			// and it is not merged so far.. so merge now
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)

			return res
		} else if newInterval[0] > intervals[i][1] {
			// new interval is to the right, leave it for future
			res = append(res, intervals[i])
		} else {
			// interesting part
			// the overlap..
			// merge to one and then continue to merge
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
			// no need to update prefix here..
			// let the next loop run get it
		}
	}

	// if by end, the new interval is not merged
	res = append(res, newInterval)
	return res
}

/*
Runtime: 11 ms, faster than 48.11% of Go online submissions for Insert Interval.
Memory Usage: 4.7 MB, less than 41.43% of Go online submissions for Insert Interval.

wow.. a little surprised it worked and passed
now let me get to the binary search

what I will search is
for first 'end', that is bigger or equal to the new interval's 'start'
for first 'start', that is smaller or equal to the new interval's 'end'

usually binary search I know it naturally does the cloest bigger/equal
but how to find the smaller//equal one?

maybe the way how I shift r or l? let me test it out
*/

func intervalBinSearch(intervals [][]int, newInterval []int) (int, int) {
	// return:
	// 1, first end that is bigger/equal than/to the new interval's start
	// 2, first start that is smaller or equal to the new interval's end
	endIdx := -1
	startIdx := -1

	// find the bigger/equal end index first
	// r don't len()-1, so l can go to len, that is when to insert at the end

	l, r := 0, len(intervals)
	// but however, the r need to minus 1 in the bottom half
	// I wonder insert at the biginning
	for l < r {
		m := l + (r-l)/2
		if intervals[m][1] == newInterval[0] {
			endIdx = m
			break
		} else if intervals[m][1] > newInterval[0] {
			r = m
		} else {
			l = m + 1
		}
	}
	if endIdx == -1 {
		endIdx = l
	}

	// find the smaller/equal start index to new interval's end
	// haha.. l start with -1, so r can over the entire array and to the left
	l, r = -1, len(intervals)-1

	for l < r {
		m := r - (r-l)/2
		// m := l + (r-l)/2, this will dead loop.. interesting
		// hmm.. so looking for first smaller, bascially I can reverse the technique
		// instead of contracting l, I contracting r
		// also the m should be calculated use m as basis... which would has bias

		if intervals[m][0] == newInterval[1] {
			startIdx = m
			break
		} else if intervals[m][0] > newInterval[1] {
			r = m - 1
		} else {
			l = m
		}
	}
	if startIdx == -1 {
		startIdx = r
	}

	return endIdx, startIdx
}

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := intervalBinSearch(intervals, newInterval)
	res := [][]int{}
	if left > right {
		res = append(res, intervals[:left]...)
		res = append(res, newInterval)
		res = append(res, intervals[left:]...)
	} else {
		res = append(res, intervals[:left]...)
		for i := left; i <= right; i++ {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
		res = append(res, newInterval)
		res = append(res, intervals[right+1:]...)
	}

	return res
}

/*
Runtime: 9 ms, faster than 65.92% of Go online submissions for Insert Interval.
Memory Usage: 4.7 MB, less than 41.43% of Go online submissions for Insert Interval.

wow.. passed on first attempt..
I am happy enough
*/

func testIntervalBinarySearch() {
	// 0 -1, insert at pos 0
	fmt.Println(intervalBinSearch([][]int{{4, 5}, {8, 10}, {12, 16}}, []int{1, 2}))
	// 4 3, insert at pos 4, which is appending actually
	fmt.Println(intervalBinSearch([][]int{{1, 2}, {4, 5}, {8, 10}, {12, 16}}, []int{17, 18}))
	// 0 3, merge all and new
	fmt.Println(intervalBinSearch([][]int{{1, 2}, {4, 5}, {8, 10}, {12, 16}}, []int{0, 17}))
	// 1 1, merge intervals[1] and new
	fmt.Println(intervalBinSearch([][]int{{1, 2}, {4, 5}, {8, 10}, {12, 16}}, []int{3, 7}))
	// 2 1, insert at index 2
	fmt.Println(intervalBinSearch([][]int{{1, 2}, {3, 5}, {8, 10}, {12, 16}}, []int{6, 7}))
	fmt.Println(intervalBinSearch([][]int{{1, 2}, {3, 5}, {8, 10}, {12, 16}}, []int{4, 7}))
	// 1 3, merge intervals[1:4] and new
	fmt.Println(intervalBinSearch([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))

	// this is cool exercise but I found this binary search very very slippy
	// will be intreseting to watch how other solved it
	// I might be just mentally farting
}
