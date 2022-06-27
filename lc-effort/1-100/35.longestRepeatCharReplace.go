// https://leetcode.com/problems/longest-repeating-character-replacement/

package main

import (
	"fmt"
	"math"
)

/*
actually I just come from another problem that requires a sliding window solution
so this one is a typical sliding window
window identified by [i,j], both end close i can equal to j
i,i = 0

window length j-i+1
the count of A-Z in this window will be in an array count[]

len(window) - max(count) <= k, if this is true, then we can flip this window to the most frequent char
this is a legit window, increasing j

otherwise, this is non-legit, increasing i..

so on and so forth, until j go over the limit...


*/

func maxCount(count []int) int {
	maxC := math.MinInt
	for _, c := range count {
		maxC = max(maxC, c)
	}

	return maxC
}

func characterReplacement(s string, k int) int {
	count := make([]int, 26)

	for i := range count {
		count[i] = 0
	}

	res := 0
	count[s[0]-'A'] = 1
	for i, j := 0, 0; i <= j && j < len(s); {
		maxCount := maxCount(count)
		windowLen := j - i + 1

		if windowLen-maxCount <= k {
			// legit window
			res = max(res, windowLen)
			if j < len(s)-1 {
				count[s[j+1]-'A']++
			}
			j++

		} else {
			// non-legit window
			count[s[i]-'A']--
			i++
		}

	}

	return res
}

func testCharacterReplacement() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}
