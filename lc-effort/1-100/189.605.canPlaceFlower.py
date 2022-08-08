"""
https://leetcode.com/problems/can-place-flowers/

not so straightforward
but I am thinking use 1s to divide this into windows and check how much can be planted in each window
"""


from typing import List


class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], t: int) -> bool:
        windS = 0
        able = 0
        for i, n in enumerate(flowerbed):
            if n == 1:
                # what is the left?
                windE = i-2
                if windE >= windS:
                    # some trick here to compute both even/odd window size
                    # (windSize-1)/2+1
                    able += (windE-windS)//2+1
                windS = i+2
        windE = len(flowerbed)-1
        if windE >= windS:
            able += (windE-windS)//2+1
        return able >= t


"""
Runtime: 210 ms, faster than 72.71% of Python3 online submissions for Can Place Flowers.
Memory Usage: 14.5 MB, less than 29.57% of Python3 online submissions for Can Place Flowers.
"""

if __name__ == '__main__':
    s = Solution()
    print(s.canPlaceFlowers([1, 0, 0, 0, 1], 2))
