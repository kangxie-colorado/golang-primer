"""
https://leetcode.com/problems/evaluate-reverse-polish-notation/


"""


from typing import List


class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []

        def apply(op, x1, x2):
            if op == '+':
                return x1+x2
            if op == '-':
                return x1-x2
            if op == '*':
                return x1*x2
            if op == '/':
                return int(x1/x2)
            return 0

        for t in tokens:
            if t in ('+-*/'):
                n1 = stack.pop()
                n2 = stack.pop()
                stack.append(apply(t, n2, n1))
            else:
                stack.append(int(t))

        return stack.pop()


if __name__ == "__main__":
    s = Solution()
    tokens = ["10", "6", "9", "3", "+", "-11",
              "*", "/", "*", "17", "+", "5", "+"]

    print(s.evalRPN(tokens))
