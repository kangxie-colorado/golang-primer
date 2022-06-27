// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/

/*
analysis:
	at first sight, quite unsure how to start
	then aha... the total and the parts

	the ask is not to get all the subsets? right?
	the ask is to judge if I can get this ... so at least there is a quick termination...

	if total-sum != ??, well, what
		if total-sum % k != 0, then not possible

	if total-sum % k == 0, then this becomes a game of searching parts that can yield a sum of total-sum/k

	then we sort it?
	[4,3,2,3,5,2,1] => [1,2,2,3,3,4,5], k=4, total-sum = 20
	then we are looking for parts that can do 5

	5 itself..
	1/4
	2/4
	2/3/3
	and we know there are 4 groups at most... so we keep a count, this ease the last group a little bit

	what if the choice? how to worry about that? do we need to worry about that?
	make or break situation?

	also the false-positive [1, 99] k=10... how to fight away these
	well, this is a bad example.. and actually points out another short cut... k is mis-usued as a sum, while k is the number of groups... so when len() < k, it is not possible either.

	so total-sum%k==0 && len() >= k
		actually, when len==k, there can only be one situation.. every elements are the same
		also this 1 <= k <= nums.length <= 16... so the total length is not super

		because of this 1 <= nums[i] <= 104 so also if any element is greater than total/k, then it is an termination automatically too

	then this kind of reveal the strategy
		starting from the right..
		if the rightmost > targe, false
		if ==, it is itself a subset
		if <, the game is on
			looking at the left.. if it can find a compliment number, then subset
			if it cannot, the this number no matter how you place it, the subset will be bigger than target; placing it alone, the subset will be smaller
			cool!!

		so what happened? by sorting and looking to the right... and find the compliment number for bigger number, I take away the need to think about the placement of small number, that can form better solutions or not? not sure
		but indeed it turns into a make/break... situation. the anchor is now on the bigger number
		should be enough to start coding

canPartitionKSubsets(nums, k)
	sum := sum(nums)
	sort(nums)

	if sum%k != 0
		return false
	target := sum/k

	// len()>=k is implicit, no need to worry
	fromLeft := 0
	for i:=len(nums)-1;i>=0;i--
		if nums[i] > target
			return false

		if nums[i] == target
			continue

		sumToTarget := nums[i]

		for j:=fromLeft;j<i;j++
			sumToTarget += nums[j]
			if sumToTarget > target
				return false
			if sumToTarget == target
				fromLeft = j+1
				break

	return true


*/

package main

import (
	"fmt"
	"sort"
)

func wrong_canPartitionKSubsets(nums []int, k int) bool {
	sum := sum(nums)
	if sum%k != 0 {
		return false
	}
	target := sum / k

	sort.Ints(nums)

	fromLeft := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > target {
			return false
		}

		if nums[i] == target {
			continue
		}

		sumToTarget := nums[i]
		for j := fromLeft; j < i; j++ {
			sumToTarget += nums[j]
			if sumToTarget > target {
				return false
			}

			if sumToTarget == target {
				fromLeft = j + 1
				break
			}

		}

	}

	return true
}

/**
first submission
94 / 159 test cases passed.
Input:
[4,4,6,2,3,8,10,2,10,7]
4
Output:
false
Expected:
true

sum=56, target=14

[2, 2, 3, 4, 4, 6, 7, 8, 10, 10]
okay.. so there is indeed placement choise issue

so the algorithm is not right...

I think when it blows up, I should keep searching instead of declare failure?
or change the problem into
I am looking for a 4.. then maintain the slice on the way out

canPartitionKSubsets(nums, k)
	sum := sum(nums)
	sort(nums)

	if sum%k != 0
		return false
	target := sum/k

	// len()>=k is implicit, no need to worry
	for i:=len(nums)-1;i>=0;i--
		if nums[i] > target
			return false

		if nums[i] == target
			continue

		// now the anchor is still the bigger number and priciple is everything must go
		// smaller number has the placement choise, but the big number doesn't?
		if !lookingFor(target - nums[i], nums, i)
			return false

	return true


lookingFor(target, nums, start, end) -> return nums and bool
	if target in nums:
		nums.delete(target)
		return nums, true

	// else, if target can be summed up using the parts.. then okay
	// now I think why there is a length constraint..

	for nums[start] == 0
		// this number has been already used
		// choose next number
		start++

	return lookingFor(target - nums[i], nums, start+1. end)


	// no need to shift nums around
	// change the element to 0
	// using this constraint: 1 <= nums[i] <= 10^4

**/

