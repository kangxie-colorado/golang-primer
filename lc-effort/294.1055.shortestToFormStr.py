"""
https://leetcode.com/problems/shortest-way-to-form-string/

"""


from collections import Counter


class Solution:
    def shortestWay(self, source: str, target: str) -> int:
        sCounter = Counter(source)
        for c in target:
            if c not in sCounter:
                return -1

        sIdx = 0
        tIdx = 0
        res = 0
        s = source
        while tIdx < len(target):
            c = target[tIdx]

            while sIdx < len(s) and c != s[sIdx]:
                sIdx += 1

            if sIdx >= len(s):
                res += 1
                sIdx = 0
                s = source
            else:
                tIdx += 1
                sIdx += 1

        return res+1


if __name__ == '__main__':
    s = Solution()
    print(s.shortestWay("abc", "abcbc"))
