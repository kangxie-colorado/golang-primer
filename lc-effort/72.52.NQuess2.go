// https://leetcode.com/problems/n-queens-ii/

package main

import "fmt"

/*
hahaha I just listened on neetcode about this one and I am shuffled to it

so let me solve it
that trick using a single number to represent a diag line is geius
but thinking of that

a row can be represent by a single number
a col can too

so why we cannot think of a posDiag or a negDiag this way
this should be applicable to all matrix shape.. not only square

we can think placing queens on rows while using them as column
we can also think queues as rows and placing them in column... it makes no differece
because this is suare but if it is not sqaure... it probably will

this will just working like queens naturally on row 1-n.. so picking/choosing col for them
*/

var res int

func _1_totalNQueens(n int) int {
	res = 0
	colTaken := make(map[int]struct{})
	posDiagTaken := make(map[int]struct{}) // r+c
	negDiagTaken := make(map[int]struct{}) // r-c

	backtrackingQueen(0, n, colTaken, posDiagTaken, negDiagTaken)
	return res
}

func backtrackingQueen(row, n int, colTaken, posDiagTaken, negDiagTaken map[int]struct{}) {
	// placing
	if row == n {
		res += 1
		return
	}

	// every queue has possible n choices
	// loop thru each one..
	// if one spot doesn't work.. backout and try next
	for i := 0; i < n; i++ {
		_, foundCol := colTaken[i]
		_, foundPos := posDiagTaken[row+i]
		_, foundNeg := negDiagTaken[row-i]

		if foundCol || foundPos || foundNeg {
			continue
		}

		// found one that can be placed for queens so far
		colTaken[i] = struct{}{}
		posDiagTaken[row+i] = struct{}{}
		negDiagTaken[row-i] = struct{}{}
		backtrackingQueen(row+1, n, colTaken, posDiagTaken, negDiagTaken)

		delete(colTaken, i)
		delete(posDiagTaken, row+i)
		delete(negDiagTaken, row-i)
	}

}

/*
Runtime: 11 ms, faster than 7.32% of Go online submissions for N-Queens II.
Memory Usage: 1.9 MB, less than 46.05% of Go online submissions for N-Queens II.

the map is slow
let me use array for those set

colTaken obviously n
how about diags

1 1
1 1

acutally I don't care how many
because it is r+c,r-c
let me just get the biggest range

should be really 2*n-1
and then how to mapped onto it

r+c is easy to map to 0+0 to (n-1)+(n-1)
r-c
the right-top corner will be the smallest, 0-(n-1) mapped to 0, so just += -= n-1

*/

func totalNQueens(n int) int {
	res = 0
	colTaken := make([]int, n)
	posDiagTaken := make([]int, 2*n-1) // r+c
	negDiagTaken := make([]int, 2*n-1) // r-c

	backtrackingQueen2(0, n, colTaken, posDiagTaken, negDiagTaken)
	return res
}

func backtrackingQueen2(row, n int, colTaken, posDiagTaken, negDiagTaken []int) {
	// placing
	if row == n {
		res += 1
		return
	}

	// every queue has possible n choices
	// loop thru each one..
	// if one spot doesn't work.. backout and try next
	for i := 0; i < n; i++ {
		if colTaken[i] != 0 || posDiagTaken[row+i] != 0 || negDiagTaken[row-i+n-1] != 0 {
			continue
		}

		// found one that can be placed for queens so far
		colTaken[i] = 1
		posDiagTaken[row+i] = 1
		negDiagTaken[row-i+n-1] = 1
		backtrackingQueen2(row+1, n, colTaken, posDiagTaken, negDiagTaken)

		colTaken[i] = 0
		posDiagTaken[row+i] = 0
		negDiagTaken[row-i+n-1] = 0
	}

}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for N-Queens II.
Memory Usage: 1.9 MB, less than 46.05% of Go online submissions for N-Queens II.
*/

func testNQueens2() {
	fmt.Println(totalNQueens(1))
	fmt.Println(totalNQueens(2))

}
