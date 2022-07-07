// https://leetcode.com/problems/remove-k-digits/

package main

import (
	"fmt"
	"strings"
)

/*
this seems very difficult to wrap you head around
but with soloving previous "remove duplicate letter and keep smallest string"
it gives me some idea

so basically I still maintain a non-decreasing stack, when violate I pop() on condition
the condition being the rest elements are still enough to form the required size...

when finish, just take results from the stack
if could be stack is over the required size, just taking upto the required size
[1,2,3,4,5,6] k=1
it will end up like [1,2,3,4,5,6]

then taking [1,2,3,4,5] for the answer...
*/

func removeKdigits(num string, k int) string {
	l := len(num)
	r := l - k // end required size
	if r == 0 {
		return "0"
	}

	stack := []int{} // save idx
	for i := range num {
		for len(stack) > 0 && num[stack[len(stack)-1]] > num[i] &&
			(l-i >= r-len(stack)+1) {
			// the condition is
			// 1. stack not empty, of course
			// 2. stack top is bigger than num[i]
			// 3. there is still enough to form r (requried) size; it is a bit confusing
			// so l-i-1 is the left element; r-(len(stack)-1) is after popping one element, the needed size to for requried size
			// ^ one mistake, I haven't push ith onto stack, so should be l-i to including ith element in the reminant
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	res := make([]byte, r)
	for i := range res {
		res[i] = num[stack[i]]
	}

	resStr := strings.TrimLeft(string(res), "0")
	if len(resStr) == 0 {
		return "0"
	}

	return resStr
}

/*
Runtime: 4 ms, faster than 70.49% of Go online submissions for Remove K Digits.
Memory Usage: 7.6 MB, less than 24.59% of Go online submissions for Remove K Digits.
*/
func testRemoveKDigits() {
	fmt.Println(removeKdigits("12312312312340", 10))
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
}
