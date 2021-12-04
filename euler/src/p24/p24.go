/**
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically,
we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
**/

package main

import (
	"fmt"
	"sort"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

// from this permuation, we try to swap two numbers
// after the swap if the end slice is bigger, then it becomes a candidate
// we will search all the candidates to identify the next one(the closest one)
func nextPermutation(perm []int) []int {
	perm_copy := make([]int, len(perm))
	perm_candidate := make([]int, len(perm))

	for i := range perm_candidate {
		// maxmize the initial candidate
		perm_candidate[i] = 9
	}

	right_idx := len(perm) - 1
	final_swap_idx := -1
	for ; right_idx >= 0; right_idx-- {
		swap_idx := right_idx - 1
		// start another search starting from more left index, need to restore the original
		copy(perm_copy, perm)

		for ; swap_idx >= 0; swap_idx-- {
			// swap two numbers
			perm_copy[swap_idx], perm_copy[right_idx] = perm[right_idx], perm[swap_idx]

			// if we found a bigger permutation, then stop..
			// otherwise we found a smaller permutaion, continue on
			if libs.CompareTwoNumSlices(perm_copy, perm) == 1 {
				break
			} else {
				// restore the original copy
				copy(perm_copy, perm)
			}

		}

		if swap_idx >= 0 {
			// we find a bigger permutaion candidate, there could be other candidates which can swap from more left indexs
			// they may be bigger than original perm but smaller than current candidate, so a closer next
			// so save this as a candidate and prepare to swap it
			if libs.CompareTwoNumSlices(perm_candidate, perm_copy) == 1 {
				copy(perm_candidate, perm_copy)
				final_swap_idx = swap_idx
			}

		}

		// if swap_idx < 0, then we actually are at the max permutation for swapping begging at this right_idx
		// we should then start from one index to the left
	}

	// if we indeed find such a bigger permutaion, sort the perm_copy[right:-1] to the smallest
	if final_swap_idx >= 0 {
		sorted_suffix := perm_candidate[final_swap_idx+1 : len(perm)]
		sort.Ints(sorted_suffix)
		return append(perm_candidate[0:final_swap_idx+1], sorted_suffix...)
	}

	return nil

}

func nextNPermutaion(perm []int, n int) []int {
	res := perm
	for i := 1; i <= n; i++ {
		res = nextPermutation(res)
	}

	return res
}

func main() {
	nextPermutation([]int{0, 3, 4, 2, 1})

	// start from the very bigning: very slow, bad performace
	/**
	perm := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 1000000; i++ {
		fmt.Println(nextNPermutaion(perm, i))
	}
	**/

	// think:
	// P9 = 362880, so
	// start from {1, 0-9}, only do 1m - 362880
	// start from {2, 0-9}, only do 1m - 2*362880 = 274,240
	// P8 = 40320, so
	// start from {2,1,0-9}, only do 274,240 - 40,320
	// start from {2,3,0-9}, only do 274,240 - 40,320*2
	/***
	>>> 274240//40320
	6
	>>> 274240 - 40320*6
	32320
	   start from {2,7,0-9}, only do 274,240 - 40,320*6 = 32320

	   P7 = 5040, so
	   start from {2,7,1,0-9}, only do 32320 - 5040
	   start from {2,7,3,0-9}, only do 32320 - 5040*2
	>>> 32320//5040
	6
	>>> 32320 - 6*5040
	2080

	   start from {2,7,8,0-9}, only do 32320 - 5040*6 = 2080, good enough to crack with program

	***/
	// Correct Answer: 2783915460

	perm := []int{2, 7, 8, 0, 1, 3, 4, 5, 6, 9}
	for i := 0; i < 2080; i++ {
		fmt.Println(nextNPermutaion(perm, i)) // [2 7 8 3 9 1 5 4 6 0]
	}
}
