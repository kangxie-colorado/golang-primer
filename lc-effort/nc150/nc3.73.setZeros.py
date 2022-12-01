"""
https://leetcode.com/problems/set-matrix-zeroes/

did this before but still cannot remember how to do it
so I think 

Scan for every 0, set the top row and first colum to 0.. since this is a must happen anyway
Scan for row[1:].col[1:], if first row element is 0, or first colum is 0 then set it to zero

let me try
if stuck, debug a bit

but don't get stuck here for too long
"""


from typing import List


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if matrix[i][j] == 0:
                    matrix[0][j] = 0
                    matrix[i][0] = 0

        for i in range(1, len(matrix)):
            for j in range(1, len(matrix[i])):
                if matrix[i][0] == 0 or matrix[0][j] == 0:
                    matrix[i][j] = 0


"""
easily failed
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
I got
[[0,1,2,0],[3,4,5,0],[1,3,1,0]]

so just send the signal to first row?

"""


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if matrix[i][j] == 0:
                    matrix[0][j] = 0

        for i in range(1, len(matrix)):
            for j in range(len(matrix[i])):
                if matrix[0][j] == 0:
                    matrix[i][j] = 0


"""
then I failed this
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

huh...
so what is the order?

process the first row/column first... then do the thing
"""


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        firstCol = fristRow = False
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if matrix[i][j] == 0:
                    matrix[0][j] = 0
                    matrix[i][0] = 0
                    if i == 0:
                        fristRow = True
                    if j == 0:
                        firstCol = True

        for i in range(1, len(matrix)):
            for j in range(1, len(matrix[i])):
                if matrix[i][0] == 0 or matrix[0][j] == 0:
                    matrix[i][j] = 0

        if fristRow:
            matrix[0] = [0]*len(matrix[0])
        if firstCol:
            for r in matrix:
                r[0] = 0


"""
so end up doing
1. scan and record the value in first row/col
2. if some number in first row is 0, then mark it firstRow=True (to be set to 0 later)
2.1. the same for first col.. 
3. set row[1:].col[1:]
4. come back to set first row and col

Runtime: 132 ms, faster than 94.72% of Python3 online submissions for Set Matrix Zeroes.
Memory Usage: 14.8 MB, less than 54.57% of Python3 online submissions for Set Matrix Zeroes.
"""
