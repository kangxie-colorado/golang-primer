// https://leetcode.com/problems/max-consecutive-ones-iii/

package main

import (
	"fmt"
)

/*
analysis:
	this resemblies memroy recycle I guess

	brute force easy
		flip k zeros, count
		shift first flipping to next right zero... count
		until reaches end... report
		this is going to be O(n^2)

	then I have a vision to compress the arrary
	[1,1,1,0,0,0,1,1,1,1,0] => [3,0,0,0,4,0] => [3,-3,4,-1]

	[1,1,1,0,1,0,1,1,1,1,0] => [3,0,1,0,4,0] => [3,-1,1,-1,4,-1]

	back one step

	[3,0,0,0,4,0] or [3,0,1,0,4,0] the non-zero sum by flipping k zeros
	[3,0,0,0,4,0]

	[0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1]
	[0,0,2,0,3,0,2,0,0,0,4]

Did some paper evolution
start with the k=0,1 then realize just need to keep a start/end index with end being the next zero

next shift, just move from first zero next, move start/end according maybe it will be able to run thru
what states to keep
	start: the start of section, which can be zero or one..
	firstZero
	nextFirstZero
	lastZero(end): this marks the end of section.. and not included

	maybe even not needing nextFirstZero because 1 is surrounded by 0.. so it is either 0 or 1

	do some paper evolution with examples

*/

func extractZerosIdxs(nums []int) []int {
	ret := []int{}
	for i, n := range nums {
		if n == 0 {
			ret = append(ret, i)
		}
	}

	return ret
}

func flip0(nums []int) int {
	// walk the nums, find out the longest stream
	l := 0
	rl := 0
	for _, n := range nums {
		if n == 0 {
			l = max(l, rl)
			rl = 0
			continue
		}

		rl++
	}
	l = max(l, rl) // end by 1
	return l
}

func _1_longestOnes(nums []int, k int) int {
	if k == 0 {
		return flip0(nums)
	}

	// extract zeros' index
	zeros := extractZerosIdxs(nums)
	// push the len of array into the zeros idx array
	zeros = append(zeros, len(nums))
	if len(zeros)-1 < k {
		return len(nums)
	}

	//fmt.Println(len(zeros))

	l := 0

	start := 0
	for i := k + 1; i <= len(zeros); i++ {
		filps := zeros[i-k-1 : i]
		l = max(l, filps[k]-start) //the last flip is the marker of end, non-inclusive

		//fmt.Println("i:", i, "l vs rl:", l, filps[k]-start, "start:", start, "flips[0,1,-1]", filps[0], filps[1], filps[k])

		if nums[start] == 0 {
			start++
		} else {
			start = filps[0] + 1
		}
	}

	return l
}

