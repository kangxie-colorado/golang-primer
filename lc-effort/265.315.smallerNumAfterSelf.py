"""
https://leetcode.com/problems/count-of-smaller-numbers-after-self/

this is a hard problem
not much idea...

but brute force is very clear O(n^2)

then I am thinking, this has some DP property
solve it using sub-problem

not quite but at least I can start with right.. which is 0.. 
then going to left

then I think maybe I can maintain an insert list

[4,5,3,4,3,2,1]
the list grows like 
1                   -> so it will 0..
1,2                 -> 2 will inserted at index-1, so it will be 1
1,2,3               -> 3 will be inserted at index-2， so it will be 2
1,2,3,4             -> 4:3
1,2,3,4 enter 3     -> so I need to bisect_right... it will be 3
1,2,3,3,4 enter 5   -> 5:5
1,2,3,3,4,5 enter 4 -> 4:4

so roughly it kind of check out
"""


from typing import List
from sortedcontainers import SortedList


class Solution:
    def countSmaller(self, nums: List[int]) -> List[int]:
        sl = SortedList()

        i = len(nums)-1
        while i >= 0:
            idx = sl.bisect_left(nums[i])
            sl.add(nums[i])
            nums[i] = idx
            i -= 1

        return nums
