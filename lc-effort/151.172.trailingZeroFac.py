"""
https://leetcode.com/problems/factorial-trailing-zeroes/

10^4 is not a big deal but the number will overflow anything
so I need to think

so read other's discussions and understand it now
every 5 will making a zero.. actually every 5 and 2 together 
but for every 5 to appear, there will be at least two 2s

then for every 25 to appear, there will be at least 12 2s.. and 
so you will not run out of 2.. the 5 is the dominate factors..

then now notice 25 is not only one 5, but 2
125 is 3 5*5*5
625 is 4... so on and so forth

because the factorail number can be very large so you don't go up to meet the factorial 
but work on the factors which is from n down...

n/5 to get a number the dividend will carry on to /5 /5 /5.. note not %5
e.g. 777(/5 155 5s) -> 155(/5 31 25s) -> 31(/5 6 125s) -> 6(/5 1 625s) ->1(/5 0 3125s)
so total is 155+31+6+1..


so this forms a recursive structure 


"""


class Solution(object):
    def trailingZeroes(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n == 0:
            return 0

        return n/5 + self.trailingZeroes(n/5)


"""
Runtime: 11 ms, faster than 98.16% of Python online submissions for Factorial Trailing Zeroes.
Memory Usage: 13.5 MB, less than 35.42% of Python online submissions for Factorial Trailing Zeroes.

"""
