"""
https://leetcode.com/problems/word-break/

I had forgotten how it was done
let me see
"""


from re import T
from typing import List


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:

        def helper(s):
            if not s:
                return True
            for word in wordDict:
                if len(s) >= len(word) and word == s[:len(word)] and helper(s[len(word):]):
                    return True
            return False

        return helper(s)


"""
I know it is going to TLE but kind of unable to figure out where

35 / 45 test cases passed.

"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]


"""


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        dp = [0] * (len(s)+1)
        dp[0] = 1
        for i in range(1, len(s)+1):
            for word in wordDict:
                if i >= len(word) and s[i-len(word):i] == word:
                    dp[i] |= dp[i-len(word)]
                    if dp[i]:
                        break

        return dp[len(s)]


"""
Runtime: 88 ms, faster than 14.18% of Python3 online submissions for Word Break.
Memory Usage: 13.9 MB, less than 70.39% of Python3 online submissions for Word Break.
Runtime: 75 ms, faster than 31.00% of Python3 online submissions for Word Break.
Memory Usage: 14.1 MB, less than 43.38% of Python3 online submissions for Word Break.
hey... cool..

that "aaaaaaa..." case used maybe 68ms... 
can it be optimized?

what optimization? I think I can replace/remove "aaa" because 'a' existence 
perhaps... but maybe I see another 
when dp[i] is true, I can just break
and that

Runtime: 40 ms, faster than 93.28% of Python3 online submissions for Word Break.
Memory Usage: 14.1 MB, less than 43.38% of Python3 online submissions for Word Break.
O(n) - O(n^2)
also if dp[i] is false... hmm... cannot really return here...
"""
