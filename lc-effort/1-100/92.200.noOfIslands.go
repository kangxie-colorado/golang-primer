// https://leetcode.com/problems/number-of-islands/

package main

/*
	this should be a BFS right
	traverse all the reach nodes.. flip it to -1

	finish.. then count
*/

func _bfs_numIslands(grid [][]byte) int {
	count := 0
	bfs := func(r, c int) {
		q := []point{{r, c}}

		for len(q) > 0 {
			node := q[0]
			q = q[1:]

			if grid[node.x][node.y] == '2' {
				continue
			}
			grid[node.x][node.y] = '2'

			for _, next := range getNeighbors2(node, len(grid), len(grid[0])) {
				if grid[next.x][next.y] == '1' {
					q = append(q, next)
				}
			}
		}
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '2' {
				continue
			}

			if grid[r][c] == '1' {
				count++
				bfs(r, c)
			}
		}
	}

	return count
}

/*
Runtime: 11 ms, faster than 32.64% of Go online submissions for Number of Islands.
Memory Usage: 6.8 MB, less than 19.04% of Go online submissions for Number of Islands.

good enough

turns out dfs is quicker..
*/

func numIslands(grid [][]byte) int {
	count := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		stack := []point{{r, c}}

		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if grid[node.x][node.y] == '2' {
				continue
			}
			grid[node.x][node.y] = '2'

			for _, next := range getNeighbors2(node, len(grid), len(grid[0])) {
				if grid[next.x][next.y] == '1' {
					stack = append(stack, next)
				}
			}
		}
	}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '2' {
				continue
			}

			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count
}
