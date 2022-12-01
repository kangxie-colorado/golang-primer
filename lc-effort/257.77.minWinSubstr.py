"""
https://leetcode.com/problems/minimum-window-substring/

I admit this is beyond my capability
I just practice for fun and leave this to LP
"""


from collections import Counter


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if len(s) < len(t):
            return ""
        if len(s) == len(t):
            return s if Counter(s) == Counter(t) else ""

        def contain(pc, tc):
            for k, v in tc.items():
                if pc[k] < v:
                    return False
            return True

        tCounter = Counter(t)
        resLen = len(s)+1
        res = ""
        i, j = 0, len(t)
        while j <= len(s):
            pCounter = Counter(s[i:j])
            if contain(pCounter, tCounter):
                if j-i < resLen:
                    res = s[i:j]
                    resLen = j-i
                i += 1
            else:
                j += 1

        while i < len(s):
            pCounter = Counter(s[i:j])
            if contain(pCounter, tCounter):
                if j-i < resLen:
                    res = s[i:j]
                    resLen = j-i
                i += 1
            else:
                break

        return res


"""
265 / 266 test cases passed

TLE.. at last case..
really..


"""


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if len(s) < len(t):
            return ""
        if len(s) == len(t):
            return s if Counter(s) == Counter(t) else ""

        def contain(pc, tc):
            for k, v in tc.items():
                if pc[k] < v:
                    return False
            return True

        tCounter = Counter(t)
        resLen = len(s)+1
        res = ""
        i, j = 0, len(t)-1
        pCounter = Counter(s[i:j])
        while j < len(s):
            pCounter[s[j]] += 1
            while contain(pCounter, tCounter):
                if j-i+1 < resLen:
                    res = s[i:j+1]
                    resLen = j-i+1
                pCounter[s[i]] -= 1
                i += 1
            j += 1

        while i < len(s) and contain(pCounter, tCounter):
            if j-i < resLen:
                res = s[i:j]
                resLen = j-i
            pCounter[s[i]] -= 1
            i += 1

        return res


"""
Runtime: 1710 ms, faster than 5.66% of Python3 online submissions for Minimum Window Substring.
Memory Usage: 14.7 MB, less than 36.37% of Python3 online submissions for Minimum Window Substring.
...

oh?

next optimization thoughts
            while contain(pCounter, tCounter):
                if j-i+1 < resLen:
                    res = s[i:j+1]
                    resLen = j-i+1
                pCounter[s[i]] -= 1
                i += 1
        I could go thru the p string and remove until pCounter doesn't contain tCounter anymore
        what can be safely removed..
        1. this char doesn't exists in tCounter
        2. this char has at least tCounter[c]+1 in p so far
"""


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if len(s) < len(t):
            return ""
        if len(s) == len(t):
            return s if Counter(s) == Counter(t) else ""

        def contain(pc, tc):
            for k, v in tc.items():
                if pc[k] < v:
                    return False
            return True

        def skipTo(pc, tc, p):
            for i, c in enumerate(p):
                if pc[c] < tc[c]+1:
                    return i, pc
                pc[c] -= 1

        tCounter = Counter(t)
        resLen = len(s)+1
        res = ""
        i, j = 0, len(t)-1
        pCounter = Counter(s[i:j])
        while j < len(s):
            pCounter[s[j]] += 1
            if contain(pCounter, tCounter):
                skip, pCounter = skipTo(pCounter, tCounter, s[i:j+1])
                if j-i-skip+1 < resLen:
                    res = s[i+skip:j+1]
                    resLen = j-i-skip+1
                pCounter[s[i+skip]] -= 1
                i += skip+1
            j += 1

        while i < len(s) and contain(pCounter, tCounter):
            if j-i < resLen:
                res = s[i:j]
                resLen = j-i
            pCounter[s[i]] -= 1
            i += 1

        return res


"""
failed here

"bbaac"
"aba"

okay, skip function is not right

        def skipTo(pc, tc, p):
            for i, c in enumerate(p):
                if c in tc or pc[c] < tc[c]+1: # no need the first one
                    return i, pc
                pc[c] -= 1



Runtime: 602 ms, faster than 18.05% of Python3 online submissions for Minimum Window Substring.
Memory Usage: 14.8 MB, less than 11.35% of Python3 online submissions for Minimum Window Substring.
"""


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if len(s) < len(t):
            return ""
        if len(s) == len(t):
            return s if Counter(s) == Counter(t) else ""

        def contain(pc, tc):
            for k, v in tc.items():
                if pc[k] < v:
                    return False
            return True

        def skipTo(pc, tc, p):
            for i, c in enumerate(p):
                if c not in pc:
                    continue
                if pc[c] < tc[c]+1:
                    return i, pc
                pc[c] -= 1

        tCounter = Counter(t)
        resLen = len(s)+1
        res = ""
        i, j = 0, 0
        pCounter = Counter()
        while j < len(s):
            if s[j] not in tCounter:
                j += 1
                continue
            pCounter[s[j]] += 1
            if contain(pCounter, tCounter):
                skip, pCounter = skipTo(pCounter, tCounter, s[i:j+1])
                if j-i-skip+1 < resLen:
                    res = s[i+skip:j+1]
                    resLen = j-i-skip+1
                pCounter[s[i+skip]] -= 1
                i += skip+1
            j += 1

        while i < len(s) and contain(pCounter, tCounter):
            if j-i < resLen:
                res = s[i:j]
                resLen = j-i
            pCounter[s[i]] -= 1
            i += 1

        return res


"""
ah.. right
instead of go thru two counters...

why don't I focus on the one counter..
and use an auxilary variable 

"""


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        tCounter = Counter(t)
        required = len(t)
        resLen = len(s)+1
        res = ""
        i, j = 0, 0
        while j < len(s):
            # only concernng the overlapping chars
            # non-relevant chars we don't care
            if s[j] in tCounter:
                # if there is not as many char in t from the sub-s, and the char in question is in t
                # required count can decrease by one
                # if already the needed count appears, then what appears is more than needed
                # so required will not change again, e.g. AAA vs 1 needed A...
                # but the counter should keep tracking of how many A has appeared so it won't release the lock
                # until all 3 As go out of scope...
                if tCounter[s[j]] > 0:
                    required -= 1
                tCounter[s[j]] -= 1

            while required == 0:
                if j-i+1 < resLen:
                    res = s[i:j+1]
                    resLen = j-i+1
                if s[i] in tCounter:
                    # only when the required char's count back to zero, which is a balance
                    # so upon leaving of this char in question, the required will need +=1
                    if tCounter[s[i]] == 0:
                        required += 1
                    tCounter[s[i]] += 1
                i += 1

            j += 1

        return res


"""
Do need to ponder about the relationship of required and tCounter

tCounter keeps track of the dynamics... from the initialized values 
and go down, and on more-than-needed supply, it can go negative

then with window contracting.. it replenished..
note.. when the counter[c]>0, we can decrease requried; when counter[c] == 0, we cannot decrease it anymore..
e.g. AAAB(vs AB) to decrease the tCounter[A], first A does the job... 2nd/3rd only enhances it.. by decreasing tCounter[A] to -2
but it cannot decrease the required count for B

coming back, only when it is time to cross the line where tCounter[c] == 0, which means balance
the required will need to increase
i.e. t=AA, when one A leaves window, the required will still be -1 so A is still in the window.. cannot increase required yet

"""

if __name__ == '__main__':
    print(Solution().minWindow("bbaac", "aba"))
    print(Solution().minWindow("ADOBECODEBANC", "ABC"))