/*
	Wrong Answer
	Details
	Input
	[1,0,0,0,1,1,0,0,1,1,0,0,0,0,0,0,1,1,1,1,0,1,0,1,1,1,1,1,1,0,1,0,1,0,0,1,1,0,1,1]
	8
	Output
	23
	Expected
	25

	the iteration on zeros.. need to cover the last so until i=len(zeros)
	then..

	51 / 52 test cases passed.
	Status: Wrong Answer
	Submitted: 0 minutes ago
	Input:
	[1,0,1,1,0,1,0,0,1,1,0,1,1,1,1,1,0,0,1,0,1,1,0,0,0,1,1,0,0,0,0,0,0,1,0,0,1,0,0,1,0,1,1,0,0,0,1,1,1,1,0,0,0,1,1,0,1,0,0,0,0,1,1,1,1,0,1,1,1,1,1,0,1,1,0,0,0,0,0,1,1,1,1,0,0,0,1,1,1,1,1,0,1,1,1,0,1,1,0,0,0,0,0,0,1,0,0,1,0,1,1,1,1,1,1,0,1,0,0,0,1,1,0,1,1,1,0,1,1,0,1,1,1,0,1,0,1,0,1,1,0,0,0,0,1,0,0,0,1,1,1,1,0,1,0,0,0,1,0,1,0,1,1,1,1,1,0,1,0,1,1,1,1,0,1,1,0,0,0,1,0,0,1,1,1,1,0,1,0,0,1,1,0,1,1,0,0,1,1,1,0,1,1,0,1,1,1,0,0,1,0,0,0,1,0,0,1,0,0,1,1,1,1,1,1,1,1,0,0,0,1,0,0,0,1,0,0,1,1,1,1,0,1,1,0,0,0,0,1,0,1,0,0,1,1,0,0,1,0,1,0,0,1,1,0,0,0,1,1,1,1,1,0,0,1,1,0,0,0,0,1,1,1,0,0,1,1,1,0,0,1,1,1,1,0,0,0,0,1,0,1,0,1,1,0,1,0,0,1,0,1,1,0,1,0,1,1,0,1,0,1,1,0,0,1,0,0,0,1,1,0,1,1,0,0,1,0,1,1,0,1,0,1,0,1,0,1,1,1,0,0,1,1,0,1,1,0,1,1,0,0,1,1,1,1,1,1,1,0,0,1,0,0,0,0,0,0,0,1,0,1,0,1,0,1,1,0,0,1,1,0,0,1,0,1,1,1,0,0,0,0,0,0,1,0,1,1,0,1,0,0,1,1,0,0,0,0,0,0,0,1,1,1,0,0,1,0,1,1,0,1,1,0,1,0,1,1,0,1,1,1,0,0,0,0,1,0,0,0,0,0,0,1,1,1,1,0,0,0,1,0,1,1,0,1,1,0,0,1,1,0,1,0,0,1,1,1,1,1,0,1,0,0,0,0,1,0,1,0,0,1,0,1,1,0,1,0,1,1,1,1,1,0,1,0,1,1,1,0,0,1,1,0,1,1,0,1,0,0,1,0,0,1,1,0,1,0,1,1,0,1,1,1,0,1,1,1,0,0,1,1,0,0,0,0,1,0,0,1,0,1,0,1,0,0,0,1,0,1,1,1,0,0,1,0,0,1,1,1,1,1,1,1,1,0,1,1,0,1,1,0,1,1,1,1,0,1,1,1,0,1,1,1,0,1,0,0,0,0,0,0,1,0,0,1,1,1,0,1,1,0,0,0,1,0,1,0,1,0,1,1,1,0,1,0,0,1,1,1,1,0,1,1,0,1,0,1,1,0,1,1,0,1,0,0,0,1,1,1,0,0,1,1,0,1,0,1,0,0,1,1,1,0,0,0,0,1,1,1,1,1,1,1,0,1,1,1,0,0,0,1,1,0,1,0,0,1,1,1,1,0,0,0,0,1,0,0,1,1,1,0,0,0,1,0,1,1,1,1,0,0,0,0,0,0,1,0,1,0,0,0,0,0,1,0,0,1,1,0,1,1,0,0,0,0,1,1,1,1,1,0,0,0,1,0,0,0,0,1,1,0,1,1,1,1,0,0,1,1,1,0,0,1,0,1,1,0,0,0,0,0,1,1,1,0,0,0,1,1,1,1,1,0,1,0,0,1,1,0,1,1,1,0,1,0,0,1,0,0,1,1,0,1,0,1,0,0,1,1,0,0,0,0,1,1,1,0,0,0,1,1,1,1,0,1,0,1,0,0,0,1,0,1,1,0,0,0,1,0,1,0,0,1,1,0,0,1,1,1,0,1,1,0,1,0,1,1,1,0,1,0,0,1,1,1,0,1,1,1,0,1,1,0,0,0,1,0,1,0,1,1,0,1,1,1,1,1,1,0,1,0,1,0,0,0,0,1,1,0,1,0,0,0,1,0,1,1,0,0,0,1,0,0,0,0,1,0,0,0,1,0,1,1,0,1,0,0,0,1,1,0,0,1,1,0,1,1,0,1,1,0,1,0,1,0,1,0,0,1,1,1,0,1,1,1,1,1,0,1,0,0,0,0,0,1,0,1,0,1,0,1,1,1,1,0,0,1,0,0,1,1,0,1,0,1,0,1,1,0,1,1,0,1,1,0,0,1,0,1,0,1,1,0,1,0,1,0,1,0,1,0,0,1,1,1,0,1,0,0,1,1,0,1,1,0,1,1,0,1,0,1,0,1,1,0,0,1,1,0,1,1,0,1,1,1,1,1,0,1,1,1,1,0,1,1,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,1,0,0,0,1,0,1,1,0,0,1,1,1,0,1,1,1,1,0,0,0,0,0,0,0,0,1,0,1,0,0,1,1,0,0,1,1,1,1,0,1,1,0,1,1,0,1,0,0,1,1,0,1,1]
	225
	Output:
	494
	Expected:
	495


	okay!
	I see where I had a mental trap...
	it is not necessary you need to flip consective zeros.. you skip and choose...

	I am only doing consective zeros..
	also I missed an important piece... you can flip at most k 0's. (AT MOST)
	so you can flip from 0 to k zeros...

	but the fact I passed 51/52 cases.. means mostly
	more zeros.. it will be longer.
	and consective zeros flipping 99% yield the right answere

	however, if it is not right 100%, then it is not

there are 4 hints..

1. One thing's for sure, we will only flip a zero if it extends an existing window of 1s. Otherwise, there's no point in doing it, right? Think Sliding Window!

2. Since we know this problem can be solved using the sliding window construct, we might as well focus in that direction for hints. Basically, in a given window, we can never have > K zeros, right?
	<== this makes think of something like sumSoFar...
		I will explore it

3. We don't have a fixed size window in this case. The window size can grow and shrink depending upon the number of zeros we have (we don't actually have to flip the zeros here!).
4. The way to shrink or expand a window would be based on the number of zeros that can still be flipped and so on.

===

but my algorithm is indeed sliding window no doubt
so what is wrong? must be missing a beat

>>> for k in range(496,1125):
...   a1=a[k-495:k]
...   if sum(a1) >=270:
...     print(k, sum(a1), len(a1))

1070 270 495
>>> a[1070-495:1070]

so okay... I pinpointed to be this sequence
>>> a[565:577]
[0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1]
	^						   ^  ^
when start is here.. it moves off to next zero... and that skipped the highlighted two 1s...
I jumped too much

instead I should jump to first zero's next but not the next zero...
why 51 test cases passed... pure luck

so okay..
passed

Success
Details
Runtime: 85 ms, faster than 39.38% of Go online submissions for Max Consecutive Ones III.
Memory Usage: 8.6 MB, less than 7.50% of Go online submissions for Max Consecutive Ones III.
Next challenges:

change to
	for i := k + 1; i <= len(zeros); i++ {
		//filps := zeros[i-k-1 : i]
        fZero := zeros[i-k-1]
        end := zeros[i-1]
		l = max(l, end-start) //the last flip is the marker of end, non-inclusive
		if nums[start] == 0 {
			start++
		} else {
			start = fZero+1
		}
	}

Success
Details
Runtime: 76 ms, faster than 50.00% of Go online submissions for Max Consecutive Ones III.
Memory Usage: 7.4 MB, less than 60.63% of Go online submissions for Max Consecutive Ones III.
Next challenges:

*/
// sliding window
func _2_longestOnes(nums []int, k int) int {

	maxLen := 0
	sum := 0
	i, j := 0, 0
	for i <= j && j < len(nums) {
		sum += nums[j]
		for sum+k < j-i+1 && i <= j {
			sum -= nums[i]
			i++
		}
		maxLen = max(maxLen, j-i+1)
		j++
	}

	return maxLen

}

