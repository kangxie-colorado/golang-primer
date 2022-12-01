"""
https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

I think I did this before but kind of forget
or not

idea is still to process and memorize the results
for each node, use heapq (-lis, val) to traverse the array.. and see how far it can go
memorize the results and this will be like an O(n)
"""


import heapq
from typing import List


class Solution:
    def longestIncreasingPath(self, matrix: List[List[int]]) -> int:
        lis = [[]]*len(matrix)
        for i in range(len(lis)):
            lis[i] = [0]*len(matrix[0])

        def getLis(r, c, lastVal):
            if r < 0 or r >= len(matrix) or c < 0 or c >= len(matrix[0]) or matrix[r][c] <= lastVal:
                return 0

            if lis[r][c] != 0:
                return lis[r][c]

            lis[r][c] = 1 + max(getLis(r-1, c, matrix[r][c]),
                                getLis(r+1, c, matrix[r][c]),
                                getLis(r, c+1, matrix[r][c]),
                                getLis(r, c-1, matrix[r][c]))
            return lis[r][c]

        res = 0
        for r in range(len(matrix)):
            for c in range(len(matrix[r])):
                lis[r][c] = getLis(r, c, -1)
                res = max(res, lis[r][c])

        return res


"""
problem
1. use heaqp, it is like a queue... cannot memorize?
so I use dfs instead and code is pretty pretty
Runtime: 1013 ms, faster than 25.02% of Python3 online submissions for Longest Increasing Path in a Matrix.
Memory Usage: 14.9 MB, less than 75.20% of Python3 online submissions for Longest Increasing Path in a Matrix.


2. let me see using something else
https://leetcode.com/problems/longest-increasing-path-in-a-matrix/discuss/78308/15ms-Concise-Java-Solution/82964

I was thinking about using heap
but didn't connect thru the fogs, this post clears it
so the idea is 
 - maxHeap by the values
 - pop by order then looking around to see how long this can extend to
    -- this is the key idea, the bigger element are all popped out so their lis[i][j] is final 
 - maintain the result


"""


class Solution:
    def longestIncreasingPath(self, matrix: List[List[int]]) -> int:
        lis = [[]]*len(matrix)
        h = []  # (-val, i,j)
        for r in range(len(matrix)):
            lis[r] = [1]*len(matrix[r])
            for c in range(len(matrix[r])):
                heapq.heappush(h, (-matrix[r][c], r, c))

        res = 0
        while h:
            val, r, c = heapq.heappop(h)
            if r > 0 and matrix[r][c] < matrix[r-1][c]:
                lis[r][c] = max(lis[r-1][c]+1, lis[r][c])
            if r < len(matrix)-1 and matrix[r][c] < matrix[r+1][c]:
                lis[r][c] = max(lis[r+1][c]+1, lis[r][c])
            if c > 0 and matrix[r][c] < matrix[r][c-1]:
                lis[r][c] = max(lis[r][c-1]+1, lis[r][c])
            if c < len(matrix[0])-1 and matrix[r][c] < matrix[r][c+1]:
                lis[r][c] = max(lis[r][c+1]+1, lis[r][c])

            res = max(res, lis[r][c])

        return res


"""
Runtime: 1316 ms, faster than 11.74% of Python3 online submissions for Longest Increasing Path in a Matrix.
Memory Usage: 15.9 MB, less than 65.37% of Python3 online submissions for Longest Increasing Path in a Matrix.

okay.. take a nap break
and continue
"""
