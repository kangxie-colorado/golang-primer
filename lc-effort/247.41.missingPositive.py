"""
https://leetcode.com/problems/first-missing-positive/

did two old problems without much trouble
so pick-one this

hard! but I'll just try and if not solvable by myself.
I'll just read and understand 

1hr top time


why there is negative number?
if I see a number... 

>> You must implement an algorithm that runs in O(n) time and uses constant extra space.
if I can use extra space, just use a hash table
and test 0/1/2/3... then it shows itself


the first hint
Think about how you would solve the problem in non-constant space. Can you apply that logic to the existing space?


yes, if with extra memory
I'll just go do a set()
then check if 1 is in the set, then 2.. then 3..

also I saw the relationship of the number of missing can only be in [1..n+1], where n is the size of nums 
because when the arr is 5, it must contain 1 to 5 in order not to miss any..
if it contains any number>5, then 1 to 5 must miss one...

therefor this gives us a solution
1st pass, just get rid of all negatives, set them to 0; also get rid of number >=n because they should be considered 
    useless, if a number >=n appears, then it means there must be a number in <n missing
2nd pass, use the index has the hash, put 1 to index 0 (why not 1, think [1,2,3,4] where to to put 4...)
3rd pass, go thru the array again and check the first out-of-place number

"""

import enum
from typing import List


class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        m = len(nums)
        # pass 1: get rid of useless numbers
        for i in range(len(nums)):
            if nums[i] < 0 or nums[i] >= m:
                nums[i] = 0

        # pass 2: put number into its position
        for i, n in enumerate(nums):
            # [3,0,2] -- there could be more than one swap?
            # 3 to idx-2, swap 2 to idx-2, hmm.. maybe not
            if n != 0 and n != i+1:
                nums[n-1] = n

        # pass 3: scan for first missing
        for i, n in enumerate(nums):
            if i+1 != n:
                return i+1

        return m+1


"""
failed at [1]
return 1 exp 2
aha. I should return m+1

hmm.. not that...
let me think over
this is positive number... so 1..n.. I >=m, I should >m

but I do that I messed up somewhere else... let me see
maybe the swapping?
"""


class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        m = len(nums)
        # pass 1: get rid of useless numbers
        for i in range(len(nums)):
            if nums[i] < 0 or nums[i] > m:
                nums[i] = 0

        # pass 2: put number into its position
        for i, n in enumerate(nums):
            # [3,0,2] -- there could be more than one swap?
            # 3 to idx-2, swap 2 to idx-2, hmm.. maybe not
            #  nums[i] != nums[nums[i]-1] this is to deal with duplicates...and dead loop
            while nums[i] != 0 and nums[i] != i+1 and nums[i] != nums[nums[i]-1]:
                n = nums[i]
                nums[n-1], nums[i] = nums[i], nums[n-1]

        # pass 3: scan for first missing
        for i, n in enumerate(nums):
            if i+1 != n:
                return i+1

        return m+1


"""
Runtime: 691 ms, faster than 74.53% of Python3 online submissions for First Missing Positive.
Memory Usage: 27.2 MB, less than 82.64% of Python3 online submissions for First Missing Positive.
"""

if __name__ == "__main__":

    s = Solution()
    print(s.firstMissingPositive([1]))
    print(s.firstMissingPositive([1, 2, 0]))
    print(s.firstMissingPositive([3, 4, 1, 1]))
    print(s.firstMissingPositive([3, 4, -1, 1]))
    print(s.firstMissingPositive([7, 8, 9, 11, 12, 13]))
