package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fishNumAfterDays(fishes []int, days int) int {

	for d := 0; d < days; d++ {
		fishNum := len(fishes)
		for f := 0; f < fishNum; f++ {
			if fishes[f] == 0 {
				fishes[f] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[f]--
			}
		}
	}

	return len(fishes)
}

func fishNumAfterDaysStdin(days int) int {
	fishes := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fishStrs := strings.Split(scanner.Text(), ",")
		for _, fs := range fishStrs {
			fishes = append(fishes, strToNum(fs))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return fishNumAfterDays(fishes, days)
}

const BatchSize int = 1_000_000

func calculateThisSwarm(swarm *[]int, newSwarmChan chan *[]int) {
	fishesNum := len(*swarm)
	for f := 0; f < fishesNum; f++ {
		if (*swarm)[f] == 0 {
			(*swarm)[f] = 6
			*swarm = append(*swarm, 8)
		} else {
			(*swarm)[f]--
		}
	}

	var newSwarm []int = nil
	if len(*swarm) > BatchSize {
		newSwarm = (*swarm)[BatchSize:]
		*swarm = (*swarm)[:BatchSize]
	}

	newSwarmChan <- &newSwarm
}

// okay, so even compute in parallel will not solve the huge memory pressure
// it would needs more than 100Gb memory
// so the solution is actually using a map
// fishByDay[0]: some-number
// fishByDay[8]: some-number
// and it would be pretty easy
func fishNumAfterDaysParallel(fishes []int, days int) int64 {
	swarms := [][]int{}
	swarms = append(swarms, fishes)

	for d := 0; d < days; d++ {
		swarmNum := len(swarms)
		newSwarms := [][]int{}
		newSwarmChann := make(chan *[]int, swarmNum)
		for s := 0; s < swarmNum; s++ {
			go calculateThisSwarm(&swarms[s], newSwarmChann)
		}

		for s := 0; s < swarmNum; s++ {
			newSwarm := <-newSwarmChann
			if newSwarm != nil {
				newSwarms = append(newSwarms, *newSwarm)
			}
		}

		// combine newSwarms
		toEnterSwarm := []int{}
		for ns := 0; ns < len(newSwarms); ns++ {
			if len(toEnterSwarm)+len(newSwarms[ns]) <= BatchSize {
				toEnterSwarm = append(toEnterSwarm, newSwarms[ns]...)
			} else {
				toEnterSwarm = append(toEnterSwarm, newSwarms[ns][:BatchSize-len(toEnterSwarm)]...)
				swarms = append(swarms, toEnterSwarm)
				toEnterSwarm = newSwarms[ns][BatchSize-len(toEnterSwarm):]
			}

		}
		swarms = append(swarms, toEnterSwarm)

	}

	var total int64 = 0
	for _, sw := range swarms {
		total += int64(len(sw))
	}

	return total
}

func fishNumAfterDaysMap(fishes map[int]int64, days int) int64 {
	for d := 0; d < days; d++ {
		copy := map[int]int64{}
		for k, v := range fishes {
			copy[k] = v
		}
		for fishday := range copy {

			if fishday == 0 {
				fishes[6] += copy[fishday]
				fishes[8] += copy[fishday]
			} else {
				fishes[fishday-1] += copy[fishday]
			}

			fishes[fishday] -= copy[fishday]

		}

	}

	var totalFish int64 = 0
	for _, num := range fishes {
		totalFish += num
	}

	return totalFish
}

func fishNumAfterDaysMapStdin(days int) int64 {
	fishes := map[int]int64{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fishStrs := strings.Split(scanner.Text(), ",")
		for _, fs := range fishStrs {
			fishes[strToNum(fs)]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Printf("%v", fishes)
	return fishNumAfterDaysMap(fishes, days)
}

func debugFishes() {
	fishes := map[int]int64{
		1: 1,
		2: 1,
		3: 2,
		4: 1,
	}

	fishNumAfterDaysMap(fishes, 256)
}
