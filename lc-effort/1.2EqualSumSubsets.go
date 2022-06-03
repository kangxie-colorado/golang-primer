// https://leetcode.com/problems/partition-equal-subset-sum/

/*
analysis:
	at first sight, I came from k-equal-sum-subsets problem
	I thought I just provide a k=2 to that..

	but however the nums can be 200 in length so the space is much bigger
	it will time out

	then I think, for 2 subsets.. it is really to look for one.. found one, the other one is settled.
	can actually use my lookfor, which is also a backtracking

	I named it sumTo
*/

package main

import (
	"fmt"
	"sort"
)

func _sumTo(target, start, end int, nums []int, level int) bool {
	if nums[start] > target {
		return false
	}

	fmt.Println(level, "   ", target, start, end)

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
		if _sumTo(target-nums[start], next, end, nums, level+1) {
			nums[start] = 0
			ret = true
			break
		}
	}

	return ret
}

func removeRedundantNums(nums []int) []int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n]++
	}

	ret := []int{}
	for k, v := range m {
		if v%2 == 1 {
			ret = append(ret, k)
		}
	}

	return ret
}

func _canPartition(nums []int) bool {

	totalSum := sum(nums)
	if totalSum%2 != 0 {
		return false
	}
	target := totalSum / 2

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	return _sumTo(target, 0, len(nums), nums, 0)
}

/**

36 / 117 test cases passed.
Status: Time Limit Exceeded
Submitted: 0 minutes ago
Last executed input:
[100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,99,97]

not working on long arrays...
but this array really has something that can be optimized away.. if there are two same number.. then remove it

^^^ stupid idea
	you cannot do it.. you are taking away the eligible placement..
	lesson learned: tricks are usually less useful in these problems


so as great as backtracking is... it can only take us this far
what else to do? early termination... really not applicable

so I thought harder on the constraint 1 <= nums[i] <= 100
and also the case study.. that it has many duplicated values...

so lets divide the numbers into buckets, number and the occurences of it

100 : 198
99 : 1
97 :1

start from the biggest,
>>> sum(a)
19996
>>> sum(a)/2
9998.0

9998 at most can use 99*100, then use 99.. find the rest which is 98.. not able to
then use 98, find the rest 198... (and notice not in the 100 buckets...), not able to

if this exhausted... then no
==

hmm... no backtracking in this
so try a different one

11 : 1
5 : 2
1 : 1

this is naturally

==

[1,2,3,4]

4:1
3:1
2:1
1:1

yeah// it would need some backtrack on the rest part?

...

anyway... this seems more feasible
AND NOW: DISPLINE...
THE DISPLINE TO DO, and equally important!!! is DISPLINE TO REST...

NOW REST>....

The basketball is too ugly...
come to do some code

type NumCountPair struct {
	num int
	count int
}


sumTo(sofar, target, i, nums)
	if sofar == target
		return true

	use_of_i = min(nums[i].count, (target-sofar)/nums[i].num)
	for use_of_i >= 0
		sofar +=  use_of_i * nums[i].num
		if sofar < target and sumTo(sofar, target, i+1, nums)
			return true
		sofar -= use_of_i * nums[i].num
		use_of_i--

	return false

can_partition(nums)
	sum = sum(nums)
	if sum%2 != 0
		return false

	target = sum/2

	nums_count_sorted := nums_to_bucket(nums) # turn into something like [pair{100,198}, pair{99,1}, pair{997,1}]
	sumTo(0, target, 0, nums)



**/

type numCountPair struct {
	num   int
	count int
}

type byNum []numCountPair

func (s byNum) Len() int {
	return len(s)
}

func (s byNum) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byNum) Less(i, j int) bool {
	return s[i].num < s[j].num
}

type byCount []numCountPair

func (s byCount) Len() int {
	return len(s)
}

func (s byCount) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byCount) Less(i, j int) bool {
	return s[i].count < s[j].count
}

func sumTo(sofar, target, i int, nums []numCountPair) bool {
	if sofar == target {
		return true
	}

	if i >= len(nums) {
		return false
	}

	useOfI := min(nums[i].count, (target-sofar)/nums[i].num)
	// must use >=, because use 0 count is also a valid choice, otherwise, it jumps to return false, and disrupt the whole algorithm
	for useOfI >= 0 {
		sofar += useOfI * nums[i].num
		if sofar <= target {

			if sumTo(sofar, target, i+1, nums) {
				return true
			}
			// this is not equal to
			// return sumTo(sofar, target, i+1, nums)
			// on false, we are not returning... we need to backtrack

		}

		sofar -= useOfI * nums[i].num
		useOfI--

	}

	return false
}

