// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

package main

/*
You must write an algorithm that runs in O(log n) time.
^ because of this, it has to be binary search

but when thinking of binary search, the thing is it has to compute mid on some condition f()
what would be the condition?

we are looking for minimum number, so the nums[0] gives up a starting point, it is at least minimum
in some part, 0..i-1, if ith being the true minimum

if nums[mid] > nums[0](lastMin), then the true minimum is
- either nums[0]
- or to the right of mid
	because between nums[0] and nums[mid], there cannot be anything <nums[0]

if nums[mid] < nums[0], then the true minimum is
- either nums[mid]
- or to the left of mid
	because whatever to the right of nums[mid] is greater than nums[mid]
	and to the left, there is possible smaller ones than nums[mid]

then we update r/l to mid/mid+1 accordingly
also update lastMin accordingly


*/

func _1_alitte_messy_findMin(nums []int) int {
	lastMin := nums[0]
	// r should be one element out of the range..
	// otherwise, when l=r=len(nums)-1
	// I actually will miss it e.g.[2,1]
	// so what do I need this??? let me chec neetcode guy
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= lastMin {
			// == case is basically mid landed on l
			l = mid + 1
		} else {
			lastMin = nums[mid]
			r = mid
		}
	}

	return lastMin
}

/*
Wrong Answer
Details
Input
[2,1]
Output
2
Expected
1

yeah an edge case
nums[mid] == nums[0]

Runtime: 3 ms, faster than 65.37% of Go online submissions for Find Minimum in Rotated Sorted Array.
Memory Usage: 2.5 MB, less than 5.37% of Go online submissions for Find Minimum in Rotated Sorted Array.

[2, 3, 4, 5, 0, 1]
 l     m        l
hmm.. so his view point is to decide m is in the left sorted part or right sorted part
if nums[m] >= nums[l] -- then it is in the left sorted part, then looking in the right
	which is what I did
if nums[m] < nums[l] -- then search to the left

it has a shortcut
if l<r, that means this l:r is already sorted and nums[l] is the smallest in this section

he also do r=mid-1... hmm.. with for l<=r

non-standard so some variation is there..

ah.. okay.. I think my issue is I bookkeeping lastMin
instead I should stick to return nums[l]
*/

func findMin(nums []int) int {

	l, r := 0, len(nums)-1

	for l < r {
		if nums[l] < nums[r] {
			break
		}

		mid := l + (r-l)/2
		if nums[mid] >= nums[l] {
			// == case is basically mid landed on l
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}

/*
Wrong Answer
Runtime: 2 ms
Your input
[3,4,5,1,2]
Output
2
Expected
1

so at 1, it >=, then it moved to 2...
I still need to bookkeeping???

so the problem is I missed using if nums[l] < nums[r], it is already sorted
		if nums[l] < nums[r] {
			break
		}


Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Minimum in Rotated Sorted Array.
Memory Usage: 2.5 MB, less than 5.37% of Go online submissions for Find Minimum in Rotated Sorted Array.

weird how is this much better than my previous solution
my solution doesn't need to know... the shortcut....
it is actually anti-intuitive
11 12 13 15

first it steps onto 12, then it looks to right
13 15, it continues to look to right.. but will never find a smaller one... (wasteful indeed)

so yeah, kind of require knowing if sorted already, return the first..

also this
public int FindMin(int[] nums) {
    int left = 0, right = nums.Length - 1, mid = 0;
    while(left < right){
        mid = (left + right) >> 1;
        if(nums[mid] > nums[right]) left = mid + 1;
        else right = mid;
    }
    return nums[right];
}

instead of comparing to left, it does the right
and also the same... shift strategy.. this is quite smart...
no need to book keep, and right will converge with left

return left/right the same thing

I think this is the best version

yeah.. comparing to right seems actually more intuitive..

*/

func _compare_to_right_findMin(nums []int) int {

	l, r := 0, len(nums)-1

	for l < r {

		mid := l + (r-l)/2
		if nums[mid] >= nums[r] {
			// == case is basically mid landed on l
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[r]
}

/*
Runtime: 3 ms, faster than 65.37% of Go online submissions for Find Minimum in Rotated Sorted Array.
Memory Usage: 2.5 MB, less than 5.37% of Go online submissions for Find Minimum in Rotated Sorted Array.
*/
