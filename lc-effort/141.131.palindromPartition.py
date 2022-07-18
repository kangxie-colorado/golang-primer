"""
https://leetcode.com/problems/palindrome-partitioning/

not much idea
let me just start all signle letter array and start combining them..
"""


from gc import set_debug
from multiprocessing.sharedctypes import copy


class Solution(object):
    def partition(self, s):
        """
        :type s: str
        :rtype: List[List[str]]
        """
        seed = [a for a in s]
        seeds = [seed]
        res = [seed]
        while True:
            newseeds = []
            for seed in seeds:
                for i in range(len(seed)):
                    j = i+1
                    k = i
                    while k >= 0 and j < len(seed) and seed[k] == seed[j]:
                        # can form a new one..
                        p = seed[:k] + [''.join(seed[k:j+1]), ] + seed[j+1:]
                        res.append(p)
                        newseeds.append(p)
                        k, j = k-1, j+1

                    j = i+1
                    k = i-1
                    while k >= 0 and j < len(seed) and seed[k] == seed[j]:
                        # also can form a new one
                        p = seed[:k] + [''.join(seed[k:j+1]), ] + seed[j+1:]
                        res.append(p)
                        newseeds.append(p)
                        k, j = k-1, j+1

            if len(newseeds) == 0:
                break
            seeds = newseeds
        s = set()
        for p in res:
            s.add(tuple(p))

        res = []
        for p in s:
            res.append(list(p))
        return res


"""
TLE on "iiiiiiiiiiiii"
I know it is going to TLE...

I might be able to see the backtracking one
let me try
"""


class Solution(object):
    def partition(self, s):
        res = []

        def isPalindrom(s):
            l, r = 0, len(s)-1
            while l < r:
                if s[l] != s[r]:
                    return False
                l, r = l+1, r-1
            return True

        def backtracking(p, s):
            # p: List[str] -> to hold the partial results
            # s: the remaining str
            if len(s) == 0:

                res.append(p.copy())
                return

            for i in range(1, len(s)+1):
                if isPalindrom(s[:i]):
                    p.append(s[:i])
                    backtracking(p, s[i:])
                    p.pop()

        backtracking([], s)
        return res


"""

Runtime: 677 ms, faster than 94.86% of Python3 online submissions for Palindrome Partitioning.
Memory Usage: 30.4 MB, less than 42.96% of Python3 online submissions for Palindrome Partitioning.

wow!!!!
"""

if __name__ == '__main__':
    s = Solution()
    # print(s.partition("aab"))
    # print(s.partition("aaab"))
    print(s.partition("aaaab"))
