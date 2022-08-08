
"""
https://leetcode.com/problems/replace-the-substring-for-balanced-string/


hmmm... 
I come here from that genius atMost(k) - atMost(k-1) solution
but I don't see that applicable here easily  so I have some doubts...

let me see.. 
for this one, if I count the occrences of each char, then I can get the required replacement
for example if QQER
Q:2
E:1
R:1

Q>1, 2- 4/4 = 1.. so I need to replace 1 Q, in shortest window
it becomes to look for a shortest window that contains one Q

e.g. QQWERWEQ
Q:3 W:2 E:2 R:1
still only need to replace 1 Q

e.g. QQWERWQQ
Q:4 W:2 E:1 R:1
I need to replace shortest window with 2 Qs

"""


import collections
from typing import Collection


class Solution:
    def balancedString(self, s: str) -> int:
        charToIdx = {
            'Q': 0,
            'W': 1,
            'E': 2,
            'R': 3,
        }

        m = collections.Counter(s)
        toReplace = [0]*4
        for k, v in m.items():
            if v > len(s)//4:
                toReplace[charToIdx[k]] = v-len(s)//4

        if max(toReplace) == 0:
            return 0

        i, j = 0, 0
        minLen = len(s)
        while j < len(s):
            toReplace[charToIdx[s[j]]] -= 1

            while max(toReplace) <= 0:
                minLen = min(minLen, j-i+1)
                toReplace[charToIdx[s[i]]] += 1
                i += 1

            j += 1

        return minLen


"""
Runtime: 293 ms, faster than 71.59% of Python3 online submissions for Replace the Substring for Balanced String.
Memory Usage: 14.7 MB, less than 80.68% of Python3 online submissions for Replace the Substring for Balanced String.

then I began to think what others do..
I cannot think of a very fine way but I see a kind of queue based way
I need at most 3 queue to store the to-replace chars index..

then when queue len > required.. I can do a calculation
then I pop the smallest one...and do that again..

it might not be brilliant at all, but somehow I keep coming to this place.. so why don't I go code this up?
"""


class Solution:
    def balancedString(self, s: str) -> int:
        charToQueues = dict()

        '''
            'Q': [0, []],  # required replace count and the queue to keep the idxs
            'W': [0, []],
            'E': [0, []],
            'R': [0, []],
        '''

        m = collections.Counter(s)
        if max(m.values()) == len(s)//4:
            return 0

        for k, v in m.items():
            if v > len(s)//4:
                charToQueues[k] = [v-len(s)//4, []]

        def satisfy():
            for v in charToQueues.values():
                if v[0] > len(v[1]):
                    return False
            return True

        def popMinLeft():
            minLeft = len(s)
            toPop = None
            for k, v in charToQueues.items():
                if v[1][0] < minLeft:
                    toPop = k
                    minLeft = v[1][0]
            charToQueues[toPop][1] = charToQueues[toPop][1][1:]
            return minLeft

        minLen = len(s)
        for i in range(len(s)):
            if s[i] in charToQueues:
                charToQueues[s[i]][1].append(i)
                while satisfy():
                    minLeft = popMinLeft()
                    minLen = min(minLen, i - minLeft+1)

        return minLen


"""
39 / 40 test cases passed.
and TLE as last one
>>> len(s)
100000

so I am not O(N)?
yes... it is O(N^2) because the satisfy()
it go thru the queue? no.. it only need to the left most..


okay.. still passed
Runtime: 7659 ms, faster than 5.30% of Python3 online submissions for Replace the Substring for Balanced String.
Memory Usage: 18 MB, less than 6.82% of Python3 online submissions for Replace the Substring for Balanced String.

so it definitely is O(N) to O(N^2)... I am not sure where it slows me down
maybe I can change this to go and give a try
"""
"""
    def balancedString(self, s):
        count = collections.Counter(s)
        res = n = len(s)
        i = 0
        for j, c in enumerate(s):
            count[c] -= 1
            while i < n and all(n / 4 >= count[c] for c in 'QWER'):
                res = min(res, j - i + 1)
                count[s[i]] += 1
                i += 1
        return res

okay.. this Lee's code is kind of structurally like mine but as always he thinks in the opposite way
I went thru it.. but unable to catch it 

head spinning for a while but then I got it
let me put comment inline
"""


class Solution:
    def balancedString(self, s):
        count = collections.Counter(s)
        res = n = len(s)
        i = 0
        for j, c in enumerate(s):
            # so this is grow j, thus grow the window
            # because we count out-of-window, so put c into window we minus its count
            count[c] -= 1
            while i < n and all(n / 4 >= count[c] for c in 'QWER'):
                # this while loop is shrinking the window
                # when can we shrink the window? when the condition meets
                # after you shrink the window, the out-of-window will grow accordingly
                # thus add the count back (because we are counting the out-of-window)
                # note i does can get ahead of j, e.g. "QWER"
                # when j,i both start with 0, the window size is actually 1
                # then move j=0 into window, the condition meets... shrink the window to 0
                # then, j=0, i=1, j-i+1 => 0
                # then count[w] becomes 2 so it won't no longer be able to shrink further..
                # and j will follow i to the end..

                # this shrinking window with increasing count (out-of-window) is kind of counter-intuitive to me...
                res = min(res, j - i + 1)
                count[s[i]] += 1
                i += 1
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.balancedString("QQWRQQQQ"))
    print(s.balancedString("QQWQ"))
    print(s.balancedString("QQWR"))
    print(s.balancedString("QQWRQQQQ"))
    print(s.balancedString("QWER"))
