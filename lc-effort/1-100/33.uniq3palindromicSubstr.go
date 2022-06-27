// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/

package main

import "math"

/*
	of course there is the brute force - 3 loop

	then I thought about bitmap, not easy
	the counter tracking not easy

then I thought
	1. build the counter map
	2. for each char in the map
		if m[char] >= 2, then there is possibilities
		find left-char, find right-char,
		the substr in middle... any char makes such a subsequece
	O(n^2)

	let me do this then think more



*/

func helper(char rune, s string) int {
	left := len(s) - 1
	right := 0

OUTER:
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == char {
			left = i
			for j := len(s) - 1; j > i; j-- {
				if rune(s[j]) == char {
					right = j
					break OUTER
				}
			}
		}
	}

	// must at least contain on char between left:right
	counter := make(map[byte]int)
	if left < right-1 {
		for i := left + 1; i < right; i++ {
			counter[s[i]]++
		}
	}

	return len(counter)
}

func _1_countPalindromicSubsequence(s string) int {

	counter := make(map[rune]int)
	for _, c := range s {
		counter[c]++
	}

	res := 0
	for k, v := range counter {
		if v < 2 {
			continue
		}

		res += helper(k, s)
	}

	return res
}

/*
Runtime: 482 ms, faster than 50.00% of Go online submissions for Unique Length-3 Palindromic Subsequences.
Memory Usage: 6 MB, less than 100.00% of Go online submissions for Unique Length-3 Palindromic Subsequences.

so it passes...
let me think something more

if I start with the very middle 3, nope..

maybe I can use the recursion to return the substr.. but nope... some pairs might just appear in front/end, it could be middle/end
not much new thinking

let me check the hint

What is the maximum number of length-3 palindromic strings?
How can we keep track of the characters that appeared to the left of a given position?

alright.. the hint doesn't make sense to me
I must be missing some fundemental stuff here

let me check discussions
most people just do what I did...

but there is something ... let me debug it
*/

func idx(ch byte) int {
	return int(ch) - int('a')
}

var list []M

func initTest(s string) {
	list = make([]M, 26)
	for i := range list {
		list[i].min = math.MaxInt
		list[i].max = math.MinInt
	}
	for i, r := range s {
		ch := idx(byte(r))
		list[ch].min = Min(list[ch].min, i)
		list[ch].max = Max(list[ch].max, i)
	}
}

// so it create an array, get the first occurence and last occurence of any char
// then it goes thru the array, pair by pair and count whats in between
// still same to my solution
// why my time is so bad?

func countPalindromicSubsequence(s string) int {
	initTest(s)

	result := 0
	str := []byte(s)
	for _, m := range list {
		count := make([]int, 26)
		for i := m.min + 1; i < m.max; i++ {
			ch := str[i]
			count[idx(ch)]++
		}
		for _, v := range count {
			if v > 0 {
				result++
			}
		}
	}
	return result
}

type M struct {
	min int
	max int
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
