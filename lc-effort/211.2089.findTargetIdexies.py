"""
https://leetcode.com/problems/find-target-indices-after-sorting-array/


"""


from bisect import bisect_left, bisect_right
from typing import List


class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        nums.sort()

        left = bisect_left(nums, target)
        right = bisect_right(nums, target)
        if left == len(nums) or nums[left] != target:
            return []
        else:
            return range(min(left, right-1), max(left, right-1)+1)


"""
Runtime: 40 ms, faster than 99.65% of Python3 online submissions for Find Target Indices After Sorting Array.
Memory Usage: 13.9 MB, less than 17.89% of Python3 online submissions for Find Target Indices After Sorting Array.


after many tries..
but the idea is that..

find first <= number, find first > number
bisect_left is actually the idx to insert to the left
bisect_right is the idx to insert to the right

[1,2,4]... insert 2 to left, it should be idx-1 and bisect_left(a,2) returns 1
>>> bisect_left([1,2,4],2)
1
[1,2,4]... insert 2 to the right, it should be indx-2 and bisect_right(a,2) returns 2
>>> bisect_right([1,2,4],2)
2

but I wonder if I need to differentiate the return like above
        if left == len(nums) or nums[left] != target:
            return []
        else:
            return range(min(left, right-1), max(left, right-1)+1)
    can I do
        return range(left, right+1)
turns out I can just do below
"""


class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        nums.sort()

        left = bisect_left(nums, target)
        right = bisect_right(nums, target)
        return range(left, right)


"""
Runtime: 54 ms, faster than 86.85% of Python3 online submissions for Find Target Indices After Sorting Array.
Memory Usage: 13.9 MB, less than 66.67% of Python3 online submissions for Find Target Indices After Sorting Array.

if left==right, no such number, then range is empty
if left<right, then it takes care of itself; right at the one pos ahead

if left>right... that is not possible...

let me write my own bisect_left/right
"""


class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        nums.sort()

        def bisect_leftx(nums, target):
            l, r = 0, len(nums)
            while l < r:
                mid = l + (r-l)//2
                if nums[mid] < target:
                    l = mid + 1
                else:
                    # >= I need
                    r = mid
            return l

        def bisect_rightx(nums, target):
            l, r = 0, len(nums)
            while l < r:
                mid = l + (r-l)//2
                if nums[mid] <= target:
                    l = mid + 1
                else:
                    # > I need
                    r = mid
            return l

        left = bisect_leftx(nums, target)
        right = bisect_rightx(nums, target)
        return range(left, right)


"""
Runtime: 59 ms, faster than 79.08% of Python3 online submissions for Find Target Indices After Sorting Array.
Memory Usage: 13.9 MB, less than 17.89% of Python3 online submissions for Find Target Indices After Sorting Array.

okay.. I got it right..
"""
