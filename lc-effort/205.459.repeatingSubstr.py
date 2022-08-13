"""
https://leetcode.com/problems/repeated-substring-pattern/discuss/94382/From-intuitive-but-slow-to-really-fast-but-a-little-hard-to-comprehend.

KMP really interesting..
"""


class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:

        for k in range(len(s)//2, 0, -1):
            if len(s) % k != 0:
                continue
            S = set()
            result = True
            for m in range(1, len(s)//k+1):
                S.add(s[(m-1)*k:m*k])
                if len(S) == 2:
                    result = False
                    break

            if result:
                return True
        return False


print(Solution().repeatedSubstringPattern('abab'))
