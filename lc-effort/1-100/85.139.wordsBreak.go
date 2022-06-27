// https://leetcode.com/problems/word-break/

package main

import (
	"fmt"
	"sort"
)

/*
first, establish this truth
if a part is not in the dict, then it will go over and blow up
so just need to break it part by part

some optimization in place
1. sort the dict and binary search
2. get the minL, maxL for quick start/early termination
*/

func _start_with_shortest_wrong_wordBreak(s string, wordDict []string) bool {

	sort.Strings(wordDict)

	minL, maxL := 10000, 1
	for _, w := range wordDict {
		minL = min(minL, len(w))
		maxL = max(maxL, len(w))
	}

	binarySearch := func(dict []string, str string) bool {
		l, r := 0, len(dict)-1

		for l < r {
			m := l + (r-l)/2
			if dict[m] == str {
				return true
			} else if dict[m] < str {
				l = m + 1
			} else {
				r = m
			}
		}

		return dict[l] == str
	}

	l, r := 0, minL // [l:r)
	for r <= len(s) {
		for r-l <= maxL && r <= len(s) {
			if binarySearch(wordDict, s[l:r]) {
				l = r
				r = r + minL
			} else {
				r++
			}
		}
		break
	}

	return l == len(s)

}

/*
	fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aaa"}))

	fails this one
	because I will be search aaa,aaa, and left 1 an 'a' out...

	hmm...
	so instead of match from minL, I should start with maxL
	that does bring some more complexity
*/

func _still_wrong_from_longest_wordBreak(s string, wordDict []string) bool {
	sort.Strings(wordDict)

	minL, maxL := 10000, 1
	for _, w := range wordDict {
		minL = min(minL, len(w))
		maxL = max(maxL, len(w))
	}

	binarySearch := func(dict []string, str string) bool {
		l, r := 0, len(dict)-1

		for l < r {
			m := l + (r-l)/2
			if dict[m] == str {
				return true
			} else if dict[m] < str {
				l = m + 1
			} else {
				r = m
			}
		}

		return dict[l] == str
	}

	l, r := 0, min(maxL, len(s)) // [l:r)
	for r-l >= minL && r <= len(s) {
		if binarySearch(wordDict, s[l:r]) {
			l = r
			r = min(r+maxL, len(s))
		} else {
			r--
		}
	}

	return l == len(s)

}

/*
	l, r := 0, maxL // [l:r)
	for r-l >= minL && r <= len(s) {
		if binarySearch(wordDict, s[l:r]) {
			l = r
			r = min(r+maxL, len(s))
		} else {
			r--
		}
	}

adjust the code to look from maxL to minL
"bb"
["a","b","bbb","bbbb"]

hmm.. this seems should be okay
let me see

	l, r := 0, min(maxL, len(s)) // [l:r)
	^ this line needs to limit the r to len(s) too


ah damn..

Wrong Answer
Details
Input
"abcd"
["a","abc","b","cd"]
Output
false
Expected
true


ah.. damn.. so my thesis of walking thru and get answers is not theoratical right

just try the luck
try both
shortest->longest
longest->shortest
*/

func _long_short_still_wrong_wordBreak(s string, wordDict []string) bool {
	sort.Strings(wordDict)

	minL, maxL := 10000, 1
	for _, w := range wordDict {
		minL = min(minL, len(w))
		maxL = max(maxL, len(w))
	}

	binarySearch := func(dict []string, str string) bool {
		l, r := 0, len(dict)-1

		for l < r {
			m := l + (r-l)/2
			if dict[m] == str {
				return true
			} else if dict[m] < str {
				l = m + 1
			} else {
				r = m
			}
		}

		return dict[l] == str
	}

	l, r := 0, min(maxL, len(s)) // [l:r)
	for r-l >= minL && r <= len(s) {
		if binarySearch(wordDict, s[l:r]) {
			l = r
			r = min(r+maxL, len(s))
		} else {
			r--
		}
	}

	if l == len(s) {
		return true
	}

	l, r = 0, minL // [l:r)
	for r-l <= maxL && r <= len(s) {
		if binarySearch(wordDict, s[l:r]) {
			l = r
			r = min(r+minL, len(s))
		} else {
			r++
		}
	}

	return l == len(s)
}

/*
"goalspecial"
["go","goal","goals","special"]

failed here..

so interesting..

so just start from shortest, search all possibilites
*/

