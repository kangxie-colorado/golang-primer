// https://leetcode.com/problems/remove-duplicate-letters/
// https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/

package main

import "fmt"

/*
I came here from a stack topic post but right now I didn't see that clearly
instead I saw a hashmap based solution

it maybe one order slower than stack.. but it should work
tomorrow work up the stack.. will see

ah... the hashmap. won't work
Input: s = "cbacdcbc"
Output: "acdb"

it is not "abcd"... the order must be maintained..
so okay... not a line of code for tonight..
*/

func _wont_work_removeDuplicateLetters(s string) string {
	chars := make([]int, 26)
	for i := range s {
		chars[s[i]-'a']++
	}
	/*
	   ah... the hashmap. won't work
	   Input: s = "cbacdcbc"
	   Output: "acdb"

	   it is not "abcd"... the order must be maintained..
	   so okay... not a line of code for tonight..
	*/

	return ""
}

/*
https://leetcode.com/problems/remove-duplicate-letters/discuss/1859410/JavaC%2B%2B-DETAILED-%2B-VISUALLY-EXPLAINED-!!

this is god damn too hard for me
probably my subconcious thought this all night and I got zero clue...

waking up and my head hurts...

so just fuck giving up, read someone brilliant's solution and understand it...

okay... I was brushing this but never pan it out
so the idea is to maintain a increasing stack
when it violates.. check if there is more top to choose from... if not, stick to it, continue the stack as is
if yes, pop/unmark and continue...

yes, shite... so this is not a pure increasing stack.. but sort of ..
hard to think out

how to prove it?
once again, focus on one char at a time, if it is only one remaining available... has to use it, end of story
otherwise, if a smaller one comes up, you have to use the later one..

considering local vs global...
*/

func removeDuplicateLetters(s string) string {
	lastIdx := make([]int, 26)
	for i, c := range s {
		lastIdx[c-'a'] = i
	}

	stack := []int{}
	inStack := make([]int, 26)

	for i := range s {
		if inStack[s[i]-'a'] == 1 {
			continue
		}

		for len(stack) > 0 && s[stack[len(stack)-1]] > s[i] {
			top := stack[len(stack)-1]
			if lastIdx[s[top]-'a'] < i {
				break
			}
			stack = stack[:len(stack)-1]
			inStack[s[top]-'a'] = 0
		}

		stack = append(stack, i)
		inStack[s[i]-'a'] = 1
	}

	res := make([]byte, len(stack))
	for i := range stack {
		res[i] = s[stack[i]]
	}

	return string(res)
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Duplicate Letters.
Memory Usage: 2.1 MB, less than 42.25% of Go online submissions for Remove Duplicate Letters.
*/

func testRemoveDuplicateLetters() {
	fmt.Println(removeDuplicateLetters("bcabc"))
	fmt.Println(removeDuplicateLetters("cbacdcb"))
}
