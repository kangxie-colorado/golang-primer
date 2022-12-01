"""
https://leetcode.com/problems/range-sum-query-mutable/


I don't how to make the most efficient data structure
but I think
use a partial sum to deliver the range sum quickly.. that is of course the way
but the trouble is with the update.. it change all the partital sums coming after this item

so how to balance these two operatios...
a bit difficult...


"""


from typing import List
from sortedcontainers import SortedList


class NumArray:

    def __init__(self, nums: List[int]):
        self.nums = nums
        self.deltaList = SortedList()
        self.rangeSums = [0] + nums
        prefix = 0
        for i, rs in enumerate(self.rangeSums):
            self.rangeSums[i] += prefix
            prefix += rs

    def update(self, index: int, val: int) -> None:
        insertAt = self.deltaList.bisect_left((index, -100))
        if insertAt < len(self.deltaList) and self.deltaList[insertAt][0] == index:
            del self.deltaList[insertAt]

        if val-self.nums[index] != 0:
            self.deltaList.add((index, val-self.nums[index]))

    def sumRange(self, left: int, right: int) -> int:
        if right+1 >= len(self.rangeSums):
            return 0

        deltaLeft = self.deltaList.bisect_left((left-0.1, 0))
        deltaRight = self.deltaList.bisect_right((right+0.1, 0))
        deltaSum = sum([i[1] for i in self.deltaList[deltaLeft:deltaRight]])

        return self.rangeSums[right+1] - self.rangeSums[left] + deltaSum


"""
pass 9/16 and TLE
so sign.. I don't know an efficient way

just looking at the solutions
binary-indexed tree and segment tree

okay... I never heard about it...
"""

if __name__ == '__main__':
    s = NumArray([1, 2, 3, 4, 5])

    print(s.sumRange(0, 2))
    print(s.sumRange(0, 3))
    print(s.sumRange(3, 4))

    s.update(3, 2)
    print(s.sumRange(0, 2))
    print(s.sumRange(0, 3))
    print(s.sumRange(3, 4))

    s.update(3, 4)
    print(s.sumRange(0, 2))
    print(s.sumRange(0, 3))
    print(s.sumRange(3, 4))

    s.update(3, 5)
    s.update(4, 6)
    print(s.sumRange(0, 2))
    print(s.sumRange(0, 3))
    print(s.sumRange(3, 4))

    s.update(3, 5)
    s.update(4, 6)
    print(s.sumRange(0, 2))
    print(s.sumRange(0, 3))
    print(s.sumRange(3, 4))
