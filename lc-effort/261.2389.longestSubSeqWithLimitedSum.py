"""
https://leetcode.com/problems/longest-subsequence-with-limited-sum/


this is an easy question
what is the angle?
"""


from bisect import bisect
from typing import List


class Solution:
    def answerQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        nums.sort()
        prefix = 0
        for i in range(len(nums)):
            nums[i] += prefix
            prefix = nums[i]

        res = []
        for q in queries:
            res.append(bisect.bisect_right(nums, q))

        return res
