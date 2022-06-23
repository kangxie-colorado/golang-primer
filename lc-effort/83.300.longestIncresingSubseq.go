// https://leetcode.com/problems/longest-increasing-subsequence/

package main

import "fmt"

/*
I don't have much idea
but drawing on whiteboard leads to looking at DP in some way

then I see
I can do looking-to-right, the max and len
then do looking-to-left, the min and len

			0 1 0 3 2 3
to r
	max		0 1 1 3 3 3
	len		1 2 2 3 3 3			<== meaning at any pos, this is the max I can see and len of increasing subseq I can have

to l
			0 0 0 2	2 3	min
			3 3	3 2	2 1	len

as you can see, 3+2 = 5..
doesn't mean using it to minus 1, I get the answer?
somehow it can be argued, at one of the spot, like the first 1,
	leading to me from left, I can see len-2 increasing subseq
	leading to me from right, I can see len-3 decreasing subseq

	so I a counted twice, 5-2 = 4
	other 3+2 doesn't work out this nicely, but that is just how DP being mystery

this would also work on the other two examples
let me try

*/

func _naive_greedy_and_wrong_lengthOfLIS(nums []int) int {
	toRightMax := make([]int, len(nums))
	toLeftMin := make([]int, len(nums))

	maxN := -10001
	runningLen := 0
	for i, n := range nums {
		if n > maxN {
			maxN = n
			runningLen++
		}

		toRightMax[i] = runningLen
	}

	minN := 10001
	runningLen = 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < minN {
			minN = nums[i]
			runningLen++
		}

		toLeftMin[i] = runningLen
	}

	res := 0
	for i := range toRightMax {
		res = max(res, toRightMax[i]+toLeftMin[i])
	}

	return res - 1
}

/*

24 / 54 test cases passed.
Status: Wrong Answer
Submitted: 0 minutes ago
Input:
[4,10,4,3,8,9]
Output:
4
Expected:
3

yeah... cannot be that easy, right?

and they asked
Follow up: Can you come up with an algorithm that runs in O(n log(n)) time complexity?

while I was doing O(N)
I being two greedy here..


so this O(n logn)
hmm... pq?

saving the

hmm..
I then drew the monotonical stack on whiteboad it also seems to work out on examples and the failed case
the rule
keep it increasing
when there is smaller one incoming, pop the bigger one
if it hits an equal one, pop itself.. let the furthest stay in the stack

anyway, worth a try...
no need to try..
it will fail..

4 5, 0, 6,7,8
you pop it out.. it will 0,6,7,8
but 4,5,6,7,8 is the answer..

oh shit...

okay.. back to brute force
just with some early termination

4 5 6 0 1 7 8 9

4,keep going to 4 5 6; 0<6, skip, so it ends 4 5 6 7 8 9 => 6
then 5, 5>4... and 5 is to the right of 4.. skip it

(suddenly I realize this can be viewed as a graph... l can go l+1,l+2,l+3..
that laters.. )
*/

func _miserable_lengthOfLIS(nums []int) int {
	res := 0
	processed := []int{}
	for i := 0; i < len(nums); i++ {
		for _, p := range processed {
			if p < nums[i] {
				continue
			}
		}

		maxSofar := nums[i]
		runningLen := 1

		for j := i + 1; j < len(nums); j++ {
			if nums[j] > maxSofar {
				maxSofar = nums[j]
				runningLen++
			}
		}

		res = max(res, runningLen)
		processed = append(processed, nums[i])
	}

	return res
}

