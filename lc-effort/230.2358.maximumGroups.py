"""
https://leetcode.com/problems/maximum-number-of-groups-entering-a-competition/

very interesting
I can see the binary search attribute
at max-m groups that can be formed, any n<max-m can be formed too

but how to decide given a m, that m groups can or cannot be formed
1. number requirement
the most compact way to form m groups, would be 1, 1+1, ... 1+m-1
so m(m+1)*2 <= len(nums)
because I can just put whatever left into last group... 

3. then the sum requirement... 
that is kind of hard?
maybe after a sort.. that is naturally taken care of.. so it is most the number dictating the results?

edge case being two same numbers e..g [8,8]

because 3 same number [8,8,8] can be done as [8] [8 8]
4 same number can be done as [8] [8 8 8]
5 same number can be done as [8] [8 8 8 8] or [8 8] [8 8 8]
6 can be done as [8] [8 8] [8 8 8]

"""


from typing import List


class Solution:
    def maximumGroups(self, grades: List[int]) -> int:

        def canForm(m):
            return m*(m+1) <= len(grades)*2

        l, r = 1, len(grades)
        while l < r:
            m = r-(r-l)//2
            if canForm(m):
                l = m
            else:
                r = m-1

        return l


"""
Runtime: 569 ms, faster than 74.15% of Python3 online submissions for Maximum Number of Groups Entering a Competition.
Memory Usage: 27.3 MB, less than 79.87% of Python3 online submissions for Maximum Number of Groups Entering a Competition.

haha... just like that?
"""
