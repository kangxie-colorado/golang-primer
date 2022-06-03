// https://leetcode.com/problems/divide-two-integers/

package main

import "fmt"

func divide(dividend int, divisor int) int {
	sign := 1
	dividend64 := int64(dividend)
	divisor64 := int64(divisor)

	if dividend < 0 {

		dividend64 = 0 - int64(dividend)

		sign = 0 - sign
	}
	if divisor < 0 {
		divisor64 = 0 - int64(divisor)
		sign = 0 - sign
	}

	quotient64 := int64(0)

	for dividend64 > 0 {
		num := divisor64
		shift := 0
		for num <= dividend64 {
			shift++
			num = divisor64 << shift

		}

		//fmt.Println(shift)

		if shift == 0 {
			break
		}

		dividend64 -= divisor64 << (shift - 1)
		quotient64 += 1 << (shift - 1)

	}

	res := int(quotient64) * sign
	if int64(sign)*quotient64 > 2147483647 {
		res = int(2147483647)
	}

	if int64(sign)*quotient64 < -2147483648 {
		res = int(-2147483648)
	}

	return res

}

/*
Success
Details
Runtime: 10 ms, faster than 45.15% of Go online submissions for Divide Two Integers.
Memory Usage: 2.4 MB, less than 55.22% of Go online submissions for Divide Two Integers.
Next challenges:

Success
Details
Runtime: 5 ms, faster than 59.33% of Go online submissions for Divide Two Integers.
Memory Usage: 2.4 MB, less than 55.22% of Go online submissions for Divide Two Integers.
Next challenges:
*/

func testDivide() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
	fmt.Println(divide(1, 1))
	fmt.Println(divide(2147483648, 1))
	fmt.Println(divide(-2147483648, 1))
	fmt.Println(divide(-2147483648, -1))
}
