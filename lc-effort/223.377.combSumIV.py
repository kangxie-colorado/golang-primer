"""
https://leetcode.com/problems/combination-sum-iv/

dp?
"""


import enum
from os import umask
from tkinter import N
from turtle import back
from typing import List


class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        dp = [0]*(target+1)

        S = set(nums)
        for i in range(1, target+1):
            if i in S:
                dp[i] += 1
            for j in range(1, i):
                dp[i] += dp[j] * dp[i-j]

        return dp[target]


"""
I tried DP, but onto a wrong path
then I tried backtracking (divide and conquere..) with no victory...

I looked at discussions and now I see what went wrong with my DP thinkings
I was thinking f(i) = f(j) * f(i-j)... but no way to deal with the duplicates
(wonder if it is all distinct number and only use once, then likely can do that way)

this is infinite of number and can use as many as possible... ugh
so the formula is 

f(i) = sum(f[i-n]) for n i nums])
when some target' plus a n (in nums) to get the target, its total solution carries over
so do a sum we get it... 

notice the subtle difference with that coin change2 problem... I'll do after this 
"""


class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        dp = [0]*(target+1)
        dp[0] = 1

        for i in range(1, target+1):
            for n in nums:
                if i >= n:
                    dp[i] += dp[i-n]

        return dp[target]


"""
Runtime: 38 ms, faster than 95.93% of Python3 online submissions for Combination Sum IV.
Memory Usage: 13.9 MB, less than 81.72% of Python3 online submissions for Combination Sum IV.

that coin change 2 can was evovled from 2-d, is this too?
"""


class Solution:
    def combinationSum4(self, coins: List[int], amount: int) -> int:
        dp = [[]]*(len(coins)+1)
        for r in range(len(dp)):
            dp[r] = [0]*(amount+1)
            dp[r][0] = 1

        for j in range(1, amount+1):
            for i in range(1, len(coins)+1):

                # exclude and include
                dp[i][j] += dp[i-1][j]
                if j >= coins[i-1]:
                    dp[i][j] += dp[i][j-coins[i-1]]

        return dp[len(coins)][amount]


"""
ah.. naha... 
it is permutation... it can look back
there is no such cascade structure here... 

but the one-d DP solution is kind of more natural..
"""

if __name__ == '__main__':
    s = Solution()
    print(s.combinationSum4([1, 2], 3))
    print(s.combinationSum4([1, 2, 3], 4))
