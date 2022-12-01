"""


"""


from collections import defaultdict
from typing import List


class Solution:
    def longestLine(self, mat: List[List[int]]) -> int:
        # row sums
        m, n = len(mat), len(mat[0])
        res = 0
        for i in range(m):
            prefix = 0
            for j in range(n):
                runOfOnes = prefix + mat[i][j]
                prefix = runOfOnes if mat[i][j] else 0
                res = max(res, runOfOnes)

        for j in range(n):
            prefix = 0
            for i in range(m):
                runOfOnes = prefix + mat[i][j]
                prefix = runOfOnes if mat[i][j] else 0
                res = max(res, runOfOnes)

        # edge cases... edge cases...
        # when the matrix is extremely unbalanced...
        # pay attention...
        if m >= 2 and n >= 2 and res < min(m, n):
            for diag in range(1-m, n):
                prefix = 0
                for i in range(m):
                    j = i+diag
                    if 0 <= j < n:
                        runOfOnes = prefix + mat[i][j]
                        prefix = runOfOnes if mat[i][j] else 0
                        res = max(res, runOfOnes)

            for antiDiag in range(m+n-1):
                prefix = 0
                for i in range(m):
                    j = antiDiag-i
                    if 0 <= j < n:
                        runOfOnes = prefix + mat[i][j]
                        prefix = runOfOnes if mat[i][j] else 0
                        res = max(res, runOfOnes)

        return res


if __name__ == '__main__':
    mat = [[1, 1, 0, 0, 1, 0, 0, 1, 1, 0], [1, 0, 0, 1, 0, 1, 1, 1, 1, 1], [1, 1, 1, 0, 0, 1, 1, 1, 1, 0], [0, 1, 1, 1, 0, 1, 1, 1, 1, 1], [0, 0, 1, 1, 1, 1, 1, 1, 1, 0], [
        1, 1, 1, 1, 1, 1, 0, 1, 1, 1], [0, 1, 1, 1, 1, 1, 1, 0, 0, 1], [1, 1, 1, 1, 1, 0, 0, 1, 1, 1], [0, 1, 0, 1, 1, 0, 1, 1, 1, 1], [1, 1, 1, 0, 1, 0, 1, 1, 1, 1]]

    print(Solution().longestLine(mat))
