package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func readTemplateAndInstructionsStdin() (string, map[string]string) {
	scanner := bufio.NewScanner(os.Stdin)

	// first row: the template
	scanner.Scan()
	template := scanner.Text()

	// read off the blank/separate line
	scanner.Scan()

	pairInsertMap := map[string]string{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		pairInsertMap[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return template, pairInsertMap
}

func makePolymer(template string, pairInsertMap map[string]string, steps int) string {
	for s := 0; s < steps; s++ {

		inserts := []string{}
		for i := 1; i < len(template); i++ {
			pair := template[i-1 : i+1]
			insert := pairInsertMap[pair]
			inserts = append(inserts, insert)
		}

		// combine the original template and the inserts
		// notice the weirdness of rune vs string here...
		combine := template[0:1]
		for i := 1; i < len(template); i++ {
			combine = combine + inserts[i-1] + template[i:i+1]
		}

		template = combine
	}

	return template
}

func polymerAnalysis(polymer string) (int, int) {
	matCountMap := map[rune]int{}

	for _, r := range polymer {
		matCountMap[r]++
	}

	most, least := 0, math.MaxInt32
	for _, v := range matCountMap {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	fmt.Println(matCountMap)

	return most, least
}

func polymerDriver() {
	template, pairInsertMap := readTemplateAndInstructionsStdin()
	//fmt.Printf("template: %v\n, the pair insert rules: %v\n", template, pairInsertMap)

	output := makePolymer(template, pairInsertMap, 10)
	most, least := polymerAnalysis(output)

	fmt.Printf("Most - Least: %v\n", most-least)
}

// part 2 thinking: okay, 40 steps, the memory and cpu won't be enough
// difficulty is it seems you must retain the information of what the string looks like, to know the next one
// observation is if you cut the string in half, you can do the calculation separately and you would be able to
// combine the final results... but still, it would be more than the computer can provide
// then spark! you just need to book-keep the pairs information
// pairCount[BB] = x; pairCount[BC] = y
// when split... BB become 1*BC+1*CN... no need to care their position...
// so yeah, lets do it

func getInitialPairCountMap(template string) map[string]int64 {
	pairCountMap := map[string]int64{}
	for i := 1; i < len(template); i++ {
		pair := template[i-1 : i+1]
		pairCountMap[pair]++
	}

	return pairCountMap
}

func makePolymer2(template string, pairInsertMap map[string]string, steps int) map[string]int64 {
	pairCountMap := getInitialPairCountMap(template)

	for s := 0; s < steps; s++ {

		copyMap := map[string]int64{}
		for k, v := range pairCountMap {
			copyMap[k] = v
		}

		for pair, count := range copyMap {
			insert := pairInsertMap[pair]
			newPair1, newPair2 := pair[:1]+insert, insert+pair[1:]

			// for a pair replace with two pairs
			// and notice the pair has a count, replace it with two count of pairs
			// this also generalize count==0 special case
			pairCountMap[newPair1] += count
			pairCountMap[newPair2] += count
			pairCountMap[pair] -= count
		}
	}

	return pairCountMap
}

func polymerAnalysis2(pairCountMap map[string]int64, firstRune, lastRune rune) (int64, int64) {
	matCountMap := map[rune]int64{}

	for pair, count := range pairCountMap {
		for _, r := range pair {
			matCountMap[r] += count
		}
	}

	// NOTICE THAT: every rune but the start/end is counted twice, BCHD I will be count like BC:1 CH:1, HD:1
	// so C is 2, H is 2, B and D is 1
	// so yeah... make sure this is factored in the calculation

	// adjust the firstRune and lastRune count, then we can divide by 2
	matCountMap[firstRune]++
	matCountMap[lastRune]++

	var most, least int64 = 0, math.MaxInt64
	for _, v := range matCountMap {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most / 2, least / 2
}

func polymerDriver2() {
	template, pairInsertMap := readTemplateAndInstructionsStdin()
	//fmt.Printf("template: %v\n, the pair insert rules: %v\n", template, pairInsertMap)

	output := makePolymer2(template, pairInsertMap, 40)

	fmt.Println(output)

	// need to know the first and last rune
	// luckily, they won't change along the insertion
	var firstRune, lastRune rune = rune(template[0]), rune(template[len(template)-1])
	most, least := polymerAnalysis2(output, firstRune, lastRune)

	fmt.Printf("Most - Least: %v\n", most-least)

}
