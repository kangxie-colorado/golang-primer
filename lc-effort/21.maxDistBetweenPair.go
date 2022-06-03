// https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/

package main

func maxDistance(nums1 []int, nums2 []int) int {
	dist := 0
	lastj := 0
	for i := 0; i < len(nums1); i++ {
		/*
			// no need for this, because lastj is always one step further
			if i > lastj {
				lastj = i
			}
		*/
		for j := lastj; j < len(nums2); j++ {
			if nums2[j] < nums1[i] {
				break
			}

			if j-i > dist {
				dist = j - i
			}

			lastj++
		}

		if lastj == len(nums2) {
			break
		}
	}

	return dist
}

/*
	naive n2 works but of course too slow
	the array is non-increasing..

	then it can end early and also remember the start point of next time
	when first i in nums1 end up with a j in nums2

	the next i+1 can start with j+1, because nums2[i] will be for sure be bigger than nums[i+1] -- the non-increasing array
	but the dist will be less than j to i: (j - (i+1)) < (j-i) of course

	so it can start with j+1, hence my above solution works

	Success
	Details
	Runtime: 200 ms, faster than 46.81% of Go online submissions for Maximum Distance Between a Pair of Values.
	Memory Usage: 9.3 MB, less than 76.60% of Go online submissions for Maximum Distance Between a Pair of Values.

	next optimization would be when lastj have reached the end.. then it is done..
		if lastj == len(nums2) {
			break
		}

	Success
	Details
	Runtime: 156 ms, faster than 80.85% of Go online submissions for Maximum Distance Between a Pair of Values.
	Memory Usage: 9.8 MB, less than 39.36% of Go online submissions for Maximum Distance Between a Pair of Values.

	overall this is not a super complicated problem
func maxDistance(nums1 []int, nums2 []int) int {
    i := 0
    j := 0
    maxDist := 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] <= nums2[j]{
            temp := j-i
            if temp > maxDist{
                maxDist =temp
            }
            j++
        }else{
            i++
        }
    }
    return maxDist
}

this is actually what I did but more elegant  O(m+n)
another solution is binary search to find first number smaller than you, then the number before will be >=... O(nlgn)


*/
