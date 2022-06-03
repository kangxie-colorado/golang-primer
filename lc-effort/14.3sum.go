package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/3sum/

/*
	I remember my previous attempt didn't go well enough
	so I looked at the hints.
	1. So, we essentially need to find three numbers x, y, and z such that they add up to the given value. If we fix one of the numbers say x, we are left with the two-sum problem at hand!
	2. For the two-sum problem, if we fix one of the numbers, say
		x
		, we have to scan the entire array to find the next number
		y
		which is
		value - x
		where value is the input parameter. Can we change our array somehow so that this search becomes faster?
	3. The second train of thought for two-sum is, without changing the array, can we use additional space somehow? Like maybe a hash map to speed up the search?

	this hash map hits my sparking point let me do that

threeSums(nums)
	# build the map
	for n in nums:
		m[n]++

	for k,v in m:
		m[k]--
		sumfor2:=0-k
		for k2,v2 in m:
			if v2==0:
				continue
			m[k2]--
			if (sumfor2-k2) in m and m[sumfor2-k2]!=0:
				# a triple if found

			m[k2++] # put it back since it is reusable but not reusable on the same triplet

		m[k++]

*/

type Triple struct {
	n1, n2, n3 int
}

func _1_threeSum(nums []int) [][]int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	tripleMap := make(map[Triple]struct{})

	for n := range numMap {
		numMap[n]--
		sumFor2 := 0 - n

		for n2, c2 := range numMap {
			if c2 == 0 {
				continue
			}

			numMap[n2]--
			if c3, found := numMap[sumFor2-n2]; found && c3 != 0 {
				triplet := []int{n, n2, sumFor2 - n2}
				fmt.Println(triplet)
				sort.Ints(triplet)
				tripleMap[Triple{triplet[0], triplet[1], triplet[2]}] = struct{}{}
			}
			numMap[n2]++

		}
		//numMap[n]++
		delete(numMap, n)
	}

	triplets := [][]int{}
	for k, _ := range tripleMap {
		triplets = append(triplets, []int{k.n1, k.n2, k.n3})
	}

	return triplets

}

/*
Success
Details
Runtime: 588 ms, faster than 17.17% of Go online submissions for 3Sum.
Memory Usage: 9.3 MB, less than 13.74% of Go online submissions for 3Sum.
Next challenges:

wow.. I never solved this problem in 2017... how bad was I?
maybe I can just remove the first number... since it has been used and appeared in all possible combinations to reduce the space

okay.. marginally better
Success
Details
Runtime: 403 ms, faster than 19.89% of Go online submissions for 3Sum.
Memory Usage: 9.2 MB, less than 15.43% of Go online submissions for 3Sum.
Next challenges:


without println
Success
Details
Runtime: 374 ms, faster than 21.03% of Go online submissions for 3Sum.
Memory Usage: 10.3 MB, less than 7.05% of Go online submissions for 3Sum.
Next challenges:

so there must lie a better solution that reduce the complextiy by one more order

on nex thought, I didn't combine sorted-array and hashmap together
let me do that


*/

func threeSum(nums []int) [][]int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	sort.Ints(nums)
	triplets := [][]int{}
	tripleMap := make(map[Triple]struct{})

	for i, n := range nums {
		if numMap[n] == 0 {
			continue
		}

		if n == -43 {
			fmt.Println("Debug")
		}

		sumFor2 := 0 - n
		numMap[n]--

		for j := i + 1; j < len(nums); j++ {
			numMap[nums[j]]--
			if j > i+1 && nums[j] == nums[j-1] {
				numMap[nums[j]]++
				continue
			}

			if c3, found := numMap[sumFor2-nums[j]]; found && c3 > 0 {
				triplet := []int{n, nums[j], sumFor2 - nums[j]}
				fmt.Println(triplet)
				sort.Ints(triplet)
				tripleMap[Triple{triplet[0], triplet[1], triplet[2]}] = struct{}{}

			}

			numMap[nums[j]]++
		}
		numMap[n] = 0

	}

	for k, _ := range tripleMap {
		triplets = append(triplets, []int{k.n1, k.n2, k.n3})
	}

	return triplets

}

/*

111 / 318 test cases passed.
Status: Wrong Answer
Submitted: 2 minutes ago
Input:
[34,55,79,28,46,33,2,48,31,-3,84,71,52,-3,93,15,21,-43,57,-6,86,56,94,74,83,-14,28,-66,46,-49,62,-11,43,65,77,12,47,61,26,1,13,29,55,-82,76,26,15,-29,36,-29,10,-70,69,17,49]
Output:
[[-49,21,28],[-14,-3,17],[-82,21,61],[-66,-11,77],[-66,10,56],[-70,13,57],[-49,-3,52],[-49,1,48],[-49,2,47],[-49,15,34],[-82,17,65],[-82,34,48],[-70,-6,76],[-14,1,13],[-14,2,12],[-43,10,33],[-29,-14,43],[-49,13,36],[-43,12,31],[-82,13,69],[-66,-3,69],[-66,17,49],[-29,1,28],[-70,1,69],[-70,15,55],[-43,-6,49],[-82,-11,93],[-70,21,49],[-70,34,36],[-49,-6,55],[-29,12,17],[-82,33,49],[-70,-14,84],[-66,1,65],[-43,17,26],[-11,-6,17],[-11,1,10],[-3,1,2],[-82,26,56],[-82,36,46],[-43,-14,57]]
Expected:
[[-82,-11,93],[-82,13,69],[-82,17,65],[-82,21,61],[-82,26,56],[-82,33,49],[-82,34,48],[-82,36,46],[-70,-14,84],[-70,-6,76],[-70,1,69],[-70,13,57],[-70,15,55],[-70,21,49],[-70,34,36],[-66,-11,77],[-66,-3,69],[-66,1,65],[-66,10,56],[-66,17,49],[-49,-6,55],[-49,-3,52],[-49,1,48],[-49,2,47],[-49,13,36],[-49,15,34],[-49,21,28],[-43,-14,57],[-43,-6,49],[-43,-3,46],[-43,10,33],[-43,12,31],[-43,15,28],[-43,17,26],[-29,-14,43],[-29,1,28],[-29,12,17],[-14,-3,17],[-14,1,13],[-14,2,12],[-11,-6,17],[-11,1,10],[-3,1,2]]

fixed a bug..
			if j > i+1 && nums[j] == nums[j-1] {
				numMap[nums[j]]++		// missing this, so some num's count go negative
				continue
			}

but still ain't that good

Success
Details
Runtime: 475 ms, faster than 18.74% of Go online submissions for 3Sum.
Memory Usage: 9.4 MB, less than 13.74% of Go online submissions for 3Sum.
Next challenges:

without print
Success
Details
Runtime: 311 ms, faster than 22.90% of Go online submissions for 3Sum.
Memory Usage: 8.2 MB, less than 27.37% of Go online submissions for 3Sum.

*/

func test3Sum() {
	fmt.Println(threeSum([]int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}))
}
