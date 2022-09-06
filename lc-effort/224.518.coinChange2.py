"""
https://leetcode.com/problems/coin-change-2/

okay... I came here from 377 combSum 4
I know the subtle behind it

but kind of hard to explain (this or that)
that might seems natural... 
"""


from typing import List


class SolutionDP:
    def change(self, amount: int, coins: List[int]) -> int:
        dp = [0] * (amount+1)
        dp[0] = 1

        for c in coins:
            for i in range(1, amount+1):
                if i >= c:
                    dp[i] += dp[i-c]
        return dp[amount]


"""
Runtime: 194 ms, faster than 88.34% of Python3 online submissions for Coin Change 2.
Memory Usage: 14 MB, less than 74.91% of Python3 online submissions for Coin Change 2.


of course this would work
notice that solution 
class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        dp = [0]*(target+1)
        dp[0] = 1

        for i in range(1, target+1):
            for n in nums:
                if i >= n:
                    dp[i] += dp[i-n]

        return dp[target]


just switching the order of lopp to take care of the duplicates.. 
hmm... very particular... 

so the coin change2 loop is kinda of saying
for i, I can use up to c-th coins.. (1) or (1,2) or (1,2,5) -- of course the use order doesn't matter

for the comb sum4 loop is kinda of saying
for target i, I can use all numbers each time..

logically it is like so.. 
the combination is very hard to solve it use backtracking
maybe the coin change2 which don't allowing backforward looking can be solved it that way
but could TLE... I'd like to try
"""


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        res = 0

        def backtrack(remain, i):
            nonlocal res
            if remain == 0:
                res += 1
                return
            if remain < 0 or i == len(coins):
                return

            for j in range(i, len(coins)):
                for count in range(1, remain//coins[j]+1):
                    backtrack(remain-coins[j]*count, j+1)

        backtrack(amount, 0)
        return res


"""
as expected 14 / 28 test cases passe and TLE    

not too far really
500
[3,5,7,8,9,10,11]

but at least this is correct I think 


            for j in range(i, len(coins)):
                for count in range(1, remain//coins[j]+1):
                                   ^ this is important
                    backtrack(remain-coins[j]*count, j+1


so what is critical to get rid of duplicates 
1. of course, don't look backwards, j starts with i (i is passed in as last_j+1)
2. count start with 1... not 0

if you start 0, you are double counting the latter combinations 
e.g.
[1,2] to get 2

0 1s, you got [2]
1 1s.. nothing
2 1s.. you got [1 1]

then move on to 2.. (top loop)
0 2s, nothing
1 2s.. pfft.. this is [2] again
..

so you cannot use 0 1s.. 0 1s is kind of implicitly done by later backtracking..

I wonder do I need a top loop???

"""


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        res = 0

        def backtrack(remain, i):
            nonlocal res
            if remain == 0:
                res += 1
                return
            if remain < 0 or i == len(coins):
                return

            for count in range(0, remain//coins[i]+1):
                backtrack(remain-coins[i]*count, i+1)

        backtrack(amount, 0)
        return res


"""
right.. so I remove that loop
rewrite like this.. it also TLE after 14 cases

just verify the correctness
100
[3,5,7,8,9,10,11]

6606 vs 6606.. so albeit very low performant.. still correct
this is the the brute force god damn it.. how fast you can ask?

but kind of means.. for such problem.. this is defintely a DP problem..
"""

"""
reading 
https://leetcode.com/problems/coin-change-2/discuss/99212/Knapsack-problem-Java-solution-with-thinking-process-O(nm)-Time-and-O(m)-Space

so this DP is kind of evovlved from 2-D DP

dp[i][j] : the number of combinations to make up amount j by using the first i types of coins
State transition:

1. not using the ith coin, only using the first i-1 coins to make up amount j, then we have dp[i-1][j] ways.
2. using the ith coin, since we can use unlimited same coin, we need to know how many ways to make up amount j - coins[i-1] by using first i coins(including ith), which is dp[i][j-coins[i-1]]
Initialization: dp[i][0] = 1

for a while I didn't do DP problems
kind of forgot the tricks.. it is scary so no matter how hard you do you cannot remember
but also assuring.. looking at it for 2 secs, I understood

hmm.. explains why I feel this one-d is sort of out of context.. but kind of make sense..
let me the do 2-d as well.

also this dancing lady is kind of brilliant too

https://leetcode.com/problems/coin-change-2/discuss/141076/Unbounded-Knapsack

and that tells me the backtrack (which is actually not a backtracking can be solved by memorization)
let me do both then off to lunch
"""


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        cache = [[]]*(len(coins)+1)
        for i in range(len(cache)):
            cache[i] = [0] * (amount+1)

        def backtrack(remain, i):
            if remain == 0:
                return 1
            if remain < 0 or i == len(coins):
                return 0

            if cache[i][remain] != 0:
                return cache[i][remain]

            for count in range(0, remain//coins[i]+1):
                cache[i][remain] += backtrack(remain-coins[i]*count, i+1)

            return cache[i][remain]

        return backtrack(amount, 0)


"""
hmm... still TLE.. passed 19 tests
failed here.. 

I see.. what I am doing.. differectly that hers
        // Recursive call after selecting the coin at the currentIndex
        int sum1 = changeFrom(amount - coins[currentIndex], coins, currentIndex);

        // Recursive call after excluding the coin at the currentIndex
        int sum2 = changeFrom(amount, coins, currentIndex + 1);

        dp[currentIndex][amount] = sum1 + sum2;

logically the same? but I didn't cache enough.. because I jumped in a few places?
so this gives me a lesson.. 

this is subtle difference between backtracking and two-choices..

"""


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        cache = [[]]*(len(coins)+1)
        for i in range(len(cache)):
            cache[i] = [0] * (amount+1)

        def backtrack(remain, i):
            if remain == 0:
                return 1
            if remain < 0 or i == len(coins):
                return 0

            if cache[i][remain] != 0:
                return cache[i][remain]

            # not to pick i-th coin
            cache[i][remain] += backtrack(remain, i+1)
            # pick i-th coin, then I can use more i-th coin so not to move i
            cache[i][remain] += backtrack(remain-coins[i], i)

            return cache[i][remain]

        return backtrack(amount, 0)


"""
yet.. still TLE.. with one more passed cases 20 / 28 test cases passed.
Status: Time Limit Exceeded

but her way of thinking is kind of more generic.. I learn 

pretty/sexy/smart women.. admire 

now the two D DP
"""


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        dp = [[]]*(len(coins)+1)
        for r in range(len(dp)):
            dp[r] = [0]*(amount+1)
            dp[r][0] = 1

        for i in range(1, len(coins)+1):
            for j in range(1, amount+1):
                # exclude and include
                # is not j-i but j-count[i]; weights.. value...
                dp[i][j] += dp[i-1][j]
                if j >= coins[i-1]:
                    dp[i][j] += dp[i][j-coins[i-1]]

        return dp[len(coins)][amount]


"""
Runtime: 570 ms, faster than 46.12% of Python3 online submissions for Coin Change 2.
Memory Usage: 34.7 MB, less than 29.33% of Python3 online submissions for Coin Change 2.

I see what is wrong with the memorization technique..
the init value I used is 0, but 0 is kind of a valid value
I discarded so many

so use -1 let me try
"""


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        cache = [[]]*(len(coins)+1)
        for i in range(len(cache)):
            cache[i] = [-1] * (amount+1)

        def backtrack(remain, i):
            if remain == 0:
                return 1
            if remain < 0 or i == len(coins):
                return 0

            if cache[i][remain] != -1:
                return cache[i][remain]

            cache[i][remain] = 0
            for count in range(0, remain//coins[i]+1):
                cache[i][remain] += backtrack(remain-coins[i]*count, i+1)

            return cache[i][remain]

        return backtrack(amount, 0)


"""
Runtime: 3830 ms, faster than 5.01% of Python3 online submissions for Coin Change 2.
Memory Usage: 32.6 MB, less than 30.38% of Python3 online submissions for Coin Change 2.

bearly..
and hers
"""


class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        cache = [[]]*(len(coins)+1)
        for i in range(len(cache)):
            cache[i] = [-1] * (amount+1)

        def backtrack(remain, i):
            if remain == 0:
                return 1
            if remain < 0 or i == len(coins):
                return 0

            if cache[i][remain] != -1:
                return cache[i][remain]

            cache[i][remain] = 0
            # not to pick i-th coin
            cache[i][remain] += backtrack(remain, i+1)
            # pick i-th coin, then I can use more i-th coin so not to move i
            cache[i][remain] += backtrack(remain-coins[i], i)

            return cache[i][remain]

        return backtrack(amount, 0)


"""
Runtime: 1365 ms, faster than 16.93% of Python3 online submissions for Coin Change 2.
Memory Usage: 38.5 MB, less than 22.72% of Python3 online submissions for Coin Change 2.

hmm.. indeed better...
"""

if __name__ == '__main__':
    s = Solution()
    print(s.change(amount=1, coins=[1, 2]))
    print(s.change(amount=2, coins=[1, 2]))
    print(s.change(amount=100, coins=[3, 5, 7, 8, 9, 10, 11]))

    print(s.change(amount=3, coins=[1, 2]))
    print(s.change(amount=5, coins=[1, 2, 5]))
