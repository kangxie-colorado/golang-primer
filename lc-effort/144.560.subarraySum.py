"""
https://leetcode.com/problems/subarray-sum-equals-k/


this feels very much like the path sum III 
and there is only one branch.. 

so actually easier..
"""


from collections import defaultdict


class Solution(object):
    def subarraySum(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        presum = 0
        count = 0
        pathSumMap = defaultdict(int)
        pathSumMap[0] = 1
        for i in enumerate(nums):
            nums[i] += presum
            presum = nums[i]

            count += pathSumMap[nums[i]-k]
            pathSumMap[nums[i]] += 1

        return count


"""
Runtime: 358 ms, faster than 72.47% of Python3 online submissions for Subarray Sum Equals K.
Memory Usage: 17.9 MB, less than 14.11% of Python3 online submissions for Subarray Sum Equals K.
"""
