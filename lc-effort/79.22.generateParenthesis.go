// https://leetcode.com/problems/generate-parentheses/

package main

import "fmt"

/*
not thinking too much
just use backtracking+stack to see if I can do it

if I cannot' , I'll go sleeping
*/

func gpHelper(l int, r int, s string, stack []string) {
	// n-i "(", and n ")"
	if l == 0 {
		// means the rest must be r or it will not work
		if len(stack) != r {
			return
		}

		for r > 0 {
			s += ")"
			r--
		}
		results[s] = struct{}{}
		return
	}

	// now you either put a l or a r
	// put a left ( and back out...
	s += "("
	stack = append(stack, "(")
	gpHelper(l-1, r, s, stack)
	stack = stack[:len(stack)-1]
	s = s[:len(s)-1]

	if len(stack) > 0 {
		// can try put a right )
		s += ")"
		stack = stack[:len(stack)-1]
		gpHelper(l, r-1, s, stack)
		stack = append(stack, "(")
		s = s[:len(s)-1]
	}

}

var results map[string]struct{}

func _1_generateParenthesis(n int) []string {
	results = make(map[string]struct{})

	for i := 1; i <= n; i++ {
		stack := []string{}
		s := ""

		for j := 0; j < i; j++ {
			s += "("
			stack = append(stack, "(")
		}
		/*
			yeah.. so I drivie from n-1..
			it is not complete on its own

			I found these duplciates created by myself hard to track
			so be humble... never assume yourself mastered anything
		*/
		gpHelper(n-i, n, s, stack)
	}
	res := []string{}
	for s := range results {
		res = append(res, s)
	}
	return res
}

/*
Runtime: 8 ms, faster than 10.03% of Go online submissions for Generate Parentheses.
Memory Usage: 3.2 MB, less than 27.76% of Go online submissions for Generate Parentheses.


okay.. watched the neetcode code
expanded the n=3 in a tree expansion and I notice it really has no duplicates to worry about

it only observes two priciples..
1. open <=3 && close <=3
	meaning when open<3, we can apply an open
2. close < open, meaning only when close<open we can apply a close
	this is a rather beautiful rule since it actually prevents ')' appears before '('
*/

func _2_generateParenthesis(n int) []string {
	var parentheses []string
	var stack []byte

	parentheses = make([]string, 0)
	stack = make([]byte, 0)

	var gp2 func(open, close int)
	gp2 = func(open, close int) {
		if open == n && close == n {
			parentheses = append(parentheses, string(stack))
			return
		}

		if open < n {
			stack = append(stack, '(')
			gp2(open+1, close)
			stack = stack[:len(stack)-1]
		}

		// meaning used close is smaller than used open
		// stack/string is left heavy
		if close < open {
			stack = append(stack, ')')
			gp2(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	gp2(0, 0)
	return parentheses
}

/*
Runtime: 4 ms, faster than 43.32% of Go online submissions for Generate Parentheses.
Memory Usage: 2.7 MB, less than 74.68% of Go online submissions for Generate Parentheses.

then I think why I have to deal with duplicates and he doesn't
one important point is I have multiple driving point and he only has one...

I think maybe we code wasn't actually necessary to do that many driving point
let me change my code this way but n->0
*/

func __copiedhere_for_referene_gpHelper(l int, r int, s string, stack []string) {
	// n-i "(", and n ")"
	/*
		apparently this is not generic
		I would rather let the two rules to deal with what is after l==0

		which is l<n..
		when l==n or 0, just adding the right one
	*/
	if l == 0 {
		// means the rest must be r or it will not work
		if len(stack) != r {
			return
		}

		for r > 0 {
			s += ")"
			r--
		}
		results[s] = struct{}{}
		return
	}

	// now you either put a l or a r
	// put a left ( and back out...
	s += "("
	stack = append(stack, "(")
	gpHelper(l-1, r, s, stack)
	stack = stack[:len(stack)-1]
	s = s[:len(s)-1]

	if len(stack) > 0 {
		// can try put a right )
		s += ")"
		stack = stack[:len(stack)-1]
		gpHelper(l, r-1, s, stack)
		stack = append(stack, "(")
		s = s[:len(s)-1]
	}

}

func generateParenthesis(n int) []string {
	var parentheses []string
	parentheses = make([]string, 0)
	s := ""

	var helper func(open, close int)
	helper = func(open, close int) {
		if open == 0 && close == 0 {
			parentheses = append(parentheses, s)
			return
		}

		if open > 0 {
			s += "("
			helper(open-1, close)
			s = s[:len(s)-1]
		}

		// if s has more '(' than ')
		// or I can use a stack to maintain the order
		// but it ain't really necessary
		// it can be told by s.count('(') > s.count(')')
		// or by close>open... (the left close ')' is more than open '(')
		if close > open {
			s += ")"
			helper(open, close-1)
			s = s[:len(s)-1]
		}
	}
	helper(n, n)
	return parentheses
}

/*
Runtime: 4 ms, faster than 43.32% of Go online submissions for Generate Parentheses.
Memory Usage: 2.8 MB, less than 45.37% of Go online submissions for Generate Parentheses.
*/

func testGenerateParenthesis() {
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}
