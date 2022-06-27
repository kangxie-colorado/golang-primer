// https://leetcode.com/problems/sum-of-even-numbers-after-queries/

package main

import "fmt"

/*
this seems pretty easy
just keep an evenSum, update it on the go
*/

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	evenSum := 0
	for _, n := range nums {
		if n%2 == 0 {
			evenSum += n
		}
	}

	res := []int{}
	for _, q := range queries {
		old := nums[q[1]]
		new := old + q[0]
		nums[q[1]] = new

		if old%2 == 0 && new%2 == 0 {
			// even to even
			evenSum += q[0]
		} else if old%2 != 0 && new%2 == 0 {
			// odd to even
			evenSum += new
		} else if old%2 == 0 && new%2 != 0 {
			// even to odd
			evenSum -= old
		}
		res = append(res, evenSum)
	}

	return res
}

/*
tricky point, -1%2 == -1
not 1

Runtime: 81 ms, faster than 75.00% of Go online submissions for Sum of Even Numbers After Queries.
Memory Usage: 8 MB, less than 25.00% of Go online submissions for Sum of Even Numbers After Queries.
*/

func testSumEvenAfterQueries() {
	//[1,2,3,4]
	// [[1,0],[-3,1],[-4,0],[2,3]]
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}))
}
