"""
https://leetcode.com/problems/wiggle-sort/

O(n)?
buckets? because 0 <= nums[i] <= 10^4

"""


from typing import List


class Solution:
    def wiggleSort(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        buckets = [0]*10001
        for n in nums:
            buckets[n] += 1

        i = 0
        minBucket, maxBucket = 0, 10000
        while i+1 < len(nums):
            # pay attention to here... because you need to operate two slots, so i+1
            while buckets[minBucket] == 0 and minBucket < 10000:
                minBucket += 1
            nums[i] = minBucket
            buckets[minBucket] -= 1

            while buckets[maxBucket] == 0 and maxBucket >= 0:
                maxBucket -= 1
            nums[i+1] = maxBucket
            buckets[maxBucket] -= 1

            i += 2

        if i < len(nums):
            while buckets[minBucket] == 0 and minBucket < 10000:
                minBucket += 1
            for j in range(buckets[minBucket]):
                nums[i] = minBucket
                i += 1


"""
edge cases are plenty!!!!!

1 <= nums.length <= 5 * 10^4
^ tells you the min length can be 1

public void wiggleSort(int[] nums) {
    for (int i = 0; i < nums.length - 1; i++) {
        if (((i % 2 == 0) && nums[i] > nums[i + 1])
                || ((i % 2 == 1) && nums[i] < nums[i + 1])) {
            swap(nums, i, i + 1);
        }
    }
}

okay
parts impact the whole...
"""
