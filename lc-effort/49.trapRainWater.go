// https://leetcode.com/problems/trapping-rain-water/

package main

/*
I actually solved this before but maybe with many pains..
now I know it is just to use stack

also the trick is to divide the calculation into horizanal rects
do't over do.. keep it simple and generic...

luckily I am able to see this thru now...

*/

func trap(height []int) int {
	stack := []int{} // maintain the indexes
	res := 0

	for i, n := range height {
		// check if anything to the left can be popped out
		// meaning, any low point that could trap water
		// note the stack is keeping the indexes.. so stack[len(stack)-1] which is peek() is basically index
		// need to height[] the index to get the height
		for len(stack) > 0 && height[stack[len(stack)-1]] < n {
			lowHeight := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				nextLow := height[stack[len(stack)-1]]
				// nextLow must be greater than or equal to lowHeight
				// but if it is equal, we don't need to calculate just yet
				// wait for its turn (in next loop run, actually this internal for loop)
				if nextLow > lowHeight {
					// i - stack[len(stack)-1] - 1 is the horizonal rect width
					// the height is min(nextLow-lowHeight, n-lowHeight)
					// the width*height is the area, the water trapped
					res += min(nextLow-lowHeight, n-lowHeight) * (i - stack[len(stack)-1] - 1)
				}
			}

		}

		stack = append(stack, i)
	}

	return res
}

/*
Runtime: 8 ms, faster than 87.37% of Go online submissions for Trapping Rain Water.
Memory Usage: 6.7 MB, less than 5.13% of Go online submissions for Trapping Rain Water.
*/
