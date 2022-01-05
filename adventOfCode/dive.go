package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func diveWithAimStdin() int {
	scanner := bufio.NewScanner(os.Stdin)

	posX := 0
	posY := 0
	aim := 0
	for scanner.Scan() {
		elems := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(elems[1])
		switch elems[0] {
		case "forward":
			posX += num
			posY += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		default:
			fmt.Println("Something is wrong!")
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return posX * posY
}

func diveDataProcessStdin() int {
	scanner := bufio.NewScanner(os.Stdin)

	forwards := []int{}
	downs := []int{}
	ups := []int{}
	for scanner.Scan() {
		elems := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(elems[1])
		switch elems[0] {
		case "forward":
			forwards = append(forwards, num)
		case "down":
			downs = append(downs, num)
		case "up":
			ups = append(ups, num)
		default:
			fmt.Println("Something is wrong!")
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	upSum := 0
	for _, u := range ups {
		upSum += u
	}

	downSum := 0
	for _, d := range downs {
		downSum += d
	}

	forwardSum := 0
	for _, f := range forwards {
		forwardSum += f
	}

	return forwardSum * (downSum - upSum)
}
