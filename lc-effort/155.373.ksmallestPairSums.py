"""
https://leetcode.com/problems/find-k-pairs-with-smallest-sums/

coming from kth element, this should be solvable

use A as row, B as columns
A[i]+B[j] as the values..


"""


import heapq


class Solution(object):
    def kSmallestPairs(self, nums1, nums2, k):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :type k: int
        :rtype: List[List[int]]
        """
        h = []
        A, B = nums1, nums2
        k = min(len(A)*len(B), k)

        for i in range(len(A)):
            heapq.heappush(h, (A[i]+B[0], i, 0))  # value, row, col

        res = []
        while k > 0:
            node = heapq.heappop(h)
            res.append([A[node[1]], B[node[2]]])

            if node[2]+1 < len(B):
                heapq.heappush(
                    h, (A[node[1]] + B[node[2]+1], node[1], node[2]+1))
            k -= 1

        return res


"""
Runtime: 1179 ms, faster than 74.30% of Python online submissions for Find K Pairs with Smallest Sums.
Memory Usage: 37.9 MB, less than 22.54% of Python online submissions for Find K Pairs with Smallest Sums.
"""
