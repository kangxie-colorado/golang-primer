// https://leetcode.com/problems/subarray-product-less-than-k/

package main

import "fmt"

/*
	gosh.. come on! don't jump straight into calculate the longest/maxest subarray length
	read the description

	so this is asking all the contiguous subarray that meet the criterea
		(this contiguous is definitely a sign for sliding window)

	so this is not asking for biggest..
	this need to grow and shrink, basically need to cover all the element in single element fashion

	it is a bit tricky
	so maybe this will work
	 	nums = [10,5,2,6], k = 100
		[10] +1									[10]
		[10,5] +2 why? [10,5] is, [5] is    	[10,5] [5]
		[10,5,2] false, shrink [5,2] +2, why 	[5,2] [2]
		[5,2,6] +3, why 						[5,2,6] [2,6] [6]
		so you can see, it is to plus the length of the subarray that meets the condition
		if one more 1
		[5,2,6,1] +4 							[5,2,6,1] [2,6,1] [6,1] [1]

		yes, this verifies using the console run (leetcode)



*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	prod := 1
	i, j := 0, 0
	for j < len(nums) {
		prod *= nums[j]

		for prod >= k && i < len(nums) {

			prod /= nums[i]
			i++
		}
		if i <= j {
			count += j - i + 1
		}

		j++

	}

	return count
}

/*
Wrong Answer
Details
Input
[10,9,10,4,3,8,3,3,6,2,10,10,9,3]
19
Output
24
Expected
18

ugh.. it should be for loop
		for prod >= k {

			prod /= nums[i]
			i++
		}

		I typed a "if" there

[1,2,3]
0
still missing this and maybe some other case..

need to guard it

		for prod >= k && i < len(nums) {

			prod /= nums[i]
			i++
		}
		if i <= j {
			count += j - i + 1
		}

	how ugly... but it passes

Runtime: 141 ms, faster than 38.02% of Go online submissions for Subarray Product Less Than K.
Memory Usage: 7.9 MB, less than 39.58% of Go online submissions for Subarray Product Less Than K.
Next challenges:

Runtime: 83 ms, faster than 90.10% of Go online submissions for Subarray Product Less Than K.
Memory Usage: 7.3 MB, less than 84.90% of Go online submissions for Subarray Product Less Than K.

the ugliness can go away like

		for prod >= k && i < len(nums) {

			prod /= nums[i]
			i++
		}
		if i <= j {
			count += j - i + 1
		}


	to
		for prod >= k && i <=j {

			prod /= nums[i]
			i++
		}

		count += j - i + 1
*/

func testNumSubarrayProductLessThanK() {

	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3, 4}, 0))
	fmt.Println(numSubarrayProductLessThanK([]int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, 19))
	fmt.Println(numSubarrayProductLessThanK([]int{10, 9, 10, 4, 3, 8, 3, 3, 6, 2, 10, 10, 9, 3}, 0))
}
