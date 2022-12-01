"""
https://leetcode.com/problems/longest-increasing-subsequence/

did this in matrix 
now I think let me do this with dfs 
"""


from tkinter import N
from typing import List


class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        lis = [0]*len(nums)

        def dfs(idx):
            if idx >= len(nums):
                return 0

            if lis[idx] != 0:
                return lis[idx]

            lis[idx] = 1
            minRight = 100001
            for i in range(idx+1, len(nums)):
                if nums[i] >= minRight:
                    continue
                if nums[i] > nums[idx]:
                    minRight = min(minRight, nums[i])
                    lis[idx] = max(lis[idx], dfs(i)+1)
            return lis[idx]

        for i in range(len(nums)):
            if lis[i] == 0:
                dfs(i)
        return max(lis)


""""
TLE...

yeah.. many duplicates comparson even memorization is used

[1,2,3] from 1, 2>1, then continue to check 3>1... which is a waste because 3>2
so I reviewed my go solution, I used a checked min... which is checked min that is bigger than me


Runtime: 6027 ms, faster than 25.37% of Python3 online submissions for Longest Increasing Subsequence.
Memory Usage: 17.5 MB, less than 5.19% of Python3 online submissions for Longest Increasing Subsequence.
barely passed

this is conceptually like swapping the cards
but it is only changing the 2nd card..
there can be the optimization of swapping 3rd/4th card... 
hard to write code for that --- 

then I found this article interesting
giving it a read 
https://leetcode.com/problems/longest-increasing-subsequence/discuss/1326552/Optimization-From-Brute-Force-to-Dynamic-Programming-Explained!
"""


if __name__ == "__main__":
    s = Solution()
    nums = [10, 9, 2, 5, 3, 7, 101, 18]
    print(s.lengthOfLIS(nums=nums))
