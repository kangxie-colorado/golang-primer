"""
https://leetcode.com/problems/minimum-size-subarray-sum/

Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).

O(N) is actually very visible..
but O(nlogn)? haha

I think I figured it out
so do a presum

nums:   [2,3,1,2,4, 3]  k=7
presum: [2,5,6,8,12,15]

then going thru the nums...
2.. delta = 0, target = 7, search for >=7 first idx.. I got 8(idx-3), so the len is 3-0+1=4
3.. delta = 2, target = 9, search and find 12(idx-4), 4-1+1=4
1..
2..
4.. delta = 8, target = 15, search and find 15(idx-5),5-4+1=2
...

binary search will cost O(lgN)
so it will be O(n*lgn)

"""


from bisect import bisect_left
from typing import List


class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        presum = 0
        presums = nums
        for i in range(len(nums)):
            presums[i] += presum
            presum = presums[i]

        res = len(nums)+1
        for i in range(len(nums)):
            if i == 0:
                k = target
            else:
                k = target+presums[i-1]

            j = bisect_left(presums, k)
            if j == len(nums):
                break
            res = min(res, j-i+1)

        if res == len(nums)+1:
            return 0
        return res


if __name__ == '__main__':
    s = Solution()
    print(s.minSubArrayLen(7, [2, 3, 1, 2, 4, 3]))
    print(s.minSubArrayLen(7, [1, 1, 1]))