func getNumCountSorted(nums []int) []numCountPair {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	numCountPairs := make([]numCountPair, len(m))
	i := 0
	for k, v := range m {
		numCountPairs[i] = numCountPair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(byNum(numCountPairs)))
	//sort.Sort(sort.Reverse(byCount(numCountPairs)))
	//sort.Sort(byNum(numCountPairs))
	return numCountPairs
}

func backtracking_canPartition(nums []int) bool {
	sumTotal := sum(nums)
	if sumTotal%2 != 0 {
		return false
	}
	target := sumTotal / 2

	numCountSorted := getNumCountSorted(nums)
	return sumTo(0, target, 0, numCountSorted)
}

/**

111 / 117 test cases passed.
Status: Wrong Answer
Submitted: 0 minutes ago
Input:
[23,13,11,7,6,5,5]
Output:
false
Expected:
true

^^^ logically right, although it passed more cases. it is worse than the logically right and passes 37/117 case
okay... this problem is super interesting and I truly have no thoughts...

there is a great editorial explaining everything
https://leetcode.com/problems/partition-equal-subset-sum/discuss/462699/Whiteboard-Editorial.-All-Approaches-explained.

let me learn!!!
Patience is now

OKAY.. only this case 38 I cannot pass, by tricking it... this algorithm is working fine on other test cases
Success
Details
Runtime: 7 ms, faster than 98.41% of Go online submissions for Partition Equal Subset Sum.
Memory Usage: 2.8 MB, less than 78.49% of Go online submissions for Partition Equal Subset Sum.
Next challenges:


**/

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 > 0 {
		return false
	}

	target := sum / 2
	for range nums {
		ans := 0
		fmt.Println("1", nums)
		for _, num := range nums {
			fmt.Println(num, ans)
			if num+ans <= target {
				ans += num
			}
			fmt.Println(num, ans)

			if ans == target {
				return true
			}
		}
		fmt.Println("2", nums)

		nums = append(nums, nums[0])
		nums = nums[1:]
		fmt.Println("3", nums)

	}

	return false
}

/**
How does this one work?
**/

func test2Paritions() {

	//fmt.Println(canPartition([]int{1, 2, 5}))
	//fmt.Println(canPartition([]int{1, 2, 1}))

	fmt.Println(canPartition([]int{4, 4, 4, 4, 4, 4, 4, 4, 8, 8, 8, 8, 8, 8, 8, 8, 12, 12, 12, 12, 12, 12, 12, 12, 16, 16, 16, 16, 16, 16, 16, 16, 20, 20, 20, 20, 20, 20, 20, 20, 24, 24, 24, 24, 24, 24, 24, 24, 28, 28, 28, 28, 28, 28, 28, 28, 32, 32, 32, 32, 32, 32, 32, 32, 36, 36, 36, 36, 36, 36, 36, 36, 40, 40, 40, 40, 40, 40, 40, 40, 44, 44, 44, 44, 44, 44, 44, 44, 48, 48, 48, 48, 48, 48, 48, 48, 52, 52, 52, 52, 52, 52, 52, 52, 56, 56, 56, 56, 56, 56, 56, 56, 60, 60, 60, 60, 60, 60, 60, 60, 64, 64, 64, 64, 64, 64, 64, 64, 68, 68, 68, 68, 68, 68, 68, 68, 72, 72, 72, 72, 72, 72, 72, 72, 76, 76, 76, 76, 76, 76, 76, 76, 80, 80, 80, 80, 80, 80, 80, 80, 84, 84, 84, 84, 84, 84, 84, 84, 88, 88, 88, 88, 88, 88, 88, 88, 92, 92, 92, 92, 92, 92, 92, 92, 96, 96, 96, 96, 96, 96, 96, 96, 97, 99}))

	// true
	fmt.Println(canPartition([]int{23, 13, 11, 7, 6, 5, 5}))
	//fmt.Println(canPartition([]int{1, 2, 3, 4}))
	//fmt.Println(canPartition([]int{1, 2, 3, 4, 4, 4, 4}))

	// false
	//fmt.Println(canPartition([]int{1, 2, 3, 5}))

	//fmt.Println(canPartition([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97}))
}