func _works_top_down_wordBreak(s string, wordDict []string) bool {
	sort.Strings(wordDict)

	minL, maxL := 10000, 1
	for _, w := range wordDict {
		minL = min(minL, len(w))
		maxL = max(maxL, len(w))
	}

	binarySearch := func(dict []string, str string) bool {
		l, r := 0, len(dict)-1

		for l < r {
			m := l + (r-l)/2
			if dict[m] == str {
				return true
			} else if dict[m] < str {
				l = m + 1
			} else {
				r = m
			}
		}

		return dict[l] == str
	}

	var helper func(str string, dict []string) bool
	m := make(map[string]bool)
	helper = func(str string, dict []string) bool {
		if v, found := m[str]; found {
			return v
		}

		// base case, whole string can be found here...
		if len(str) <= maxL && len(str) >= minL && binarySearch(wordDict, str) {
			return true
		}

		l, r := 0, minL // [l:r)
		for r-l <= maxL && r <= len(str) {
			if binarySearch(wordDict, str[l:r]) {
				if helper(str[r:], dict) {
					m[str] = true
					return true
				}
			}
			r++
		}
		m[str] = false
		return false
	}

	return helper(s, wordDict)
}

/*
Runtime: 2 ms, faster than 61.82% of Go online submissions for Word Break.
Memory Usage: 2.2 MB, less than 49.53% of Go online submissions for Word Break.

Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Break.
Memory Usage: 2.3 MB, less than 26.44% of Go online submissions for Word Break.

so this is already fast enough

I was wondering if I can do better
1 <= s.length <= 300

apparently I can replace map with array
but I don't want to try that now..

let me thinking the bottom up DP..
there might be this

given length
if its sub-length is true, the rest part is also in the dict, then it can be true too
e.g.
abcd vs [a,b,abc,cd]
0 1 2 3 4 5 6 7 8
  t t t
looking at len-4

"abcd"[4-1] = d, d not in dict cannot use dp[3]
"abcd"[4-2] = cd, in dict, can use dp[2], if that is true and that is
*/

func _works_dp_go_thru_string_wordBreak(s string, wordDict []string) bool {
	m := make(map[string]struct{})
	minL, maxL := 10000, 1
	for _, w := range wordDict {
		minL = min(minL, len(w))
		maxL = max(maxL, len(w))
		m[w] = struct{}{}
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			_, found := m[s[j:i]]
			if dp[j] != 0 && i-j <= maxL && i-j >= minL && found {
				dp[i] = 1
				break
			}
		}
	}

	return dp[len(s)] != 0
}

/*
Runtime: 3 ms, faster than 47.96% of Go online submissions for Word Break.
Memory Usage: 2.1 MB, less than 67.47% of Go online submissions for Word Break.

nice.. this is simpler to understand but harder to come up

Runtime: 2 ms, faster than 61.71% of Go online submissions for Word Break.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Word Break.

oh damn right, search a string in the dict... you could simply use map..
why I am thinking to use binaray search...

somehow poisoned

Runtime: 2 ms, faster than 61.71% of Go online submissions for Word Break.
Memory Usage: 2.2 MB, less than 49.63% of Go online submissions for Word Break.

maybe I have enough of leetcode today..
go back working a bit..

ah.. smart idea of going thru the dict instead of the string itself
and from the back to front.. this way.. it is slighly quicker

let me do that
*/

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]struct{})
	for _, w := range wordDict {
		m[w] = struct{}{}
	}

	dp := make([]int, len(s)+1)
	dp[len(s)] = 1

	for j := len(s) - 1; j >= 0; j-- {
		for _, d := range wordDict {

			if j+len(d) <= len(s) && s[j:j+len(d)] == d {
				dp[j] = dp[j+len(d)]
				// break is important
				// because abcd vs [a,b,abc,d]
				// when at idx:0, going thru dic and a will be setting it to true
				// abc will set it to false.......
				// and only break when it is true...
				// corner case..tricky part
				if dp[j] != 0 {
					break
				}

			}
		}
	}

	return dp[0] != 0
}

/*
"cars"
["car","ca","rs"]

failed here
should only break if
the dp[j] is true already...

*/

func testWordBreak() {
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))

	fmt.Println(wordBreak("abcd", []string{"a", "b", "abc", "cd"}))

	fmt.Println(wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
	fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aaa"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))

}