// sliding window 2
func _3_longestOnes(nums []int, k int) int {

	maxLen := 0
	sum := nums[0]
	i, j := 0, 0
	for j < len(nums) {

		// this is tricky; you need to let i to go beyond j to break the tie sometime
		// that depends on what kind of state (and some time the end result) you are tracking
		// here. I am tracking the sum (num of 1s) then sum+k vs windLen
		// but you can track number of 0s, vs k..

		windLen := j - i + 1
		if sum+k >= windLen {
			maxLen = max(maxLen, windLen)
			if j < len(nums)-1 {
				sum += nums[j+1]
			}
			j++
		} else {
			sum -= nums[i]
			i++
		}
	}

	return maxLen

}

// this guy's template is actually very cool
// https://leetcode.com/problems/max-consecutive-ones-iii/discuss/1504260/C%2B%2B-Sliding-Window-(%2B-Cheat-Sheet)

// sliding window 3: tracking 0s count
func _4_longestOnes(nums []int, k int) int {
	maxLen := 0
	count := 0
	for i, j := 0, 0; j < len(nums); {
		if nums[j] == 0 {
			count++

		}

		for count > k {
			if nums[i] == 0 {
				count--

			}
			i++
		}

		j++
		// because j++ happens before, so j-i
		maxLen = max(maxLen, j-i)

		// or written this way, the same
		// maxLen = max(maxLen, j-i+1)
		// j++
	}

	return maxLen
}

// there is yet another non-shrinkable solution
// which is even more slippy
// it use the k as a quota or something like
// when it establish a window and when the condition is violated it will stop growing (i++, and j++ and the same time)
// i++, j++ the window keeps the same size
// if at later time, it can grow again... j++ without i++...
// just alient-like smart

/*
think the scenarios when count>k
	nums[j] is 1, and nums[i] is 1 -> window won't grow or shrink and count stable
	nums[j] is 0, count++, nums[i] is 1 -> window won't grow or shrink and count + 1 -- more debt
	nums[j] is 1, nums[i] is 0, window won't grow and count - 1, -- less debt

	nums[j] is 1, when count <=k, window just grow, count stable -- no debt yet...

	so when it found a window, it can shrink or grow, but the max so far is j-i..
	and from there on, j and i increase in parallel... so window size won't change
	but the count could change...

	depending on how it changes. it could grow more... window bigger
	if it go over k, then window won't grow anymore..

	the debt has to be paid off first

	damn... no matter how I try to explain.. this feels very alien..

*/

func longestOnes(nums []int, k int) int {
	i, j := 0, 0
	count := 0
	for j < len(nums) {
		if nums[j] == 0 {
			count++
		}

		if count > k {
			if nums[i] == 0 {
				count--

			}
			i++
		}

		j++

	}

	return j - i
}

func testLongestOnes() {
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 1, 0, 0}, 0))

	fmt.Println(longestOnes([]int{1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1}, 225))

	fmt.Println(longestOnes([]int{1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1}, 8))

	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 0))
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1}, 0))
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 1))

	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))

}