/*
failed miserable here
Input: nums = [0,1,0,3,2,3]
Output: 4

but one little pattern is revealed
		for _, p := range processed {
			if p < nums[i] {
				continue
			}
		}
if anything to my right is bigger than me, then that can be skipped as a starting point...
if anything to my right is equal to me, then that can be skipped as a starting point...

but still zero thoughts..

this is not a easy one for sure..
back to the graph thinking...




*/
//  [10,9,2,5,3,7,101,18]
func _1_lengthOfLIS(nums []int) int {
	var dfs func(i int) int
	var skipChecking func(checking []int, num int) bool
	m := make(map[int]int)

	skipChecking = func(checking []int, num int) bool {
		for _, c := range checking {
			if num > c {
				return true
			}
		}
		return false
	}

	dfs = func(i int) int {
		if v, found := m[i]; found {
			return v
		}

		maxSubLen := 0
		checking := []int{}
		for j := i + 1; j < len(nums); j++ {
			// can use a little heap here..
			if j > i+1 && skipChecking(checking, nums[j]) {
				continue
			}

			if nums[j] > nums[i] {
				checking = append(checking, nums[j])
				maxSubLen = max(dfs(j), maxSubLen)
			}
		}
		m[i] = maxSubLen + 1
		return maxSubLen + 1
	}

	proceesedMin := 10001
	maxLen := 0
	for i := range nums {
		if nums[i] > proceesedMin {
			continue
		}
		maxLen = max(maxLen, dfs(i))
		proceesedMin = min(proceesedMin, nums[i])
	}

	return maxLen
}

/*
37 / 54 test cases passed.

okay TLE..
let me see I got it right at least

of course, there can be memorization

with it, I got 33 too.. so algorithm is probably right
okay..

barely I passed
Runtime: 256 ms, faster than 5.09% of Go online submissions for Longest Increasing Subsequence.
Memory Usage: 6.8 MB, less than 5.25% of Go online submissions for Longest Increasing Subsequence.

I can use that processedMin instead of skipChecking
*/

func _dfs_lengthOfLIS(nums []int) int {
	var dfs func(i int) int
	m := make(map[int]int)

	dfs = func(i int) int {
		if v, found := m[i]; found {
			return v
		}

		maxSubLen := 0
		checkingMin := 10001
		for j := i + 1; j < len(nums); j++ {
			// can use a little heap here..
			if j > i+1 && nums[j] > checkingMin {
				continue
			}

			if nums[j] > nums[i] {
				checkingMin = min(checkingMin, nums[j])
				maxSubLen = max(dfs(j), maxSubLen)
			}
		}
		m[i] = maxSubLen + 1
		return maxSubLen + 1
	}

	proceesedMin := 10001
	maxLen := 0
	for i := range nums {
		if nums[i] > proceesedMin {
			continue
		}
		maxLen = max(maxLen, dfs(i))
		proceesedMin = min(proceesedMin, nums[i])
	}

	return maxLen
}

/*
Runtime: 102 ms, faster than 11.27% of Go online submissions for Longest Increasing Subsequence.
Memory Usage: 6.3 MB, less than 5.25% of Go online submissions for Longest Increasing Subsequence.

okay.. 2x faster... but I know there is an order better solution
eat first

Runtime: 57 ms, faster than 61.27% of Go online submissions for Longest Increasing Subsequence.
Memory Usage: 6.5 MB, less than 5.25% of Go online submissions for Longest Increasing Subsequence.

hmm...

previously cpp solution

class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        if(nums.size() == 0)
            return 0;
        vector<int> dp(nums.size(),1);
        int res=1;

        for(int i=1;i<nums.size();++i) {
            for(int j=0; j<i; ++j) {
                if(nums[i] > nums[j])
                    dp[i] = max(dp[i], dp[j]+1);
            }

            res = max(dp[i], res);
        }

        return res;
    }
};

it is simple DP yes.. I have thought about it actullly
but now it runs
Runtime: 340 ms, faster than 52.85% of C++ online submissions for Longest Increasing Subsequence.
Memory Usage: 10.5 MB, less than 42.70% of C++ online submissions for Longest Increasing Subsequence.
*/

/*
wow.. the binary search...
replace with smaller value and binary search...
*/

