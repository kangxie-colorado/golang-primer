"""
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

I did this one I know
also this is kind of a sub problem in that binary search in rotated array

only need to compare with the right most

if m<r, then m could be the min, r=m
if m>r, then m cannot be the min, l=m+1
"""


from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        l, r = 0, len(nums)-1

        while l < r:
            m = l+(r-l)//2
            if nums[m] < nums[r]:
                r = m
            else:
                l = m+1

        return nums[l]


"""
Runtime: 86 ms, faster than 31.31% of Python3 online submissions for Find Minimum in Rotated Sorted Array.
Memory Usage: 14.2 MB, less than 62.9
"""
