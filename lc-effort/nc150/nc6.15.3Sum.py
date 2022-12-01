"""
https://leetcode.com/problems/3sum/

hello old friend
"""


from tkinter.messagebox import RETRY
from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()

        def twoSum(S, A):
            l, r = 0, len(A)-1
            res = []
            while l < r:
                if A[l] + A[r] == S:
                    res.append([A[l], A[r]])
                    while l+1 < len(A)-1 and A[l+1] == A[l]:
                        l += 1
                    while r-1 > 0 and A[r-1] == A[r]:
                        r -= 1

                    l, r = l+1, r-1
                elif A[l] + A[r] < S:
                    l += 1
                else:
                    r -= 1

            return res

        res = []
        i = 0
        while i < len(nums):
            n = nums[i]
            for two in twoSum(0-n, nums[i+1:]):
                res.append([n]+two)
            while i+1 < len(nums)-1 and nums[i+1] == nums[i]:
                i += 1
            i += 1

        return res


if __name__ == "__main__":
    s = Solution()
    print(s.threeSum([-1, 0, 1, 2, -1, -4]))
    print(s.threeSum([0, 0, 0]))
