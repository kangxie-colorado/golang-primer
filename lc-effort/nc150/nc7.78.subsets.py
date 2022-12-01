"""
https://leetcode.com/problems/subsets/

backtrack or build up?

backtrack: two choices.. select or not
"""


from tkinter import N
from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []

        def helper(run, i):
            if i == len(nums):
                res.append(run.copy())
                return

            helper(run, i+1)
            helper(run+[nums[i]], i+1)

        helper([], 0)
        return res


"""
Runtime: 73 ms, faster than 9.95% of Python3 online submissions for Subsets.
Memory Usage: 14.2 MB, less than 36.16% of Python3 online submissions for Subsets.

so this is not back track.. just a recursive
with backtrack usually it comes with for loop
this has just two branches : choose or not-choose

I'd better recetify that

Runtime: 39 ms, faster than 86.15% of Python3 online submissions for Subsets.
Memory Usage: 14.2 MB, less than 36.16% of Python3 online submissions for Subsets.

what are other solutions?

indeed there is this backtrack 
backtrack usually won't open two branches.. just open/close let me do that
"""


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []

        def backtrack(run, i):
            # every run is valid, just add
            # no need to test the terminal condition.. it will end by itself
            res.append(run.copy())

            for j in range(i, len(nums)):
                # on the loop.. include nums[j] for a run
                # the not include.. for a run of rest
                # this is basically the same thing of the dfs-helper
                run.append(nums[j])
                backtrack(run, j+1)
                run.pop()

        backtrack([], 0)
        return res


"""
this works
yeah... feel the subtle differences...
"""

"""
also I felt like an iterative solution but didn't pan out
I wasn't thinking about the right middle results so it should be like this

Iterative

Using [1, 2, 3] as an example, the iterative process is like:

Initially, one empty subset [[]]
Adding 1 to []: [[], [1]];
Adding 2 to [] and [1]: [[], [1], [2], [1, 2]];
Adding 3 to [], [1], [2] and [1, 2]: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]].
"""


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = [[]]

        for n in nums:
            nextRes = []
            for r in res:
                nextRes.append([n]+r)
            res.extend(nextRes)

        return res


"""
ah yes, the bit manipulation 
[1,2,3] will be eight 
0-7
000 001 010 ... 111
"""


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []
        p = 2**len(nums)

        for i in range(p):
            run = []
            for j in range(len(nums)):
                if (i >> j) & 1:
                    run.append(nums[j])
            res.append(run)
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.subsets([1, 2, 3]))
