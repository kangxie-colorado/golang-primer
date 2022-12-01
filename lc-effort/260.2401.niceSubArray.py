"""
https://leetcode.com/problems/longest-nice-subarray/

example here
Input: nums = [1,3,8,48,10]
Output: 3
Explanation: The longest nice subarray is [3,8,48]. This subarray satisfies the conditions:
- 3 AND 8 = 0.
- 3 AND 48 = 0.
- 8 AND 48 = 0.
It can be proven that no longer nice subarray can be obtained, so we return 3.


this is a sliding window, of course 
random thoughts:
you cannot use 3 and 8, then "and 48" ... when 3 and 8 == 0, 0 and anything is 0
so you need to see the new number ands everything in the window to be zero.. 

(3 ,8) 
enter 48, if 48 & 3 = 0 and 48&8 = 0, then it can enter the window
which I think it equivilent to 48 & (3|8)

then how to slide the window
when the window doesn't meet the NICE condition.. then it means there must be two numbers that are not & to 0
e.g. 1 3
then contract the window.. until it is nice( or length 1)
when contracting, how to maintain the or-value??
>>> 3|8
11
>>> 11^3
8

so it should be xor(^)? nah.. see this 
>>> 1|3
3
>>> 3^1
2

actually let me think
when some number enter the window and make it not-nice.. 
I could jump straight to this number... because whatever I do, the max window is also accounted for?

.. NICE-WINDOW. n1..
when n1 enter the window.. it becomes not-nice..
then even I remove the first number it becomes nice, the max is still len(nice-window)
so I can jump to n1???
nah... 

what if Nice-Window[1:n1+d] is nice again.. 
so cannot jump

back to think how to maintain the or-value
maybe it is xor afterall, for non-nice window, you cannot do that..
but for nice-window values. because a1&a2 is zero, they don't have overlap 1s
so xor that will work

>>> 3|8|48
59
>>> 59^3
56
>>> 56^8
48

yeah.. looke like so
okay, now I can code a bit
"""


from typing import List


class Solution:
    def longestNiceSubarray(self, nums: List[int]) -> int:
        res = 1
        i, j = 0, 0
        orValue = 0
        while j < len(nums):
            if orValue & nums[j] == 0:
                res = max(res, j-i+1)
                orValue |= nums[j]
            else:
                while orValue & nums[j] != 0:
                    orValue ^= nums[i]
                    i += 1
                orValue |= nums[j]
            j += 1
        return res


"""
Runtime: 1671 ms, faster than 45.75% of Python3 online submissions for Longest Nice Subarray.
Memory Usage: 29.3 MB, less than 46.63% of Python3 online submissions for Longest Nice Subarray.
"""


class Solution:
    def longestNiceSubarray(self, nums: List[int]) -> int:
        res = 1
        i, j = 0, 0
        orValue = 0
        while j < len(nums):
            if orValue & nums[j] == 0:
                res = max(res, j-i+1)
                orValue |= nums[j]
                j += 1
            else:
                orValue ^= nums[i]
                i += 1

        return res


"""
change the logical a bit
Runtime: 2770 ms, faster than 19.97% of Python3 online submissions for Longest Nice Subarray.
Memory Usage: 29.4 MB, less than 46.63% of Python3 online submissions for Longest Nice Subarray.


what is a bit tricky is 
we contract when the window is not NICE to NICE again...
not like the regular contract until it become not-satisfying-condition


okay.. Lee's Code
https://leetcode.com/problems/longest-nice-subarray/discuss/2527496/JavaC%2B%2BPython-Sliding-Window

    def longestNiceSubarray(self, A):
        res = AND = i = 0
        for j in range(len(A)):
            while AND & A[j]:
                AND ^= A[i]
                i += 1
            AND |= A[j]
            res = max(res, j - i + 1)
        return res

I wanted to do this but cannot clear my path 
"""
