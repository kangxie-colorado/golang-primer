"""
https://leetcode.com/problems/maximize-distance-to-closest-person/

did this before
I am seeing left-right and the right-left two pass solution
"""


from typing import List


class Solution:
    def maxDistToClosest(self, seats: List[int]) -> int:
        n = len(seats)
        leftToRight = [n]*n
        rightToLeft = [n]*n

        for i, s in enumerate(seats):
            if s == 1:
                leftToRight[i] = 0
            else:
                if i > 0:
                    leftToRight[i] = leftToRight[i-1]+1

        i = n-1
        while i >= 0:
            s = seats[i]
            if s == 1:
                rightToLeft[i] = 0
            else:
                if i < n-1:
                    rightToLeft[i] = rightToLeft[i+1]+1
            i -= 1

        return max([min(a, b) for a, b in zip(leftToRight, rightToLeft)])


if __name__ == "__main__":
    s = Solution()
    seats = [1, 0, 0, 0, 1, 0, 1]
    print(s.maxDistToClosest(seats))
