// https://leetcode.com/problems/minimize-the-difference-between-target-and-chosen-elements/

package main

import (
	"fmt"
	"sort"
)

/*
analysis...
it isn't very clear how to tackle it

but maybe I have a solution
I do a paritial column sum then binary search

1,2,3
4,5,6
7,8,9

will become
1,2,3
5,7,9,
12,15,18

then binary search 13 in last row.. find the min-abs-diff, which will be idx0 (7)
13-A[2][0] = 6

continue this up.. and we can find the min-abs-diff
but if there negative value.. then

aha 1 <= mat[i][j] <= 70
no need to worry about that

*/

func _1_logically_wrong_minimizeTheDifference(mat [][]int, target int) int {
	partialColSum := make([][]int, len(mat))
	for r := range partialColSum {
		sort.Ints(mat[r])
		partialColSum[r] = make([]int, len(mat[r]))

		for c := range partialColSum[r] {
			if r == 0 {
				partialColSum[r][c] = mat[r][c]
			} else {
				partialColSum[r][c] = partialColSum[r-1][c] + mat[r][c]
			}
		}
	}

	// now we have the paritial sum, we can go serch row by row for the closes diff
	chosenElementsSum := 0
	rowTarget := target
	for r := len(mat) - 1; r >= 0; r-- {

		// this binary search doesn't work on searching the min abs diff..
		// because it is not sorted by abs diff
		// let me just do a linear search
		// and maybe I can sort the array by absdiff in first place (next try)
		/*
			left, right := 0, len(mat[r])
			for left < right {
				mid := left + (right-left)/2
				absDiff := abs(partialColSum[r][mid] - rowTarget)

				if absDiff == 0 {
					left = mid
					break
				} else if absDiff > 0 {
					right = mid
				} else {
					left = mid + 1
				}
			}
			chosenElementsSum += mat[r][left]
			rowTarget -= mat[r][left]
		*/
		// 1 <= mat[i][j] <= 70
		// 1 <= target <= 800
		// this should be enough given above constraints
		minAdsDiff := 100000
		minIdx := -1
		for i := 0; i < len(mat[r]); i++ {
			if abs(partialColSum[r][i]-rowTarget) < minAdsDiff {
				minIdx = i
				minAdsDiff = abs(partialColSum[r][i] - rowTarget)
			}
		}

		chosenElementsSum += mat[r][minIdx]
		rowTarget -= mat[r][minIdx]
	}

	return abs(chosenElementsSum - target)

}

/*
60 / 81 test cases passed.
[[4,2,6],[2,1,8],[3,9,10],[7,8,9],[6,3,6],[5,5,10],[7,1,9],[3,1,5],[1,3,3],[3,2,8]]
61
Output:
1
Expected:
0

my algorithm is not full? maybe when two are at the same absdiff.. I need to consider two possibilities?
push this thought down.. for now
let me sort by absdiff at first place
hmm... not really easy.. discard the thought for now

cont to fix this code
same code... but will try to fix

okay.. admit it although it passed 40 cases.. it is still logically wrong
rethink...

probably still back-tracking
*/

func minimizeTheDifference(mat [][]int, target int) int {
	partialColSum := make([][]int, len(mat))
	for r := range partialColSum {
		sort.Ints(mat[r])
		partialColSum[r] = make([]int, len(mat[r]))

		for c := range partialColSum[r] {
			if r == 0 {
				partialColSum[r][c] = mat[r][c]
			} else {
				partialColSum[r][c] = partialColSum[r-1][c] + mat[r][c]
			}
		}
	}

	// now we have the paritial sum, we can go serch row by row for the closes diff
	chosenElementsSum := 0
	rowTarget := target
	for r := len(mat) - 1; r >= 0; r-- {
		minAdsDiff := 100000
		minIdx := -1
		for i := 0; i < len(mat[r]); i++ {
			if abs(partialColSum[r][i]-rowTarget) <= minAdsDiff {
				minIdx = i
				minAdsDiff = abs(partialColSum[r][i] - rowTarget)
			}
		}

		chosenElementsSum += mat[r][minIdx]
		rowTarget -= mat[r][minIdx]
	}

	// if abs(chosenElementsSum - target) is 0, then this actually can finish now
	if abs(chosenElementsSum-target) == 0 {
		return 0
	}

	// if not, I have to go back to see
	sum := make(map[int]struct{})
	res := 10000
	sum[0] = struct{}{}

	for r := range mat {
		sumTemp := make(map[int]struct{})
		for c := range mat[r] {
			for k := range sum {

				if k+mat[r][c] < target+abs(chosenElementsSum-target) {
					sumTemp[k+mat[r][c]] = struct{}{}
				}
				if r == len(mat)-1 {
					res = min(res, abs(k+mat[r][c]-target))
				}
			}
		}
		sum = sumTemp

	}

	return res

}

/*
I must stick to my own pricinples
- stuck 30 minutes, check hints
- stuck 1 hr, give up and check answers

So the hints
- The sum of chosen elements will not be too large. Consider using a hash set to record all possible sums while iterating each row.
- Instead of keeping track of all possible sums, since in each row, we are adding positive numbers, only keep those that can be a candidate, not exceeding the target by too much.

WTF? so basically it is brute force.. but be smart about it
so okay... previously I established a range

actually if it is a zero diff, then my solution actually covered it
if not, I had a max diff... so I can use it


ugh...
Runtime: 820 ms, faster than 40.00% of Go online submissions for Minimize the Difference Between Target and Chosen Elements.
Memory Usage: 8.1 MB, less than 20.00% of Go online submissions for Minimize the Difference Between Target and Chosen Elements.

what a stupid problem!!!

*/

func testMinimizeTheDifference() {
	fmt.Println(minimizeTheDifference([][]int{{1, 2, 9, 8, 7}}, 6))
	fmt.Println(minimizeTheDifference([][]int{{4, 2, 6}, {2, 1, 8}, {3, 9, 10}, {7, 8, 9}, {6, 3, 6}, {5, 5, 10}, {7, 1, 9}, {3, 1, 5}, {1, 3, 3}, {3, 2, 8}}, 61))
}
