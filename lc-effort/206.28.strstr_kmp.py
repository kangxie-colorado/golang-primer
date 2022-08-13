"""
https://leetcode.com/problems/implement-strstr/

kmp.. hope i can still recover it from memory
"""


from operator import ne


def buildLPS(needle):
    m = len(needle)
    lps = [0]*m
    prevLPS, i = 0, 1
    while i < m:
        # print(prevLPS, lps[prevLPS])
        if needle[i] == needle[prevLPS]:
            # prevLPS is before i, what is longest prefix/suffix length
            # the whole array is to store lps for every position
            # so when needle[i] == needle[prevLPS], that means lps can extend by 1 from prevLps
            lps[i] = prevLPS + 1
            # also prevLPS is updated
            prevLPS += 1
            i += 1
        else:
            if prevLPS == 0:
                lps[i] = 0  # default is 0 but explicitly spells it
                i += 1
            else:
                # if cannot extend from prevLPS
                # try backing up
                prevLPS = lps[prevLPS-1]

    return lps


class Solution(object):
    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        n, m = len(haystack), len(needle)
        if m == 0:
            return 0
        if m > n:
            return -1

        lps = buildLPS(needle)
        print(lps)
        i = j = 0
        while i < n:
            print(i, haystack[i], j, needle[j])
            if haystack[i] == needle[j]:
                i, j = i+1, j+1
            else:
                if j == 0:
                    # j is 0.. meaning first char of needle doesn't mather
                    # naturally this should move i forward
                    i += 1
                else:
                    # backup to last lps
                    # come to here after a mismatch with j!=0
                    # instead of backup i, it backs up j..
                    # notice it cannot move i forward either here..
                    # e.g ABCABDABCABC vs ABCABC
                    # when mismtach, j is 5, i is 5..
                    # anyway.. because AB (in ABD) still match first AB
                    # so backoff j to 2 to use that information
                    # not backing j to 4... !!!
                    """
                    s-macbook-pro:lc-effort xiekang$ python3 206.28.strstr_kmp.py 
                    [0, 0, 0, 1, 2, 3]
                    0 A 0 A
                    1 B 1 B
                    2 C 2 C
                    3 A 3 A
                    4 B 4 B
                    5 D 5 C
                    5 D 2 C
                    5 D 0 A
                    6 A 0 A
                    7 B 1 B
                    8 C 2 C
                    9 A 3 A
                    10 B 4 B
                    11 C 5 C
                    6
                    """
                    j = lps[j-1]
            if j == m:
                return i-m

        return -1


def testLPS():

    print(buildLPS('AAAA'))
    print(buildLPS('AAAC'))
    print(buildLPS('AAACAAAC'))
    print(buildLPS('AAACAAACAAAC'))
    print(buildLPS('AAACAAACBAAAC'))

    """
    this example: 
    interesting.. prevLPS never goes over 2

    0 0
    1 1
    0 0
    0 0
    0 0
    0 0
    1 1
    2 0 <== when prevLPS is 2.. the lps[2] (lps[prevLPS]) is 0.. meaning, at pos2, there is no prefix/suffix 
            meeting the criterea AAB...
            unless the incoming is B.. otherwise, it cannot exend to 3... so it backs up
            prevLPS-1 => 1, lps[1] is also 1... 
            interesting...
    1 1
    0 0
    0 0
    0 0
    0 0
    0 0
    1 1
    2 0
    1 1
    0 0
    0 0
    0 0
    0 0
    1 1
    [0, 1, 0, 0, 0, 1, 2, 0, 0, 0, 0, 1, 2, 0, 0, 0, 1, 2]
    """
    print(buildLPS('AABCDAAEDFGAAHIJAA'))


"""
try writing this in another form
"""


class Solution(object):
    def strStr(self, haystack, needle):
        if len(needle) == 0:
            return 0
        if len(needle) > len(haystack):
            return -1

        lps = buildLPS(needle)
        i = j = 0
        while i+len(needle) <= len(haystack):
            while j < len(needle) and haystack[i+j] == needle[j]:
                j += 1

            if j == len(needle):
                return i

            if j == 0:
                i += 1
            else:
                i += max(1, j-lps[j-1])
                j = lps[j-1]

        return -1


"""
Runtime: 17 ms, faster than 88.48% of Python online submissions for Implement strStr().
Memory Usage: 13.5 MB, less than 30.29% of Python online submissions for Implement strStr().
"""

if __name__ == '__main__':
    s = Solution()
    print(s.strStr('ABCABDABCABC', 'ABCABC'))
