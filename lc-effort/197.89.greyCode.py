"""
https://leetcode.com/problems/gray-code/

form 
1: 0 1
2: 0 1 3 2
I took a guess for 3
3: 0 1 3 2  6 7 5 4
verify 
000 001 011 010   110 111 101 100
so naturally add 3rd bit, 6 7 5 4 is just like 2 3 1 0... 
then to keep the first/last diff by only 1 bit.. it has to be 4 for n=3 

so this is easy to get result now..

"""


from typing import List


class Solution:
    def grayCode(self, n: int) -> List[int]:
        res = [0, 1]
        for i in range(1, n):
            res = res + [e + 2**i for e in res[::-1]]

        return res


"""
Runtime: 176 ms, faster than 52.90% of Python3 online submissions for Gray Code.
Memory Usage: 21.3 MB, less than 72.90% of Python3 online submissions for Gray Code.

Runtime: 141 ms, faster than 76.16% of Python3 online submissions for Gray Code.
Memory Usage: 21.2 MB, less than 91.27% of Python3 online submissions for Gray Code.
"""
