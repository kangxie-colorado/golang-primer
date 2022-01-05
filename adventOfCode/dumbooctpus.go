package main

import (
	"bufio"
	"fmt"
	"os"
)

func readOctuposStdin() [][]int {
	scanner := bufio.NewScanner(os.Stdin)

	matrix := [][]int{}
	for scanner.Scan() {
		row := []int{}
		line := scanner.Text()
		for _, r := range line {
			row = append(row, int(r-'0'))
		}

		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return matrix
}

func oneStep(matrix *[][]int) int {
	flashes := 0
	shouldFlash := 0

	// first add 1 for all of them
	for r := range *matrix {
		for c := range (*matrix)[r] {
			(*matrix)[r][c]++
			if (*matrix)[r][c] > 9 {
				shouldFlash++
			}
		}
	}

	// now check if any ocupos > 9, whic could trigger chain reaction
	// until there is no more > 9
	// how to check it? scan thru 100 numbers seems ok
	// but would there be something easier?
	// maybe bookkeeping the number>9 and adjust it dynamically

	// so this is scan until no more > 9
	// of coz you should optmize this with BFS...
	// but only 100 numbers and lets just do it the hardest way
	for shouldFlash > 0 {
		for r := range *matrix {
			for c := range (*matrix)[r] {
				if (*matrix)[r][c] > 9 {
					// flash
					(*matrix)[r][c] = 0
					shouldFlash--
					flashes++

					nbs := neighborsWithDiags(r, c, len(*matrix), len((*matrix)[r]))
					for _, nb := range nbs {
						if (*matrix)[nb.row][nb.col] != 0 {
							// if == 0, the it just flahsed should not get increase in this step again
							(*matrix)[nb.row][nb.col]++
							if (*matrix)[nb.row][nb.col] == 10 {
								// if > 10, then it was at least 10 before this increase,
								// should not re-count it
								shouldFlash++
							}
						}
					}
				}
			}
		}
	}

	return flashes
}

func compareTwoIntMatrixes(m1, m2 *[][]int) bool {
	if len(*m1) != len(*m2) {
		return false
	}

	for r := range *m1 {
		if !compareTwoIntSlice((*m1)[r], (*m2)[r]) {
			return false
		}
	}

	return true
}

func testOneStep() {
	matrix := [][]int{
		{1, 1, 1, 1, 1},
		{1, 9, 9, 9, 1},
		{1, 9, 1, 9, 1},
		{1, 9, 9, 9, 1},
		{1, 1, 1, 1, 1},
	}

	flashes := oneStep(&matrix)
	wanted := [][]int{
		{3, 4, 5, 4, 3},
		{4, 0, 0, 0, 4},
		{5, 0, 0, 0, 5},
		{4, 0, 0, 0, 4},
		{3, 4, 5, 4, 3},
	}

	if !compareTwoIntMatrixes(&wanted, &matrix) || flashes != 9 {
		fmt.Printf("Not Right!")
	}

	flashes = oneStep(&matrix)
	wanted = [][]int{
		{4, 5, 6, 5, 4},
		{5, 1, 1, 1, 5},
		{6, 1, 1, 1, 6},
		{5, 1, 1, 1, 5},
		{4, 5, 6, 5, 4},
	}

	if !compareTwoIntMatrixes(&wanted, &matrix) || flashes != 0 {
		fmt.Printf("Not Right!")
	}

	matrix = [][]int{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}

	flashes = oneStep(&matrix)
	wanted = [][]int{
		{6, 5, 9, 4, 2, 5, 4, 3, 3, 4},
		{3, 8, 5, 6, 9, 6, 5, 8, 2, 2},
		{6, 3, 7, 5, 6, 6, 7, 2, 8, 4},
		{7, 2, 5, 2, 4, 4, 7, 2, 5, 7},
		{7, 4, 6, 8, 4, 9, 6, 5, 8, 9},
		{5, 2, 7, 8, 6, 3, 5, 7, 5, 6},
		{3, 2, 8, 7, 9, 5, 2, 8, 3, 2},
		{7, 9, 9, 3, 9, 9, 2, 2, 4, 5},
		{5, 9, 5, 7, 9, 5, 9, 6, 6, 5},
		{6, 3, 9, 4, 8, 6, 2, 6, 3, 7},
	}
	if !compareTwoIntMatrixes(&wanted, &matrix) || flashes != 0 {
		fmt.Printf("Step 1: Not Right!")
	}

	flashes = oneStep(&matrix)
	wanted = [][]int{
		{8, 8, 0, 7, 4, 7, 6, 5, 5, 5},
		{5, 0, 8, 9, 0, 8, 7, 0, 5, 4},
		{8, 5, 9, 7, 8, 8, 9, 6, 0, 8},
		{8, 4, 8, 5, 7, 6, 9, 6, 0, 0},
		{8, 7, 0, 0, 9, 0, 8, 8, 0, 0},
		{6, 6, 0, 0, 0, 8, 8, 9, 8, 9},
		{6, 8, 0, 0, 0, 0, 5, 9, 4, 3},
		{0, 0, 0, 0, 0, 0, 7, 4, 5, 6},
		{9, 0, 0, 0, 0, 0, 0, 8, 7, 6},
		{8, 7, 0, 0, 0, 0, 6, 8, 4, 8},
	}
	if !compareTwoIntMatrixes(&wanted, &matrix) || flashes != 35 {
		fmt.Printf("Step 2: Not Right!")
	}

	flashes = oneStep(&matrix)
	wanted = [][]int{
		{0, 0, 5, 0, 9, 0, 0, 8, 6, 6},
		{8, 5, 0, 0, 8, 0, 0, 5, 7, 5},
		{9, 9, 0, 0, 0, 0, 0, 0, 3, 9},
		{9, 7, 0, 0, 0, 0, 0, 0, 4, 1},
		{9, 9, 3, 5, 0, 8, 0, 0, 6, 3},
		{7, 7, 1, 2, 3, 0, 0, 0, 0, 0},
		{7, 9, 1, 1, 2, 5, 0, 0, 0, 9},
		{2, 2, 1, 1, 1, 3, 0, 0, 0, 0},
		{0, 4, 2, 1, 1, 2, 5, 0, 0, 0},
		{0, 0, 2, 1, 1, 1, 9, 0, 0, 0},
	}
	if !compareTwoIntMatrixes(&wanted, &matrix) || flashes != 45 {
		fmt.Printf("Step 3: Not Right!")
	}

	oneStep(&matrix)
	oneStep(&matrix)
	flashes = oneStep(&matrix)
	wanted = [][]int{
		{5, 5, 9, 5, 2, 5, 5, 1, 1, 1},
		{3, 1, 5, 5, 2, 5, 5, 2, 2, 2},
		{3, 3, 6, 4, 4, 4, 4, 6, 0, 5},
		{2, 2, 6, 3, 4, 4, 4, 4, 9, 6},
		{2, 2, 9, 8, 4, 1, 4, 3, 9, 6},
		{2, 2, 7, 5, 7, 4, 4, 3, 4, 4},
		{2, 2, 6, 4, 5, 8, 3, 3, 4, 2},
		{7, 7, 5, 4, 4, 6, 3, 3, 4, 4},
		{3, 7, 5, 4, 4, 6, 9, 4, 3, 3},
		{3, 3, 5, 4, 4, 5, 2, 4, 3, 3},
	}
	if !compareTwoIntMatrixes(&wanted, &matrix) || flashes != 1 {
		fmt.Printf("Step 6: Not Right!")
	}

	fmt.Println("Single Step Tests Passed!")

	matrix = [][]int{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}

	flashes = after100Steps(&matrix)
	wanted = [][]int{
		{0, 3, 9, 7, 6, 6, 6, 8, 6, 6},
		{0, 7, 4, 9, 7, 6, 6, 9, 1, 8},
		{0, 0, 5, 3, 9, 7, 6, 9, 3, 3},
		{0, 0, 0, 4, 2, 9, 7, 8, 2, 2},
		{0, 0, 0, 4, 2, 2, 9, 8, 9, 2},
		{0, 0, 5, 3, 2, 2, 2, 8, 7, 7},
		{0, 5, 3, 2, 2, 2, 2, 9, 6, 6},
		{9, 3, 2, 2, 2, 2, 8, 9, 6, 6},
		{7, 9, 2, 2, 2, 8, 6, 8, 6, 6},
		{6, 7, 8, 9, 9, 9, 8, 7, 6, 6},
	}
	if !compareTwoIntMatrixes(&wanted, &matrix) || flashes != 1656 {
		fmt.Printf("Step 6: Not Right!")
	}
	fmt.Println("100 Steps Test Passed!")
}

func after100Steps(matrix *[][]int) int {

	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += oneStep(matrix)
	}

	return flashes

}

func syncFlashes(matrix *[][]int) int {
	synced := false
	step := 0

	for !synced {
		step++
		if oneStep(matrix) == len(*matrix)*len((*matrix)[0]) {
			synced = true
		}
	}

	return step
}

func dumboOctuposDriver() {

	start := [][]int{
		{4, 1, 1, 2, 2, 5, 6, 3, 7, 2},
		{3, 1, 4, 3, 2, 5, 3, 7, 1, 2},
		{4, 5, 1, 6, 8, 4, 8, 6, 3, 1},
		{3, 7, 8, 3, 4, 7, 7, 1, 3, 7},
		{3, 7, 4, 6, 7, 2, 3, 5, 8, 2},
		{5, 8, 6, 1, 3, 5, 8, 8, 8, 4},
		{4, 8, 4, 3, 3, 5, 1, 7, 7, 4},
		{2, 3, 1, 6, 4, 4, 7, 6, 2, 1},
		{6, 6, 4, 3, 8, 1, 7, 7, 4, 5},
		{6, 3, 6, 6, 8, 1, 5, 8, 6, 8},
	}

	// flashes := after100Steps(&start)
	// fmt.Println(flashes)

	firstSyncStep := syncFlashes(&start)
	fmt.Println(firstSyncStep)
}