func lookFor(target, start, end int, nums []int) bool {
	if nums[start] > target {
		return false
	}

	for i, n := range nums[start:end] {
		if n == target {
			nums[i+start] = 0
			return true
		}
	}

	for nums[start] == 0 {
		start++
	}

	ret := false
	// backtracking
	for next := start + 1; next < end; next++ {
		if lookFor(target-nums[start], next, end, nums) {
			nums[start] = 0
			ret = true
			break
		}
	}

	return ret
}

func _1_canPartitionKSubsets(nums []int, k int) bool {
	sum := sum(nums)
	if sum%k != 0 {
		return false
	}
	target := sum / k
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			// this number has been extracted from left side
			continue
		}

		if nums[i] > target {
			return false
		}

		if nums[i] == target {
			continue
		}

		successForI := false
		for j := 0; j < i; j++ {
			if nums[j] == 0 {
				continue
			}
			if lookFor(target-nums[i], j, i, nums) {
				successForI = true
				break
			}
		}

		if !successForI {
			return false
		}

	}

	return true
}

/**
125 / 159 test cases passed.
[730,580,401,659,5524,405,1601,3,383,4391,4485,1024,1175,1100,2299,3908]
4
Output:
false
Expected:
true

okay.. must be something else I don't know about yet...

[3, 383, 401, 405, 580, 659, 730, 1024, 1100, 1175, 1601, 2299, 3908, 4391, 4485, 5524]

looking for first compliment 1643... no way..
but I started from 3... I shall retry from 2nd number

here=>
        if !lookFor(target-nums[i], 0, i, nums) {
            return false
        }


rewrite into this loop
		for j := 0; j < i; j++ {
			if lookFor(target-nums[i], j, i, nums) {
				break
			}
		}

		^^^ if solved, break.
			if not, retry from 2nd smallest number until exhausted

okay. I see where the problem is
the backtracking isn't right... you cannot simpley look over one by one... sometimes you need to jump
	ret := false
	if lookFor(target-nums[start], start+1, end, nums) {
		nums[start] = 0
		ret = true
	}

change into
	ret := false
	for next:=start+1; next<end; next++ {
		if lookFor(target-nums[start], start+lookingNext, end, nums) {
			nums[start] = 0
			ret = true
		}
	}




**/

/**

148 / 159 test cases passed.
Status: Wrong Answer
Submitted: 0 minutes ago
Input:
[2,2,2,2,3,4,5]
4
Output:
true
Expected:
false

hmm... this one, should be taken care
when 4, the smallest is 2.. it should already blow out

okay.. here

        successForI := false
		for j := 0; j < i; j++ {
			if lookFor(target-nums[i], j, i, nums) {
				successForI = true
				break
			}
		}

		if !successForI {
			return false
		}

if a number cannot be solved, then it should alraedy be false...
by far, this algorithem is very messy... think later to cleanse it and optimize

**/

/**

Wrong Answer
Details
Input
[1,1,1,1,2,2,2,2]
2
Output
false
Expected
true


okay.. missing  a break here

	// backtracking
	for next := start + 1; next < end; next++ {
		if lookFor(target-nums[start], next, end, nums) {
			nums[start] = 0
			ret = true
			break
		}
	}

when you found it... you should not continue to search other solutions and consume more numbers

**/

/**

155 / 159 test cases passed.
Status: Wrong Answer
Submitted: 0 minutes ago
Input:
[709,374,1492,1279,2848,6337,365,1923,739,1904,1938,4627,1106,5885,1202]

[365, 374, 709, 739, 1106, 1202, 1279, 1492, 1904, 1923, 1938, 2848, 4627, 5885, 6337]
4
Output:
false
Expected:
true

okay.. this is a placement problem
if to match 6337, I choose 365/374/1106, then I cannot get 5885 done

thought? maybe also start from the bigger end looking for compliment???
that way, I can see 739 is 365+374 and capture it first...

or change a way
add the bigger number to all numbers. and minuse by target, see the remainder exists or not...
mightbe a way yes...

anyway, done for today
what I learn... I quickly realize the backtracking and the ability to see on each layer of recursion ....
so not a vain


**/

