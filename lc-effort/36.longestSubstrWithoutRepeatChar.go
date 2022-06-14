// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import "fmt"

/*
this is of course a sliding window problem
I quickly jumped into 26 char... but then found
s consists of English letters, digits, symbols and spaces.
so not 26

but I need to use a map or something
yes, a map, when the count hits 2, shrink the window

*/

func lengthOfLongestSubstring(s string) int {
	count := make(map[byte]int)

	maxLen := 0
	for i, j := 0, 0; i <= j && j < len(s); {
		windLen := j - i + 1
		if _, found := count[s[j]]; !found || count[s[j]] < 1 {
			maxLen = max(maxLen, windLen)
			count[s[j]]++
			j++
		} else {
			// no need to test if it is >= 1, because when it is found, then it naturally is
			// ugh.. when count becomes 0, it will not disappear
			count[s[i]]--
			i++
		}
	}

	return maxLen
}

/*
Runtime: 12 ms, faster than 55.52% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 3 MB, less than 54.19% of Go online submissions for Longest Substring Without Repeating Characters.

not bad, not great but fine

so this
func lengthOfLongestSubstring(s string) int {
    m := make(map[rune]int)
    i := 0
    maxLen := 0
    ok := false // this asved 4 ms for total run
    ss := []rune(s)
    for j := 0; j < len(ss); j++ {
        _, ok = m[ss[j]]
        if ok {
           i =  Max(m[ss[j]], i)
        }
        maxLen = Max(maxLen, j-i+1)
        m[ss[j]] = j + 1
    }
 return maxLen
}

        if ok {
           i =  Max(m[ss[j]], i)
        }
		m[ss[j]] = j + 1

		it has this window fast forward.. slippy but I don't think I need that
*/

func testLengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
