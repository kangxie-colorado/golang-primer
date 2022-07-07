// https://leetcode.com/problems/count-servers-that-communicate/

package main

import "fmt"

/*
just bfs?
*/

func _bfs_countServers(grid [][]int) int {
	bfs := func(r, c int) int {
		count := 0
		q := []point{{r, c}}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if grid[node.x][node.y] == 0 {
				continue
			}

			grid[node.x][node.y] = 0
			count++
			for i := 0; i < len(grid); i++ {
				if grid[i][node.y] == 0 {
					continue
				}
				q = append(q, point{i, node.y})
			}
			for j := 0; j < len(grid[r]); j++ {
				if grid[node.x][j] == 0 {
					continue
				}
				q = append(q, point{node.x, j})
			}
		}

		return count
	}

	servers := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				count := bfs(r, c)
				if count > 1 {
					servers += count
				}
			}
		}
	}

	return servers
}

/*
Input [[1,0,0,1,0],[0,0,0,0,0],[0,0,0,1,0]]
Output 0
Expected 3

ah... I was too quick to jump on conclusion..
this is not plain bfs... same row/colum can communicate as well..

need to modify it
I made the stupid mistake of use r/c as node.x/y...
not being careful at all...

Runtime: 130 ms, faster than 6.25% of Go online submissions for Count Servers that Communicate.
Memory Usage: 9.3 MB, less than 6.25% of Go online submissions for Count Servers that Communicate.
yes I think dfs should be better

*/

func dfs_countServers(grid [][]int) int {
	dfs := func(r, c int) int {
		count := 0
		q := []point{{r, c}} // keep the name as q but use it as stack
		for len(q) > 0 {
			node := q[len(q)-1]
			q = q[:len(q)-1]
			if grid[node.x][node.y] == 0 {
				continue
			}

			grid[node.x][node.y] = 0
			count++
			for i := 0; i < len(grid); i++ {
				if grid[i][node.y] == 0 {
					continue
				}
				q = append(q, point{i, node.y})
			}
			for j := 0; j < len(grid[r]); j++ {
				if grid[node.x][j] == 0 {
					continue
				}
				q = append(q, point{node.x, j})
			}
		}

		return count
	}

	servers := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				count := dfs(r, c)
				if count > 1 {
					servers += count
				}
			}
		}
	}

	return servers
}

/*
Runtime: 73 ms, faster than 75.00% of Go online submissions for Count Servers that Communicate.
Memory Usage: 9.7 MB, less than 6.25% of Go online submissions for Count Servers that Communicate.

Runtime: 69 ms, faster than 81.25% of Go online submissions for Count Servers that Communicate.
Memory Usage: 9.9 MB, less than 6.25% of Go online submissions for Count Servers that Communicate.

I almost discover the trick, row+col>2 means it can communicate...
but yet I still not able to see thru

admit it.. IQ is not as high as the other guy but still let me code it up
*/

func countServers(grid [][]int) int {
	rowSum := make([]int, len(grid))
	colSum := make([]int, len(grid[0]))
	for r := range grid {
		for c := range grid[r] {
			rowSum[r] += grid[r][c]
			colSum[c] += grid[r][c]
		}
	}

	servers := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 1 {
				if rowSum[r]+colSum[c] > 2 {
					servers++
				}
			}

		}
	}

	return servers
}

/*
Runtime: 55 ms, faster than 97.92% of Go online submissions for Count Servers that Communicate.
Memory Usage: 8 MB, less than 29.17% of Go online submissions for Count Servers that Communicate.
*/

func testcountServers() {
	fmt.Println(countServers([][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}))
	fmt.Println(countServers([][]int{{1, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}}))
}
