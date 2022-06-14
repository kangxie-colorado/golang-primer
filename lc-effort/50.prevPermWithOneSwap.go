// https://leetcode.com/problems/previous-permutation-with-one-swap/

/*
	after white board evolution of next-perm I think I know how to

	1. serach from right, the decreasing sub-array, if that is the whole array, then no prev-permutation
		then the element to the left of that sub-arry(r-part) is the element to swap with
	2. swap the last element, which should be the biggest element in the sub-array with the boundary of l-part
	3. reverse sort the r-part to maint its biggest possible value
*/

package main

import (
	"fmt"
	"sort"
)

func _1_prevPermOpt1(arr []int) []int {
	// just to save up the index
	decreasing := []int{}

	for j := len(arr) - 1; j >= 0; j-- {

		if len(decreasing) == 0 || arr[j] <= arr[decreasing[len(decreasing)-1]] {
			decreasing = append(decreasing, arr[j])
		} else {
			// find j
			arr[j], arr[decreasing[0]] = arr[decreasing[0]], arr[j]
			sort.Sort(sort.Reverse(sort.IntSlice(arr[j+1:])))
		}

	}

	return arr
}

/*
okay, I over-read the problem
I don't have to reverse sort it..

and it is exactly one swap...

*/

func _2_prevPermOpt1(arr []int) []int {
	// just to save up the index
	decreasing := []int{}

	for j := len(arr) - 1; j >= 0; j-- {

		if len(decreasing) == 0 || arr[j] <= arr[decreasing[len(decreasing)-1]] {
			decreasing = append(decreasing, j)
		} else {
			// found the j
			// from the right, find first index that is smaller than arr[j]
			for d := range decreasing {
				if arr[decreasing[d]] < arr[j] {
					arr[j], arr[decreasing[d]] = arr[decreasing[d]], arr[j]
					return arr
				}
			}
		}

	}

	return arr
}

/*
Input
[3,1,1,3]
Output
[3,1,1,3]
Expected
[1,3,1,3]

I see, I need to find the smaller number but as left as I can
*/

func _3_prevPermOpt1(arr []int) []int {
	// just to save up the index
	decreasing := []int{}

	for j := len(arr) - 1; j >= 0; j-- {

		if len(decreasing) == 0 || arr[j] <= arr[decreasing[len(decreasing)-1]] {
			decreasing = append(decreasing, j)
		} else {
			// found the j
			// from the right, find first number that is smaller than arr[j]
			// then need to find the farthest one if this is a repeating number..
			for d := range decreasing {
				if arr[decreasing[d]] < arr[j] {
					swapIdx := decreasing[d]
					for dd := d + 1; dd < len(decreasing); dd++ {
						if arr[decreasing[dd]] == arr[decreasing[dd-1]] {
							swapIdx = decreasing[dd]

						} else {
							break
						}

					}
					arr[j], arr[swapIdx] = arr[swapIdx], arr[j]
					return arr
				}
			}
		}

	}

	return arr
}

/*
Runtime: 25 ms, faster than 100.00% of Go online submissions for Previous Permutation With One Swap.
Memory Usage: 7.8 MB, less than 20.00% of Go online submissions for Previous Permutation With One Swap.

even it works I feel awful
too many trivia fixing... pathcing..
not a clear path...

I'll give another try tomorrow.
maybe push some information into the decreasing

*/

func prevPermOpt1(arr []int) []int {
	// just to save up the index upto l-part
	decreasing := len(arr)

OUTER:
	for j := len(arr) - 1; j >= 0; j-- {

		if decreasing == len(arr) || arr[j] <= arr[decreasing] {
			decreasing--
		} else {
			// look from right again and see which element I should switch
			for k := len(arr) - 1; k > j; k-- {
				if arr[k] < arr[j] {
					// found the number, arr[k]
					// the find the left-most one if it is repeating
					// this part is only getting the swapidx, don't over do
					// otherwise you could be missing out

					swapIdx := k
					/*
						for kk := k - 1; kk >= j; kk-- {
							if arr[kk] != arr[kk+1] {
								swapIdx = kk + 1
								break
							}
						}
					*/
					for kk := k - 1; kk > j; kk-- {
						if arr[kk] != arr[kk+1] {
							break
						}

						swapIdx = kk
					}
					arr[j], arr[swapIdx] = arr[swapIdx], arr[j]
					break OUTER
				}
			}
		}

	}

	return arr
}

/*
Runtime: 35 ms, faster than 40.00% of Go online submissions for Previous Permutation With One Swap.
Memory Usage: 7.4 MB, less than 20.00% of Go online submissions for Previous Permutation With One Swap.

so
					for kk := k - 1; kk >= j; kk-- {
						if arr[kk] != arr[kk+1] {
							swapIdx = kk + 1
							break
						}
					}
				vs
					for kk := k - 1; kk > j; kk-- {
						if arr[kk] == arr[kk+1] {
							swapIdx = kk

						} else {
							break
						}
					}

the first one, waiting for the right moment to update, but its much easier to make a mistake and notice the ending condition for loop
kk >= j, because kk can go out of the repeating sequence... very ugly

the 2nd one, update the index each time a repeating appears, is more generic in this kind of loop
ugly but can be updated to
					for kk := k - 1; kk > j; kk-- {
						if arr[kk] != arr[kk+1] {
							break
						}

						swapIdx = kk
					}


damn, checked a few others' code, even uglier than mine
so be happy about it
*/

func testPrevPermOpt1() {
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))
	fmt.Println(prevPermOpt1([]int{3, 1, 1, 3}))
}
