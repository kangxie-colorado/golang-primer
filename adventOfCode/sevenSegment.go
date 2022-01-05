package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var DigitSigsMap map[int]int = map[int]int{
	0: 6, // to display 0, need 6 signals
	1: 2,
	2: 5,
	3: 5,
	4: 4,
	5: 5,
	6: 6,
	7: 3,
	8: 9,
	9: 6,
}

var UniqSigsDigitMap map[int]int = map[int]int{
	2: 1, // if 2 signals, it must be 1
	4: 4,
	3: 7,
	7: 8,
}

func unqiNumberCounts(outputs []string) int {
	res := 0
	for _, op := range outputs {
		if _, ok := UniqSigsDigitMap[len(op)]; ok {
			res += 1
		}
	}

	return res
}

func unqiNumberCountsStdin() int {
	scanner := bufio.NewScanner(os.Stdin)
	outputs := []string{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		outputs = append(outputs, strings.Split(strings.Trim(parts[1], " "), " ")...)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return unqiNumberCounts(outputs)
}

// rules:
// len()==2 -> 1
// len()==3 -> 7
// can be described as functions
var numStrMap map[int]string
var strNumMap map[string]int

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func ruleOne(str string) bool {
	if len(str) == 2 {
		numStrMap[1] = str
		strNumMap[sortString(str)] = 1
		return true
	}

	return false
}

func ruleSeven(str string) bool {
	if len(str) == 3 {
		numStrMap[7] = str
		strNumMap[sortString(str)] = 7
		return true
	}

	return false
}

func ruleFour(str string) bool {
	if len(str) == 4 {
		numStrMap[4] = str
		strNumMap[sortString(str)] = 4

		return true
	}

	return false
}

func ruleEight(str string) bool {
	if len(str) == 7 {
		numStrMap[8] = str
		strNumMap[sortString(str)] = 8

		return true
	}

	return false
}

func contains(str, substr string) bool {
	// not necessarily substr as a sequence
	for _, c := range substr {
		if !strings.ContainsRune(str, c) {
			return false
		}
	}

	return true
}

func commonChars(str1, str2 string) int {
	res := 0
	for _, c := range str1 {
		if strings.ContainsRune(str2, c) {
			res += 1
		}
	}

	return res
}

func ruleThree(str string) bool {
	if len(str) == 5 {
		seven, ok7 := numStrMap[7]
		one, ok1 := numStrMap[1]
		toContain := "not-possible"
		if ok7 {
			toContain = seven
		} else if ok1 {
			toContain = one
		}

		if contains(str, toContain) {
			numStrMap[3] = str
			strNumMap[sortString(str)] = 3

			return true
		}
	}

	return false
}

func ruleTwo(str string) bool {
	if len(str) == 5 {

		if four, ok := numStrMap[4]; ok && commonChars(str, four) == 2 {
			numStrMap[2] = str
			strNumMap[sortString(str)] = 2
			return true
		}
	}

	return false
}

func ruleFive(str string) bool {
	seven, ok7 := numStrMap[7]
	one, ok1 := numStrMap[1]
	notContained := ""
	if ok7 {
		notContained = seven
	} else if ok1 {
		notContained = one
	}
	if len(str) == 5 {
		if four, ok := numStrMap[4]; ok && commonChars(str, four) == 3 && (notContained != "" && !contains(str, notContained)) {
			numStrMap[5] = str
			strNumMap[sortString(str)] = 5

			return true
		}
	}

	return false
}

func ruleNine(str string) bool {
	if len(str) == 6 {
		if four, ok := numStrMap[4]; ok && contains(str, four) {
			numStrMap[9] = str
			strNumMap[sortString(str)] = 9
			return true
		}
	}

	return false
}

func ruleSix(str string) bool {
	seven, ok7 := numStrMap[7]
	one, ok1 := numStrMap[1]
	notContained := ""
	if ok7 {
		notContained = seven
	} else if ok1 {
		notContained = one
	}
	if len(str) == 6 {
		if notContained != "" && !contains(str, notContained) {
			numStrMap[6] = str
			strNumMap[sortString(str)] = 6

			return true
		}
	}

	return false
}

func ruleZero(str string) bool {
	// seems no solid rule to tell a zero
	// only need to sort everything else out and it is not 6 or 9
	if len(str) == 6 {
		seven, ok7 := numStrMap[7]
		four, ok4 := numStrMap[4]
		one, ok1 := numStrMap[1]

		if ok4 && ((ok7 && contains(str, seven)) || (ok1 && contains(str, one))) && !contains(str, four) {
			numStrMap[0] = str
			strNumMap[sortString(str)] = 0
			return true
		}
	}

	return false

}

var rules []func(string) bool = []func(string) bool{
	ruleOne,
	ruleFour,
	ruleSeven,
	ruleEight,
	ruleTwo,
	ruleThree,
	ruleFive,
	ruleSix,
	ruleNine,
	ruleZero,
}

func mapping(strs []string) {
	// this will each time re-initialize the maps
	numStrMap = make(map[int]string)
	strNumMap = make(map[string]int)

	unmapped := len(strs)
	for unmapped != 0 {
		copy := strs
		strs = []string{}
		for _, s := range copy {
			matched := false
			for _, rule := range rules {
				if rule(s) {
					matched = true
					break
				}
			}

			if !matched {
				strs = append(strs, s)
			}
		}

		unmapped = len(strs)
	}

	fmt.Printf("%v\n", numStrMap)
	fmt.Printf("%v\n", strNumMap)

}

func sevenSegmentsStdin() int {
	// this will use the maps that are populated by mapping each time
	scanner := bufio.NewScanner(os.Stdin)

	totalScore := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")

		outputs := []string{}
		trainings := []string{}
		trainings = append(trainings, strings.Split(strings.Trim(parts[0], " "), " ")...)
		mapping(trainings)

		outputs = append(outputs, strings.Split(strings.Trim(parts[1], " "), " ")...)
		fmt.Printf("%v\n", outputs)
		score := 0
		for _, op := range outputs {
			if val, ok := strNumMap[sortString(op)]; ok {
				score = score*10 + val
			}
		}

		totalScore += score
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return totalScore
}
