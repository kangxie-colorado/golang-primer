"""
https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/

did a blank board evolution 
basically think of two contributions 

self, (0 or 1)
combination with preceeding numers (last+1, or last, or 0(>right))

this is an O(n) but still very low on efficiency 

"""


from typing import List


class Solution:
    def numSubarrayBoundedMax(self, nums: List[int], left: int, right: int) -> int:
        i = j = 0
        contributions = []  # contribute as self and combinations to preceeding numbers
        while j < len(nums):
            if nums[j] > right:
                contributions.append((0, 0))
                i = j = j+1
                continue

            if left <= nums[j] <= right:
                contributions.append((1, j-i))
            else:
                comb = 0
                if j > 0:
                    if right >= nums[j-1] >= left:
                        comb = contributions[j-1][1] + 1
                    else:
                        comb = contributions[j-1][1]
                contributions.append((0, comb))
            j += 1

        return sum([i[0]+i[1] for i in contributions])


if __name__ == "__main__":
    s = Solution()

    print(s.numSubarrayBoundedMax(
        [16, 69, 88, 85, 79, 87, 37, 33, 39, 34], 55, 57))
    print(s.numSubarrayBoundedMax([1, 1, 2, 1, 1], 2, 8))
    print(s.numSubarrayBoundedMax([2, 1, 2, 5, 6, 1, 1, 2, 9, 2, 1, 1],
                                  2,
                                  8))
