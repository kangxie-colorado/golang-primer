"""
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

so obviously this must be binary search
so let me do this 
1. search for first index < target; which can be -1
2. search for first index > target; which can be len()

actually this edge case can be taken care first
then everything falls into 0, len()-1.. but do I need? let me say I need now

it will tricky
"""


def lastSmaller(nums, target):
    if target <= nums[0]:
        return -1

    l, r = 0, len(nums)-1
    while l < r:
        m = r - (r-l)//2
        if nums[m] >= target:
            r = m - 1
        else:
            l = m

    return l


def firstLarger(nums, target):
    if target >= nums[-1]:
        return len(nums)

    l, r = 0, len(nums)-1
    while l < r:
        m = l + (r-l)//2
        if nums[m] <= target:
            l = m + 1
        else:
            r = m

    return l


class Solution(object):
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        if not nums:
            return [-1, -1]
        l = lastSmaller(nums, target)
        r = firstLarger(nums, target)

        if r-l == 1:
            return [-1, -1]
        return [l+1, r-1]


"""
Runtime: 96 ms, faster than 47.88% of Python online submissions for Find First and Last Position of Element in Sorted Array.
Memory Usage: 14.5 MB, less than 60.41% of Python online submissions for Find First and Last Position of Element in Sorted Array.

Runtime: 64 ms, faster than 93.08% of Python online submissions for Find First and Last Position of Element in Sorted Array.
Memory Usage: 14.9 MB, less than 7.04% of Python online submissions for Find First and Last Position of Element in Sorted Array.
"""
if __name__ == '__main__':
    print(lastSmaller([5, 6, 7, 8, 8, 10], 4))
    print(lastSmaller([5, 6, 7, 8, 8, 10], 11))
    print(lastSmaller([5, 6, 7, 8, 8, 10], 6))
    print(lastSmaller([5, 6, 7, 8, 8, 10], 5))
    print(lastSmaller([5, 6, 7, 8, 8, 10], 7))
    print(lastSmaller([5, 6, 7, 8, 8, 10], 8))
    print(lastSmaller([5, 6, 7, 8, 8, 10], 10))
    print(lastSmaller([5, 6, 7, 8, 8, 10], 9))

    print("")

    print(firstLarger([5, 6, 7, 8, 8, 10], 4))
    print(firstLarger([5, 6, 7, 8, 8, 10], 11))
    print(firstLarger([5, 6, 7, 8, 8, 10], 6))
    print(firstLarger([5, 6, 7, 8, 8, 10], 5))
    print(firstLarger([5, 6, 7, 8, 8, 10], 7))
    print(firstLarger([5, 6, 7, 8, 8, 10], 8))
    print(firstLarger([5, 6, 7, 8, 8, 10], 10))
    print(firstLarger([5, 6, 7, 8, 8, 10], 9))
