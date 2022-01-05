package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sonarSweepWindow3(data []int) int {
	res := 0
	sumOf3 := 0
	for i, d := range data {
		if i < 3 {
			sumOf3 += d
			continue
		}

		newSumOf3 := sumOf3 - data[i-3] + d

		if newSumOf3 > sumOf3 {
			res += 1
		}
	}

	return res
}

func sonarSweepFromStdinWind3() int {
	scanner := bufio.NewScanner(os.Stdin)

	data := []int{}
	for scanner.Scan() {
		last, _ := strconv.Atoi(scanner.Text())
		data = append(data, last)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return sonarSweepWindow3(data)
}

func sonarSweep(data []int) int {
	res := 0
	for i, d := range data {
		if i == 0 {
			continue
		}

		if d > data[i-1] {
			res += 1
		}
	}

	return res
}

func sonarSweepFromStdin() int {
	scanner := bufio.NewScanner(os.Stdin)
	last := -1
	increases := 0
	for scanner.Scan() {
		if last == -1 {
			last, _ = strconv.Atoi(scanner.Text())
			continue
		}

		thisLevel, _ := strconv.Atoi(scanner.Text())
		fmt.Println(last, thisLevel)
		if thisLevel > last {
			increases += 1
		}
		last = thisLevel
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return increases
}
