// https://leetcode.com/problems/valid-sudoku/

package main

/*
is ther a better way that just brute force looking at
each row
eacho column
each 3*3??

let me see if it passes
*/

func isValidSudoku(board [][]byte) bool {

	// check every row
	for r := range board {
		counter := make([]int, 9)
		for c := range board[r] {
			if board[r][c] != '.' {
				counter[board[r][c]-'1']++
				if counter[board[r][c]-'1'] == 2 {
					return false
				}
			}
		}
	}
	// check every colomn
	for c := range board[0] {
		counter := make([]int, 9)
		for r := range board {
			if board[r][c] != '.' {
				counter[board[r][c]-'1']++
				if counter[board[r][c]-'1'] == 2 {
					return false
				}
			}
		}
	}

	// check 3*3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			counter := make([]int, 9)
			for r := i * 3; r < i*3+3; r++ {
				for c := j * 3; c < j*3+3; c++ {
					if board[r][c] != '.' {
						counter[board[r][c]-'1']++
						if counter[board[r][c]-'1'] == 2 {
							return false
						}
					}
				}
			}
		}
	}

	return true
}

/*
Runtime: 9 ms, faster than 22.21% of Go online submissions for Valid Sudoku.
Memory Usage: 2.6 MB, less than 87.06% of Go online submissions for Valid Sudoku.
*/
