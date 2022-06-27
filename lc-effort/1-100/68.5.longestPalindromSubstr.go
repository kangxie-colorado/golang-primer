// https://leetcode.com/problems/longest-palindromic-substring/

package main

import "fmt"

/*
	I did this one before but I don't remember

	looking at it, really not much clue where to start
	of cource there is a O(n^3) brute force solution

	but lets not think about that

	then I think where to start, to extend to both side?
	why don't start with mid

		mid-2	mid-1	mid mid+1 mid+2
	then left, right

	so naturally you see a recursion structure

	let me try


*/

func extendToLR(toL, toR int, s string, l, r int) (int, int) {
	if toL > toR {
		return toL, toR
	}

	for toL >= l && toR <= r && s[toL] == s[toR] {
		toL--
		toR++
	}
	return toL + 1, toR - 1
}

func extendSameChar(toL int, toR int, s string, l, r int) (int, int) {
	// only possible when it contains one same char
	for i := toL; i <= toR; i++ {
		if i > toL && s[i] != s[i-1] {
			return toL, toR
		}
	}

	toL2 := toL
	for toL >= l && s[toL] == s[toR] {
		toL--
	}
	for toR <= r && s[toL2] == s[toR] {
		toR++
	}
	return toL + 1, toR - 1
}

// return the start,end idx inclusive? nah.. follow the convention let end be open
// nah.. inclusive is better
func longestPalindromeHelper(s string, start, end int) (int, int) {
	if start >= end {
		return start, end
	}

	l, r := start, end
	mid := l + (r+1-l)/2

	// the length center around the mid one
	// either extend by two or by one
	start, end = extendToLR(mid, mid, s, l, r)

	lStart, lEnd := longestPalindromeHelper(s, l, mid-1)
	st, ed := extendToLR(lStart, lEnd, s, l, r)
	if end-start < ed-st {
		start, end = st, ed
	}
	st, ed = extendSameChar(lStart, lEnd, s, l, r)
	if end-start < ed-st {
		start, end = st, ed
	}

	rStart, rEnd := longestPalindromeHelper(s, mid+1, r)
	st, ed = extendToLR(rStart, rEnd, s, l, r)
	if end-start < ed-st {
		start, end = st, ed
	}
	st, ed = extendSameChar(rStart, rEnd, s, l, r)
	if end-start < ed-st {
		start, end = st, ed
	}

	return start, end

}

func _2_rabiit_hole_longestPalindrome(s string) string {

	start, end := longestPalindromeHelper(s, 0, len(s)-1)

	return s[start : end+1]
}

/*
this sends me down to a very deep rabit hole
I don't think I can pull it off

need to re-think
to be a palindromic str, the two sides must be the same, right?

so I create a map
a: [1,3,4,5,8]

when I scan a, I see okay.. I have these many possibilities
then I pick the farthest and see if it is palindromic, if yes.. great, no look further
otherwise, look next one

just going thru all of them

*/

func isSubStrPalindrom(s string) bool {
	for l, r := 0, len(s)-1; l <= r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}

	return true
}

func longestPalindrome(s string) string {
	m := make(map[rune][]int)
	for i, c := range s {
		m[c] = append(m[c], i)
	}

	maxLen := 0
	resStr := ""
	for i, c := range s {
		for j := len(m[c]) - 1; j >= 0; j-- {
			lastPos := m[c][j]
			if lastPos+1-i < maxLen {
				break
			}

			if isSubStrPalindrom(s[i : lastPos+1]) {
				if lastPos+1-i > maxLen {
					maxLen = lastPos + 1 - i
					resStr = s[i : lastPos+1]
				}
			}

		}
	}

	return resStr
}

/*
Runtime: 71 ms, faster than 44.23% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 5.2 MB, less than 27.99% of Go online submissions for Longest Palindromic Substring.

I don't really need hashmap for this...

looking at the time distribution of the submissions..
there seems to be a solution with one less order of complexity
*/

func testLongestPalindrome() {
	fmt.Println(longestPalindrome("aabbbbbb"))

	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("cbbbbbbxyziosadd"))
	fmt.Println(longestPalindrome("cbabcbllbcbabxyziosadd"))
	fmt.Println(longestPalindrome("cbabcbllbcbabcxyziosadd"))
	fmt.Println(longestPalindrome("cbabcbllbcbabxyziosaddddddddddddddd"))

}