func insertIndex(nums []int, num int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == num {
			return mid
		} else if nums[mid] > num {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func lengthOfLIS(nums []int) int {
	subseq := []int{}

	for i := range nums {
		if len(subseq) == 0 || nums[i] > subseq[len(subseq)-1] {
			subseq = append(subseq, nums[i])
		} else {
			idx := insertIndex(subseq, nums[i])
			subseq[idx] = nums[i]
		}

	}

	return len(subseq)
}

/*
Runtime: 11 ms, faster than 74.38% of Go online submissions for Longest Increasing Subsequence.
Memory Usage: 3.6 MB, less than 59.57% of Go online submissions for Longest Increasing Subsequence.

ref:
https://leetcode.com/problems/longest-increasing-subsequence/discuss/1326308/C%2B%2BPython-DP-Binary-Search-BIT-Solutions-Picture-explain-O(NlogN)

https://leetcode.com/problems/longest-increasing-subsequence/discuss/1326308/C++Python-DP-Binary-Search-BIT-Solutions-Picture-explain-O(NlogN)/1040509
^ this is actually the intuitive explanation...


*/

func testInsertIndex() {

	fmt.Println(insertIndex([]int{2, 6, 8}, 2))
	fmt.Println(insertIndex([]int{2, 6, 8}, 0))
	fmt.Println(insertIndex([]int{2, 6, 8}, 4))
	fmt.Println(insertIndex([]int{2, 6, 8}, 10))
}

func testLengthOfLIS() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	// should be 33,
	fmt.Println(lengthOfLIS([]int{-147, -171, -584, 590, 501, 13, 489, -938, 396, -544, -229, 697, 157, -933, -264, -436, -691, -469, 49, -574, 694, 50, 672, -858, -923, 974, -157, -507, -907, 429, 529, -591, 802, -351, -606, 296, -485, 454, 540, 300, -1000, 408, 923, 0, -975, -548, 62, -990, 835, 650, 733, -611, -385, -580, 330, 394, 566, -191, 612, -608, -478, -104, -425, 58, -849, 601, 851, -208, -810, 400, 412, 571, -535, -995, 627, -481, -702, 457, -29, 375, 792, -186, -921, -275, 654, -356, -322, -28, -843, 527, -266, -970, 556, 852, -890, 169, -413, 2, -958, -651, 371, 895, -994, 671, 243, -605, -556, 735, -246, 179, -104, -771, 658, -554, 932, -829, -455, -981, -731, -148, 512, -547, -946, -997, -197, 864, 870, 629, -961, 659, 574, 543, -501, 582, -799, -428, 876, -334, 115, 759, 197, -905, 275, 76, 242, 357, 694, -254, -361, -338, -57, 596, 786, -710, -51, -496, -100, 246, -969, 874, 504, 938, 931, -365, 175, -40, -616, 596, 440, 567, 999, 15, -363, -256, -578, -869, -653, 78, -352, 882, 749, -33, 462, -592, -751, 761, -96, 206, 489, 34, 367, 960, 68, 837, 37, -764, -897, 72, 639, -69, 353, 836, -67, 491, 126, -171, -532, -757, -358, 217, 806, 712, -32, 843, -790, -691, -381, -138, 6, -712, 153, -184, -544, 3, 840, -561, 917, -704, -126, -230, 468, 963, -993, 445, -892, -543, 941, -665, 58, 268, -362, 181, -529, 684, 313, -380, -712, 700, 601, -962, -886, 702, 439, 153, -87, 140, 583, -323, 70, -460, -863, -859, -784, 571, 169, 44, -460, 181, 883, 600, 982, -367, -191, 815, -84, 961, -791, -713, 149, -499, 330, -351, -442, -989, -662, -183, -220, -617, -638, -916, 454, 604, 559, -304, -812, 526, -891, 984, -762, -669, -414, -481, -219, -776, 690, -72, -250, -282, -961}))
}
