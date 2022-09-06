"""
https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/

the simple way is to calculate the SUM of division one by one
I am thinking if I can do better.. by using
bisect_right(d)
bisect_right(2d)
bisect_right(3d)
to calculate the count
"""


from bisect import bisect_right
from typing import List


def getDivisionSum(nums, d):
    nums.sort()
    lastMultipleIdx = 0
    multiple = 1
    res = 0
    while d*multiple <= nums[-1]:
        idx = bisect_right(nums, d*multiple)
        res += (idx-lastMultipleIdx)*multiple
        multiple += 1
        lastMultipleIdx = idx
    res += (len(nums)-lastMultipleIdx)*multiple
    return res


class Solution:
    def smallestDivisor(self, nums: List[int], threshold: int) -> int:
        nums.sort()

        def getDivisionSum(d):
            lastMultipleIdx = 0
            multiple = 1
            res = 0
            while d*multiple <= nums[-1]:
                idx = bisect_right(nums, d*multiple)
                res += (idx-lastMultipleIdx)*multiple
                multiple += 1
                lastMultipleIdx = idx
            res += (len(nums)-lastMultipleIdx)*multiple
            return res

        l, r = 1, 10**6
        while l < r:
            m = l+(r-l)//2
            if getDivisionSum(m) <= threshold:
                r = m
            else:
                l = m+1
        return l


"""
Runtime: 315 ms, faster than 99.89% of Python3 online submissions for Find the Smallest Divisor Given a Threshold.
Memory Usage: 20.6 MB, less than 29.57% of Python3 online submissions for Find the Smallest Divisor Given a Threshold.
"""

"""
    print(getDivisionSum([1, 2, 5, 9], 1))
    print(getDivisionSum([1, 2, 5, 9], 2))
    print(getDivisionSum([1, 2, 5, 9], 4))
    print(getDivisionSum([1, 2, 5, 9], 5))
"""


class Solution:
    def smallestDivisor(self, nums: List[int], threshold: int) -> int:
        nums.sort()

        def getDivisionSum(nums, d):
            multiple = 1
            res = 0
            idx = 0
            while nums and d*multiple <= nums[-1]:
                idx = bisect_right(nums, d*multiple)
                res += idx*multiple
                multiple += 1
                nums = nums[idx:]

            res += len(nums)*multiple
            return res

        l, r = 1, nums[-1]
        while l < r:
            m = l+(r-l)//2
            if getDivisionSum(nums, m) <= threshold:
                r = m
            else:
                l = m+1
        return l


"""
Runtime: 394 ms, faster than 96.28% of Python3 online submissions for Find the Smallest Divisor Given a Threshold.
Memory Usage: 20.2 MB, less than 97.70% of Python3 online submissions for Find the Smallest Divisor Given a Threshold.
"""


if __name__ == '__main__':
    s = Solution()
    print(s.smallestDivisor([1, 2, 5, 9], 6))
    print(s.smallestDivisor(nums=[44, 22, 33, 11, 1], threshold=5))
    print(s.smallestDivisor(nums=[44, 22, 33, 11, 1, 23, 123], threshold=23))
