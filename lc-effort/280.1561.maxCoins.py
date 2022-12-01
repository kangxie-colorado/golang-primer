"""
https://leetcode.com/problems/maximum-number-of-coins-you-can-get/

so the goal is to get 2nd max and make sure bob gets the less... 
no matter what, alice will get the biggest ones... 

so a sort.. will work

[1,2,3,4,5,6,7,8,9]
9->alice 1->bob, 8->self
7->alice 2->bob, 6->self
5->alice 3->bob, 4->self

"""


from typing import List


class Solution:
    def maxCoins(self, piles: List[int]) -> int:
        piles.sort(reverse=True)

        return sum(piles[1:len(piles)-len(piles)//3:2])
