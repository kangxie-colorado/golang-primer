"""
https://leetcode.com/problems/count-and-say/

count and say

Input: n = 4
Output: "1211"
Explanation:
countAndSay(1) = "1"
countAndSay(2) = say "1" = one 1 = "11"
countAndSay(3) = say "11" = two 1's = "21"
countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"

so this is naturally recursive
countAndSay(4) = count 21 => 1 2 1 1; say.. basically say it.. 
"""


import re


class Solution(object):
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """
        def helper(n):
            if n == 1:
                return "1"

            lastSay = helper(n-1)
            say = ""
            i = 0
            while i < len(lastSay):
                count = 1
                while i+1 < len(lastSay) and lastSay[i] == lastSay[i+1]:
                    count += 1
                    i += 1
                say += str(count) + str(lastSay[i])
                i += 1

            return say

        return helper(n)


"""
Runtime: 47 ms, faster than 54.40% of Python online submissions for Count and Say.
Memory Usage: 13.5 MB, less than 66.71% of Python online submissions for Count and Say.
"""

if __name__ == '__main__':
    s = Solution()
    print(s.countAndSay(2))
    print(s.countAndSay(3))
    print(s.countAndSay(4))
