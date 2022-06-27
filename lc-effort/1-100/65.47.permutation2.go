// https://leetcode.com/problems/permutations-ii/

/*
there should be quite a few ways to do this
key word in this one is duplicate, that makes the swap/sort a bit hard?

let me try backtracking

I'll appoint the first num, swap it with 4 pos - which is fix one pos, let the rest shuffle
do this recursively, I should get some answers
*/

package main

import "fmt"

func permuteUniqueHelper(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := [][]int{}
	swapped := make(map[int]struct{})
	for i := range nums {
		if _, found := swapped[nums[i]]; found {
			continue
		}

		nums[i], nums[0] = nums[0], nums[i]
		for _, p := range permuteUniqueHelper(nums[1:]) {
			cpy := make([]int, len(nums))
			cpy[0] = nums[0]
			copy(cpy[1:], p)

			res = append(res, cpy)
		}
		nums[0], nums[i] = nums[i], nums[0]

		swapped[nums[i]] = struct{}{}
	}

	return res
}

/*
Runtime: 10 ms, faster than 29.07% of Go online submissions for Permutations II.
Memory Usage: 6.7 MB, less than 12.68% of Go online submissions for Permutations II.
Next challenges:

cool...

the use of map at each layer solved the duplicate issue

Runtime: 4 ms, faster than 75.67% of Go online submissions for Permutations II.
Memory Usage: 6.7 MB, less than 12.68% of Go online submissions for Permutations II.
*/

func permuteUnique(nums []int) [][]int {
	res := permuteUniqueHelper(nums)

	return res
}

func testPermuteUnique() {
	fmt.Println(permuteUnique([]int{1, 1, 2, 2}))
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 2, 3, 4}))
}
