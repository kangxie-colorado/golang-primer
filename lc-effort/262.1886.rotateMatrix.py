"""
https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/

at most 4 shapes... 
to rotate 90 degree, it would be a little hard to code 
so the easier way is 
swap row with column and fold
"""


from typing import List


class Solution:
    def findRotation(self, mat: List[List[int]], target: List[List[int]]) -> bool:
        def rotate90():
            nonlocal mat
            for r in range(len(mat)):
                for c in range(r):
                    mat[r][c], mat[c][r] = mat[c][r], mat[r][c]

            for r in range(len(mat)):
                i, j = 0, len(mat[r])-1
                while i < j:
                    mat[r][i], mat[r][j] = mat[r][j], mat[r][i]
                    i, j = i+1, j-1

        for i in range(4):
            if mat == target:
                return True
            rotate90()

        return False


"""
made a mistake for swap row/col
you only need to do the lower half below the diagonal line

so this
            for r in range(len(mat)):
                for c in range(len(mat[r])//2):
                    mat[r][c], mat[c][r] = mat[c][r], mat[r][c]
will be 
            for r in range(len(mat)):
                for c in range(r):
                    mat[r][c], mat[c][r] = mat[c][r], mat[r][c]


Runtime: 45 ms, faster than 93.07% of Python3 online submissions for Determine Whether Matrix Can Be Obtained By Rotation.
Memory Usage: 14 MB, less than 29.02% of Python3 online submissions for Determine Whether Matrix Can Be Obtained By Rotation.
"""


if __name__ == '__main__':
    s = Solution()
    mat = [[0, 0, 0], [0, 1, 0], [1, 1, 1]]
    target = [[1, 1, 1], [0, 1, 0], [0, 0, 0]]
    print(s.findRotation(mat, target))
