// https://leetcode.com/problems/two-city-scheduling/

package main

import "math"

/*
I feel this should be a heap problem
but I cannot make it happen yet..

then I see there might be a backtracking(brute force) solution, let me try that first
*/

func twoCitySchedCost(costs [][]int) int {
	minCost := math.MaxInt
	n := len(costs) / 2

	var backtrack func(A, B []int)
	backtrack = func(A []int, B []int) {
		if len(A) == n && len(B) == n {
			// calculate the cost and return
			cost := 0
			for i := range A {
				cost += A[i] + B[i]
			}

			minCost = min(minCost, cost)
			return
		}

		next := len(A) + len(B)
		if len(A) < n {
			A = append(A, costs[next][0])
			backtrack(A, B)
			A = A[:len(A)-1]
		}

		if len(B) < n {
			B = append(B, costs[next][1])
			backtrack(A, B)
			B = B[:len(B)-1]
		}

	}

	backtrack([]int{}, []int{})
	return minCost

}

/*
example passed.. let me see
yeah... failed here

[[555,594],[403,796],[203,847],[292,885],[525,478],[327,970],[297,483],[540,102],[855,412],[174,684],[591,837],[431,385],[107,740],[433,854],[472,205],[862,439],[961,598],[440,574],[716,156],[202,976],[170,661],[661,823],[867,44],[372,606],[8,281],[705,704],[48,6],[909,266],[675,515],[484,749]]
TLE

>>> a=[[555,594],[403,796],[203,847],[292,885],[525,478],[327,970],[297,483],[540,102],[855,412],[174,684],[591,837],[431,385],[107,740],[433,854],[472,205],[862,439],[961,598],[440,574],[716,156],[202,976],[170,661],[661,823],[867,44],[372,606],[8,281],[705,704],[48,6],[909,266],[675,515],[484,749]]
>>> len(a)
30


okay, the sort by diff is hard to see but that is cool
thinking it like sending all people to A, then get max refund (or minimal extra cost) by sending half to B...

wow... how smart it is!!!

then this p1[0] - p1[1] < p2[0] - p2[1] ==> p1[0]+p2[1] < p2[0] + p1[1]
so it means sending p1 to A and p2 to B will be cheaper

(in the sorted array, diff1( p1[0] - p1[1]) will be at first half, diff2 will be in the right half)
first half goes to city B...

mother fucker...

also I failed to see there is indeed a DP solution

but overall I think I should not take a fucking break...
stop leetcoding for a couple of days.. maybe better

then pickup here again...
*/
