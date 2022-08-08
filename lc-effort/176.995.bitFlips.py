"""
https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/

probably I cannot make it happen
but give a try

just sliding the fixed window
when it is all 1: len=sum, don't flip
otherwise: flip

actually this feels like the zero can be pressed forward... like some array problem
so leave the judgement to the end


YEAH... attempted to write some logic.. but non makes sense
"""


from typing import List


class Solution:
    def minKBitFlips(self, nums: List[int], k: int) -> int:
        flips = 0

        i, j, S = 0, k, sum(nums[:k])
        target = 1
        while j <= len(nums):
            while nums[i] == target and j < len(nums):
                S += nums[j]-1
                i += 1
                j += 1

            # i is 0.. (or target ^ 0.. target will flip between )
            if j == len(nums)-1 and S > 0:
                return -1

            flips += 1
            target ^= 1
            S = k
            i += 1
            j += 1

        return flips


if __name__ == '__main__':
    s = Solution()

    a = [0, 0, 0, 1, 0, 1, 1, 0]
    print(s.minKBitFlips(a, 3))

    print(s.minKBitFlips([1, 1, 0], 1))
    print(s.minKBitFlips([0, 1, 0], 1))
    print(s.minKBitFlips([1, 1, 0], 2))
