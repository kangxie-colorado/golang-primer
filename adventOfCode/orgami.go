package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
The first section is a list of dots on the transparent paper. 0,0 represents the top-left coordinate. The first value, x, increases to the right. The second value, y, increases downward.

<== so x is the column and y is the row...
**/

type Insturction struct {
	foldingDir  string
	foldingLine int
}

func readDotsStdin(dots *[]Loc, instuctions *[]Insturction, maxRow, maxCol *int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanDotsDone := false
	for scanner.Scan() {

		if scanner.Text() == "" {
			// read the blank line, the following with be folding instuctions
			scanDotsDone = true
			continue
		}

		if !scanDotsDone {
			parts := strings.Split(scanner.Text(), ",")
			row, col := strToNum(parts[1]), strToNum(parts[0])

			if row > *maxRow {
				*maxRow = row
			}

			if col > *maxCol {
				*maxCol = col
			}

			*dots = append(*dots, Loc{row, col})
		} else {
			// now read the instructions
			instuction := strings.Split(scanner.Text(), " ")[2]
			parts := strings.Split(instuction, "=")

			foldingDir := "row"
			if parts[0] == "x" {
				foldingDir = "col"
			}
			*instuctions = append(*instuctions, Insturction{foldingDir, strToNum(parts[1])})
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func getOrgamiField(dots *[]Loc, maxRow, maxCol int) *[][]string {
	rowLen := maxRow + 1
	colLen := maxCol + 1
	field := make([][]string, rowLen)
	for r := 0; r < rowLen; r++ {
		field[r] = make([]string, colLen)
	}

	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			field[r][c] = "."
		}
	}

	for _, d := range *dots {
		field[d.row][d.col] = "#"
	}

	return &field
}

func printOrgamiField(field *[][]string) {
	for r := 0; r < len(*field); r++ {
		fmt.Printf("%v\n", (*field)[r])
	}

	fmt.Println()
}

func printInstructions(instrs []Insturction) {
	for _, i := range instrs {
		fmt.Printf("Folding at %v:%v\n", i.foldingDir, i.foldingLine)
	}
}

func foldOneInstruction(field *[][]string, instuction Insturction) {
	switch instuction.foldingDir {
	case "row":
		foldByRow(field, instuction.foldingLine)
	case "col":
		foldByCol(field, instuction.foldingLine)
	default:
		panic("Something Not Right!")
	}
}

func combineTwoRows(field *[][]string, upperLine, lowerLine int) {
	for c := 0; c < len((*field)[upperLine]); c++ {
		if (*field)[upperLine][c] == "#" || (*field)[lowerLine][c] == "#" {
			(*field)[upperLine][c] = "#"
		}
	}
}

func combineTwoCols(field *[][]string, leftLine, rightLine int) {
	for r := 0; r < len(*field); r++ {
		if (*field)[r][leftLine] == "#" || (*field)[r][rightLine] == "#" {
			(*field)[r][leftLine] = "#"
		}
	}

}

func foldByRow(field *[][]string, foldingLine int) {
	for r := foldingLine; r < len(*field); r++ {
		offset := r - foldingLine
		if offset != 0 {
			combineTwoRows(field, foldingLine-offset, foldingLine+offset)
		}
	}

	*field = (*field)[:foldingLine]
}

func foldByCol(field *[][]string, foldingLine int) {

	for c := foldingLine; c < len((*field)[0]); c++ {
		offset := c - foldingLine
		if offset != 0 {
			combineTwoCols(field, foldingLine-offset, foldingLine+offset)
		}
	}

	for r := 0; r < len(*field); r++ {
		(*field)[r] = (*field)[r][:foldingLine]
	}
}

func countDots(field *[][]string) int {
	total := 0
	for r := range *field {
		for c := range (*field)[r] {
			if (*field)[r][c] == "#" {
				total += 1
			}
		}
	}

	return total
}

func foldPaper(field *[][]string, instuctions *[]Insturction) {
	for _, instr := range *instuctions {
		foldOneInstruction(field, instr)
		fmt.Println(countDots(field))

	}

	printOrgamiField(field)

}

func orgamiDriver() {
	//debugOrgami()
	//return

	dots := []Loc{}
	instructions := []Insturction{}
	maxRow, maxCol := 0, 0 // this is the max index, should be inclusive when iteration thru the field
	readDotsStdin(&dots, &instructions, &maxRow, &maxCol)

	field := getOrgamiField(&dots, maxRow, maxCol)

	printOrgamiField(field)
	printInstructions(instructions)

	//foldOneInstruction(field, instructions[0])

	foldPaper(field, &instructions)

}

func debugOrgami() {
	dots := []Loc{
		{10, 6},
		{14, 0},
		{10, 9},
	}

	instructions := []Insturction{
		{"row", 7},
		{"col", 5},
	}

	field := getOrgamiField(&dots, 14, 9)
	printOrgamiField(field)
	printInstructions(instructions)

	foldOneInstruction(field, instructions[0])

}
