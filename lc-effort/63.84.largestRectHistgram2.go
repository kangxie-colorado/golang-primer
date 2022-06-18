// https://leetcode.com/problems/largest-rectangle-in-histogram/

package main

/*
this is classical mono-stack
non-decreasing, maintaining how long the horizon can extend to

to left.. when popping out taller ones,
to right... until is it popped or all the way thru
*/

func _1_largestRectangleArea(heights []int) int {
	res := 0
	stack := [][]int{} // [start, val]

	for i, h := range heights {

		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res = max(res, (i-top[0])*top[1])
			start = top[0]
		}

		stack = append(stack, []int{start, h})

	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = max(res, (len(heights)-top[0])*top[1])
	}

	return res
}

/*
Runtime: 330 ms, faster than 13.76% of Go online submissions for Largest Rectangle in Histogram.
Memory Usage: 21.4 MB, less than 5.26% of Go online submissions for Largest Rectangle in Histogram.

interestingly, this is O(n)
but it looks like there are some faster solutions

I cannot think of any
but let me read

this left-next-smaller and right-next-smaller array is actually prettry cool too
so
							2 	1 5 6 3 4
idx							0 	1 2 3 4 5
next-smaller-to-right		1 	6 4 5 6 6
next-bigger-to-left 	   -1  -1 1 1 1 4

so for 5: to-left 1, to-right 4 --- (1,4) not inclusive on both ends
4-1-1 = 2,  2*5=10

to get the to-left/right we can use mono-stack (non-decreasing)
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

/*
Runtime: 244 ms, faster than 20.24% of Go online submissions for Largest Rectangle in Histogram.
Memory Usage: 15 MB, less than 7.29% of Go online submissions for Largest Rectangle in Histogram.

I see the guy talking about jumping to get brute force scanning O(n^2) to O(n)

how did he do that?
let me give a crack...



*/
