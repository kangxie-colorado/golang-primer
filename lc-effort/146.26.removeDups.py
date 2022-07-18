"""
https://leetcode.com/problems/remove-duplicates-from-sorted-array/

I think maintain two indexes

wIdx, which just going forward when i!=j... 

"""


class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        wIdx = 0
        i = 0
        while i < len(nums):
            while i+1 < len(nums) and nums[i] == nums[i+1]:
                # two ways to end this
                # 1. A[i] != A[i+1], then i is pointing at the last dup
                # 2. i+1 is len, then i is pointing at the last element
                i += 1

            nums[wIdx] = nums[i]
            wIdx += 1
            i += 1
        return wIdx