func lookForFromRight(target, start int, nums []int) bool {
	if start < 0 {
		return false
	}

	if nums[start] > target {
		return false
	}

	for i := start; i >= 0; i-- {
		if nums[i] == target {
			nums[i] = 0
			return true
		}
	}

	for start >= 0 && nums[start] == 0 {
		start--
	}

	ret := false
	// backtracking
	for next := start; next >= 0; next-- {
		if lookForFromRight(target-nums[start], next-1, nums) {
			nums[start] = 0
			ret = true
			break
		}
	}

	return ret
}

func _2_canPartitionKSubsets(nums []int, k int) bool {
	sum := sum(nums)
	if sum%k != 0 {
		return false
	}
	target := sum / k
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			// this number has been extracted from left side
			continue
		}

		if nums[i] > target {
			return false
		}

		if nums[i] == target {
			nums[i] = 0
			continue
		}

		numscpy := make([]int, len(nums))
		for i, n := range nums {
			numscpy[i] = n
		}

		successForI := false
		for j := i - 1; j >= 0; j-- {
			if nums[j] == 0 {
				continue
			}
			if lookForFromRight(target-nums[i], j, nums) {
				successForI = true
				break
			}
		}

		if !successForI {
			// look from right fails
			// maybe there is false negative because of placement...
			// let me just try my luck, add the look from left

			for j := 0; j < i; j++ {
				if numscpy[j] == 0 {
					continue
				}
				if lookFor(target-numscpy[i], j, i, numscpy) {
					successForI = true
					break
				}
			}

			if !successForI {
				return false

			}
		}

	}

	return true
}

/**


158 / 159 test cases passed.
Status: Wrong Answer
Submitted: 0 minutes ago
Input:
[4,5,9,3,10,2,10,7,10,8,5,9,4,6,4,9]
5
Output:
false
Expected:
true

okay... looking from the right is not all good
let me gambly add if it fails, look from the left again

something diabolic

		if !successForI {
			// look from right fails
			// maybe there is false negative because of placement...
			// let me just try my luck, add the look from left

			for j := 0; j < i; j++ {
				if nums[j] == 0 {
					continue
				}
				if lookFor(target-nums[i], j, i, nums) {
					successForI = true
					break
				}
			}

			if !successForI {
				return false

			}
		}

^^^ okay... from left, from right, none work
		but it has been 158/159 case..
		my god... this problem is so damn difficult

		now think this constrain
		The frequency of each element is in the range [1, 4].





**/

/*
		okay, I admit I cannot fingure this out
		let me study the working algorithm to understand what is going on tomorrow

class Solution:
    def canPartitionKSubsets(self, nums, k):
        sums = [0]*k
        subsum = sum(nums) / k
        nums.sort(reverse=True)
        l = len(nums)

        def walk(i):
            if i == l:
                return len(set(sums)) == 1
            for j in range(k):
                sums[j] += nums[i]
                if sums[j] <= subsum and walk(i+1):
                    return True
                sums[j] -= nums[i]
                if sums[j] == 0:
                    break
            return False

        return walk(0)


I stared at this for quite a while to understand
until I print a lot of debug log I wasn't at a loss

so this is the classical backtrack
pick every number (sorted reversely, so the general direction is same to mine), and see which bucket it could be put into

each bucket should become subsum(sum/k) eventually if it works out or there is no solution
so the attempt will be like

num[0] -> sums(0,1,2,3)
	there is no better choice so just follow the sequence order
	if bucket 1 can hold it, then hold
		and continue to try num[1], on bucket1
		if it can hold it, then hold
			and continue to try num[2], on bucket1
			if it cannot hold it, back it out (backtracking)

			try num[2] on bucket2...
			and such such

notice the ending condition is the nums are all tried and filled... if they recursion reached len deep, then it is a pass
also notice the backtracking... and continue to fill the still-ok buckets until it becomes not-ok

the backtracking will undo the addition, that happened on the same level.... always.. recursion is focused on the same level
the very important thing is when a solution won't work out, the backtracking will have a bucket backout to eventually 0, and that is a time to return false
also when all the buckets are tried, no solution... that is a false

only when all buckets are filled equally, it return true.. and the triggers a cascading chain of return true...
this is very very subtle but kind of beautiful...

I think I can handle this one now...
hope I can spot the related variants

*/

