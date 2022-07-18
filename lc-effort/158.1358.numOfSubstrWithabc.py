"""
https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/

seems simple but I cannot get it right quickly
frustration

"""


from collections import defaultdict


class Solution1:
    def numberOfSubstrings(self, s: str) -> int:
        i, j = 0, 0
        m = defaultdict(int)
        unique = 0
        count = 0
        while j < len(s):
            while unique == 3:
                # when abc meets the critirea, any later chars don't change it
                # so it will be pure adding
                count += len(s)-j+1
                m[s[i]] -= 1
                if m[s[i]] == 0:
                    unique -= 1
                i += 1

            m[s[j]] += 1
            if m[s[j]] == 1:
                unique += 1

            j += 1
        # the tail occurance and make i also contracting
        while i < len(s):
            if unique == 3:
                count += 1

            m[s[i]] -= 1
            if m[s[i]] == 0:
                unique -= 1
            i += 1
        return count


"""
Runtime: 456 ms, faster than 30.12% of Python3 online submissions for Number of Substrings Containing All Three Characters.
Memory Usage: 14.2 MB, less than 98.31% of Python3 online submissions for Number of Substrings Containing All Three Characters.

yeah... this tail loop reminds me the sliding policy is not the optimal at all
let me rethink
"""


class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        i, j = 0, 0
        m = [0]*3
        count = 0
        while j < len(s):
            if min(m) == 0:
                m[ord(s[j]) - ord('a')] += 1
                j += 1
            else:
                m[ord(s[i]) - ord('a')] -= 1
                i += 1

            if min(m) > 0:
                count += len(s)-j+1

        while i < len(s):
            m[ord(s[i]) - ord('a')] -= 1
            i += 1
            if min(m) > 0:
                count += len(s)-j+1

        return count


"""
aaacb failed..
still the tail effect I didn't take care of
let me see

hmm.. how can I take care of that tail effect without 2nd loop?

"""


class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        i, j = 0, 0
        m = [0]*3
        count = 0
        while j < len(s):
            while j < len(s) and min(m) == 0:
                m[ord(s[j]) - ord('a')] += 1
                j += 1

            while i < len(s) and min(m) > 0:
                count += len(s)-j+1
                m[ord(s[i]) - ord('a')] -= 1
                i += 1
        return count


"""
okay.. this is more concise but
Runtime: 882 ms, faster than 7.96% of Python3 online submissions for Number of Substrings Containing All Three Characters.
Memory Usage: 14.5 MB, less than 7.95% of Python3 online submissions for Number of Substrings Containing All Three Characters.

what is the performance
"""


class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        i, j = 0, 0
        m = [0]*3
        count = 0
        while j < len(s):
            m[ord(s[j]) - ord('a')] += 1

            while min(m) > 0:
                count += len(s)-j
                m[ord(s[i]) - ord('a')] -= 1
                i += 1
            j += 1
        return count


"""
okay.. a little more concise
Runtime: 446 ms, faster than 31.98% of Python3 online submissions for Number of Substrings Containing All Three Characters.
Memory Usage: 14.4 MB, less than 34.18% of Python3 online submissions for Number of Substrings Containing All Three Characters.
"""

"""
good read here

https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/discuss/516977/JavaC%2B%2BPython-Easy-and-Concise

this solution is kind like mine
    def numberOfSubstrings(self, s):
        res = i = 0
        count = {c: 0 for c in 'abc'}
        for j in xrange(len(s)):
            count[s[j]] += 1
            while all(count.values()):
                count[s[i]] -= 1
                i += 1
            res += i
        return res
but it focus on front i+=1 while I focus on the end: count += len(s)-j
no diff actually

this solution is very genius as always by this lee guy
    int numberOfSubstrings(string s) {
        int last[3] = {-1, -1, -1}, res = 0, n = s.length();
        for (int i = 0; i < n; ++i) {
            last[s[i] - 'a'] = i;
            res += 1 + min({last[0], last[1], last[2]});
        }
        return res;
    }

so it also focus on the front... 

now I think why I have to deal with the tail effect is because I focus on the tail
if I focus on the front... it might be more concise for first solutions --- let me try again
"""


class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        i, j = 0, 0
        m = [0]*3
        count = 0
        while j < len(s):
            if min(m) == 0:
                m[ord(s[j]) - ord('a')] += 1
                j += 1
            else:
                m[ord(s[i]) - ord('a')] -= 1
                i += 1

            if min(m) > 0:
                count += i+1
        # still needing a tail here... sigh!
        return count


"""
hmm... still the tail effect
but this actually reflects on a bit on the lee's solution, just to keep last i.. which would be at idx-2

"""


class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        i, j = 0, 0
        m = [0]*3
        count = 0
        while j < len(s):
            m[ord(s[j]) - ord('a')] += 1

            while m[ord(s[i]) - ord('a')] > 1:
                m[ord(s[i]) - ord('a')] -= 1
                i += 1
            if min(m) > 0:
                count += i+1
            j += 1

        return count


"""
Runtime: 668 ms, faster than 12.52% of Python3 online submissions for Number of Substrings Containing All Three Characters.
Memory Usage: 14.4 MB, less than 34.18% of Python3 online submissions for Number of Substrings Containing All Three Characters.

okay.. this is focusing on the front.. I got it right
the key is kind of just tracking the lastIdx after all three has appeared.
            while m[ord(s[i]) - ord('a')] > 1:
                    ^ loop until it becomes 1 or 0.. if it is already 1/0, it will not even loop
                        this is actually more obscure than lee's soluion
                        alright.. but still good to see it turns into code

alirght let me finally code that lee's solution and go for a ride
"""


class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        i, j = 0, 0
        # why -1, two folds
        # 1. 0 is an legal index;
        # 2. and when last idx is 0, that means looking to the front
        #   there is actually one
        # so -1, plus 1 get the correct number

        last = [-1]*3
        count = 0
        while j < len(s):
            last[ord(s[j]) - ord('a')] = j
            count += 1 + min(last)
            j += 1

        return count


"""
Runtime: 289 ms, faster than 70.39% of Python3 online submissions for Number of Substrings Containing All Three Characters.
Memory Usage: 14.4 MB, less than 34.18% of Python3 online submissions for Number of Substrings Containing All Three Characters.


yeah.. somehow this is the best
"""

if __name__ == '__main__':
    s = Solution()
    print(s.numberOfSubstrings("aaacb"))
    print(s.numberOfSubstrings("abcaa"))
