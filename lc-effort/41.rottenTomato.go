// https://leetcode.com/problems/rotting-oranges/

package main

import "fmt"

/*
this is obviously a bfs marking problem
one caveat is when a spot can be reached by more than one source

so the queue should start with all the rotten ones
but no need to pick/choose which one first, because if a spot can be reached by more than one source
the closes source will always reach it first or it doesn't matter if two sources are at same distance

*/

func _wrong_orangesRotting(grid [][]int) int {
	orangeNum := 0
	rotten := make(map[point]struct{})
	q := []point{}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] > 0 {
				orangeNum++
			}

			if grid[r][c] == 2 {
				q = append(q, point{r, c})
			}
		}
	}

	if orangeNum == 0 {
		return 0
	}

	steps := 0
	for len(q) > 0 {
		nextQ := []point{}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if _, found := rotten[node]; found {
				continue
			}
			rotten[node] = struct{}{}

			// not necessary because of the map
			grid[node.x][node.y] = 2
			for _, next := range getNeighbors(node, grid) {
				if grid[next.x][next.y] == 1 {
					nextQ = append(nextQ, next)
				}
			}
		}
		q = nextQ
		if len(q) > 0 {
			steps++
		}
	}

	if len(rotten) == orangeNum {
		return steps
	}

	return -1
}

/*
when bfs twisted a bit
it can lead to error
e.g. {{2, 2}, {1, 1}}

when doing [1,0], it sees {1,1}.. and {1,1} is currently set to 1.. so it gets into the next queue..
and that cause the steps to add an extra 1..

so what is the better form of this
maybe I should mark them 2, when looking them up
				if grid[next.x][next.y] == 1 {
					nextQ = append(nextQ, next)
				}

				to
				if grid[next.x][next.y] == 1 {
					nextQ = append(nextQ, next)
					grid[next.x][next.y] = 2
				}


*/

func orangesRotting(grid [][]int) int {
	orangeNum := 0
	// I am using visited to keep tracking of rotten ones
	visited := make(map[point]struct{})
	q := []point{}

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] > 0 {
				orangeNum++
			}

			if grid[r][c] == 2 {
				q = append(q, point{r, c})
			}
		}
	}

	if orangeNum == 0 {
		return 0
	}

	steps := 0
	for len(q) > 0 {
		nextQ := []point{}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if _, found := visited[node]; found {
				continue
			}
			visited[node] = struct{}{}

			for _, next := range getNeighbors(node, grid) {
				if grid[next.x][next.y] == 1 {
					nextQ = append(nextQ, next)
					// mark the next batch as rotten..
					// but add them one by one to visited when it is their turn thru the process queue
					grid[next.x][next.y] = 2
				}
			}
		}
		q = nextQ
		if len(q) > 0 {
			steps++
		}
	}

	if len(visited) == orangeNum {
		return steps
	}

	return -1
}

/*
Runtime: 11 ms, faster than 11.85% of Go online submissions for Rotting Oranges.
Memory Usage: 3.8 MB, less than 10.64% of Go online submissions for Rotting Oranges.

Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotting Oranges.
Memory Usage: 3.8 MB, less than 10.33% of Go online submissions for Rotting Oranges.

still I feel the form of code isn't pretty

*/

func testOragngeRotting() {
	fmt.Println(orangesRotting([][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0, 2}}))

}
