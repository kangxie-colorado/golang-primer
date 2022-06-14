// https://leetcode.com/problems/maximize-distance-to-closest-person/

package main

/*
analysis
	this seems to be a dfs problem
	starting from each 1 and calculate the distances to ones for those zeros

	so it stops on
	1. out of bounds left/right
	2. another 1
	3. when dist in flight is bigger than a previously calculated dist

	should be not too hard

*/

func _1_maxDistToClosest(seats []int) int {

	dists := make([]int, len(seats))
	for d := range dists {
		dists[d] = len(seats)
	}

	for s := range seats {
		if seats[s] == 1 {
			dists[s] = 0
			stack := []int{s}
			for len(stack) > 0 {
				node := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				left := node - 1
				right := node + 1

				if left >= 0 && seats[left] != 1 && dists[left] > dists[node]+1 {
					dists[left] = dists[node] + 1
					stack = append(stack, left)
				}
				if right < len(seats) && seats[right] != 1 && dists[right] > dists[node]+1 {
					dists[right] = dists[node] + 1
					stack = append(stack, right)
				}
			}

		}
	}

	maxD := 0
	for _, d := range dists {
		maxD = max(maxD, d)
	}

	return maxD
}

/*
Runtime: 23 ms, faster than 17.54% of Go online submissions for Maximize Distance to Closest Person.
Memory Usage: 6.1 MB, less than 14.04% of Go online submissions for Maximize Distance to Closest Person.

Runtime: 15 ms, faster than 56.14% of Go online submissions for Maximize Distance to Closest Person.
Memory Usage: 6.3 MB, less than 7.02% of Go online submissions for Maximize Distance to Closest Person.

or course I don't really need to do dfs... I just need to traverse the array back and forth

*/

func maxDistToClosest(seats []int) int {

	dists := make([]int, len(seats))
	for d := range dists {
		dists[d] = len(seats)
	}

	for s := range seats {
		if seats[s] == 1 {
			dists[s] = 0
			// look to left
			for j := s - 1; j >= 0; j-- {
				if seats[j] == 1 || dists[j] <= dists[j+1]+1 {
					break
				}
				dists[j] = dists[j+1] + 1
			}
			// look to right
			for i := s + 1; i < len(seats); i++ {
				if seats[i] == 1 || dists[i] <= dists[i-1]+1 {
					break
				}
				dists[i] = dists[i-1] + 1
			}
		}
	}

	maxD := 0
	for _, d := range dists {
		maxD = max(maxD, d)
	}

	return maxD
}

/*
Runtime: 11 ms, faster than 71.93% of Go online submissions for Maximize Distance to Closest Person.
Memory Usage: 6 MB, less than 28.07% of Go online submissions for Maximize Distance to Closest Person.

good enough
*/
