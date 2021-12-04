/***

The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?

***/

package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func isTriangleNum(num int) bool {
	sqrt := int(math.Sqrt(float64(num * 2)))
	return num*2 == sqrt*(sqrt+1)
}

func main() {
	dat, err := os.ReadFile("./p042_words.txt")
	if err != nil {
		fmt.Println("Failed to open the file")
	}

	words := strings.Split(string(dat), ",")
	triangleNumCount := 0
	for _, word := range words {
		word = strings.Trim(word, "\"")
		sum := 0
		for _, c := range word {
			sum += int(c-'A') + 1
		}

		if isTriangleNum(sum) {
			triangleNumCount += 1
		}
	}

	fmt.Println(triangleNumCount)

}
