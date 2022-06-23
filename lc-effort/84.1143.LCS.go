// https://leetcode.com/problems/longest-common-subsequence/

package main

import "fmt"

/*
pity i didn't do this before I see the brief video
i now know this is a DP

[a1,a2..an] lcs [b1,b2,..bn]
	= if a1==b1, 1+ [a2...an] lcs [b2...bn]
	   or max( [a2...an] lcs [b1...bn],
	   		   [a1...an] lcs [b2...bn])

let me go draw this in whiteboard
*/

/*
yeah, as I suspected
Time Limit Exceeded
Details
Last executed input
"mhunuzqrkzsnidwbun"
"szulspmhwpazoxijwbq"

so this can be
- bottom up dp
- memorization

let me do memorizing first
interesting.. even with memorization it is still TLE

then I think we just brute force..
m*n

for one string, seach every byte.. replace the byte with something unmatchable..
then repeat..

worst is m*n

that DP seems like more that m*n.. werid where is wrong?
*/

func _wrong_brute_force_longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}

	bytes1 := []byte(text1)
	bytes2 := []byte(text2)

	res1 := 0
	matchIdx := -1
	for b1 := range bytes1 {

		for _, b := range bytes1[b1:] {
			for i := range bytes2 {
				if b == bytes2[i] && i > matchIdx {
					res1++
					// only containing lower case enligsh letters, so 0 is not matchable
					bytes2[i] = 0
					matchIdx = i
					break
				}
			}
		}
	}

	bytes1, bytes2 = bytes2, bytes1
	res2 := 0
	matchIdx = -1
	for b1 := range bytes1 {

		for _, b := range bytes1[b1:] {
			for i := range bytes2 {
				if b == bytes2[i] && i > matchIdx {
					res2++
					// only containing lower case enligsh letters, so 0 is not matchable
					bytes2[i] = 0
					matchIdx = i
					break
				}
			}
		}
	}

	return max(res1, res2)

}

/*
ah.. okay..
it has to be in same order

still not right
"pmjghexybyrgzczy"
"hafcdqbgncrcbihkd"

I got 4
but the answer is 6

I see, it can start at any place..
still got wrong answer...

so have to swap

still not right..

so yeah.. this brute force is not even right
it has to be intertwined... this is harder than I thought...
*/

/*
"pmjghexybyrgzczy"
"hafcdqbgncrcbihkd"

"mhunuzqrkzsnidwbun"
"szulspmhwpazoxijwbq"
*/

func _pass_but_slow_longestCommonSubsequence(text1 string, text2 string) int {
	type StrPair struct {
		str1 string
		str2 string
	}

	m := make(map[StrPair]int)

	var helper func(s1, s2 string) int
	helper = func(s1, s2 string) int {
		if len(s1)*len(s2) == 0 {
			return 0
		}

		if v, found := m[StrPair{s1, s2}]; found {
			return v
		}

		res := 0
		if s1[0] == s2[0] {
			res = 1 + helper(s1[1:], s2[1:])
		} else {
			res = max(helper(s1, s2[1:]), helper(s1[1:], s2))
		}

		m[StrPair{s1, s2}] = res
		return res
	}

	return helper(text1, text2)
}

/*
Runtime: 1057 ms, faster than 5.08% of Go online submissions for Longest Common Subsequence.
Memory Usage: 190.6 MB, less than 5.08% of Go online submissions for Longest Common Subsequence.

albeit super slow, it passed
it was a stupid typo..

I should recursively call helper(), not the outer function

If I change the map key to ints not strings
*/

func _map_slow_longestCommonSubsequence(text1 string, text2 string) int {
	type IdxPair struct {
		idx1 int
		idx2 int
	}

	m := make(map[IdxPair]int)

	var helper func(s1, s2 string, idx1, idx2 int) int
	helper = func(s1, s2 string, idx1, idx2 int) int {
		if len(s1)*len(s2) == 0 {
			return 0
		}

		if v, found := m[IdxPair{idx1, idx2}]; found {
			return v
		}

		res := 0
		if s1[0] == s2[0] {
			res = 1 + helper(s1[1:], s2[1:], idx1+1, idx2+1)
		} else {
			res = max(helper(s1, s2[1:], idx1, idx2+1), helper(s1[1:], s2, idx1+1, idx2))
		}

		m[IdxPair{idx1, idx2}] = res
		return res
	}

	return helper(text1, text2, 0, 0)
}

/*
Runtime: 658 ms, faster than 6.30% of Go online submissions for Longest Common Subsequence.
Memory Usage: 119.4 MB, less than 9.14% of Go online submissions for Longest Common Subsequence.

better but still low...
let me change map to m*n matrix


*/

func _top_down_recursive_longestCommonSubsequence(text1 string, text2 string) int {
	rows := len(text1)
	cols := len(text2)

	m := make([][]int, rows+1)
	for r := range m {
		m[r] = make([]int, cols+1)
		for c := range m[r] {
			m[r][c] = -1
		}
	}

	var helper func(s1, s2 string, idx1, idx2 int) int
	helper = func(s1, s2 string, idx1, idx2 int) int {
		if len(s1)*len(s2) == 0 {
			return 0
		}

		if m[idx1][idx2] != -1 {
			return m[idx1][idx2]
		}

		res := 0
		if s1[0] == s2[0] {
			res = 1 + helper(s1[1:], s2[1:], idx1+1, idx2+1)
		} else {
			res = max(helper(s1, s2[1:], idx1, idx2+1), helper(s1[1:], s2, idx1+1, idx2))
		}

		m[idx1][idx2] = res
		return res
	}

	return helper(text1, text2, 0, 0)
}

/*
Runtime: 31 ms, faster than 24.59% of Go online submissions for Longest Common Subsequence.
Memory Usage: 11.3 MB, less than 26.83% of Go online submissions for Longest Common Subsequence.

this is a big jump..

Runtime: 11 ms, faster than 47.97% of Go online submissions for Longest Common Subsequence.
Memory Usage: 11.2 MB, less than 29.27% of Go online submissions for Longest Common Subsequence.

this is actually setting up the bottom-up DP already...
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	rows := len(text1)
	cols := len(text2)

	m := make([][]int, rows+1)
	for r := range m {
		m[r] = make([]int, cols+1)
		for c := range m[r] {
			m[r][c] = 0
		}
	}

	s1 := reverse(text1)
	s2 := reverse(text2)

	for r := 1; r <= rows; r++ {
		for c := 1; c <= cols; c++ {
			if s1[r-1] == s2[c-1] {
				m[r][c] = 1 + m[r-1][c-1]
			} else {
				m[r][c] = max(m[r][c-1], m[r-1][c])
			}
		}
	}

	return m[rows][cols]
}

/*
Runtime: 10 ms, faster than 51.83% of Go online submissions for Longest Common Subsequence.
Memory Usage: 11 MB, less than 74.80% of Go online submissions for Longest Common Subsequence.

*/

func testLCS() {
	fmt.Println(longestCommonSubsequence("ace", "abcdfe"))
	fmt.Println(longestCommonSubsequence("pmjghexybyrgzczy", "hafcdqbgncrcbihkd"))
	fmt.Println(longestCommonSubsequence("mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq"))
}
