"""
# https://leetcode.com/problems/interleaving-string/
thinking of dp, think the sub problem
this is obvious
"""


class Solution_topdown:
    def isInterleave(self, s1: str, s2: str, s3: str) -> bool:
        m = [[]] * (len(s1) + 1)
        for i in range(len(m)):
            m[i] = [0] * (len(s2) + 1)

        def helper(i, j, k):
            if i == len(s1) and j == len(s2) and k == len(s3):
                return True
            if m[i][j] != 0:
                return m[i][j] == 1

            s1i = s1[i] if i < len(s1) else ""
            s2j = s2[j] if j < len(s2) else ""
            s3k = s3[k] if k < len(s3) else ""

            if s1i == s3k == s2j:
                m[i][j] = 1 if helper(
                    i + 1, j, k + 1) or helper(i, j + 1, k + 1) else -1
            elif s1i == s3k:
                m[i][j] = 1 if helper(i + 1, j, k + 1) else -1
            elif s2j == s3k:
                m[i][j] = 1 if helper(i, j + 1, k + 1) else -1
            else:
                m[i][j] = -1

            return m[i][j] == 1

        return len(s1) + len(s2) == len(s3) and helper(0, 0, 0)


''' # noqa
simple but yet I fail
"aabcc"
"dbbca"
"aadbbcbcac"

stupidly missing +1 on                 m[i][j] = 1 if helper(i, j+1, k + 1) else 0

yet I still tle
"bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa"
"babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab"
"babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab"

I see... because the m[i][j]=0 is actually a definitive state but I treated it like un-calculated
I should use a neatural state
'''

""" # noqa
Runtime: 58 ms, faster than 67.62% of Python3 online submissions for Interleaving String.
Memory Usage: 14.8 MB, less than 47.88% of Python3 online submissions for Interleaving String.

finally...
rusty on python.. need some picking up for sure

before I do another version, let me do some work..
there should be a DP solution.. but why can't I see it...
"""


class Solution_dp_arrary:
    def isInterleave(self, s1: str, s2: str, s3: str) -> bool:
        if len(s1) + len(s2) != len(s3):
            return False
        dp = [[]] * (len(s1) + 1)
        for i in range(len(dp)):
            dp[i] = [False] * (len(s2) + 1)

        for i in range(len(s1) + 1):
            for j in range(len(s2) + 1):
                if i == 0 and j == 0:
                    dp[i][j] = True
                else:
                    dp[i][j] = \
                        (i > 0 and s1[i - 1] == s3[i + j - 1] and dp[i - 1][j]) \
                        or \
                        (j > 0 and s2[j - 1] == s3[i + j - 1] and dp[i][j - 1])
                """ #noqa
                elif (i > 0 and s1[i - 1] == s3[i + j - 1]) or (j > 0 and s2[j - 1] == s3[i + j - 1]):
                    dp[i][j] = dp[i - 1][j] or dp[i][j - 1]
                    
                    see where is wrong here?
                    it opens up (j > 0 and s2[j - 1] == s3[i + j - 1]) -> dp[i - 1][j]
                    when one condition matches above.. the condition on the other side sneaks in...
                    
                    e.g. s1="ab" s2='bc' s3 'bbac'
                      " b c
                    " T T F
                    a F
                    b
                    
                    at dp[1][1], because s2[0] == s3[1], and dp[i - 1][j] or dp[i][j - 1] is True
                    then dp[1][1] is set to true... wrongly!!! this is because the condition check
                    I wrote is so clever and wrong
                    
                    this is you trying to be too smart...
                    
                    how can i stop doing this!!!
                """

        return dp[len(s1)][len(s2)]


""" # noqa
Runtime: 83 ms, faster than 26.04% of Python3 online submissions for Interleaving String.
Memory Usage: 13.9 MB, less than 98.22% of Python3 online submissions for Interleaving String.
"""

'''
I feel rusty in python now..
and the editor is not as good as vs code..

okay back to vs code..
let me change this to single row extra memory which is obvious now
'''


class Solution:
    def isInterleave(self, s1: str, s2: str, s3: str) -> bool:
        if len(s1) + len(s2) != len(s3):
            return False

        dp = [False] * (len(s2) + 1)

        for i in range(len(s1) + 1):
            for j in range(len(s2) + 1):
                if i == 0 and j == 0:
                    dp[j] = True
                else:
                    dp[j] = \
                        (i > 0 and s1[i - 1] == s3[i + j - 1] and dp[j]) or \
                        (j > 0 and s2[j - 1] == s3[i + j - 1] and dp[j - 1])

        return dp[len(s2)]


"""
Runtime: 69 ms, faster than 46.61% of Python3 online submissions for Interleaving String. # noqa
Memory Usage: 13.9 MB, less than 88.03% of Python3 online submissions for Interleaving String.
"""

if __name__ == '__main__':
    s = Solution()
    print(s.isInterleave("ab", "bc", "bbac"))
