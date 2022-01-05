package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func _boilerPlateCodeStdin() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}

// only to ignore the error and make the code a little better look
func strToNum(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic("Unable to convert the string to number")
	}

	return num
}

type Loc struct {
	row, col int
}

func neighborsWithoutiags(r, c, lenR, lenC int) []Loc {
	neighbors := []Loc{}
	if r > 0 {
		neighbors = append(neighbors, Loc{r - 1, c})
	}

	if c > 0 {
		neighbors = append(neighbors, Loc{r, c - 1})
	}

	if r < lenR-1 {
		neighbors = append(neighbors, Loc{r + 1, c})
	}

	if c < lenC-1 {
		neighbors = append(neighbors, Loc{r, c + 1})

	}

	return neighbors
}

func neighborsWithDiags(r, c, lenR, lenC int) []Loc {
	return findNeighbors(r, c, lenR, lenC)
}

func findNeighbors(r, c, lenR, lenC int) []Loc {
	neighbors := []Loc{}

	for x := -1; x < 2; x += 1 {
		for y := -1; y < 2; y += 1 {
			row := r + x
			col := c + y
			if row >= 0 && row < lenR && col >= 0 && col < lenC {
				neighbors = append(neighbors, Loc{row, col})
			}
		}
	}

	return neighbors
}

func compareTwoIntSlice(intS1, intS2 []int) bool {
	if len(intS1) != len(intS2) {
		return false
	}

	for i := range intS1 {
		if intS1[i] != intS2[i] {
			return false
		}
	}

	return true
}
