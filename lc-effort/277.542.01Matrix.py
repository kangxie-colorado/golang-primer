"""
https://leetcode.com/problems/01-matrix/


this is that prim algorithm???
relax the edges 
"""


from collections import defaultdict
from typing import List


class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        m, n = len(mat), len(mat[0])

        done = [[]]*m
        for i in range(m):
            done[i] = [-1]*n

        def getNeighbors(r, c):
            res = []
            if r > 0:
                res.append((r-1, c))
            if c > 0:
                res.append((r, c-1))
            if r < m-1:
                res.append((r+1, c))
            if c < n-1:
                res.append((r, c+1))
            return res

        frontLine = defaultdict(int)
        currDist = 0
        for i in range(m):
            for j in range(n):
                if mat[i][j] == currDist:
                    frontLine[(i, j)] = currDist

        while True:
            nextFrontLine = defaultdict(int)
            for node, dist in frontLine.items():
                done[node[0]][node[1]] = min(
                    dist, done[node[0]][node[1]]) if done[node[0]][node[1]] != -1 else dist
                for nei in getNeighbors(*node):
                    if nei not in frontLine and nei not in nextFrontLine and done[nei[0]][nei[1]] == -1:
                        nextFrontLine[nei] = dist+1
            if not nextFrontLine:
                return done
            frontLine = nextFrontLine


"""
class Solution:  # 520 ms, faster than 96.50%
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        m, n = len(mat), len(mat[0])

        for r in range(m):
            for c in range(n):
                if mat[r][c] > 0:
                    top = mat[r - 1][c] if r > 0 else math.inf
                    left = mat[r][c - 1] if c > 0 else math.inf
                    mat[r][c] = min(top, left) + 1

        for r in range(m - 1, -1, -1):
            for c in range(n - 1, -1, -1):
                if mat[r][c] > 0:
                    bottom = mat[r + 1][c] if r < m - 1 else math.inf
                    right = mat[r][c + 1] if c < n - 1 else math.inf
                    mat[r][c] = min(mat[r][c], bottom + 1, right + 1)

        return mat

cool!
two scans.. 

top-left to finish
bottom-right to finish..

open my eyes
"""
