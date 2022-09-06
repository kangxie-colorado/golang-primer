"""
https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/

werid cannot get it right
"""


import heapq
from typing import List


class Solution:
    def maxSubsequence(self, nums: List[int], k: int) -> List[int]:
        h = []

        for i in range(len(nums)):
            if i < k:
                heapq.heappush(h, (nums[i], i))
            else:
                topVal, topIdx = heapq.heappop(h)
                if nums[i] > topVal:
                    heapq.heappush(h, (nums[i], i))
                else:
                    heapq.heappush(h, (topVal, topIdx))
        return [t[0] for t in sorted(h, key=lambda x: x[1])]


"""
Runtime: 66 ms, faster than 88.64% of Python3 online submissions for Find Subsequence of Length K With the Largest Sum.
Memory Usage: 14.3 MB, less than 21.94% of Python3 online submissions for Find Subsequence of Length K With the Largest Sum.
"""

if __name__ == "__main__":
    s = Solution()
    nums = [-1, -2, 3, 4]
    k = 3
    assert [-1, 3, 4] == s.maxSubsequence(nums, k)
