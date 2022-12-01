"""
https://leetcode.com/problems/maximum-points-in-an-archery-competition/

this feels like a two-dimensional knapsack DP
the tweak is any cell will be conditional

not all clear but guess I can have a try
"""


from typing import List


class Solution:
    def maximumBobPoints(self, numArrows: int, aliceArrows: List[int]) -> List[int]:
        dp = [[]]*(numArrows+1)
        for r in range(len(dp)):
            dp[r] = [-1]*12
        for i in range(0, 12):
            dp[0][i] = 0
        for r in range(numArrows+1):
            dp[r][0] = 0

        def f(n, i):
            if n <= 0 or i <= 0:
                return 0
            if dp[n][i] != -1:
                return dp[n][i]

            # take or notTake this section
            take = 0
            if n > aliceArrows[i]:
                take = f(n-aliceArrows[i]-1, i-1)+i
            notTake = f(n, i-1)

            dp[n][i] = max(take, notTake)
            return dp[n][i]

        maxPoints = f(numArrows, 11)

        # use maxPoints to deduct the arrow distributions???
        # using backtrack??
        res = []

        def backtrack(points, n, i, run):
            if n < 0 or i < 0:
                return False
            if (i == 0 or n == 0) and points > 0:
                return False
            if points == 0:
                nonlocal res
                res = run.copy()
                return True

            run[i] = aliceArrows[i]+1
            if backtrack(points - i, n-aliceArrows[i]-1, i-1, run):
                return True
            run[i] = 0

            return backtrack(points, n, i-1, run)

        backtrack(maxPoints, numArrows, 11, [0]*12)
        if sum(res) < numArrows:
            res[0] = numArrows - sum(res)
        return res


"""
Runtime: 8071 ms, faster than 5.23% of Python3 online submissions for Maximum Points in an Archery Competition.
Memory Usage: 45.3 MB, less than 8.50% of Python3 online submissions for Maximum Points in an Archery Competition.

wow...
even mirrerable still super happy!!
"""

"""
cannot get it right.

how about I take the score max
then I reverse to get a combination to reach that score??

then I think I only need to do backtrack to find all combinations

"""


class Solution:
    def maximumBobPoints(self, numArrows: int, aliceArrows: List[int]) -> List[int]:
        res = []

        def backtrack(n, i, run):
            if n < 0:
                return
            if i == 0:
                run[0] = n
                res.append(
                    (sum([i for i in range(12) if run[i] > 0]), run.copy()))
                return

            run[i] = aliceArrows[i]+1
            backtrack(n-aliceArrows[i]-1, i-1, run)
            run[i] = 0
            backtrack(n, i-1, run)

        backtrack(numArrows, 11, [0]*12)
        return max(res)[1]


if __name__ == '__main__':
    s = Solution()
    print(s.maximumBobPoints(numArrows=3, aliceArrows=[
          0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2]))
    print(s.maximumBobPoints(8,
                             [0, 1, 0, 2, 0, 0, 1, 0, 1, 0, 1, 2]))

    print(s.maximumBobPoints(7,
                             [0, 0, 1, 1, 3, 0, 0, 0, 0, 2, 0, 0]))
    print(s.maximumBobPoints(numArrows=9, aliceArrows=[
          1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0]))
    print(s.maximumBobPoints(numArrows=3, aliceArrows=[
          0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2]))
