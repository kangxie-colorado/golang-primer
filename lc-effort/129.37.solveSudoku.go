// https://leetcode.com/problems/sudoku-solver/

package main

import (
	"fmt"
	"strconv"
)

/*
except back tracking.. what can you do..
*/

func solveSudoku(board [][]byte) {
	valid := func(r, c int) bool {
		// check row
		counter := make([]int, 9)
		for j := range board[r] {
			if board[r][j] != '.' {
				counter[board[r][j]-'1']++
				if counter[board[r][j]-'1'] == 2 {
					return false
				}
			}
		}

		// check col
		counter = make([]int, 9)
		for i := range board {
			if board[i][c] != '.' {
				counter[board[i][c]-'1']++
				if counter[board[i][c]-'1'] == 2 {
					return false
				}
			}
		}

		// check 3*3
		rStart, cStart := r/3*3, c/3*3
		counter = make([]int, 9)
		for i := rStart; i < rStart+3; i++ {
			for j := cStart; j < cStart+3; j++ {
				if board[i][j] != '.' {
					counter[board[i][j]-'1']++
					if counter[board[i][j]-'1'] == 2 {
						return false
					}
				}
			}
		}
		return true
	}

	var backtracking func(r, c int) bool
	backtracking = func(r, c int) bool {
		nextRC := func() (int, int) {
			nextR := r
			nextC := c + 1
			for {
				if nextR == 9 {
					break
				}
				for ; nextC < 9; nextC++ {
					if board[nextR][nextC] == '.' {
						break
					}
				}
				if nextC == 9 {
					nextR++
					nextC = 0
				} else {
					break
				}
			}

			return nextR, nextC
		}

		for v := 1; v <= 9; v++ {
			if r == 2 && c == 0 {
				fmt.Print("")
			}
			board[r][c] = byte(v + '0')
			if valid(r, c) {
				nextR, nextC := nextRC()
				if nextR == 9 {
					return true
				}
				if backtracking(nextR, nextC) {
					return true
				}
			}

			board[r][c] = '.'
		}

		return false
	}

	for r := range board {
		for c := range board[r] {

			if board[r][c] == '.' {
				backtracking(r, c)
			}
		}
	}

}

/*
Runtime: 22 ms, faster than 23.04% of Go online submissions for Sudoku Solver.
Memory Usage: 2.1 MB, less than 32.26% of Go online submissions for Sudoku Solver.

Runtime: 12 ms, faster than 43.32% of Go online submissions for Sudoku Solver.
Memory Usage: 2.1 MB, less than 79.26% of Go online submissions for Sudoku Solver.
*/

func testSolve() {
	// b := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	b := [][]byte{{'.', '.', '9', '7', '4', '8', '.', '.', '.'}, {'7', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '2', '.', '1', '.', '9', '.', '.', '.'}, {'.', '.', '7', '.', '.', '.', '2', '4', '.'}, {'.', '6', '4', '.', '1', '.', '5', '9', '.'}, {'.', '9', '8', '.', '.', '.', '3', '.', '.'}, {'.', '.', '.', '8', '.', '3', '.', '2', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '6'}, {'.', '.', '.', '2', '7', '5', '9', '.', '.'}}

	for r := range b {
		for c := range b[r] {
			val := b[r][c] - 48
			str := strconv.Itoa(int(val))
			if b[r][c] == '.' {
				str = "."
			}
			fmt.Print(str, " ")
		}
		fmt.Println()
	}
	fmt.Println()
	solveSudoku(b)
	for r := range b {
		for c := range b[r] {
			val := b[r][c] - 48
			str := strconv.Itoa(int(val))
			if b[r][c] == '.' {
				str = "."
			}
			fmt.Print(str, " ")
		}
		fmt.Println()
	}
}

func main() {
	testSolve()
}
