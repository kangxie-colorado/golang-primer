package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getAreasStdin(areas *[][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		nums := []int{}
		for _, r := range scanner.Text() {
			nums = append(nums, int(r-'0'))
		}

		*areas = append(*areas, nums)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}

func lowPtsRisks(areas *[][]int) int {
	risks := 0
	for r := range *areas {
		for c := range (*areas)[r] {
			if r > 0 {
				// look up
				if (*areas)[r][c] >= (*areas)[r-1][c] {
					// not a low point
					continue
				}
			}

			if c > 0 {
				// look left
				if (*areas)[r][c] >= (*areas)[r][c-1] {
					// not a low point
					continue
				}
			}

			if r < len(*areas)-1 {
				// look down
				if (*areas)[r][c] >= (*areas)[r+1][c] {
					// not a low point
					continue
				}
			}

			if c < len((*areas)[r])-1 {
				if (*areas)[r][c] >= (*areas)[r][c+1] {
					// not a low point
					continue
				}
			}

			// this is a low point
			risks += (*areas)[r][c] + 1
		}
	}

	return risks
}

func findBasins(areas *[][]int) []int {
	basins := []int{}
	seen := map[Loc]bool{}

	for r := range *areas {
		for c := range (*areas)[r] {
			if (*areas)[r][c] == 9 {
				continue
			}

			if _, ok := seen[Loc{r, c}]; ok {
				continue
			}

			toProcessQueue := []Loc{}
			size := 0

			// enqueue first item
			toProcessQueue = append(toProcessQueue, Loc{row: r, col: c})

			for len(toProcessQueue) != 0 {
				// dequeue: focus on the left end
				item := toProcessQueue[0]
				toProcessQueue = toProcessQueue[1:]

				// if item is seen/processed just dequeue and not processing it more
				if _, ok := seen[Loc{item.row, item.col}]; ok {
					continue
				}

				size += 1
				seen[item] = true

				for _, n := range neighborsWithoutiags(item.row, item.col, len(*areas), len((*areas)[r])) {
					if (*areas)[n.row][n.col] == 9 {
						continue
					}

					// if you do enqueue control here: like if the neighbor is seen, don't enqueue it
					// you risk on node being neighbor of two or more other nodes
					// and when the other nodes are being processed, this one node could be added into queue more that one time but without
					// looking it twice about if it has been processed at the dequeue time, which would end badly
					// so just enqueue every neighbor but only to discard them if it is seen on the dequeue part ^

					toProcessQueue = append(toProcessQueue, Loc{n.row, n.col})
				}

			}
			basins = append(basins, size)
		}
	}

	sort.Ints(basins)
	return basins
}

func smokeBasinDriver() {
	areas := [][]int{}
	getAreasStdin(&areas)
	// print the basic to see for debug
	/*
		for _, r := range areas {
			fmt.Printf("%v\n", r)
		}
	*/

	fmt.Println(lowPtsRisks(&areas))

	fmt.Println(findBasins(&areas))
}

func testBasin(area [][]int, want []int) {
	got := findBasins(&area)
	if !compareTwoIntSlice(got, want) {
		fmt.Printf("Wanted: %v, Got: %v\n", want, got)
		panic("Something Wrong!")
	}
}

func testFindBasins() {

	area := [][]int{
		{2, 1},
		{3, 9},
	}

	testBasin(area, []int{3})

	area = [][]int{
		{2, 1, 9, 9, 9},
		{3, 9, 8, 7, 8},
	}
	testBasin(area, []int{3, 3})

	area = [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1},
		{3, 9, 8, 7, 8, 9, 4, 9, 2},
	}
	testBasin(area, []int{3, 6, 3})

	area = [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
	}
	testBasin(area, []int{3, 9, 8, 1})

	fmt.Println("Tests Passed!")
}
