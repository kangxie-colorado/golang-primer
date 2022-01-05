package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func syntaxScoreStdin() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	sum := 0
	for _, line := range lines {
		sum += processLine(line)
	}

	fmt.Println(sum)

	completionScores := []int{}
	for _, line := range lines {
		score := completeLineScore(line)
		if score != -1 {
			completionScores = append(completionScores, score)

		}
	}

	// mid-score
	sort.Ints(completionScores)
	fmt.Println(completionScores[len(completionScores)/2])

}

var rightLeftMap map[rune]rune = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var rightScoreCorruptionMap map[rune]int = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var leftScoreCompletionMap map[rune]int = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func processLine(line string) int {
	firstCorrupted := '0'

	stack := []rune{}

	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack = append(stack, r)
		case ')', ']', '}', '>':
			if len(stack) != 0 {
				// de-stack
				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				leftWanted, _ := rightLeftMap[r]
				if left != leftWanted {
					// this is corrupted
					firstCorrupted = r
					break
				}
			}
		}
	}

	if score, ok := rightScoreCorruptionMap[firstCorrupted]; ok {
		return score
	}
	return 0
}

func completeLineScore(line string) int {

	stack := []rune{}

	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack = append(stack, r)
		case ')', ']', '}', '>':
			if len(stack) != 0 {
				// de-stack
				left := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				leftWanted, _ := rightLeftMap[r]
				if left != leftWanted {
					// this is corrupted, just return -1 and the caller should discard it
					return -1
				}
			}
		}
	}

	// by here, if the stack is not empty it shall be an incomplete line, lets begin to score
	score := 0
	for len(stack) != 0 {
		// de-stack
		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		point, _ := leftScoreCompletionMap[left]
		score = score*5 + point
	}

	return score
}

func syntaxScoreDriver() {
	syntaxScoreStdin()
}
