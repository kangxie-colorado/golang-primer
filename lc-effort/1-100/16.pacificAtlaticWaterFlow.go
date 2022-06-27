// https://leetcode.com/problems/pacific-atlantic-water-flow/

package main

import "fmt"

/*
  the natural solution is BFS
	for any node, it can reach both ocean (r==0||c==0) && (r==len||c==len(r))
	or any node that can, so use a status array to track them... this saves a lot of calculation

	I don't know if it is going to be quick but at least worthy trying



*/

// looking from inside  out
// so >=
func getNeighborIslands(p point, m [][]int) []point {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	res := []point{}
	for _, d := range dirs {
		newX := p.x + d[0]
		newY := p.y + d[1]
		if newX < len(m) && newX >= 0 && newY < len(m[0]) && newY >= 0 && m[p.x][p.y] >= m[newX][newY] {
			res = append(res, point{newX, newY})
		}

	}

	return res
}

// looking from outside in
// so <=
func getNeighborIslandsInwards(p point, m [][]int) []point {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	res := []point{}
	for _, d := range dirs {
		newX := p.x + d[0]
		newY := p.y + d[1]
		if newX < len(m) && newX >= 0 && newY < len(m[0]) && newY >= 0 && m[p.x][p.y] <= m[newX][newY] {
			res = append(res, point{newX, newY})
		}

	}

	return res
}

func bfs(r, c int, m [][]int, ms [][]bool) bool {
	if ms[r][c] {
		return true
	}

	reachP := false
	reachA := false

	visited := map[point]struct{}{}

	q := []point{point{r, c}}
	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		if _, found := visited[n]; found {
			continue
		}

		visited[n] = struct{}{}

		if ms[n.x][n.y] {
			return true
		}

		if n.x == 0 || n.y == 0 {
			reachP = true
		}

		if n.x == len(m)-1 || n.y == len(m[0])-1 {
			reachA = true
		}

		if reachA && reachP {
			return true
		}

		for _, neighbor := range getNeighborIslands(n, m) {
			q = append(q, neighbor)
		}

	}

	return reachA && reachP
}

func _1_pacificAtlantic(heights [][]int) [][]int {
	status := make([][]bool, len(heights))
	for r := range heights {
		status[r] = make([]bool, len(heights[r]))
		for c := range heights[r] {
			status[r][c] = false
		}
	}

	res := [][]int{}

	for r := range heights {
		for c := range heights[r] {

			status[r][c] = bfs(r, c, heights, status)
			if status[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}

/*
Success
Details
Runtime: 269 ms, faster than 15.83% of Go online submissions for Pacific Atlantic Water Flow.
Memory Usage: 7.8 MB, less than 48.33% of Go online submissions for Pacific Atlantic Water Flow.

then I think I can do the frontline expansion on two separate problem -- reach P and reach A
for it starts
1 1 1 1 1
1
1
1
1

then expand
1 1 1 1 1
1 1 1 1 1
1 1 1
1 1
1

until it stops expanding
*/

func pacificAtlantic(heights [][]int) [][]int {
	reachP := make([][]bool, len(heights))
	reachA := make([][]bool, len(heights))

	frontLineP := make(map[point]bool)
	frontLineA := make(map[point]bool)
	for r := range heights {
		reachP[r] = make([]bool, len(heights[r]))
		reachA[r] = make([]bool, len(heights[r]))
		for c := range heights[r] {

			reachP[r][c] = false
			reachA[r][c] = false

			if r == 0 || c == 0 {
				reachP[r][c] = true
				frontLineP[point{r, c}] = true
			}

			if r == len(heights)-1 || c == len(heights[0])-1 {
				reachA[r][c] = true
				frontLineA[point{r, c}] = true
			}

		}
	}

	// reach P
	for true {
		nextFrontLine := make(map[point]bool)

		for p, _ := range frontLineP {
			for _, n := range getNeighborIslandsInwards(p, heights) {
				if reachP[n.x][n.y] == false {
					reachP[n.x][n.y] = true
					nextFrontLine[n] = true
				}
			}
		}

		if len(nextFrontLine) == 0 {
			break
		}
		frontLineP = nextFrontLine
	}

	// reach A
	for true {
		nextFrontLine := make(map[point]bool)

		for a, _ := range frontLineA {
			for _, n := range getNeighborIslandsInwards(a, heights) {
				if reachA[n.x][n.y] == false {
					reachA[n.x][n.y] = true
					nextFrontLine[n] = true
				}
			}
		}

		if len(nextFrontLine) == 0 {
			break
		}
		frontLineA = nextFrontLine
	}

	res := [][]int{}
	for r := range reachP {
		for c := range reachP[r] {
			if reachA[r][c] && reachP[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}

/*
Success
Details
Runtime: 45 ms, faster than 56.67% of Go online submissions for Pacific Atlantic Water Flow.
Memory Usage: 7.2 MB, less than 73.75% of Go online submissions for Pacific Atlantic Water Flow.
Next challenges:
*/

func testPacificAtlantic() {
	fmt.Println(pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
}
