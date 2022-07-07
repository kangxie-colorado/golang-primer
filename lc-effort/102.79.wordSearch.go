// https://leetcode.com/problems/word-search/

package main

/*
not much but backtracking
when back trcking, the char must be put back to unused? yeah, feels like so
*/

func exist(board [][]byte, word string) bool {
	var backtrack func(char int, r, c int) bool
	// visited := [][]int{}
	// 1 <= m, n <= 6, so use a 2D array to save the state
	visited := make([][]int, len(board))
	for r := range visited {
		visited[r] = make([]int, len(board[0]))
		for c := range visited[r] {
			visited[r][c] = 0
		}
	}

	backtrack = func(char int, r, c int) bool {
		if char == len(word)-1 {
			return true
		}

		visited[r][c] = 1
		for _, nei := range getNeighbors2(point{r, c}, len(board), len(board[0])) {
			if visited[nei.x][nei.y] == 1 {
				continue
			}

			if board[nei.x][nei.y] == word[char+1] {
				if backtrack(char+1, nei.x, nei.y) {
					return true
				}
			}
		}
		visited[r][c] = 0
		return false
	}

	for r := range board {
		for c := range board[r] {
			if board[r][c] == word[0] && backtrack(0, r, c) {
				return true
			}
		}
	}

	return false
}

/*
Runtime: 517 ms, faster than 23.14% of Go online submissions for Word Search.
Memory Usage: 6.8 MB, less than 13.85% of Go online submissions for Word Search.

add some memorization
hmm...

will fail at
C A A
A A A
B C D
*/
