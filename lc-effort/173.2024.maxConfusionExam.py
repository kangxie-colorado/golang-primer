"""
https://leetcode.com/problems/maximize-the-confusion-of-an-exam/


"""


class Solution:
    def maxConsecutiveAnswers(self, answerKey: str, k: int) -> int:
        def helper(c):
            i, j = 0, 0
            res = 0
            count = 0

            while j < len(answerKey):
                count += 1 if answerKey[j] == c else 0
                if count+k < j-i+1:
                    count -= 1 if answerKey[i] == c else 0
                    i += 1
                j += 1

            return j-i

        return max(helper('T'), helper('F'))


"""
Runtime: 429 ms, faster than 87.36% of Python3 online submissions for Maximize the Confusion of an Exam.
Memory Usage: 14.4 MB, less than 36.82% of Python3 online submissions for Maximize the Confusion of an Exam.
Next challenges:
"""

if __name__ == '__main__':
    s = Solution()
    print(s.maxConsecutiveAnswers("TTFF", 2))
