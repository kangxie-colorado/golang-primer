// https://leetcode.com/problems/monotone-increasing-digits/

package main

import "fmt"

/*
cannot think out a smart way to make it go fast.
let me just testing from the top down, and test it monotonicity

if this brute force won't pass then think where it breaks the monotonicity
and work from there...
*/

func monotoneIncreasingDigits(n int) int {
	isMonotone := func(n int) bool {
		current := 10
		next := n % 10

		for n > 0 && current >= next {
			n /= 10
			current = next
			next = n % 10
		}

		return n == 0
	}

	digits := []int{}
	getDigits := func(num int) {
		for num > 0 {
			digits = append(digits, num%10)
			num /= 10
		}
	}

	getDigits(n)
	nextSmallMonotone := func() int {
		i := len(digits) - 1
		for i-1 >= 0 && digits[i-1] >= digits[i] {
			i--
		}

		for i+1 < len(digits) && digits[i+1] == digits[i] {
			i++
		}

		digits[i] -= 1
		for j := 0; j < i; j++ {
			digits[j] = 9
		}

		sum := 0
		for k := len(digits) - 1; k >= 0; k-- {
			sum = sum*10 + digits[k]
		}

		return sum
	}

	if isMonotone(n) {
		return n
	}

	return nextSmallMonotone()
}

/*
ugh.. break it here..
554889396

but this number isn't particularly big or special?
okay.. purely this alone Runtime: 443 ms
okay.. I need to optimize..

find where the monotonicy is broken
apparently to find it is broken, I need to start from the left...

still TLE..

then I found I maybe doing too much
actually the break point actually reveals the answer?
1232 -> 1229

and.. yeah...
Runtime: 5 ms, faster than 15.38% of Go online submissions for Monotone Increasing Digits.
Memory Usage: 1.9 MB, less than 76.92% of Go online submissions for Monotone Increasing Digits.

Runtime: 0 ms, faster than 100.00% of Go online submissions for Monotone Increasing Digits.
Memory Usage: 2 MB, less than 76.92% of Go online submissions for Monotone Increasing Digits.

tough to beat this...
*/

func testMonotoneIncreasingDigits() {
	fmt.Println(monotoneIncreasingDigits(1234))
}
