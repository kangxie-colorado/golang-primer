from typing import List


class Solution:
    def numSubmat(self, mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        res = 0

        rowSums = [[]]*m
        colSums = [[]]*m
        for r in range(m):
            rowSums[r] = [0]*n
            colSums[r] = [0]*n

        rowSums[0][0] = colSums[0][0] = mat[0][0]

        for i in range(0, m):
            for j in range(0, n):
                if mat[i][j]:
                    if j == 0:
                        rowSums[i][j] = mat[i][j]
                    else:
                        rowSums[i][j] = rowSums[i][j-1] + 1

                    if i == 0:
                        colSums[i][j] = mat[i][j]
                    else:
                        colSums[i][j] = colSums[i-1][j] + 1

                    res += colSums[i][j]*rowSums[i][j]

        return res


""""
wrong solution

should resort back to see the acreage -- nah
still not right

so calculate the column sums and looking to the left to see how many each new cell can contribute

let me say at row 3, the column sum looks like
    1
  1 2 
1 2 3

what would be the cell[2][2] contributing ?

on column 3.. of course, it contribute 3
from column 2.. it should contribute 2.. how to get that, min(3,2)
from column 1.. it should contribute 1.. how to get that, min(3,1)

so still sort of like acreage 
"""


class Solution:
    def numSubmat(self, mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        res = 0

        for i in range(m):
            for j in range(0, n):
                if mat[i][j]:
                    # calculate the column pre-sums
                    if i > 0:
                        mat[i][j] = mat[i-1][j] + 1
                    k = j
                    minHeightRow = mat[i][j]
                    while k >= 0 and mat[i][k] != 0:
                        # calculate each column, how many this cell can contribute to
                        # note, it needs to keep at the lowest
                        # e.g. 3 1 2 --> then between height-3 and height-2.. the effective height is only 1
                        # this generize the current column too
                        minHeightRow = min(mat[i][j], mat[i][k], minHeightRow)
                        res += minHeightRow
                        k -= 1
        return res


"""
[[1,0,1,1,1,1,1],[1,1,0,0,0,1,1],[1,1,1,0,0,1,1],[1,0,1,0,1,0,1],[1,0,1,1,1,0,1],[1,1,0,1,1,1,1],[1,0,0,1,1,0,1]]

I found that for the scenario like 3 1 2.. I need to use 1 to connect 2 and 3
okay.. after this

Runtime: 2996 ms, faster than 22.52% of Python3 online submissions for Count Submatrices With All Ones.
Memory Usage: 14.7 MB, less than 73.18% of Python3 online submissions for Count Submatrices With All Ones.

Runtime: 1421 ms, faster than 53.31% of Python3 online submissions for Count Submatrices With All Ones.
Memory Usage: 14.8 MB, less than 28.48% of Python3 online submissions for Count Submatrices With All Ones.

this looking to left column by column is a bit O(n^2)
can use monotonic stack to reduce that to O(n)
"""


class Solution:
    def numSubmat(self, mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        res = 0

        for i in range(m):
            rowStack = []  # (val,idx)
            for j in range(0, n):
                if mat[i][j]:
                    # calculate the column pre-sums
                    if i > 0:
                        mat[i][j] = mat[i-1][j] + 1
                    idx = j
                    # this column itself
                    res += mat[i][j]
                    # anything bigger than me
                    while rowStack and rowStack[-1][0] >= mat[i][j]:
                        _, _, idx = rowStack.pop()
                    res += mat[i][j] * (j-idx)

                    # anything smaller than me
                    for val, origIdx, extendIdx in rowStack:
                        res += val * (origIdx-extendIdx+1)
                    rowStack.append((mat[i][j], j, idx))
                else:
                    rowStack = []
        return res


"""
Runtime: 871 ms, faster than 65.90% of Python3 online submissions for Count Submatrices With All Ones.
Memory Usage: 14.7 MB, less than 28.48% of Python3 online submissions for Count Submatrices With All Ones.

the difficulty is the reconcile between the extended and non-extended scenarios..
                    while rowStack and rowStack[-1][0] >= mat[i][j]:
                        _, _, idx = rowStack.pop()
                    res += mat[i][j] * (j-idx)

                    for val, origIdx, extendIdx in rowStack:
                        res += val * (origIdx-extendIdx+1)
                    rowStack.append((mat[i][j], j, idx))

because however without one variable you do it.. you would miss another one!
so I don't know... I used the 2nd idx.. and it seems to work

Runtime: 852 ms, faster than 67.88% of Python3 online submissions for Count Submatrices With All Ones.
Memory Usage: 14.6 MB, less than 73.18% of Python3 online submissions for Count Submatrices With All Ones.

okay... anything smaller than me

                    # anything smaller than me
                    for val, origIdx, extendIdx in rowStack:
                        res += val * (origIdx-extendIdx+1)
            this can be further simplied by saving up the sum at that point... 
            I have another idea.. let me code it up 
"""


class Solution:
    def numSubmat(self, mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        res = 0

        for i in range(m):
            rowStack = [(0, -1, 0)]  # (val,idx, dpSum)
            for j in range(0, n):
                if mat[i][j]:
                    contributions = 0
                    # calculate the column pre-sums
                    if i > 0:
                        mat[i][j] = mat[i-1][j] + 1
                    idx = j
                    # anything bigger equal than me
                    while rowStack and rowStack[-1][0] >= mat[i][j]:
                        rowStack.pop()
                    contributions += mat[i][j] * (j-rowStack[-1][1])
                    contributions += rowStack[-1][2]
                    rowStack.append((mat[i][j], j, contributions))
                    res += contributions

                else:
                    # insert another sentinel is fine
                    rowStack.append((0, j, 0))
        return res


""""
Runtime: 774 ms, faster than 74.50% of Python3 online submissions for Count Submatrices With All Ones.
Memory Usage: 14.7 MB, less than 73.18% of Python3 online submissions for Count Submatrices With All Ones.

"""


if __name__ == "__main__":

    mat = [[1, 0, 1, 1, 1, 1, 1], [1, 1, 0, 0, 0, 1, 1], [1, 1, 1, 0, 0, 1, 1], [
        1, 0, 1, 0, 1, 0, 1], [1, 0, 1, 1, 1, 0, 1], [1, 1, 0, 1, 1, 1, 1], [1, 0, 0, 1, 1, 0, 1]]
    print(Solution().numSubmat(mat))
    mat = [[0, 1, 1, 0], [0, 1, 1, 1], [1, 1, 1, 0]]
    print(Solution().numSubmat(mat))
    mat = [[1, 0, 1], [1, 1, 0], [1, 1, 0]]
    print(Solution().numSubmat(mat))
