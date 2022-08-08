"""
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

in-place.. keep the relative order
so naturally this is maintain a writeIdx and readIdx

for state tracking
last.
countOfLast.

update them according.. because this is allowing at most two duplicates
"""


from os import read
from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        readIdx = writeIdx = 1
        last = nums[0]
        lastCount = 1
        while readIdx < len(nums):
            if nums[readIdx] == last:
                lastCount += 1
                if lastCount <= 2:
                    nums[writeIdx] = nums[readIdx]
                    writeIdx += 1
            else:
                # new number appears
                last = nums[readIdx]
                lastCount = 1
                nums[writeIdx] = nums[readIdx]
                writeIdx += 1
            readIdx += 1
        return writeIdx


"""
Runtime: 54 ms, faster than 96.48% of Python3 online submissions for Remove Duplicates from Sorted Array II.
Memory Usage: 13.9 MB, less than 30.38% of Python3 online submissions for Remove Duplicates from Sorted Array II.

pretty good already but the code can be more concise
like below...
like a sliding window?

"""


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        readIdx = writeIdx = 1
        last = nums[0]
        lastCount = 1
        while readIdx < len(nums):
            if nums[readIdx] == last:
                lastCount += 1
            else:
                # new number appears
                last = nums[readIdx]
                lastCount = 1
            if lastCount <= 2:
                nums[writeIdx] = nums[readIdx]
                writeIdx += 1
            readIdx += 1
        return writeIdx


"""
yeah.. like a sliding window..
only the state for the window if lastCount
"""


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        last = nums[0]
        lastCount = 1
        readIdx = writeIdx = 1
        while readIdx < len(nums):
            lastCount = (lastCount+1) if nums[readIdx] == last else 1
            if lastCount <= 2:
                nums[writeIdx] = nums[readIdx]
                writeIdx += 1

            last = nums[readIdx]
            readIdx += 1
        return writeIdx


"""
huh.. this failed?
I forget to update last...

Runtime: 61 ms, faster than 87.27% of Python3 online submissions for Remove Duplicates from Sorted Array II.
Memory Usage: 13.8 MB, less than 74.15% of Python3 online submissions for Remove Duplicates from Sorted Array II.

and this one

def removeDuplicates(self, nums):
    i = 0
    for n in nums:
        if i < 2 or n > nums[i-2]:
            nums[i] = n
            i += 1
    return i

what the heck? he just has one more eye?

"""

if __name__ == '__main__':

    print(Solution().removeDuplicates([1, 1, 1, 2, 2, 3]))
    a = [0, 0, 1, 1, 1, 1, 2, 3, 3]
    print(Solution().removeDuplicates(a))
