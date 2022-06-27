// https://leetcode.com/problems/n-queens/

package main

import "fmt"

/*
same as the other problem but to keep a tracking of the combination
*/

func solveNQueensHelper(row, n int, colTaken, posDiagTaken, negDiagTaken []int, comb [][]byte, res *[][]string) {
	if row == n {
		res1 := []string{}
		for _, c := range comb {
			res1 = append(res1, string(c))
		}
		*res = append(*res, res1)

		return
	}

	// every row could have a possible n choices of colum
	for i := 0; i < n; i++ {

		if colTaken[i] != 0 || posDiagTaken[row+i] != 0 || negDiagTaken[row-i+n-1] != 0 {
			continue
		}

		comb[row][i] = 'Q'
		colTaken[i] = 1
		posDiagTaken[row+i] = 1
		negDiagTaken[row-i+n-1] = 1
		solveNQueensHelper(row+1, n, colTaken, posDiagTaken, negDiagTaken, comb, res)

		comb[row][i] = '.'
		colTaken[i] = 0
		posDiagTaken[row+i] = 0
		negDiagTaken[row-i+n-1] = 0

	}

}

func solveNQueens(n int) [][]string {
	res := [][]string{}
	comb := make([][]byte, n)
	for i := range comb {
		comb[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			comb[i][j] = '.'
		}
	}

	colTaken := make([]int, n)
	posDiagTaken := make([]int, 2*n-1) // r+c
	negDiagTaken := make([]int, 2*n-1) // r-c (need to apply a  +(n-1))

	solveNQueensHelper(0, n, colTaken, posDiagTaken, negDiagTaken, comb, &res)
	return res
}

/*
Runtime: 4 ms, faster than 81.95% of Go online submissions for N-Queens.
Memory Usage: 3.4 MB, less than 49.80% of Go online submissions for N-Queens.
*/

func testNQueens1() {
	fmt.Println(solveNQueens(1))
	fmt.Println(solveNQueens(2))
	fmt.Println(solveNQueens(4))

}
