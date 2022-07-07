// https://leetcode.com/problems/palindromic-substrings/
package main

/*
if I don't know the answer to a previous palindromic substring problem
which is solving by expanding from one point or two points..

it would be really difficult
now knowing that, this should be quickly solvable
*/

func countSubstrings(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		// single point expansion
		count++
		for j, k := i-1, i+1; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] != s[k] {
				break
			}
			count++
		}

		// double point expansion
		if i+1 < len(s) && s[i] == s[i+1] {
			count++
			for j, k := i-1, i+2; j >= 0 && k < len(s); j, k = j-1, k+1 {
				if s[j] != s[k] {
					break
				}
				count++
			}

		}

	}

	return count
}

/*
"aaaaaa" failed..

so duplicates...
damn..
			for j, k := i-1, i+2; j >= 0 && k < len(s); j, k = j+1, k+1 {
				j+1 vs j-1
			for j, k := i-1, i+2; j >= 0 && k < len(s); j, k = j-1, k+1 {

"leetcode" failed again..

oh fuck, I should just break when s[j]!=s[k]
instead it let it go thru..

			if s[j] != s[k] {
                break
            }
			count++

	little error kills

passes
Runtime: 4 ms, faster than 51.91% of Go online submissions for Palindromic Substrings.
Memory Usage: 1.9 MB, less than 65.05% of Go online submissions for Palindromic Substrings.

on an if... at least ask youself what is expected by either branch...
*/
