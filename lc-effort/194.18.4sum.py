"""
https://leetcode.com/problems/4sum/

maybe I can use 3sum.. to jump over the used numbers
or maybe I can do what I did in combSum2

just try
"""


from collections import Counter
from typing import List


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        for i in range(len(nums)):
            nums[i] += 10**9
        C = Counter(nums)
        nums.sort()
        res = []
        target += 4*10**9

        def helper(run, S, i):
            nonlocal res
            if len(run) == 4 and S == target:
                res.append(run.copy())
                return
            if len(run) > 4 or S > target or i == len(nums):
                return

            for j in range(i, len(nums)):
                # use i to de-duplicate?
                # how?
                if j > i and nums[j] == nums[j-1]:
                    continue
                run.append(nums[j]-10**9)
                helper(run, S+nums[j], j+1)
                run.pop()

        helper([], 0, 0)
        return res


"""
I know it is going to TLE..
and it passed 220 / 291 test cases passed.

so at least the duplicate is taken care of..
how?
                if j > i and nums[j] == nums[j-1]:
                    continue
            so conceptually I kind of grab something...
            but not very clearly..
            so if this is only one layer of loop.. j>0 and nums[j] == nums[j-1]:

            but on layer i, it should be j>i..
            this is probably appliable to prev-combSum problem too..

failed TLE here
[-500,-481,-480,-469,-437,-423,-408,-403,-397,-381,-379,-377,-353,-347,-337,-327,-313,-307,-299,-278,-265,-258,-235,-227,-225,-193,-192,-177,-176,-173,-170,-164,-162,-157,-147,-118,-115,-83,-64,-46,-36,-35,-11,0,0,33,40,51,54,74,93,101,104,105,112,112,116,129,133,146,152,157,158,166,177,183,186,220,263,273,320,328,332,356,357,363,372,397,399,420,422,429,433,451,464,484,485,498,499]
2139

>>> len(a)
89

okay.. now let me do the 3sum ways..

"""


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()

        def twoSumSorted(nums, target):
            i, j = 0, len(nums)-1
            res = []

            while i < j:
                S = nums[i]+nums[j]
                if S == target:
                    res.append([nums[i], nums[j]])
                    while i < j and nums[i] == nums[i+1]:
                        i += 1
                    while i < j and nums[j] == nums[j-1]:
                        j -= 1

                    # this also deals with the situation where there are no duplicates
                    # so it looks weird to +=1 again for dupliates.
                    # but it is actually more generic way to deal with both
                    i += 1
                    j -= 1

                elif S < target:
                    i += 1
                else:
                    j -= 1
            return res

        def threeSumSorted(nums, target):
            res = []
            i = 0
            while i < len(nums):
                n = nums[i]
                for twoSum in twoSumSorted(nums[i+1:], target-n):
                    if len(twoSum) == 2:
                        res.append([n] + twoSum)
                while i < len(nums)-1 and nums[i] == nums[i+1]:
                    i += 1
                i += 1

            return res

        res = []
        i = 0
        while i < len(nums):
            n = nums[i]
            for threeSum in threeSumSorted(nums[i+1:], target-n):
                if len(threeSum) == 3:
                    res.append([n]+threeSum)
            while i < len(nums)-1 and nums[i] == nums[i+1]:
                i += 1
            i += 1

        return res


"""
Runtime: 994 ms, faster than 63.00% of Python3 online submissions for 4Sum.
Memory Usage: 13.9 MB, less than 64.57% of Python3 online submissions for 4Sum.

notice how the code is so the same between four/three-sum
wonder if I can improve that
"""


class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        nums.sort()

        def oneSumSorted(nums, target, nextLevelFunc=None):
            if target in nums:
                return [[target]]
            return [[]]

        def twoSumSorted(nums, target, nextLevelFunc):
            res = []
            i = 0
            while i < len(nums):
                n = nums[i]
                for paritialSum in nextLevelFunc(nums[i+1:], target-n):
                    if len(paritialSum) == 1:
                        res.append([n] + paritialSum)
                while i < len(nums)-1 and nums[i] == nums[i+1]:
                    i += 1
                i += 1

            return res

        def threeSumSorted(nums, target, nextLevelFunc):
            res = []
            i = 0
            while i < len(nums):
                n = nums[i]
                for paritialSum in nextLevelFunc(nums[i+1:], target-n, oneSumSorted):
                    if len(paritialSum) == 2:
                        res.append([n] + paritialSum)
                while i < len(nums)-1 and nums[i] == nums[i+1]:
                    i += 1
                i += 1

            return res

        def fourSumSorted(nums, target, nextLevelFunc):

            res = []
            i = 0
            while i < len(nums):
                n = nums[i]
                for paritialSum in nextLevelFunc(nums[i+1:], target-n, twoSumSorted):
                    if len(paritialSum) == 3:
                        res.append([n]+paritialSum)
                while i < len(nums)-1 and nums[i] == nums[i+1]:
                    i += 1
                i += 1
            return res

        return fourSumSorted(nums, target, threeSumSorted)


"""
first I refactor the four/three/two.. to be very the same

albeit slower still pass
Runtime: 4556 ms, faster than 5.03% of Python3 online submissions for 4Sum.
Memory Usage: 14.1 MB, less than 33.64% of Python3 online submissions for 4Sum.

2nd, unify some of the variable names.. 
threeSum
twoSum
oneSum

to paritialSum
now there is only one difference, the func it calls..
extrace a variable to nextLevelFunc

yeah.. it beautifies a bit
but then what?

maybe some chain?
not sure.. pause here for now
"""


if __name__ == '__main__':

    print(Solution().fourSum([1, 0, -1, 0, -2, 2], 0))
    print(Solution().fourSum([2, 2, 2, 2, 2], 8))
