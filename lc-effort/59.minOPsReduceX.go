// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/

package main

import (
	"fmt"
	"math"
)

/*
at first sight, really no idea how to tackle it
then this minimum key word reminds me binary search possible

when K operations can do it...
then k+1 definitely can... well... not necessarily

so there is some monotonicity

so given K operations, I can go greedy to find a sum for X
a few possible outcomes

1. bingo.. then K becomes a candidate, try next  (don't have to be exact k ops, anything <= k)
2. end up smaller
3. end up bigger??? probably no solution

hmm... doesn't seem so right
maybe just a backtracking

now fucking rest now
*/

func _1_minOperations(nums []int, x int) int {
	if x == 0 {
		return 0
	}

	if len(nums) == 0 || x < 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	opLeft := minOperations(nums[l+1:], x-nums[l])
	opRight := minOperations(nums[:r], x-nums[r])

	if opLeft == -1 && opRight == -1 {
		return -1
	}

	if opLeft == -1 {
		return 1 + opRight
	}

	if opRight == -1 {
		return 1 + opLeft
	}

	return 1 + min(opLeft, opRight)
}

/*
Wrong Answer
Details
Input
[8828,9581,49,9818,9974,9869,9991,10000,10000,10000,9999,9993,9904,8819,1231,6309]
134365
Output
-1
Expected
16

didn't go too far
>>> a=[8828,9581,49,9818,9974,9869,9991,10000,10000,10000,9999,9993,9904,8819,1231,6309]
>>>
>>> sum(a)
134365

so.. there is an edge case somewhere
the order of early returns

	if x == 0 {
		return 0
	}

	if len(nums) == 0 || x < 0 {
		return -1
	}

should test the success first
cause when nums are finished, x becomes zero that should be a success

now I got TLE


10 / 93 test cases passed.
Status: Time Limit Exceeded
Submitted: 0 minutes ago
Last executed input:
[1241,8769,9151,3211,2314,8007,3713,5835,2176,8227,5251,9229,904,1899,5513,7878,8663,3804,2685,3501,1204,9742,2578,8849,1120,4687,5902,9929,6769,8171,5150,1343,9619,3973,3273,6427,47,8701,2741,7402,1412,2223,8152,805,6726,9128,2794,7137,6725,4279,7200,5582,9583,7443,6573,7221,1423,4859,2608,3772,7437,2581,975,3893,9172,3,3113,2978,9300,6029,4958,229,4630,653,1421,5512,5392,7287,8643,4495,2640,8047,7268,3878,6010,8070,7560,8931,76,6502,5952,4871,5986,4935,3015,8263,7497,8153,384,1136]
894887480

of course, maybe memorization can help a bit
*/

func minOperationsHelper(nums []int, x int) int {
	if x == 0 {
		return 0
	}

	if len(nums) == 0 || x < 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	if v, found := m[pair{l, r}]; found {
		return v
	}

	opLeft := minOperations(nums[l+1:], x-nums[l])
	opRight := minOperations(nums[:r], x-nums[r])

	if opLeft == -1 && opRight == -1 {
		return -1
	}

	if opLeft == -1 {
		return 1 + opRight
	}

	if opRight == -1 {
		return 1 + opLeft
	}

	return 1 + min(opLeft, opRight)
}

var m map[pair]int

func _1_backtracking_with_memorization_minOperations(nums []int, x int) int {
	m = make(map[pair]int)
	return minOperationsHelper(nums, x)

}

/*
okay.. this is not going to be okay
O(n!)

let me think otherwise...
back to the binary search...

I think I am stuck enough
check the hints

- Think in reverse; instead of finding the minimum prefix + suffix, find the maximum subarray.
- Finding the maximum subarray is standard and can be done greedily.

WTF...

think in reverse....

(The 2nd day)
must admit, I failed to come up with any efficient solution to this one
I was trapped deep and cannot see

I briefly touch partial sum... but after one week of grinding and some real-time distractions I had nothing left
if I had more energy, I don't know if I am able to see the solution a bit better

but anyway... where I fall, where I get up

let me evolve these solutions and code it up...


*/
func _2_fix_one_end_minOperations(nums []int, x int) int {

	lSumMap := make(map[int]int)
	lSumMap[0] = 0 // this generize right end
	res := math.MaxInt

	sum := 0
	for i := 0; i < len(nums) && sum < x; i++ {
		sum += nums[i]
		lSumMap[sum] = i + 1
		if sum == x {
			res = i + 1
		}
	}

	sum = 0
	for j := len(nums) - 1; j >= 0 && sum < x; j-- {
		sum += nums[j]
		if c, found := lSumMap[x-sum]; found {
			res = min(res, c+len(nums)-j)
		}
	}

	// ah.. cannot let both left/right reuse a same one
	if res > len(nums) {
		return -1
	}

	return res
}

/*
Runtime: 515 ms, faster than 7.69% of Go online submissions for Minimum Operations to Reduce X to Zero.
Memory Usage: 14.9 MB, less than 7.69% of Go online submissions for Minimum Operations to Reduce X to Zero.

*/

/*
okay, another is from the hints, the sliding window...
which I wasn't able to come up too...

actually I think I solved a max consecutive 2 problem??? not sure
anyway, it seems that non-shrinkable template would not be able to work well here...

becase when the window <= X, it might be < not =, so cannot measure
but anyway the shrinkable solution

- grow the window as long as the sum < x
- when =, take a candidate
- when >, shink to until it is smaller again
*/

func minOperations(nums []int, x int) int {
	res := math.MaxInt
	sum := 0

	for _, n := range nums {
		sum += n
	}
	target := sum - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}

	sum = 0
	for i, j := 0, 0; j < len(nums); {
		sum += nums[j]
		for sum > target {
			sum -= nums[i]
			i++
		}

		j++
		if sum == target {
			res = min(res, len(nums)-(j-i))
		}

	}

	if res > len(nums) {
		return -1
	}

	return res
}

/*
Runtime: 168 ms, faster than 76.92% of Go online submissions for Minimum Operations to Reduce X to Zero.
Memory Usage: 9.9 MB, less than 38.46% of Go online submissions for Minimum Operations to Reduce X to Zero.
*/

func testMinOperations() {
	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	//fmt.Println(minOperations([]int{1241, 8769, 9151, 3211, 2314, 8007, 3713, 5835, 2176, 8227, 5251, 9229, 904, 1899, 5513, 7878, 8663, 3804, 2685, 3501, 1204, 9742, 2578, 8849, 1120, 4687, 5902, 9929, 6769, 8171, 5150, 1343, 9619, 3973, 3273, 6427, 47, 8701, 2741, 7402, 1412, 2223, 8152, 805, 6726, 9128, 2794, 7137, 6725, 4279, 7200, 5582, 9583, 7443, 6573, 7221, 1423, 4859, 2608, 3772, 7437, 2581, 975, 3893, 9172, 3, 3113, 2978, 9300, 6029, 4958, 229, 4630, 653, 1421, 5512, 5392, 7287, 8643, 4495, 2640, 8047, 7268, 3878, 6010, 8070, 7560, 8931, 76, 6502, 5952, 4871, 5986, 4935, 3015, 8263, 7497, 8153, 384, 1136}, 894887480))
}
