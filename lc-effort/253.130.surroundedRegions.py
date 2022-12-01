"""
https://leetcode.com/problems/surrounded-regions/

okay.. this is like that sink islands I think
only catch is the border cannot be sunk

think I start from the border then keep all the unsunk alive and then sink the rest
"""


from typing import List


class Solution:
    def solve(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        m, n = len(board), len(board[0])

        markers = [[]]*m
        for r in range(m):
            markers[r] = [0]*n

        def helper(i, j):
            if board[i][j] == 'O':
                q = [(i, j)]
                while q:
                    x, y = q[0]
                    q = q[1:]
                    if x < 0 or y < 0 or x >= m or y >= n:
                        continue
                    if markers[x][y] != 0:
                        continue
                    if board[x][y] == 'O':
                        markers[x][y] = 1
                        q.append((x-1, y))
                        q.append((x+1, y))
                        q.append((x, y-1))
                        q.append((x, y+1))
                    else:
                        markers[x][y] = -1

        for i in (0, m-1):
            for j in range(n):
                helper(i, j)

        for j in (0, n-1):
            for i in range(m):
                helper(i, j)

        for i in range(m):
            for j in range(n):
                board[i][j] = 'X'
                if markers[i][j] == 1:
                    board[i][j] = 'O'


"""
Runtime: 150 ms, faster than 89.30% of Python3 online submissions for Surrounded Regions.
Memory Usage: 15.2 MB, less than 98.03% of Python3 online submissions for Surrounded Regions.
"""

if __name__ == '__main__':
    B = [["O", "O", "O"], ["O", "O", "O"], ["O", "O", "O"]]
    print(Solution().solve(B))
