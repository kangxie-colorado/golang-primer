package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func sumToN(n int) int {
	return n * (n + 1) / 2
}

func calcFuel(crabs []int, align int) int {
	fuel := 0
	for _, c := range crabs {
		fuel += sumToN(int(math.Abs(float64(c - align))))
	}

	return fuel
}

func leastFuel(crabs []int) int {
	// sort the crabs and the align point must be between min and max
	sort.Ints(crabs)
	fuel := math.MaxInt32
	for pt := crabs[0]; pt <= crabs[len(crabs)-1]; pt++ {
		fuel = int(math.Min(float64(fuel), float64(calcFuel(crabs, pt))))
	}

	return fuel
}

func leastFuelStdin() int {
	scanner := bufio.NewScanner(os.Stdin)
	crabs := []int{}
	for scanner.Scan() {
		crabStr := strings.Split(scanner.Text(), ",")
		for _, cs := range crabStr {
			crabs = append(crabs, strToNum(cs))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return leastFuel(crabs)
}
