"""

https://leetcode.com/problems/dice-roll-simulation/


I would not be able to figure out this myself
I worked thru the angle of deduction -- mis-lead by the examples..

when n>3, that angle will run into trouble of dealing with duplicates
1.5 hrs passed and I wasn't able to figure out

so I had to read up
it was pretty hard - almost a 3D DP, 2D but sum of row.. or multiple elements(5 or 6) are contributing to next row

also put it in the DP way is less visible to me
so here I go, just write out the code after walking thru the discussions and white board evolutions 

dp[i][j]: after rolling i-th time, total number of combinations that end with dice-value-j
i: 0->n.. rolling 0 times, the combination ending with 1-6 can only be zero.. but.. the total sum needs to be 1.. 
that is a weird definition to have and kind of need to have it to make the algorithm work

(that 1 can also be viewed as a seed to make the first row to work without specializing it)
"""


from typing import List


class Solution:
    def dieSimulator(self, n: int, rollMax: List[int]) -> int:
        mod = 10**9+7
        dp = [[]]*(n+1)
        for i in range(n+1):
            dp[i] = [0]*7

        # or understand the row-1 are all seeds
        dp[0][6] = 1

        for i in range(1, n+1):
            rowTotal = 0
            for j in range(6):
                k = rollMax[j]
                dp[i][j] = dp[i-1][-1]
                if i-1-k >= 0:
                    # this is to get rid of the k consecutive dice-(j+1)
                    dp[i][j] -= dp[i-1-k][-1] - dp[i-1-k][j]

                rowTotal += dp[i][j]
            dp[i][6] = rowTotal % mod

        return dp[n][6] % (10**9+7)


"""
Runtime: 183 ms, faster than 98.16% of Python3 online submissions for Dice Roll Simulation.
Memory Usage: 45.6 MB, less than 61.96% of Python3 online submissions for Dice Roll Simulation.

python allow you to store very big int.. but for other language you would need add the mode in the middle
I see the trick is add a mod then %: i.e. (xx + mod)%mod.. think this is to deal with the minus... 
"""


class Solution:
    def dieSimulator(self, n: int, rollMax: List[int]) -> int:
        mod = 10**9+7

        dp = [[]]*(n+1)
        for i in range(n+1):
            dp[i] = [0]*7

        # or understand the row-1 are all seeds
        dp[0][6] = 1

        for i in range(1, n+1):
            rowTotal = 0
            for j in range(6):
                k = rollMax[j]
                dp[i][j] = dp[i-1][-1]
                if i-1-k >= 0:
                    dp[i][j] -= dp[i-1-k][-1] - dp[i-1-k][j]
                dp[i][j] %= mod
                rowTotal += dp[i][j] % mod
            dp[i][6] = rowTotal % mod

        return dp[n][6] % (10**9+7)


"""
Runtime: 138 ms, faster than 99.39% of Python3 online submissions for Dice Roll Simulation.
Memory Usage: 15.5 MB, less than 83.44% of Python3 online submissions for Dice Roll Simulation.

"""
