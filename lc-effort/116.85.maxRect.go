// https://leetcode.com/problems/maximal-rectangle/

package main

/*
so there is a previous proble that solves the max rect in histogram
so this is an extention of that..

basically what can be done is go thru the row
and calculate on this row so far the max rect there is

of cource, you have pre-process the matrix to stack the numbers up
*/

func largestRectangleArea(heights []int) int {
	stack := []int{} // index
	smallerToRight := make([]int, len(heights))
	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			smallerToRight[top] = i
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		smallerToRight[top] = len(heights)
	}

	stack = []int{} // index
	smallerToLeft := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			smallerToLeft[top] = i
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		smallerToLeft[top] = -1
	}

	res := 0
	for i := range heights {
		res = max(res, (smallerToRight[i]-smallerToLeft[i]-1)*heights[i])
	}

	return res
}

func maximalRectangle(matrix [][]byte) int {
	prefix := make([]int, len(matrix[0]))
	numMatrix := make([][]int, len(matrix))
	maxRec := 0

	for r := range matrix {
		numMatrix[r] = make([]int, len(matrix[0]))
		for c := range matrix[r] {
			if matrix[r][c] == '0' {
				numMatrix[r][c] = 0
			} else {
				numMatrix[r][c] = 1
			}

			if numMatrix[r][c] != 0 {
				numMatrix[r][c] += prefix[c]
			}
			prefix[c] = numMatrix[r][c]
		}
		maxRec = max(maxRec, largestRectangleArea(numMatrix[r]))
	}

	return maxRec
}
