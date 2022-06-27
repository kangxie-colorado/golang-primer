// https://leetcode.com/problems/spiral-matrix/

package main

/*
it seems like a simple dfs
dir is right, down, left, up.. and rotating
*/

func spiralOrder(matrix [][]int) []int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n := len(matrix), len(matrix[0])
	dir := 0

	res := []int{}
	next := []int{0, 0}
	for len(res) < m*n {
		res = append(res, matrix[next[0]][next[1]])
		matrix[next[0]][next[1]] = 1000 // walked here

		newRow := next[0] + dirs[dir][0]
		newCol := next[1] + dirs[dir][1]

		if newRow >= m || newCol >= n || newCol < 0 || matrix[newRow][newCol] == 1000 {
			dir = (dir + 1) % 4
			newRow = next[0] + dirs[dir][0]
			newCol = next[1] + dirs[dir][1]
		}

		next = []int{newRow, newCol}

	}

	return res
}

/*
Runtime: 1 ms, faster than 58.87% of Go online submissions for Spiral Matrix.
Memory Usage: 1.9 MB, less than 95.36% of Go online submissions for Spiral Matrix.

Runtime: 0 ms, faster than 100.00% of Go online submissions for Spiral Matrix.
Memory Usage: 1.9 MB, less than 74.19% of Go online submissions for Spiral Matrix.
*/
