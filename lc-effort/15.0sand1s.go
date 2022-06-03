// https://leetcode.com/problems/ones-and-zeroes/

package main

import (
	"fmt"
	"strings"
)

/*
	largest subsets with at most m 1s and n 0s


	first observation:
	at most m 1s and  n 0s, means at most m+n characters
	meaning this favors shorter strs

	so the choice is again, choose ith or not..
	if able to choosee... or not able to..

	sounds like some transition formula can be established?

*/

func maxForm(weights []int, total int) int {
	sums := make([]int, total+1)
	for i := range sums {
		sums[i] = 0
	}

	for _, w := range weights {
		for j := len(sums) - 1; j >= 0; j-- {
			if w == 0 {
				sums[j]++
			} else {
				if j >= w {
					// if w==0, sums[j-w] is sums[j], and auto +1
					sums[j] = max(sums[j], sums[j-w]+1)
				}
			}

		}
	}

	return sums[total]
}

func findMaxForm(strs []string, m int, n int) int {

	totals := make([]int, len(strs))
	ones := make([]int, len(strs))
	zeros := make([]int, len(strs))

	for i, s := range strs {
		totals[i] = len(s)
		ones[i] = strings.Count(s, "1")
		zeros[i] = strings.Count(s, "0")
	}

	maxTotal := maxForm(totals, m+n)
	maxZero := maxForm(zeros, m)
	maxOne := maxForm(ones, n)

	return min(maxOne, min(maxZero, maxTotal))

}

/*
Success
Details
Runtime: 5 ms, faster than 100.00% of Go online submissions for Ones and Zeroes.
Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Ones and Zeroes.
*/

func testMaxForm() {

	fmt.Println(maxForm([]int{2, 1, 1}, 2))
	fmt.Println(maxForm([]int{0, 1, 1}, 1))

	fmt.Println(maxForm([]int{2, 4, 6, 1, 1}, 8))
	fmt.Println(maxForm([]int{1, 1, 4, 1, 0}, 3))
	fmt.Println(maxForm([]int{1, 3, 2, 0, 1}, 5))

}

func testFindMaxForm() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"0", "0", "1", "1"}, 2, 2))

}
