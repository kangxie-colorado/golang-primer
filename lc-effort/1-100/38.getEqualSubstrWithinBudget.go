// https://leetcode.com/problems/get-equal-substrings-within-budget/

package main

import "fmt"

/*
	tracking state: cost

*/

func _1_equalSubstring(s string, t string, maxCost int) int {
	maxLen := 0
	cost := 0
	for i, j := 0, 0; j < len(s); {

		cost += abs(int(t[j]) - int(s[j]))

		for cost > maxCost {
			cost -= abs(int(t[i]) - int(s[i]))
			i++
		}

		j++
		maxLen = max(maxLen, j-i)
	}

	return maxLen

}

/*
Wrong Answer
Details
Input
"pxezla"
"loewbi"
25
Output
1
Expected
4

hmm... how can be wrong?
missing... abs()
and the byte minus ops is funny becasue it has only 256.. so a minus go over limit too easy

abs(int(t[j]) - int(s[j]))
convert to int then minus... don't minus then convert

and yeah... great
Runtime: 4 ms, faster than 61.11% of Go online submissions for Get Equal Substrings Within Budget.
Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Get Equal Substrings Within Budget.

try the non-shrinkable..

*/
func equalSubstring(s string, t string, maxCost int) int {
	i, j := 0, 0
	cost := 0
	for j < len(s) {
		cost += abs(int(t[j]) - int(s[j]))
		if cost > maxCost {
			cost -= abs(int(t[i]) - int(s[i]))
			i++
		}
		j++
	}

	return j - i

}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Get Equal Substrings Within Budget.
Memory Usage: 2.9 MB, less than 83.33% of Go online submissions for Get Equal Substrings Within Budget.


why it is called non-shinkable.. because the window never shrinks...
just when the condition is not met, it won't grow... just sliding/shifting/moving
*/

func testEqualSubstring() {
	fmt.Println(equalSubstring("pxezla", "loewbi", 25))
}
