""""
https://leetcode.com/problems/paths-in-matrix-whose-sum-is-divisible-by-k/

the path count is a trivial DP problem
this divisible by k makes things complicated 

but there are at most K mod value when you mod a number by k... so maybe that is a place to dig
"""


from collections import defaultdict
from typing import List


class Solution:
    def numberOfPaths(self, grid: List[List[int]], k: int) -> int:
        # buckets = {} # 0 - k-1
        m, n = len(grid), len(grid[0])
        dp = [[]]*m

        for i in range(m):
            dp[i] = [{}]*n
            dp[i][0] = defaultdict(int)
            if i == 0:
                dp[i][0][grid[0][0] % k] = 1
            else:
                dp[i][0][(grid[i][0] + list(dp[i-1][0].keys())[0]) % k] = 1

        for j in range(n):

            if j != 0:
                dp[0][j] = defaultdict(int)
                dp[0][j][(list(dp[0][j-1].keys())[0] + grid[0][j]) % k] = 1

        for i in range(1, m):
            for j in range(1, n):
                dp[i][j] = defaultdict(int)
                for s, c in dp[i-1][j].items():
                    dp[i][j][(s+grid[i][j]) % k] += c
                for s, c in dp[i][j-1].items():
                    dp[i][j][(s+grid[i][j]) % k] += c

        return dp[m-1][n-1][0] % (10**9+7)


"""
Runtime: 8912 ms, faster than 15.79% of Python3 online submissions for Paths in Matrix Whose Sum Is Divisible by K.
Memory Usage: 213.7 MB, less than 10.15% of Python3 online submissions for Paths in Matrix Whose Sum Is Divisible by K.

oops... I thought it is going to TLE or MLE

although passed
this is a bottom-up, although I could save some memory by using one dimension array but
I wonder if there is a top-down and that would be targetting the wanted sum


"""


class Solution:
    def numberOfPaths(self, grid: List[List[int]], k: int) -> int:
        m, n = len(grid), len(grid[0])
        cache = {}

        def helper(i, j, toModValue):
            if i < 0 or j < 0:
                return 0
            if i == j == 0:
                return 1 if (grid[0][0] % k) == toModValue else 0
            if (i, j, toModValue) in cache:
                return cache[(i, j, toModValue)]

            selfMod = grid[i][j] % k
            cache[(i, j, toModValue)] = helper(i-1, j, (toModValue+k-selfMod) %
                                               k) + helper(i, j-1, (toModValue+k-selfMod) % k)

            cache[(i, j, toModValue)] %= (10**9+7)
            return cache[(i, j, toModValue)]

        return helper(m-1, n-1, 0)


"""
okay.. I thought this ought to be better 
but not really.. TLE all the time
"""


if __name__ == "__main__":
    s = Solution()
    grid = [[5, 2, 4], [3, 0, 5], [0, 7, 2]]
    k = 3
    print(s.numberOfPaths(grid, k))
    grid = [[1, 5, 3, 7, 3, 2, 3, 5]]
    k = 29
    print(s.numberOfPaths(grid, k))
