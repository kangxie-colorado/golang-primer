"""
https://leetcode.com/problems/different-ways-to-add-parentheses/

solution is not obvious

"""


from ast import operator
from itertools import count
from typing import List


class Solution:
    def diffWaysToCompute(self, expression: str) -> List[int]:

        def helper(exp):
            m = []
            for i, c in enumerate(exp):
                if c in '+-*':
                    m.append(i)
            if len(m) == 0:
                return [int(exp)]
            res = []
            for i in m:
                left = helper(exp[:i])
                right = helper(exp[i+1:])
                for n1 in left:
                    for n2 in right:
                        op = exp[i]
                        if op == '+':
                            res.append(n1+n2)
                        if op == '-':
                            res.append(n1-n2)
                        if op == '*':
                            res.append(n1*n2)
            return res
        return helper(expression)


"""
Runtime: 45 ms, faster than 76.26% of Python3 online submissions for Different Ways to Add Parentheses.
Memory Usage: 14.1 MB, less than 17.36% of Python3 online submissions for Different Ways to Add Parentheses.

wowo... not bad...

Runtime: 62 ms, faster than 40.01% of Python3 online submissions for Different Ways to Add Parentheses.
Memory Usage: 13.9 MB, less than 84.66% of Python3 online submissions for Different Ways to Add Parentheses.
"""

if __name__ == '__main__':
    s = Solution()
    print(s.diffWaysToCompute('2-1-1'))
    print(s.diffWaysToCompute('2*3-4*5'))
