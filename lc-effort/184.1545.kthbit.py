"""
https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/

of course I can get the whole string and then get the kth bit...
and it is pretty obvious (that invert reverse is just distraction)

left+right is 1..
left 1 right..

e.g.

S2 = "0 1 1"
S3 = "011 1 001"
S4 = "0111001 1 0110001"

therefore it is easy to get the string
"""


import math


class Solution:
    def findKthBit(self, n: int, k: int) -> str:
        s = '0'
        t = s
        for i in range(2, n+1):
            t2 = ''
            for c in t:
                t2 += '1' if c == '0' else '0'
            s = t + '1' + t2[::-1]
            t = s

        return s[k-1]


"""
Runtime: 945 ms, faster than 51.03% of Python3 online submissions for Find Kth Bit in Nth Binary String.
Memory Usage: 17 MB, less than 64.38% of Python3 online submissions for Find Kth Bit in Nth Binary String.

then maybe I don't have to calculate the string
use n=4 k=11

1. decide which part it is in
k=11, index = 10

0111001 1 0110001
            ^
because 10 >8, 10<15
so 10 is at the right part of n=4 full string
so 10 is just the reverse bit of index-4

001 vs 011 (reversed or before-reverse 110)
10-4
    10 >= 2^3 (8), 8-1=7, 7*2 - 10 =4       10 = invert(4)
        4>= 2^2 (4), 4-1=3, 3*2 - 4 = 2      4 = invert(2)
            2>=2^1, 2-1=1, 2*1-2=0 -===> 0   2 = invert(0)


9-5
    9 = invert(5)
    5 = invert(1)
    1 ==> base-case 1 --> 5:0 ---> 9:1

"""


class Solution:
    def findKthBit(self, n: int, k: int) -> str:
        inverts = 0

        def getToBase(K):
            nonlocal inverts
            if K == 0:
                return 0
            lowPow = int(math.log2(K))
            if K == 2**(lowPow+1)-1:
                return 1
            inverts += 1
            next = (2**lowPow - 1)*2 - K
            return getToBase(next)

        base = getToBase(k-1)
        inverts %= 2
        return str(base ^ inverts)


"""
failed
4
12
-1 vs 0

12 -> idx-11 -> idx-3... aha... 3 is the middle 1
just as 1 is the base case 
3 should be too

so it is this 
            if K == 2**(lowPow+1)-1:
                return 1

Runtime: 48 ms, faster than 83.56% of Python3 online submissions for Find Kth Bit in Nth Binary String.
Memory Usage: 13.9 MB, less than 75.34% of Python3 online submissions for Find Kth Bit in Nth Binary String.

then I guess 
        if K <= 1:
            return K
can be removed to let the generic base case to dominate
            if K == 2**(lowPow+1)-1:
                return 1

hmm.. case-0 cannot be taken away..
"""

if __name__ == '__main__':
    s = Solution()
    print(s.findKthBit(4, 12))
    print(s.findKthBit(4, 10))
    print(s.findKthBit(4, 9))
    print(s.findKthBit(4, 11))
