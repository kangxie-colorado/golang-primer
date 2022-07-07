// https://leetcode.com/problems/maximal-square/

package main

/*
https://leetcode.com/problems/maximal-square/
https://leetcode.com/problems/maximal-rectangle/


sqaure vs rect..
only one little difference, the solution is totally different

you can solve it like the rect, but it will be even harder than that already hard problem
so there is of course the DP solution to square...

I was wondering why the rect problem cannot be solved by DP?
yeah.. it invovles two dimensional states.. and for squre somehow it only has to maintain one dimesional state?

anyway the idea is start from right-bottom (or left-up, let me do this... less naturall actually but practice)

if m[r][c] == 0, then max squre ending here can be only 0
if m[r][c] == 1
	then if m[r-1][c] == 1 and m[r][c-1] == 1
	max sqare here could be max(1, (sqrt(m[r-1][c-1])+1)^2) # 1 is the 1 itself, and 2nd part is too complicated, so just keep the length
	^ actually this hypothesis is not right
	to form a 3*3, the r-1,c-1 should be 2*2; r-1,c should be 2*2; r,c-1 should be 2*2
	so it should be min(m[r-1][c-1], m[r-1][c], m[r][c-1]) + 1 really

*/

func _DP_maximalSquare(matrix [][]byte) int {
	maxSquare := make([][]int, len(matrix)+1)
	maxSide := 0
	for r := range maxSquare {
		maxSquare[r] = make([]int, len(matrix[0])+1)
		if r == 0 {
			continue
		}
		for c := range maxSquare[r] {
			if c == 0 {
				continue
			}

			if matrix[r-1][c-1] == '1' {
				maxSquare[r][c] = 1 // by itself
				if maxSquare[r-1][c] > 0 && maxSquare[r][c-1] > 0 && maxSquare[r-1][c-1] > 0 {
					maxSquare[r][c] = min(maxSquare[r-1][c], min(maxSquare[r][c-1], maxSquare[r-1][c-1])) + 1
				}

				maxSide = max(maxSide, maxSquare[r][c])
			}
		}
	}
	return maxSide * maxSide
}

/*
Runtime: 7 ms, faster than 76.25% of Go online submissions for Maximal Square.
Memory Usage: 6.5 MB, less than 60.54% of Go online submissions for Maximal Square.

notice we only need this row and last row..
so space can be saved a bit more

let me try

it actaully follows a pattern when I got the matrix vs matrix version under my pocket
thisRow --> replace maxSquare[r]
lastRow --> replace maxSquare[r-1]

rotate thisRow/lastRow like this
		lastRow := maxSquare[r%2]
		thisRow := maxSquare[(r+1)%2]

*/

func maximalSquare(matrix [][]byte) int {
	maxSquare := make([][]int, 2)
	maxSquare[0] = make([]int, len(matrix[0])+1)
	maxSquare[1] = make([]int, len(matrix[0])+1)
	maxSide := 0
	for r := range matrix {
		lastRow := maxSquare[r%2]
		thisRow := maxSquare[(r+1)%2]

		for c := range thisRow {
			if c == 0 {
				continue
			}

			if matrix[r][c-1] == '1' {
				thisRow[c] = 1 // by itself
				if lastRow[c] > 0 && thisRow[c-1] > 0 && lastRow[c-1] > 0 {
					thisRow[c] = min(lastRow[c], min(thisRow[c-1], lastRow[c-1])) + 1
				}

				maxSide = max(maxSide, thisRow[c])
			} else {
				thisRow[c] = 0
			}
		}

	}
	return maxSide * maxSide
}

/*
Runtime: 3 ms, faster than 96.17% of Go online submissions for Maximal Square.
Memory Usage: 5.5 MB, less than 96.17% of Go online submissions for Maximal Square.

although it works but it is much more complicated to handle correct
1. you need to iterate thru the matrix space
	because there are only two rows so literally you cannot do the auxiliary matrix
	for r := range matrix {
2. you must explicitly set the value for each c in thisRow[c]
		else {
				thisRow[c] = 0
			}
	otherwise, lastRow's infor will be carried over.. and may fuck you up


*/
