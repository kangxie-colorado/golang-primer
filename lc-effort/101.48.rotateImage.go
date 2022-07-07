// https://leetcode.com/problems/rotate-image/

package main

/*
the memory comes back
because this is in-place, so it is difficult to manage the rotating
so we need to do swapping

the direct swapping is not that easy
and notice this is a n*n matrix

so we can first do swap between row/col
then we fold

m[r][c] = m[c][r]

then for any row
m[r][c] = m[r][cols-c-1]
*/

func rotate(matrix [][]int) {
	for r := range matrix {
		for c := 0; c < r/2; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

	for row := range matrix {
		for l, r := 0, len(matrix[row])-1; l < r; l, r = l+1, r-1 {
			matrix[row][l], matrix[row][r] = matrix[row][r], matrix[row][l]
		}
	}
}

/*
Wrong Answer
Runtime: 3 ms
Your input
[[1,2,3],[4,5,6],[7,8,9]]
Output
[[3,2,1],[6,5,4],[9,8,7]]
Expected
[[7,4,1],[8,5,2],[9,6,3]]

hahah...

first step swapped twice and it went back
only need to swap the lower-left half (with the upper-top half)

		for c := 0; c < r; c++ {

previsouly tried
		for c := 0; c < len(matrix); c++ { <== double swap
		for c := 0; c < len()/2; c++ {		<== no idea where it comes from



Runtime: 5 ms, faster than 13.51% of Go online submissions for Rotate Image.
Memory Usage: 2.3 MB, less than 72.50% of Go online submissions for Rotate Image.

people doing this

            topLeft := matrix[top][l+i]
            matrix[top][l+i] = matrix[bottom-i][l]
            matrix[bottom-i][l] = matrix[bottom][r-i]
            matrix[bottom][r-i] = matrix[top+i][r]
            matrix[top+i][r] = topLeft

	apparently that is not smart

2017 I used the upper half
    	for(int i=0; i<matrix.size(); ++i) {
            for(int j=i+1; j<matrix.size(); ++j) {
                swap(matrix[i][j],matrix[j][i]);
            }
    	}
*/
