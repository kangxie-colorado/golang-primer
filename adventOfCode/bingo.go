package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readStdin() ([]int, [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	chosenNums := []int{}
	boards := [][]int{}
	for scanner.Scan() {
		if len(chosenNums) == 0 {
			// assume the first line is the chosen numers
			numStrs := strings.Split(scanner.Text(), ",")
			for _, ns := range numStrs {
				n, _ := strconv.Atoi(ns)
				chosenNums = append(chosenNums, n)
			}
		}

		if scanner.Text() == "" {
			// scan next 5 lines for a board
			// yeah, here I just assume the format
			// that is not important here
			board := []int{}
			for row := 0; row < 5; row++ {
				scanner.Scan()
				numStrs := strings.Split(strings.ReplaceAll(strings.Trim(scanner.Text(), " "), "  ", " "), " ")
				for _, ns := range numStrs {
					n, _ := strconv.Atoi(ns)
					board = append(board, n)
				}
			}
			boards = append(boards, board)

		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return chosenNums, boards
}

func printBingo(chosenNums []int, boards [][]int) {
	for _, cn := range chosenNums {
		fmt.Print(cn, ",")
	}

	for _, board := range boards {
		fmt.Println()

		for i, n := range board {
			if i%5 == 0 {
				fmt.Println()
			}

			fmt.Print(n, " ")
		}
	}
}

func numIdxInSlice(num int, board []int) int {
	res := -1
	for i, n := range board {
		if num == n {
			res = i
		}
	}

	return res
}

func bingoComplete(auxBoard [5][5]int) bool {
	for row := 0; row < 5; row++ {
		sumOfRow := 0
		for col := 0; col < 5; col++ {
			sumOfRow += auxBoard[row][col]
		}

		if sumOfRow == 5 {
			return true
		}
	}

	for col := 0; col < 5; col++ {
		sumOfCol := 0
		for row := 0; row < 5; row++ {
			sumOfCol += auxBoard[row][col]
		}

		if sumOfCol == 5 {
			return true
		}
	}

	return false

}

// return
// winning number: the last number from chosen number to complete the bingo
// the board as a flat list converted to 5*5 board with some verification?
func bingoPlayWinner(chosenNums []int, boards [][]int) (int, []int, [5][5]int) {
	auxBoards := [][5][5]int{}
	for i := 0; i < len(boards); i++ {
		// making the auxililary boards for bookkeeping
		auxBoard := [5][5]int{}
		auxBoards = append(auxBoards, auxBoard)
	}

	for _, cn := range chosenNums {
		for bi, board := range boards {
			matchIdx := numIdxInSlice(cn, board)
			if matchIdx != -1 {
				// some number matches, mark it in auxBoard
				auxBoards[bi][matchIdx/5][matchIdx%5] = 1
				// check if this board completes bingo, if yes, terminate now
				if bingoComplete(auxBoards[bi]) {
					return cn, board, auxBoards[bi]
				}
			}
		}
	}

	return -1, nil, [5][5]int{}
}

func playToTheLast(chosenNums []int, boards [][]int) (int, []int, [5][5]int) {
	auxBoards := [][5][5]int{}
	for i := 0; i < len(boards); i++ {
		// making the auxililary boards for bookkeeping
		auxBoard := [5][5]int{}
		auxBoards = append(auxBoards, auxBoard)
	}

	completedBoards := make(map[int]bool)
	lastMatchedBoardIdx := -1
	lastNum := -1
	for _, cn := range chosenNums {
		if len(completedBoards) == len(boards) {
			break
		}

		for bi, board := range boards {
			if _, ok := completedBoards[bi]; ok {
				continue
			}

			matchIdx := numIdxInSlice(cn, board)
			if matchIdx != -1 {
				// some number matches, mark it in auxBoard
				auxBoards[bi][matchIdx/5][matchIdx%5] = 1
				// check if this board completes bingo, if yes, terminate now
				if bingoComplete(auxBoards[bi]) {
					completedBoards[bi] = true
					lastMatchedBoardIdx = bi
					lastNum = cn
				}
			}
		}
	}

	if lastMatchedBoardIdx != -1 {
		return lastNum, boards[lastMatchedBoardIdx], auxBoards[lastMatchedBoardIdx]
	}

	return -1, nil, [5][5]int{}
}

func winnerScore(board []int, auxBoard [5][5]int) int {
	score := 0
	for row := range auxBoard {
		for col := range auxBoard[row] {
			if auxBoard[row][col] == 0 {
				score += board[row*5+col]
			}
		}
	}

	return score
}

func printWinner(board []int, auxBoard [5][5]int) {
	for row := range auxBoard {
		str1 := ""
		str2 := ""
		for col := range auxBoard[row] {
			str1 += fmt.Sprintf("%v ", auxBoard[row][col])
			str2 += fmt.Sprintf("%2d ", board[row*5+col])
		}

		fmt.Printf("%s     %s\n", str1, str2)

	}
}
