// https://leetcode.com/problems/maximum-product-subarray/

package main

import (
	"fmt"
	"math"
)

/*
this is one of blind 75
I think the smart DP is keeping the min/max.. along with the prefix
meeintg 0, the prefix will becomes 1 again

I would be interested to see what others come up with
*/

func _missing_cases_maxProduct(nums []int) int {
	prefix := 1
	prod := 1

	minP := math.MaxInt
	maxP := math.MinInt

	for _, num := range nums {
		prod = prefix * num
		minP = min(minP, prod)
		maxP = max(maxP, prod)
		prefix = prod
		if num == 0 {
			prefix = 1
		}
	}

	return maxP
}

/*
Wrong Answer
Details
Input
[3,-1,4]
Output
3
Expected
4

okay... I am not being careful
hahaha.. even I watched video I cannot solve this on one go
*/

func _still_wrong_maxProduct(nums []int) int {
	prefix, posPrefix := 1, 1
	prod, posProd := 1, 1

	minP := math.MaxInt
	maxP := math.MinInt

	for _, num := range nums {
		prod = prefix * num
		posProd = posPrefix * num // can be negative even I name it posProd

		minP = min(minP, prod)
		maxP = max(maxP, max(posProd, prod))
		prefix = prod
		posPrefix = max(prefix, 1)
		if num == 0 {
			prefix = 1
			posPrefix = 1
		}
	}

	return maxP
}

/*
Wrong Answer
Details
Input
[2,-5,-2,-4,3]
Output
20
Expected
24

gosh.. am I stupid?
I am being lazy...

I thought I watched/listened the video.. and I knew the answer and I let my alert down

well... you have slide the window all the same...
probably means I need to use extra memory..

after some working
I found I was being too lazy and thought I knew it by heart

then I found
to keep a min and max so far, which will be initialized to nums[0] and start with nums[1]

minP = min(minP*num, maxP*num, num)
maxP = max(minP*num, maxP*num, num)

it will cover all the cases.
*/

func maxProduct(nums []int) int {
	minP, maxP := nums[0], nums[0]
	res := maxP

	for i := 1; i < len(nums); i++ {
		oldMin := minP
		minP = min(minP*nums[i], min(maxP*nums[i], nums[i]))
		maxP = max(oldMin*nums[i], max(maxP*nums[i], nums[i]))
		res = max(res, maxP)
	}

	return res
}

/*
Runtime: 8 ms, faster than 56.88% of Go online submissions for Maximum Product Subarray.
Memory Usage: 3.5 MB, less than 18.74% of Go online submissions for Maximum Product Subarray.

and this legendary one

https://leetcode.com/problems/maximum-product-subarray/discuss/183483/JavaC%2B%2BPython-it-can-be-more-simple

prefix
suffix..

but how it works?

   def maxProduct(self, A):
        B = A[::-1]
        for i in range(1, len(A)):
            A[i] *= A[i - 1] or 1
            B[i] *= B[i - 1] or 1
        return max(A + B)
			^ note it is not A[i]+B[i], it is the biggest in two array

This is really great solution. But I don't recommand anybody to use it in interivew. It will waste of you lots of time to explain to interviewer why it is right.

I was asked this question today by a Stupid Amazon interviewer. I give two solutions for both DP and this one. Neighter of two solutions he can understand. Write code 10 mins, explain for 40 mins.

but hey.. this..
so you do have to explain to the interviewer... so the drawing is a necessary part
be patient to drive to the end when practicing

*/

func testMaxProduct() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))      // 6
	fmt.Println(maxProduct([]int{2, 0, -1}))         // 2
	fmt.Println(maxProduct([]int{-2, 0, -1}))        // 0
	fmt.Println(maxProduct([]int{3, -1, 4}))         // 4
	fmt.Println(maxProduct([]int{2, -5, -2, -4, 3})) // 24
}
