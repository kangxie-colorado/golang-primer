// https://leetcode.com/problems/sum-of-subarray-minimums/
package main

import "fmt"

/*
this is quite hard
brute force, of course O(n^3), that may or may not cross the accept line

why don't I try??
since I also evolved the stack-based contribution style algorithm
I can hold off doing that..

let me do the brute force
yeah.. it is TLE..
*/

func _brute_force_sumSubarrayMins(arr []int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			minK := 0
			for k := i; k < j; k++ {
				minK = min(minK, arr[k])
			}

			sum += minK
			sum %= 1_000_000_007
		}
	}

	return sum
}

func sumSubarrayMins(arr []int) int {
	ltBdry := make([]int, len(arr))
	rtBdry := make([]int, len(arr))

	getBoundary := func(start, end, step int, boundary []int) {
		// get the left boundary that i-th can be the minimum value
		stack := [][]int{} // 0: idx, 1: lt-boundary-idx
		for i := start; i != end; i += step {
			newBoundary := i
			for len(stack) > 0 && (arr[stack[len(stack)-1][0]] > arr[i] ||
				(step == 1 && arr[stack[len(stack)-1][0]] == arr[i])) { // on extend on left boundary
				stackTop := stack[len(stack)-1]
				boundary[stackTop[0]] = stackTop[1]
				stack = stack[:len(stack)-1]
				newBoundary = stackTop[1]
			}

			stack = append(stack, []int{i, newBoundary})
		}
		for len(stack) > 0 {
			stackTop := stack[len(stack)-1]
			boundary[stackTop[0]] = stackTop[1]
			stack = stack[:len(stack)-1]
		}
	}

	getBoundary(0, len(arr), 1, ltBdry)
	getBoundary(len(arr)-1, -1, -1, rtBdry)

	sum := 0
	for i, n := range arr {
		l, r := ltBdry[i], rtBdry[i]
		sum += (i - l + 1) * (r - i + 1) * n
		sum %= 1_000_000_007
	}

	return sum
}

func testSumSubarrayMins() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
	fmt.Println(sumSubarrayMins([]int{3, 3, 2, 4}))
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}

/*
[3,3,2,4] failed to be correct 25, for which I got 28
I double calculated 3,3...
[3,3] alone should be 9.. which is obvious
but my algorithem will do it 12...

because I extend to right and left on equal double times..
it should be only one time

Runtime: 78 ms, faster than 54.64% of Go online submissions for Sum of Subarray Minimums.
Memory Usage: 7.7 MB, less than 54.64% of Go online submissions for Sum of Subarray Minimums.
*/
