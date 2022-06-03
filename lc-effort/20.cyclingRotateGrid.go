// https://leetcode.com/problems/cyclically-rotating-a-grid/

package main

import "fmt"

/*
analysis
	rotate a layer doesn't seem so hard
	save up the values.. of left-top cell, assign the right value to this cell

	then walk to right/down.. ...
	notice it change direction use a dir
	notice it will also assign the saved value to last one
	notice it can be saved by %(m+n)

	then m-=2;n-=2;
	when m or n reaches 2, then this is the last..
	but no matter just let it go on and exit on hitting 0


rotateGrind(grid, k)
	m:=len(grid)
	n:=len(grid[0])

	layer:=0
	for m>0 && n>0 {
		rotateLayer(m,n,layer, grid)
		m-=2
		n-=2
		layer+=1
	}

rotateLayer(m,n,layer, grid)
	start:={layer,layer}
	dirs = {
		r: {0, 1}
		d: {1,0}
		l: {0,-1}
		u: {-1,0}
	}

	dir=0
	buffer:=grid[start]

	# walt to right

	next:=start
	dirIdx=0

	for true

		if next+dir = start
			# full circle completed
			grid[next] = grid[start]
			break

		if next+dir isOutofLayer()
			dirIdx++

		dir := dirs[dirIdx]
		grid[next] = grid[next+dir]
		next = next+dir

uh. k steps.. you can let it roate k times or you can look into k steps beyond..
rotate k times.. apparenetly wasteful..
look into k steps... k%(m=n) really

a bit hard.. but let me try

rotateLayer(m,n,layer, grid, k)
	start:={layer,layer}
	steps = {
		{0, 1} * m -1
		{1,0} * n-1
		l: {0,-1} * m-1
		u: {-1,0}* n-1
	}

	buffer = make(int[], k)

	next:=start
	for i:=0;i<k;i++ {
		# save k values in buffer
		buffer[i] = next

		# walk steps to k
		next = next + steps[i]

	}

	# begin the rotate
	pos:=start
	total:=2m+2n-4
	for i:=0;i< total; i++ {
		if k+i >= total {
			grid[pos] = buffer[0]
			buffer = buffer[1:]
			pos = pos + steps[i]
			cont
		}

		# when roate by m+n-k you need to use buffer
		grid[pos] = grid[next]

		# walk from the k steps
		next = next+steps[k+i]
		pos = pos + steps[i]
	}



*/

func rotateLayer(m, n, layer, k int, grid [][]int) {
	start := point{layer, layer}
	total := 2*m + 2*n - 4
	k = k % total

	steps := make([]point, total)
	si := 0

	for i := 0; i < n-1; i++ {
		steps[si] = point{0, 1}
		si++
	}

	for i := 0; i < m-1; i++ {
		steps[si] = point{1, 0}
		si++
	}

	for i := 0; i < n-1; i++ {
		steps[si] = point{0, -1}
		si++
	}

	for i := 0; i < m-1; i++ {
		steps[si] = point{-1, 0}
		si++
	}

	buffer := make([]int, k)
	next := start
	for i := 0; i < k; i++ {
		buffer[i] = grid[next.x][next.y]

		next.x = next.x + steps[i].x
		next.y = next.y + steps[i].y
	}

	// begin rotate
	pos := start
	for i := 0; i < total; i++ {
		if k+i >= total {
			grid[pos.x][pos.y] = buffer[0]
			buffer = buffer[1:]
			pos.x = pos.x + steps[i].x
			pos.y = pos.y + steps[i].y
			continue
		}

		grid[pos.x][pos.y] = grid[next.x][next.y]

		next.x = next.x + steps[k+i].x
		next.y = next.y + steps[k+i].y

		pos.x = pos.x + steps[i].x
		pos.y = pos.y + steps[i].y
	}

}

func rotateGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	layer := 0

	for m > 0 && n > 0 {
		rotateLayer(m, n, layer, k, grid)
		m -= 2
		n -= 2
		layer += 1
	}

	return grid

}

/*
Success
Details
Runtime: 42 ms, faster than 33.33% of Go online submissions for Cyclically Rotating a Grid.
Memory Usage: 7.7 MB, less than 33.33% of Go online submissions for Cyclically Rotating a Grid.
Next challenges:

wow... not super but accepted..

*/

func testRotateGrid() {

	fmt.Println(rotateGrid([][]int{{40, 10}, {30, 20}}, 1))
	fmt.Println(rotateGrid([][]int{{40, 10}, {30, 20}}, 2))

	fmt.Println(rotateGrid([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, 2))

}
