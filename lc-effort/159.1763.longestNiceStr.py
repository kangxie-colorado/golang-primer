"""
https://leetcode.com/problems/longest-nice-substring/

this is an easy?
I am struck and have no idea actually

what is the sliding policy?
1. use the out loop to slide.. still keep two maps?
    upper and lower cases.. if both appears
2. actually this has something similar to that k repeating
    if there is a char, that cannot meet the criterea, it needs to divide and conquer?
    what is the other solution to that problem? yeah, sliding window.. keep unique and the repeating frequency
    and I missed the very most important one.. it looped 26 times over there are at most 26 possible unique characters

"""


from collections import defaultdict


class Solution:
    def longestNiceSubstring(self, s: str) -> str:
        maxlen = 0
        res = ""
        for N in range(1, 27):
            uniques = 0
            bothCases = 0
            m = defaultdict(int)
            i, j = 0, 0

            while j < len(s):

                m[s[j]] += 1
                if m[s[j]] == 1:  # 0 -> 1
                    uniques += 1
                    if s[j].isupper() and m[s[j].lower()] > 0 or \
                            s[j].islower() and m[s[j].upper()] > 0:
                        bothCases += 1
                j += 1

                while uniques > 2*N:
                    m[s[i]] -= 1
                    if m[s[i]] == 0:
                        uniques -= 1
                        if s[i].isupper() and m[s[i].lower()] > 0 or \
                                s[i].islower() and m[s[i].upper()] > 0:
                            bothCases -= 1
                    i += 1

                if uniques == 2*N and bothCases == N:
                    if j-i > maxlen:
                        maxlen = j-i
                        res = s[i:j]

        return res


"""
"HkhBubUYy"
Output "BubU"
Expected "BubUYy"

tail effect...?
then I adjusted the sliding policy..
- alway process j first
- then shrink i until the window is eligible again.. 

^^ this might be a way to deal with my tail effect?
always process the tail first.. because I am looking into the tail..

then it failed

"dDzeE"
Output "eE"
Expected "dD"


    if j-i+1 > maxlen:
          ^^ shit.. what is this +1 doing here.. 

then it passes..
Runtime: 90 ms, faster than 57.82% of Python3 online submissions for Longest Nice Substring.
Memory Usage: 13.9 MB, less than 36.21% of Python3 online submissions for Longest Nice Substring.
"""

if __name__ == '__main__':
    s = Solution()
    print(s.longestNiceSubstring("dDzeE"))
