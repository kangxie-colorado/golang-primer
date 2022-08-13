"""
https://leetcode.com/problems/subsets-ii/

backtracking?
how to deal with duplicates?

maybe like previous problem combSum..
j>i.. (but it depends on how many duplicates are allowed)
like 2 2 2...
2
2 2
2 2 2 are valid

but [2] [2]
or [2 2] [ 2 2] are not..
thinking the number of elements are kind of a condition/state to track/check too

let me see.
because 1 <= nums.length <= 10
so I will just enuermate [] to full

actually [] corresponds to full
1 coorsponds to 9...

so only [] to len()//2 then you get another half..

"""


from re import sub
from typing import List


class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        subsets = [[]] * (len(nums)+1)
        for i in range(len(nums)+1):
            subsets[i] = []
        nums.sort()

        def backtracking(size, run, i):
            if len(run) == size:
                subsets[size].append(run.copy())
                return

            for j in range(i, len(nums)):
                if j > i and nums[j] == nums[j-1]:
                    continue

                run.append(nums[j])
                backtracking(size, run, j+1)
                run.pop()

        def getOpposite(nums, c):
            for n1 in c:
                for i in range(len(nums)):
                    if nums[i] == n1:
                        nums[i] += 100
                        break

            return [n for n in nums if n <= 10]

        for s in range(0, len(nums)//2+1):
            backtracking(s, [], 0)
            if s*2 != len(nums):
                for c in subsets[s]:
                    subsets[len(nums)-s].append(getOpposite(nums.copy(), c))

        return [s for subset in subsets for s in subset]


"""
when I am troubleshooting
it gives a slew of insight.. to deal with duplicates.. just focus on one layer at a time
"""

"""
Runtime: 47 ms, faster than 78.25% of Python3 online submissions for Subsets II.
Memory Usage: 14.2 MB, less than 49.24% of Python3 online submissions for Subsets II.

turns out the criterea to rid of duplicates are still
                if j > i and nums[j] == nums[j-1]:
                    continue

                j>i.. if the layer its on
                nums[j] == nums[j-1].. basically means it is duplicate number

                I over-think sometimes..
        basically, the duplicates we want to remove is for the element to start at its layer
        so i is the first element in the layer.. j>i, then j can only be 2nd element 
        nums[j] == nums[j-1].. this simplifies thing a lot.. bascially just it means nums[j] has appeared before

Runtime: 35 ms, faster than 97.32% of Python3 online submissions for Subsets II.
Memory Usage: 14.1 MB, less than 49.24% of Python3 online submissions for Subsets II.

code is efficient but ugly..
"""


if __name__ == '__main__':
    print(Solution().subsetsWithDup([1]))
    print(Solution().subsetsWithDup([1, 2, 2]))
