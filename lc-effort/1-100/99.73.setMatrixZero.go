// https://leetcode.com/problems/set-matrix-zeroes/

package main

/*
either I don't understand this or this is too simple
i just need keep a row prod and col prod, if it is 0 then on next iteration
setting the whole row/colomn to zero
*/

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	rowProds := make([]int, rows)
	colProds := make([]int, cols)

	for r := 0; r < rows; r++ {
		rowProds[r] = 1
	}
	for c := 0; c < cols; c++ {
		colProds[c] = 1
	}

	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == 0 {
				rowProds[r] = 0
				colProds[c] = 0
			}
		}
	}

	for r := range matrix {
		for c := range matrix[r] {
			if rowProds[r] == 0 || colProds[c] == 0 {
				matrix[r][c] = 0
			}
		}
	}
}

/*
hmm... error...
so it is actually overflow... when the matrix gets really big

change to
	for r := range matrix {
		for c := range matrix[r] {
            if matrix[r][c] == 0 {
                rowProds[r]=0
                colProds[c]=0
            }

		}
	}

Runtime: 25 ms, faster than 20.43% of Go online submissions for Set Matrix Zeroes.
Memory Usage: 6.6 MB, less than 30.88% of Go online submissions for Set Matrix Zeroes.

Runtime: 8 ms, faster than 96.20% of Go online submissions for Set Matrix Zeroes.
Memory Usage: 6.3 MB, less than 70.55% of Go online submissions for Set Matrix Zeroes.

ah right
the mark can just use first row/col
            for(int col =1; col<n; col++) {
                if(matrix[row][col] == 0) {
                    matrix[row][0] = 0;
                    matrix[0][col] = 0;
                }
            }
save some space...

but much more trickier to get it right
first pass, top-down
2nd pass, bottom-up

use first row and first col
but easy to mess up.. laster zero or previously set zeros..
*/
