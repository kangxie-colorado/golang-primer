package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func readCaveStdin() [][]int {
	scanner := bufio.NewScanner(os.Stdin)

	cave := [][]int{}
	for scanner.Scan() {
		row := []int{}
		for _, b := range scanner.Bytes() {
			row = append(row, int(b-'0'))
		}

		cave = append(cave, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return cave
}

var locRiskMap [][]int

func copyVisited(visited [][]int) [][]int {

	visistedCopy := [][]int{}
	for r := range visited {
		rowCopy := []int{}
		for c := range visited[r] {
			rowCopy = append(rowCopy, visited[r][c])
		}

		visistedCopy = append(visistedCopy, rowCopy)
	}

	return visistedCopy
}

var lowRisk int = 719
var runs int = 0

func riskFromPoint(row, col, row2, col2 int, cave *[][]int, visited [][]int, level int, riskSofar int) {

	runs++
	visited[row][col] = 1

	if runs%10000 == 0 {
		fmt.Println("Run:", runs, "Level", level, "Visited So Far!")
		for r := range visited {
			fmt.Printf("%v\n", visited[r])
		}
	}

	/*
		fmt.Println("Level", level, "Visited So Far!")
		for r := range visited {
			fmt.Printf("%v\n", visited[r])
		}
	*/

	riskPath := riskSofar + (*cave)[row][col]
	if riskPath > lowRisk {
		// this path has already greater risk, even it is not ended yet
		return
	}

	if row == row2 && col == col2 {

		fmt.Printf("Found a path with risk %v, currently lowRisk %v\n", riskPath, lowRisk)
		if riskPath < lowRisk {
			lowRisk = riskPath
		}
		/*
			fmt.Println("Level", level, "Visited So Far!")
			for r := range visited {
				fmt.Printf("%v\n", visited[r])
			}
		*/

		return
	}

	// wanted to generalize the out of bound
	// but it becomes one true-end condition, meaning when step out of bound, the whole function could terminate?
	// better way is not to step out of bounds;
	// key is not to bear a value for the invalid location

	// calculate two risks: going right or down
	if col < len((*cave)[0])-1 && row < len(*cave)-1 && (*cave)[row][col+1] < (*cave)[row+1][col] {
		if col < len((*cave)[0])-1 && visited[row][col+1] == 0 {
			riskFromPoint(row, col+1, row2, col2, cave, copyVisited(visited), level+1, riskPath)
		}

		if row < len(*cave)-1 && visited[row+1][col] == 0 {
			riskFromPoint(row+1, col, row2, col2, cave, copyVisited(visited), level+1, riskPath)
		}
	} else {

		if row < len(*cave)-1 && visited[row+1][col] == 0 {
			riskFromPoint(row+1, col, row2, col2, cave, copyVisited(visited), level+1, riskPath)
		}

		if col < len((*cave)[0])-1 && visited[row][col+1] == 0 {
			riskFromPoint(row, col+1, row2, col2, cave, copyVisited(visited), level+1, riskPath)
		}

	}

	if col > 0 && visited[row][col-1] == 0 {
		riskFromPoint(row, col-1, row2, col2, cave, copyVisited(visited), level+1, riskPath)
	}

	if row > 0 && visited[row-1][col] == 0 {
		riskFromPoint(row-1, col, row2, col2, cave, copyVisited(visited), level+1, riskPath)
	}

	/*
		if col > 0 && row > 0 && (*cave)[row][col-1] < (*cave)[row-1][col] {
			if col > 0 && visited[row][col-1] == 0 {
				riskFromPoint(row, col-1, row2, col2, cave, copyVisited(visited), level+1, riskPath)
			}

			if row > 0 && visited[row-1][col] == 0 {
				riskFromPoint(row-1, col, row2, col2, cave, copyVisited(visited), level+1, riskPath)
			}

		} else {
			if row > 0 && visited[row-1][col] == 0 {
				riskFromPoint(row-1, col, row2, col2, cave, copyVisited(visited), level+1, riskPath)
			}

			if col > 0 && visited[row][col-1] == 0 {
				riskFromPoint(row, col-1, row2, col2, cave, copyVisited(visited), level+1, riskPath)
			}

		}
	*/

}

func lowestRisk(cave *[][]int, visited [][]int) {
	// hybrid
	riskFromPointHybrid(0, 0, 0, 0, len(*cave)-1, len((*cave)[0])-1, cave, visited, 0, 0, nil, &lowRisk)
	fmt.Printf("lowestRisk: %v\n", lowRisk-(*cave)[0][0])

	// parallel + dfs
	/*
		done := make(chan bool)
		go riskFromPointParallel(0, 0, len(*cave)-1, len((*cave)[0])-1, cave, visited, 0, 0, done, &lowRisk)
		<-done
	*/

	// parallel + dfs + using middle point
	/*
		lowestRiskSofar := math.MaxInt32

		lowestByPoint := [][]int{}
		for i := 0; i < len(*cave); i++ {
			riskRow := make([]int, len((*cave)[0]))
			for j := 0; j < len(riskRow); j++ {
				riskRow[j] = 0
			}
			lowestByPoint = append(locRiskMap, riskRow)
		}

		deltas := [][]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		}

		r := len(*cave) / 2
		{
			c := len((*cave)[0]) / 2
			{
				stack := []Loc{}
				stack = append(stack, Loc{r, c})

				for len(stack) != 0 {

					r, c = stack[len(stack)-1].row, stack[len(stack)-1].col
					stack = stack[:len(stack)-1]

					if r < 0 || r > len(*cave)-1 || c < 0 || c > len((*cave)[0])-1 {
						continue
					}

					if (r == 0 && c == 0) || (r == len(*cave)-1 && c == len((*cave)[0])-1) {
						continue
					}

					if lowestByPoint[r][c] >= lowestRiskSofar {
						continue
					}

					for d := range deltas {
						stack = append(stack, Loc{r + deltas[d][0], c + deltas[d][1]})
					}

					risk1, risk2 := math.MaxInt32, math.MaxInt32

					var visited [][]int

					for i := 0; i < len(*cave); i++ {
						riskRow := make([]int, len((*cave)[0]))
						visitedRow := make([]int, len((*cave)[0]))
						for j := 0; j < len(riskRow); j++ {
							visitedRow[j] = 0
						}
						visited = append(visited, visitedRow)
					}
					copyVis := copyVisited(visited)

					chan1, chan2 := make(chan bool), make(chan bool)

					fmt.Println("part1---")
					go riskFromPointParallel(0, 0, 0, 0, r, c, cave, visited, 0, 0, chan1, &risk1)

					fmt.Println("part2---")
					go riskFromPointParallel(r, c, r, c, len(*cave)-1, len((*cave)[0])-1, cave, copyVis, 0, 0, chan2, &risk2)

					<-chan1
					<-chan2

					risk1 = risk1 - (*cave)[0][0]
					risk2 = risk2 - (*cave)[r][c]

					fmt.Printf("r:%v, c:%v, risk1: %v, risk2: %v, total risk: %v\n", r, c, risk1, risk2, risk1+risk2)

					if risk1+risk2 < lowestRiskSofar {
						lowestRiskSofar = risk1 + risk2
					}

					lowestByPoint[r][c] = risk1 + risk2
				}

			}
		}

		fmt.Println(lowestRiskSofar)
		fmt.Println(lowestByPoint)
	*/

	// plain dfs
	//riskFromPoint(0, 0, len(*cave)-1, len((*cave)[0])-1, cave, visited, 0, 0)

	// use a middle point, plain dfs
	/*
		lowestRiskSofar := math.MaxInt32

		for r := range *cave {
			for c := range (*cave)[0] {
				if (r == 0 && c == 0) || (r == len(*cave)-1 && c == len((*cave)[0])-1) {
					continue
				}

				fmt.Println("part1---")
				risk1, risk2 := -1, -1
				lowRisk = lowestRiskSofar
				riskFromPoint(0, 0, r, c, cave, visited, 0, 0)
				risk1 = lowRisk - (*cave)[0][0]

				fmt.Println("part2---")
				lowRisk = lowestRiskSofar
				copyVis := copyVisited(visited)
				riskFromPoint(r, c, len(*cave)-1, len((*cave)[0])-1, cave, copyVis, 0, 0)
				risk2 = lowRisk - (*cave)[r][c]

				fmt.Printf("r:%v, c:%v, risk1: %v, risk2: %v, total risk: %v\n", r, c, risk1, risk2, risk1+risk2)

				if risk1+risk2 < lowestRiskSofar {
					lowestRiskSofar = risk1 + risk2

				}
			}
		}

		fmt.Println(lowestRiskSofar)
	*/
}

func chitonDriver() {
	cave := readCaveStdin()

	var visited [][]int

	for i := 0; i < len(cave); i++ {
		riskRow := make([]int, len(cave[0]))
		visitedRow := make([]int, len(cave[0]))
		for j := 0; j < len(riskRow); j++ {
			riskRow[j] = -1
			visitedRow[j] = 0
		}
		locRiskMap = append(locRiskMap, riskRow)
		visited = append(visited, visitedRow)
	}

	lowestRisk(&cave, visited)

}

// okay I was misled by the example
// it can also move uppward and left
// this is not a dynamic programming problem
// this is dfs/least weight problem

func riskFromPointParallel(row, col, row1, col1, row2, col2 int, cave *[][]int, visited [][]int, level int, riskSofar int, done chan bool, lowestRisk *int) {

	runs++
	visited[row][col] = 1

	if runs%3000 == 0 {
		fmt.Println("Run:", runs, "Level", level, "risksofar:", riskSofar, "row, col:", row, col)
		if runs%60000 == 0 {
			for r := range visited {
				fmt.Printf("%v\n", visited[r])
			}
		}

	}

	/*
		fmt.Println("Level", level, "Visited So Far!")
		for r := range visited {
			fmt.Printf("%v\n", visited[r])
		}
	*/

	riskPath := riskSofar + (*cave)[row][col]
	if riskPath > *lowestRisk {
		// this path has already greater risk, even it is not ended yet
		done <- true

		return
	}

	if row == row2 && col == col2 {

		//fmt.Printf("Found a path with risk %v, currently lowRisk %v\n", riskPath, *lowestRisk)
		if riskPath < lowRisk {
			*lowestRisk = riskPath
		}
		/*
			fmt.Println("Level", level, "Visited So Far!")
			for r := range visited {
				fmt.Printf("%v\n", visited[r])
			}
		*/

		done <- true
		return
	}

	// wanted to generalize the out of bound
	// but it becomes one true-end condition, meaning when step out of bound, the whole function could terminate?
	// better way is not to step out of bounds;
	// key is not to bear a value for the invalid location

	// calculate two risks: going right or down
	dones := []chan bool{}

	/*
		// some tweak - useless probably
		if col < len((*cave)[0])-1 && row < len(*cave)-1 && (*cave)[row][col+1] < (*cave)[row+1][col] {
			if col < len((*cave)[0])-1 && visited[row][col+1] == 0 {
				doneOne := make(chan bool)
				dones = append(dones, doneOne)
				go riskFromPointParallel(row, col+1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)

			}

			if row < len(*cave)-1 && visited[row+1][col] == 0 {
				doneOne := make(chan bool)
				dones = append(dones, doneOne)
				go riskFromPointParallel(row+1, col, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)
			}
		} else {

			if row < len(*cave)-1 && visited[row+1][col] == 0 {
				doneOne := make(chan bool)
				dones = append(dones, doneOne)
				go riskFromPointParallel(row+1, col, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)
			}

			if col < len((*cave)[0])-1 && visited[row][col+1] == 0 {
				doneOne := make(chan bool)
				dones = append(dones, doneOne)
				go riskFromPointParallel(row, col+1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)
			}

		}
	*/

	if col < col2 && visited[row][col+1] == 0 {
		doneOne := make(chan bool)
		dones = append(dones, doneOne)
		go riskFromPointParallel(row, col+1, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)

	}

	if row < row2 && visited[row+1][col] == 0 {
		doneOne := make(chan bool)
		dones = append(dones, doneOne)
		go riskFromPointParallel(row+1, col, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)
	}

	if col > col1 && visited[row][col-1] == 0 {
		doneOne := make(chan bool)
		dones = append(dones, doneOne)
		go riskFromPointParallel(row, col-1, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)
	}

	if row > row1 && visited[row-1][col] == 0 {
		doneOne := make(chan bool)
		dones = append(dones, doneOne)
		go riskFromPointParallel(row-1, col, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)
	}

	for d := range dones {
		<-dones[d]
	}

	done <- true

}

func riskFromPointHybrid(row, col, row1, col1, row2, col2 int, cave *[][]int, visited [][]int, level int, riskSofar int, done chan bool, lowestRisk *int) {

	runs++
	visited[row][col] = 1

	riskPath := riskSofar + (*cave)[row][col]

	if runs%3000 == 0 {
		fmt.Println("Run:", runs, "Level", level, "riskUptoHere:", riskPath, "row:", row, "col:", col)

		if runs%60000 == 0 {
			for r := range visited {
				fmt.Printf("%v\n", visited[r])
			}
		}
	}

	if riskPath > *lowestRisk {
		// this path has already greater risk, even it is not ended yet
		if done != nil {
			done <- true
		}

		return
	}

	if row == row2 && col == col2 {

		//fmt.Printf("Found a path with risk %v, currently lowRisk %v\n", riskPath, *lowestRisk)
		if riskPath < lowRisk {
			*lowestRisk = riskPath
		}
		/*
			fmt.Println("Level", level, "Visited So Far!")
			for r := range visited {
				fmt.Printf("%v\n", visited[r])
			}
		*/

		if done != nil {
			done <- true
		}
		return
	}

	// wanted to generalize the out of bound
	// but it becomes one true-end condition, meaning when step out of bound, the whole function could terminate?
	// better way is not to step out of bounds;
	// key is not to bear a value for the invalid location

	// calculate two risks: going right or down
	dones := []chan bool{}

	if level > 110 {
		if col < col2 && visited[row][col+1] == 0 {
			doneOne := make(chan bool)
			dones = append(dones, doneOne)
			go riskFromPointHybrid(row, col+1, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)

		}

		if row < row2 && visited[row+1][col] == 0 {
			doneOne := make(chan bool)
			dones = append(dones, doneOne)
			go riskFromPointHybrid(row+1, col, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)
		}

		if col > col1 && visited[row][col-1] == 0 {
			doneOne := make(chan bool)
			dones = append(dones, doneOne)
			go riskFromPointHybrid(row, col-1, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)
		}

		if row > row1 && visited[row-1][col] == 0 {
			doneOne := make(chan bool)
			dones = append(dones, doneOne)
			go riskFromPointHybrid(row-1, col, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, doneOne, lowestRisk)
		}

		for d := range dones {
			<-dones[d]
		}

		if done != nil {
			done <- true
		}

	} else {
		if col < col2 && visited[row][col+1] == 0 {
			riskFromPointHybrid(row, col+1, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, nil, lowestRisk)

		}

		if row < row2 && visited[row+1][col] == 0 {
			riskFromPointHybrid(row+1, col, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, nil, lowestRisk)
		}

		if col > col1 && visited[row][col-1] == 0 {
			riskFromPointHybrid(row, col-1, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, nil, lowestRisk)
		}

		if row > row1 && visited[row-1][col] == 0 {
			riskFromPointHybrid(row-1, col, row1, col1, row2, col2, cave, copyVisited(visited), level+1, riskPath, nil, lowestRisk)
		}

	}

}

func riskFromPointPriorityQueue(row, col int, cave [][]int) {
	var visited [][]int
	deltas := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for i := 0; i < len(cave); i++ {
		visitedRow := make([]int, len(cave[0]))
		for j := 0; j < len(cave[0]); j++ {
			visitedRow[j] = 0
		}
		visited = append(visited, visitedRow)
	}

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		value:    Loc{0, 0},
		priority: 0,
		index:    0,
	}

	heap.Init(&pq)

	for pq.Len() != 0 {
		item := heap.Pop(&pq)
		node := item.(*Item)

		if visited[node.value.row][node.value.col] == 1 {
			continue
		}

		fmt.Printf("item is %v\n", *node)

		if node.value.row == len(cave)-1 && node.value.col == len(cave[0])-1 {
			fmt.Println(node.priority)
			//for r := range visited {
			//	fmt.Printf("%v\n", visited[r])
			//}
			break
		}

		visited[node.value.row][node.value.col] = 1

		for d := range deltas {
			nextRow := node.value.row + deltas[d][0]
			nextCol := node.value.col + deltas[d][1]

			// watch out for one node being  added multiple times by its neighbors
			// and it can create a loop...
			// control at this end verse control at the top
			// or both
			// thinking DFS/BFS/Shortest-Path... subtly different
			if nextRow >= 0 && nextRow < len(cave) && nextCol >= 0 && nextCol < len(cave[0]) {
				nextNode := Item{
					value:    Loc{nextRow, nextCol},
					priority: node.priority + cave[nextRow][nextCol],
				}
				heap.Push(&pq, &nextNode)
			}

		}

	}
}

func cave5by5(cave [][]int) [][]int {
	origWidth := len(cave[0])
	origLength := len(cave)

	// 1*5: col*5
	for r := range cave {

		for i := 1; i < 5; i++ {
			extend := []int{}
			for c := len(cave[r]) - origWidth; c < len(cave[r]); c++ {
				base := 0
				if cave[r][c] == 9 {
					base = 0
				} else {
					base = cave[r][c]
				}

				extend = append(extend, base+1)
			}

			cave[r] = append(cave[r], extend...)
		}
	}

	// 5*1: row*5
	for i := 1; i < 5; i++ {
		currentLen := len(cave)
		for r := currentLen - origLength; r < currentLen; r++ {
			newRow := []int{}
			for c := range cave[r] {
				base := 0
				if cave[r][c] == 9 {
					base = 0
				} else {
					base = cave[r][c]
				}

				newRow = append(newRow, base+1)
			}

			cave = append(cave, newRow)
		}
	}

	return cave
}

func chitonDriverPQ() {
	cave := readCaveStdin()

	// make it five*five
	//newCave := cave
	newCave := cave5by5(cave)

	riskFromPointPriorityQueue(0, 0, newCave)

}

func debugChiton() {
	cave := [][]int{
		{1, 1, 6, 3, 7, 5, 1, 7, 4, 2},
		{1, 3, 8, 1, 3, 7, 3, 6, 7, 2},
		{2, 1, 3, 6, 5, 1, 1, 3, 2, 8},
		{3, 6, 9, 4, 9, 3, 1, 5, 6, 9},
		{7, 4, 6, 3, 4, 1, 7, 1, 1, 1},
		{1, 3, 1, 9, 1, 2, 8, 1, 3, 7},
		{1, 3, 5, 9, 9, 1, 2, 4, 2, 1},
		{3, 1, 2, 5, 4, 2, 1, 6, 3, 9},
		{1, 2, 9, 3, 1, 3, 8, 5, 2, 1},
		{2, 3, 1, 1, 9, 4, 4, 5, 8, 1},
	}

	newCave := cave5by5(cave)
	for r := range newCave {
		fmt.Printf("%v\n", newCave[r])
	}

	riskFromPointPriorityQueue(0, 0, newCave)

	//fmt.Println(lowestRisk(&cave, visited))
}
