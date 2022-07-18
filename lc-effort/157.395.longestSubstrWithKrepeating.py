"""
https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/

I don't have any idea but I kind of see
if I can count the numbers of each char.. then replace those with smaller than k to " "
then aabcccdee k=2 will become "aa ccc ee".. do a split..

each part is actually an eligible substr...
nah.. that is not necessarily right

aabcbaa k=
although b is 2 but after split
aab baa it doesn't meet the eligibility

so still a recursive fuck... that is ok.. let me just

"""


from collections import defaultdict
from enum import unique


class Solution(object):
    def longestSubstring(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: int
        """
        def helper(s):
            if len(s) < k:
                return 0

            m = defaultdict(int)
            for c in s:
                m[c] += 1

            unqaulified = set()
            for c, v in m.items():
                if v < k:
                    unqaulified.add(c)

            if len(unqaulified) == 0:
                return len(s)

            aslist = list(s)
            lastIdx = 0
            maxLen = 0
            for i in range(len(s)):
                if s[i] in unqaulified:
                    maxLen = max(maxLen, helper("".join(aslist[lastIdx:i])))
                    lastIdx = i+1

            # qualify thru the end.. need to calculate it separately
            maxLen = max(maxLen, helper("".join(aslist[lastIdx:])))
            return maxLen

        return helper(s)


"""
"bbaaacbd"
3

failed here.. missed a scenario, the ending part is all qualified.. need to include that calculation

and after fixing that
Runtime: 50 ms, faster than 82.91% of Python3 online submissions for Longest Substring with At Least K Repeating Characters.
Memory Usage: 14 MB, less than 57.56% of Python3 online submissions for Longest Substring with At Least K Repeating Characters.

... shit... what problem is this.. I am brute force but still good

so the better solution is sliding window
I scratched a little surface on that. I was thinking
maybe try a, try b... but I never came close to think try 1 unique numbers.. 2 uniqe numbers...
there are at most 26 unique numbers..

so the sliding policy is going to be
given N unique char are wanted..
when there isn't enough uniqs... grow
there there is more than enough.. shrink

and each step take care of the states
what states?
1. how many unique chars being tracked (from 0->1)
2. how many chars appeared k+ times.. (from k-1 -> k)

and when the state is eligible update the max..
so it will be O(26*N)

pay attention the sliding policy is around how many unique
not around how many kOrMoreTime...

that is resovling around not only the char I am after.. so it won't be easy to use that centering
"""


class Solution(object):
    def longestSubstring(self, s, k):
        maxLen = 0
        for N in range(1, 27):
            unique = 0
            kOrMoreTime = 0
            m = defaultdict(int)
            i, j = 0, 0
            while j < len(s):
                while unique > N:
                    m[s[i]] -= 1
                    if m[s[i]] == 0:
                        unique -= 1
                    if m[s[i]] == k-1:
                        kOrMoreTime -= 1
                    i += 1

                m[s[j]] += 1
                if m[s[j]] == 1:
                    unique += 1
                if m[s[j]] == k:
                    kOrMoreTime += 1

                # just uniq N chars; and all of them are kOrMoreTime
                # it is eligible, update the max
                if unique == N and kOrMoreTime == N:
                    maxLen = max(maxLen, j-i+1)

                j += 1
        return maxLen


"""
Runtime: 203 ms, faster than 34.24% of Python3 online submissions for Longest Substring with At Least K Repeating Characters.
Memory Usage: 13.9 MB, less than 89.15% of Python3 online submissions for Longest Substring with At Least K Repeating Characters.

let me see that non-shrinking template also applies here
"""


class Solution(object):
    def longestSubstring(self, s, k):
        maxLen = 0
        for N in range(1, 27):
            unique = 0
            kOrMoreTime = 0
            m = defaultdict(int)
            i, j = 0, 0
            while j < len(s):
                if unique > N:
                    m[s[i]] -= 1
                    if m[s[i]] == 0:
                        unique -= 1
                    if m[s[i]] == k-1:
                        kOrMoreTime -= 1
                    i += 1

                m[s[j]] += 1
                if m[s[j]] == 1:
                    unique += 1
                if m[s[j]] == k:
                    kOrMoreTime += 1
                if unique == N and kOrMoreTime == N:
                    maxLen = max(maxLen, j-i+1)
                j += 1
            while i < len(s):
                if unique == N and kOrMoreTime == N:
                    maxLen = max(maxLen, j-i)
                m[s[i]] -= 1
                if m[s[i]] == 0:
                    unique -= 1
                if m[s[i]] == k-1:
                    kOrMoreTime -= 1
                i += 1

        return maxLen


"""
notice I just change the while to if.. and adjust the place to update max(which isn't really necessary)
make it back

and I just missed this test case after passing 33
"aaaaaaaaabbbcccccddddd"
5
output - 9; expected: 10
for which I think is because I didn't take care of the tail instance.. so actually I should not update maxLen
each time.. just do it at the end of the loop...

but how can I gurantee after loop it is stil meeting uniqe=N and kOrMoreTime == N:


Runtime: 466 ms, faster than 17.76% of Python3 online submissions for Longest Substring with At Least K Repeating Characters.
Memory Usage: 14 MB, less than 57.56% of Python3 online submissions for Longest Substring with At Least K Repeating Characters.

so the trick is after J is done, cont to walk i to j

            while i < len(s):
                if unique == N and kOrMoreTime == N:
                    maxLen = max(maxLen, j-i)
                m[s[i]] -= 1
                if m[s[i]] == 0:
                    unique -= 1
                if m[s[i]] == k-1:
                    kOrMoreTime -= 1
                i += 1

with this patch.. this non-shrinking window which actually shirnk at the end would also work..
but how ugly.. I don't know how it should be

but the shrinking is kind of more obvious as always..

"""


if __name__ == '__main__':
    s = Solution()
    print(s.longestSubstring("aaabbb", k=3))
    print(s.longestSubstring("aaaaaaaaabbbcccccddddd", k=5))
    print(s.longestSubstring("bbaaacbd", k=3))
    print(s.longestSubstring("ababbc", k=2))
    print(s.longestSubstring("aabac", k=2))
