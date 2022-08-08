"""
https://leetcode.com/problems/binary-subarrays-with-sum/

Without much thinking
I feel this is a solvable using that atMost(k)?

I also see that extract-ones-indexes then apply left*right..
which is actually pretty simple let me try
"""


from collections import defaultdict
from tkinter import N
from typing import List


class Solution:
    def numSubarraysWithSum(self, nums: List[int], goal: int) -> int:
        def goal0():
            # count 0s, and for each consecutive 0, add the weight by 1
            # 0: 1
            # 0,0.. it add 2 actually [,0] and [0,0]
            # on 1, reset weight
            res = 0
            weight = 0
            for n in nums:
                if n == 0:
                    weight += 1
                else:
                    weight = 0

                res += weight
            return res

        def postiveGoal():
            onesIdx = [-1]
            for i, n in enumerate(nums):
                if n == 1:
                    onesIdx.append(i)
            onesIdx.append(len(nums))

            res = 0

            for j in range(goal, len(onesIdx)-1):
                i = j-goal+1
                res += (onesIdx[i]-onesIdx[i-1])*(onesIdx[j+1]-onesIdx[j])

            return res

        return postiveGoal() if goal > 0 else goal0()


"""
Input: nums = [1,0,1,0,1], goal = 2
hmm.. this one right..

Input: nums = [0,0,0,0,0], goal = 0
this one wrong...
yeah, this solution is for solving 1s... not all zeros..

so not really widely applicable..haha

it probably can solve goal>0
but cannot get to goal == 0

hmm.. but interesting to make it work as edge case


then I also notice another point is goal can be bigger than sum(nums)
but somehow for j in range(goal, len(onesIdx)-1): takes care of that... but this is an unreliable point
>>> [i for i in range(5,4)]
[]

unreliable... huh.. maybe not.. the loop is just going to end before start.. not an invalid concept

allright.. it passed
Runtime: 430 ms, faster than 53.75% of Python3 online submissions for Binary Subarrays With Sum.
Memory Usage: 15.9 MB, less than 65.05% of Python3 online submissions for Binary Subarrays With Sum.

now let me think another solution

so I am thinking still apply the typical sliding window but on zeros..
let me do special handling?
hmm.. 

still.. [0,1,0,1], the last one is also like need to add i+1 (last-one-idx +1)
so a queue.. O(n) space...

so there is no special handling for zeros.. but when goal is 0...
this acutally returns to above solution... shit...

hmm.. why not I count zeros and make it prefixes.. which will be reflected on the res
prefix=1 then each subarray it prefixes will add 1

when to reset prefix? when it > goal, right? still not correct... think [1,0,1,0,0,0,0,1,1] k=2


"""


class Solution:
    def numSubarraysWithSum(self, nums: List[int], goal: int) -> int:

        def atMost(k):
            i, j = 0, 0
            res, sum = 0, 0

            while j < len(nums):
                sum += nums[j]
                while i < len(nums) and sum > k:
                    sum -= nums[i]
                    i += 1

                res += max(j-i+1, 0)
                j += 1

            return res

        return atMost(goal) - atMost(goal-1)


"""
Runtime: 721 ms, faster than 5.08% of Python3 online submissions for Binary Subarrays With Sum.
Memory Usage: 14.9 MB, less than 88.81% of Python3 online submissions for Binary Subarrays With Sum.

oh.. I am a little fuzzy now but it passes?

lee's code is 
    def numSubarraysWithSum(self, A, S):
        def atMost(S):
            if S < 0: return 0
            res = i = 0
            for j in xrange(len(A)):
                S -= A[j]
                while S < 0:
                    S += A[i]
                    i += 1
                res += j - i + 1
            return res
        return atMost(S) - atMost(S - 1)

        but can it deal with -1
        ah... if S < 0: return 0

so what is special about this is 
res += max(j-i+1, 0)

because it is at most
so when [1,0,1] sums to 2.. there is actually 6 subarray..
[1] [1,0] [0] [1,0,1] [0,1] [1]

so each time you add the length.. you will end up like 6 
it is actually 1+2+3

because the whole is at most, so any sub-sum is also at-most..

"""


"""
then I see that pathsum solution
I kind of touched on that for a moment but didn't dive deeper into that idea

let me do that
"""


class Solution:
    def numSubarraysWithSum(self, nums: List[int], goal: int) -> int:
        pathSumCount = defaultdict(int)
        # this is a crucial init state
        # without it, it will not be okay.. how to explain?
        # zero sum is always a case.. because no element is needed
        pathSumCount[0] = 1
        res = 0
        pathsum = 0
        for n in nums:
            pathsum += n
            res += pathSumCount[pathsum-goal]
            pathSumCount[pathsum] += 1

        return res


"""
    print(s.numSubarraysWithSum([1, 0, 1, 0, 1], 2))
    print(s.numSubarraysWithSum([0, 0, 0, 0, 0], 0))

    shit I got 2 and 10 (should be 4 15)
    so really walk this thru on paper or board first

Runtime: 442 ms, faster than 50.52% of Python3 online submissions for Binary Subarrays With Sum.
Memory Usage: 17.8 MB, less than 25.72% of Python3 online submissions for Binary Subarrays With Sum.

the important m[0]=1
"""

if __name__ == '__main__':
    s = Solution()
    print(s.numSubarraysWithSum([1, 0, 1, 0, 1], 2))
    print(s.numSubarraysWithSum([0, 0, 0, 0, 0], 0))
