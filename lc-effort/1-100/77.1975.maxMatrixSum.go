// https://leetcode.com/problems/maximum-matrix-sum/

package main

import "math"

/*
at first sight, there is zero idea
and jumped into analysis of the 4 ways.. dynamic.. blabla
and getting no where

so had to check the hints
- Try to use the operation so that each row has only one negative number.
- If you have only one negative element you cannot convert it to positive.

so realize
1. divide and conquer, and it should not work in row/col way; not the ocuptus way..
2. the negative sign can be passed along, anywhere
3. two negative sign will cancel eachother eventually

so I can do a column or row based swith..
force the negative to one element onto a single element or no element(if even numbers are negative)

then from another direction, check thur all slot where the negative numbers are forced onto.. and see
possible biggest..

but now when I am typing this
I realize the negative sign can transition in a row, in a column, it can actually travel to anywhere
so they can also cancel each other...


so very cool indeed,
just going thru all the numbers, adding their abs together and remember the smallest one..
if it has to zero, minus 2*that-number-abs
*/

func maxMatrixSum(matrix [][]int) int64 {
	res := int64(0)
	sign := 1
	smallest := math.MaxInt

	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] < 0 {
				sign = -sign
			}
			smallest = min(smallest, abs(matrix[r][c]))
			res += int64(abs(matrix[r][c]))
		}
	}

	if sign < 0 {
		res -= int64(2 * smallest)
	}

	return res
}

/*
Runtime: 240 ms, faster than 22.22% of Go online submissions for Maximum Matrix Sum.
Memory Usage: 8.6 MB, less than 55.56% of Go online submissions for Maximum Matrix Sum.
*/
