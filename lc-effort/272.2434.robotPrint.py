"""
https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/


I can see a decision tree but what about the memorization?
but at input data of 10^5 ... possible to pass????


"""


from collections import Counter
from functools import lru_cache


class Solution:
    def robotWithString(self, s: str) -> str:
        res = "z"*len(s)
        m = {}

        def backtrack(start, robot, paper):
            nonlocal res
            if len(paper) == len(s):
                res = min(res, paper)
                return

            if start:
                backtrack(start[1:], robot+start[0], paper)

            if robot:
                backtrack(start, robot[:-1], paper+robot[-1])

        backtrack(s, "", "")
        return res


"""
as expected TLE and only 6 / 62 test cases passed.
faled even here

"mmuqezwmomeplrtskz"
so I am thinking for better or worse.. the min char must go to the result first
in this case it is e... 

there are two e here.. 
so 
mmuqe
mmuqezwmome

it cannot be better than this..
then I think.. is there a recursive relationship?

zwmomeplrtskz mmuqe ""
after I put a e.. 
zwmomeplrtskz mmuq "e"
I need another e to be here

plrtskz mmuqzwmome "e' 
plrtskz mmuqzwmom "ee'

at this time... I would need to search the min in S, which is also smaller than m in T, which is k..

so 
z mmuqzwmomplrtsk "ee'
z mmuqzwmomplrts "eek'

now I search min of S which is smaller than p.. and there is none..
so it has to be s
z mmuqzwmomplrts "eek'
z mmuqzwmomplrt "eeks'

so the choice is min-of-S vs the last-of-T.. 
"" is the biggest, and on equal case.. move from T to P



"""


class Solution:
    def robotWithString(self, s: str) -> str:
        t = p = ''
        while s or t:
            minS = min(s) if s else ""
            if not t or (minS and minS < t[-1]):
                idx = s.index(minS)
                t += s[:idx+1]
                s = s[idx+1:]

            else:
                p += t[-1]
                t = t[:-1]
        return p


"""
34 / 62 test cases passed.

TLE...
okay.. so I think... 

need further optimization 
"""


class Solution:
    def robotWithString(self, s: str) -> str:
        @lru_cache
        def lastMinIdx(s):
            if not s:
                return "", -1

            minS = 'z'
            lastIdx = -1
            for i, c in enumerate(s):
                if c <= minS:
                    minS = c
                    lastIdx = i

            return minS, lastIdx

        t = p = ''
        while s or t:
            minS, lastIdx = lastMinIdx(s)
            if not t or (minS and minS < t[-1]):
                move2T = s[:lastIdx+1].replace(minS, "")
                t += move2T
                s = s[lastIdx+1:]
                p += minS * (lastIdx+1 - len(move2T))
            else:
                p += t[-1]
                t = t[:-1]
        return p


"""
Runtime: 2339 ms, faster than 50.03% of Python3 online submissions for Using a Robot to Print the Lexicographically Smallest String.
Memory Usage: 16.4 MB, less than 95.52% of Python3 online submissions for Using a Robot to Print the Lexicographically Smallest String.


sometimes it passes.. sometimes it doesn't
so I checked the hints

- If there are some character “a” ’ s in the string, they can be written on paper before anything else.
- Every character in the string before the last “a” should be written in reversed order.
- After the robot writes every “a” on paper, the same holds for other characters “b”, ”c”, …etc.

interesting... 
"""


class Solution:
    def robotWithString(self, s: str) -> str:
        t = p = ""
        chars = sorted(list(set(s)))

        for c in chars:
            while t and t[-1] <= c:
                p += t[-1]
                t = t[:-1]
            i = 0
            rIdx = s.rfind(c)
            while i <= rIdx:
                if s[i] != c:
                    t += s[i]
                else:
                    p += c
                i += 1
            s = s[rIdx+1:]

        return p


class Solution:
    def robotWithString(self, s: str) -> str:
        t = p = ""
        chars = sorted(list(set(s)))

        for c in chars:
            i = 0
            while len(t)-1-i >= 0 and t[len(t)-1-i] <= c:
                p += t[len(t)-i-1]
                i += 1
            t = t[:len(t)-i]
            i = 0
            rIdx = s.rfind(c)
            while i <= rIdx:
                if s[i] != c:
                    t += s[i]
                else:
                    p += c
                i += 1
            s = s[rIdx+1:]

        return p


"""
Runtime: 1241 ms, faster than 79.93% of Python3 online submissions for Using a Robot to Print the Lexicographically Smallest String.
Memory Usage: 15.9 MB, less than 99.09% of Python3 online submissions for Using a Robot to Print the Lexicographically Smallest String.

Runtime: 1149 ms, faster than 82.59% of Python3 online submissions for Using a Robot to Print the Lexicographically Smallest String.
Memory Usage: 16 MB, less than 99.09% of Python3 online submissions for Using a Robot to Print the Lexicographically Smallest String.
"""


if __name__ == '__main__':
    s = Solution()
    # print(s.robotWithString('zza'))
    print(s.robotWithString('mmuqezwmomeplrtskz'))  # "eekstrlpmomwzqummz"
