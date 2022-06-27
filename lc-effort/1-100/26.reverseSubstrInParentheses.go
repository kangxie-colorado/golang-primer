// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/

/*
analysis
	naturally this is a recursive thing...
	each () return a str

	on '(' into stack
		then push chars
		on ')', coming out of stack..

	possibly missing some optimization but let me try

*/

package main

import (
	"fmt"
	"strings"
)

func reverseParenthesesHelper(pos int, s string) ([]rune, int) {
	stack := []rune{}

	for i := pos; i < len(s); i++ {
		if s[i] == '(' {
			substr, rightP := reverseParenthesesHelper(i+1, s)
			ReverseSlice(substr)

			for _, r := range substr {
				stack = append(stack, r)
			}

			i = rightP

		} else if s[i] == ')' {
			return stack, i
		} else {
			stack = append(stack, rune(s[i]))
		}

	}

	return stack, 0
}

func _1_reverseParentheses(s string) string {

	str, _ := reverseParenthesesHelper(0, s)

	return string(str)

}

/*
Success
Details
Runtime: 3 ms, faster than 35.71% of Go online submissions for Reverse Substrings Between Each Pair of Parentheses.
Memory Usage: 2.3 MB, less than 28.57% of Go online submissions for Reverse Substrings Between Each Pair of Parentheses.

okay... at least accepted
the difficulty is to manage the i... don't let walk less, neither let it walk more
it must walk to the right position and let the i++ walk the next

also the 3 branches.. must make sure, only one branch happens... don't make 2 happens
that will be inviting trouble.. there is no saving by doing sometime 1 sometime 2 in a loop.. only chaos.

and now.. apparently, I know with multiple layer of parenthese, there are duplicate calculations..
if we can optimize that, then it might improve...


Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Substrings Between Each Pair of Parentheses.
Memory Usage: 2.4 MB, less than 28.57% of Go online submissions for Reverse Substrings Between Each Pair of Parentheses.

aha...

I was think to optimize to use the layer odd layer, even layer, didn't fly
then run it again.. maybe I am already the best solution?

Hints
. Find all brackets in the string.
. Does the order of the reverse matter ?
. The order does not matter.

let me take a break..

by the hints, I think I can scan and get each pair of parenthese
then I go on to reverse... and each time I reverse what is inside the pair of "()" including the nested "()" and remove one layer
I don't think the performance will be good

one optimization is adjacent (()) can be eliminated natrually

let me code it up



*/

type ParenthesesPair struct {
	left  int
	right int
}

func reverseIgnoreParenthese(s []byte) {
	size := len(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		if s[i] == '(' {
			s[i] = '_'
		}
		if s[j] == ')' {
			s[j] = '_'
		}
		s[i], s[j] = s[j], s[i]
	}
}

func reverseParentheses(s string) string {
	pStack := []int{}
	pairsOfPar := []ParenthesesPair{}

	for i, c := range s {
		if c == '(' {
			pStack = append(pStack, i)
		}

		if c == ')' {
			leftP := pStack[len(pStack)-1]
			pStack = pStack[:len(pStack)-1]
			pairsOfPar = append(pairsOfPar, ParenthesesPair{leftP, i})
		}

	}
	// aha.. this actually build up the pair inside out
	// the most inside pair will be firstly added to pairsOfPar
	// if I reverse the parisOfPar
	// ReverseSlice(pairsOfPar)
	// it will be logically wrong

	for _, p := range pairsOfPar {
		tmp := []byte(s[p.left : p.right+1])
		reverseIgnoreParenthese((tmp))
		s = s[:p.left] + string(tmp) + s[p.right+1:]
	}

	return strings.ReplaceAll(s, "_", "")
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Substrings Between Each Pair of Parentheses.
Memory Usage: 2.8 MB, less than 21.43% of Go online submissions for Reverse Substrings Between Each Pair of Parentheses.

but still I am not using the hints...
*/

/*
Wrong Answer
Details
Input
"a(bcdefghijkl(mno)p)q"
Output
"aponm(lkjihcdefgbq"
Expected
"apmnolkjihgfedcbq"

reverse the pairs then something wrong.. interesting
also I didn't see the case of "()"" not at the front/end


okay.. it is easy to see even without this case
reverse the order is wrong
it is still better to be from inside out...

okay, let me see others' solutions
*/

func testReverseParentheses() {
	fmt.Println(reverseParentheses("a(bcdefghijkl(mno)p)q"))
}
