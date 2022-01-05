package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Line struct {
	x1, y1 int
	x2, y2 int
}

func (l Line) repr() string {
	return fmt.Sprintf("Line{%v,%v -> %v,%v}", l.x1, l.y1, l.x2, l.y2)
}

func getLinesStdin() []Line {
	lines := []Line{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		p1Dims := strings.Split(points[0], ",")
		p2Dims := strings.Split(points[1], ",")

		lines = append(lines, Line{
			strToNum(p1Dims[0]),
			strToNum(p1Dims[1]),
			strToNum(p2Dims[0]),
			strToNum(p2Dims[1]),
		})

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return lines
}

func printLines(lines []Line) {
	for _, l := range lines {
		fmt.Println(l.x1, l.y1, "-->", l.x2, l.y2)
	}
}

const Size int = 1000

func markHAndVLines(line Line, theField *[Size][Size]int) {
	if line.x1 == line.x2 || line.y1 == line.y2 {
		startX, startY := -1, -1
		endX, endY := -1, -1
		if line.x1 < line.x2 || line.y1 < line.y2 {
			startX, startY = line.x1, line.y1
			endX, endY = line.x2, line.y2
		} else {
			startX, startY = line.x2, line.y2
			endX, endY = line.x1, line.y1
		}

		// this is covering a rectangle area
		// but because x1==x2 or y1==y2, so this is effectively a line
		// so its a hack to simplify the code
		if endX < startX || endY < startY {
			fmt.Printf("%v\n", line.repr())
			fmt.Printf("Starting: %v,%v, Ending: %v,%v\n", startX, startY, endX, endY)
			panic("Not Right!")
		}

		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				theField[x][y] += 1
			}
		}
	}
}

func markDiagLines(line Line, theField *[Size][Size]int) {
	if math.Abs(float64(line.x1-line.x2)) == math.Abs(float64(line.y1-line.y2)) {
		// this is 45 degree diagonal
		stepX := 1
		if line.x1 > line.x2 {
			stepX = -1
		}

		stepY := 1
		if line.y1 > line.y2 {
			stepY = -1
		}

		endReached := false
		x := line.x1
		y := line.y1
		for !endReached {
			if x == line.x2 {
				// reached the end point
				endReached = true
			}

			theField[x][y] += 1

			x += stepX
			y += stepY
		}

	}
}

// only mark horizontal and vertical lines
func markVentField(lines []Line) [Size][Size]int {
	// arbitarily 1000*1000, seems big enough for all the input
	theField := [Size][Size]int{}

	for _, line := range lines {
		markHAndVLines(line, &theField)
	}

	return theField
}

func debugMarkVentField() {
	l1 := Line{0, 9, 5, 9}
	l2 := Line{0, 9, 2, 9}

	markVentField([]Line{l1, l2})
}

// mark horzontal/vertical and diagonal lines
func markVentFieldWithDiag(lines []Line) [Size][Size]int {
	// arbitarily 1000*1000, seems big enough for all the input
	theField := [Size][Size]int{}

	for _, line := range lines {
		markHAndVLines(line, &theField)
		markDiagLines(line, &theField)
	}

	return theField
}

func printField(theField [Size][Size]int) {
	for row := range theField {
		for col := range theField[row] {
			fmt.Print(theField[row][col], " ")
		}
		fmt.Println()
	}
}

func countOverlaps(theField [Size][Size]int, atleast int) int {
	howMany := 0
	for row := range theField {
		for col := range theField[row] {
			if theField[row][col] >= atleast {
				howMany += 1
			}
		}
	}

	return howMany
}
