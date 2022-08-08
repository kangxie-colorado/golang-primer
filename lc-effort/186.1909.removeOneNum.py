"""
https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/

I did make it happen and see two possibilities
but I did in two passes..

now let me do it in one pass
the key is to test if idx i-2 is bigger or smaller than idx i
if it is bigger.. the i offends two index, i should be at least removed
"""

# this is my first solution.. pretty neaty actuall


import re
from typing import List


class Solution:
    def canBeIncreasing(self, nums: List[int]) -> bool:
        def isStrictIncreasing(A):
            if len(A) == 0:
                return True

            last = A[0]
            for i in range(1, len(A)):
                if A[i] <= last:
                    return False
                last = A[i]
            return True

        last = nums[0]
        for i in range(1, len(nums)):
            if nums[i] <= last:
                return isStrictIncreasing(nums[:i-1]+nums[i:]) or isStrictIncreasing(nums[:i]+nums[i+1:])
            last = nums[i]

        return True


class Solution:
    def canBeIncreasing(self, nums: List[int]) -> bool:
        i = 1
        removed = False
        while i < len(nums):
            if nums[i] <= nums[i-1]:
                if removed:
                    return False
                if i == 1:
                    # this is just remove first element
                    # and continue without another chance
                    removed = True

                else:
                    if nums[i] <= nums[i-2]:
                        # in this case I remove nums[i]
                        removed = True
                        nums[i] = nums[i-1]

                    else:
                        # remove nums[i-1], use the chance
                        # draw up the nums[i-2] to nums[i-1]
                        # to signify the removal of nums[i-1]
                        removed = True
                        nums[i-1] = nums[i-2]

            i += 1
        return True


"""
Runtime: 93 ms, faster than 51.24% of Python3 online submissions for Remove One Element to Make the Array Strictly Increasing.
Memory Usage: 14 MB, less than 83.88% of Python3 online submissions for Remove One Element to Make the Array Strictly Increasing.

it is quite messy..
so I see why other use a side state variable to help
"""


class Solution:
    def canBeIncreasing(self, nums: List[int]) -> bool:
        i = 1
        removed = 0
        while i < len(nums):
            if nums[i] <= nums[i-1]:
                removed += 1

                if i > 1 and nums[i] <= nums[i-2]:
                    # in this case I remove nums[i]
                    # I draw up the nums[i-1] to nums[i] to conceptually remove nums[i]
                    # and let the array to continue with nums[i-1]'s information
                    nums[i] = nums[i-1]

            i += 1
        return removed < 2


"""
Runtime: 94 ms, faster than 50.25% of Python3 online submissions for Remove One Element to Make the Array Strictly Increasing.
Memory Usage: 14.1 MB, less than 49.65% of Python3 online submissions for Remove One Element to Make the Array Strictly Increasing.

then I actually see I can re-organize my previous submission like this
because after all, you only need to deal with removing nums[i] scenario

removing nums[i-1] scenario is naturally done
"""


class Solution:
    def canBeIncreasing(self, nums: List[int]) -> bool:
        i = 1
        removed = False
        while i < len(nums):
            if nums[i] <= nums[i-1]:
                if removed:
                    return False
                removed = True

                if i > 1 and nums[i] <= nums[i-2]:
                    # in this case I remove nums[i]
                    removed = True
                    nums[i] = nums[i-1]

            i += 1
        return True