// for simplicity, use some global varirables
var l int
var sums []int
var target int
var K int
var nums []int

func canParHelper(i int) bool {
	if i == l {
		// reached the end
		return true
	}

	for k := 0; k < K; k++ {
		sums[k] += nums[i]
		if sums[k] <= target && canParHelper(i+1) {
			// the equal to target is also an important bit to set up moving to next
			// if only <, then = will be become false.. and everything will not work
			return true
		}
		sums[k] -= nums[i]

		if sums[k] == 0 {
			break
		}
	}

	return false
}

func correct_canPartitionKSubsets(ns []int, k int) bool {
	// set up the global variables
	l = len(ns)
	totalSum := sum(ns)
	target = totalSum / k
	K = k

	sums = make([]int, k)
	nums = ns

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	/*
		with sort, wow
		Success
		Details
		Runtime: 2 ms, faster than 100.00% of Go online submissions for Partition to K Equal Sum Subsets.
		Memory Usage: 2.2 MB, less than 38.30% of Go online submissions for Partition to K Equal Sum Subsets.
	*/
	if totalSum%k != 0 {
		return false
	}

	return canParHelper(0)
}

/**
	okay, after submission I saw many more different solution..
	interesting
	lets pull some down and study




**/

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	sum /= k

	set := make(map[int][]int)
	set[0] = []int{0}

	fmt.Println(nums)
	for i, n := range nums {

		sumsWithN := make(map[int][]int)
		for s := range set {

			sumsWithN[s+n] = make([]int, 0)
			for _, v := range set[s] {
				fmt.Println("n s v i", n, s, v, i)

				if s+n <= sum {
					sumsWithN[s+n] = append(sumsWithN[s+n], 1<<i|v)
				}
			}
		}

		// fmt.Println("i:", i, "n:", n, "sumsWithN:", sumsWithN, "set", set)
		fmt.Println("sumWithN", sumsWithN)
		for sn := range sumsWithN {
			for _, v := range sumsWithN[sn] {
				set[sn] = append(set[sn], v)
			}
		}
		fmt.Println("set", set)
		fmt.Println("")

	}

	//fmt.Println(set)
	fmt.Println("")
	if sums, ok := set[sum]; ok {
		set := make(map[int]struct{})
		set[0] = struct{}{}
		for _, v := range sums {
			sumsWithN := make(map[int]struct{})
			for s := range set {
				//fmt.Println("v:", v, "s:", s, "sumsWithN:", sumsWithN, "s|v", s|v)
				if s&v == 0 {
					sumsWithN[s|v] = struct{}{}
				}
			}
			for sn := range sumsWithN {
				set[sn] = struct{}{}
			}
			fmt.Println(set)
		}
		_, ok = set[(1<<len(nums))-1]
		return ok
	}
	return false
}

/**
WHAT THE FUCK IS THIS??????
完全没有办法看懂
**/

func testPartitionKSubsets() {
	// true expected
	// [2, 2, 3, 4, 4, 6, 7, 8, 10, 10]
	fmt.Println(canPartitionKSubsets([]int{1, 3}, 2))

	/*
		fmt.Println(canPartitionKSubsets([]int{2, 2, 3, 4, 4, 6, 7, 8, 10, 10}, 4)) // true


		fmt.Println(canPartitionKSubsets([]int{730, 580, 401, 659, 5524, 405, 1601, 3, 383, 4391, 4485, 1024, 1175, 1100, 2299, 3908}, 4))
		fmt.Println(canPartitionKSubsets([]int{709, 374, 1492, 1279, 2848, 6337, 365, 1923, 739, 1904, 1938, 4627, 1106, 5885, 1202}, 4))
		fmt.Println(canPartitionKSubsets([]int{1, 1, 1, 1, 2, 2, 2, 2}, 2))
		fmt.Println(canPartitionKSubsets([]int{4, 5, 9, 3, 10, 2, 10, 7, 10, 8, 5, 9, 4, 6, 4, 9}, 5))

		// false expected
		fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 4))
		fmt.Println(canPartitionKSubsets([]int{2, 2, 2, 2, 3, 4, 5}, 4))
	*/

}
