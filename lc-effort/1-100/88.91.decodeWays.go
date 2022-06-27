// https://leetcode.com/problems/decode-ways/

package main

import "fmt"

/*
so this should also be sub-problem... dp
f(s1s2s3) = f(s2s3) # take s1 alone if s2!=0
			+ f(s3) # taking s1s2, if s1s2 <= 26

then it can use memorization..

from bottom up can also be DP


*/

func _brute_force_numDecodings(s string) int {
	m := make([]int, len(s))
	var helper func(s string, start int) int
	helper = func(s string, start int) int {
		if len(s) == 0 {
			return 1
		}
		if s[0] == '0' {
			return 0
		}
		if len(s) == 1 {
			return 1
		}

		if m[start] != 0 {
			return m[start]
		}

		res := 0
		if s[1] != '0' && s[0:2] <= "26" {
			res = helper(s[1:], start+1) + helper(s[2:], start+2)
		} else if s[1] != '0' && s[0:2] > "26" {
			res = helper(s[1:], start+1)
		} else if s[1] == '0' && s[0:2] <= "26" {
			res = helper(s[2:], start+2)
		} else {
			res = 0
		}

		m[start] = res
		return res
	}

	return helper(s, 0)
}

/*
ugly.. and
"111111111111111111111111111111111111111111111" TLE

okay memorization
Runtime: 23 ms, faster than 5.43% of Go online submissions for Decode Ways.
Memory Usage: 2.1 MB, less than 45.11% of Go online submissions for Decode Ways.

okay.. thinking DP
from right to left

it actually is a fibnacci series again
 		1 1 1 1 1
first two must be calculated manually
		1 1 1 1 1
			  2	1
the 3rd 1, will be 2+1=3
the 2nd (from right 4th) 1, will be 2+3 = 5

run the program
"1"
"11"
"111"
"1111"

1
2
3
5

but of course it cannot be the ordinary fibnacci..
it has to meet the pre-condition
if it can go double, or only alone or only double..

so it is a dp..

  2 1 0 1
		1
	  0
	1			<== it can only go in 10.. so = dp[i+2]
  1				<== it can go alone or double, so dp[1] + dp[2], but dp[2] will be 0

  2 2 6
	  1 1
	2
  3

  2 2 7
	  1 1
	1
  2

  can I put a seed 1 at the behind? hmm cannot
  or seed 0
  2 1 0 1
        1 0
	  ?		<== no matter what, this cannot go.. so on zero, zero.. but other cases can be helped

  2 2 7
	  1	1
	2			1+1=2

   2 2 0
	   0  1
	 1			1+0=1
*/

func numDecodings(s string) int {
	l := len(s)
	dp := make([]int, l+1)
	dp[l] = 1
	dp[l-1] = 1
	if s[l-1] == '0' {
		dp[l-1] = 0
	}

	for i := l - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else if s[i:i+2] <= "26" {
			dp[i] = dp[i+1] + dp[i+2]
		} else {
			dp[i] = dp[i+1]
		}
	}

	return dp[0]
}

/*
		if s[i] == '0' {
			dp[i] = 0
		} else if s[i:i+2] <= "26" {
			dp[i] = dp[i+1] + dp[i+2]
		} else {
			dp[i] = dp[i+1]
		}
this is the complete logic...

Runtime: 0 ms, faster than 100.00% of Go online submissions for Decode Ways.
Memory Usage: 2 MB, less than 45.11% of Go online submissions for Decode Ways.
*/

func testNumDecoding() {
	fmt.Println(numDecodings("2101"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("06"))
	fmt.Println(numDecodings("111111111111111111111111111111111111111111111"))
}
