"""
https://leetcode.com/problems/jump-game-ii/

I did this before but cannot recall anything
I experimented the closes and farthest then use that to build DP table
but not able to make it right   
e.g. [4,1,1,3,1,1,1]

the immediatary results.
[7, 7, 7, 7, 7, 1, 0]
[7, 7, 7, 7, 2, 1, 0]
[7, 7, 7, 1, 2, 1, 0]
[7, 7, 2, 1, 2, 1, 0]
[7, 3, 2, 1, 2, 1, 0]
[3, 3, 2, 1, 2, 1, 0]
the problem is with the first 4.. it jumps over 3 but I didn't capture this information at all

but looking at this I think I should thus maintain a monotonic increasing stack...
or queue... or anything...

"""


from collections import deque
from typing import Deque, List


class Solution:
    def jump(self, nums: List[int]) -> int:
        n = len(nums)
        dp = [n]*n
        dp[-1] = 0

        for i in range(n-2, -1, -1):

            if nums[i]:
                dp[i] = min(dp[i+1]+1, dp[min(i+nums[i], n-1)]+1)
                while dp[i] < dp[i+1]:
                    dp[i+1] = dp[i]
            print(dp)
        return dp[0]


"""
Runtime: 377 ms, faster than 34.66% of Python3 online submissions for Jump Game II.
Memory Usage: 15 MB, less than 58.70% of Python3 online submissions for Jump Game II.

huh... 

Runtime: 171 ms, faster than 79.34% of Python3 online submissions for Jump Game II.
Memory Usage: 14.9 MB, less than 90.60% of Python3 online submissions for Jump Game II.

huh... not bad ?

Runtime: 169 ms, faster than 80.16% of Python3 online submissions for Jump Game II.
Memory Usage: 15.1 MB, less than 58.70% of Python3 online submissions for Jump Game II.

okay... after reading the previous code and discussions
the jump by steps/intervals came back to me

the thing is to go thru a begin:end interval and see the furthese it can jump to
when i reachs the end... set begin to end and end to furthest to jump another round..

I think this might be used to solve jump game 1 as well
"""


class Solution:
    def jump(self, nums: List[int]) -> int:
        begin = end = furthest = 0
        step = 0
        while end < len(nums)-1:
            step += 1
            for i in range(begin, end+1):
                furthest = max(furthest, nums[i]+i)

            if furthest >= len(nums)-1:
                return step

            begin, end = end+1, furthest

        return step


"""
debug/modify and whatever... guessing
finally right

Runtime: 289 ms, faster than 39.78% of Python3 online submissions for Jump Game II.
Memory Usage: 15.1 MB, less than 58.70% of Python3 online submissions for Jump Game II.
"""


if __name__ == "__main__":
    s = Solution()
    print(s.jump([1, 2, 3]))
