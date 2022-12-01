from typing import List


class Solution:
    def getMaximumGold(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        res = 0

        def maxGolds(x, y, gold):
            nonlocal res
            if x < 0 or y < 0 or x >= m or y >= n or grid[x][y] == 0:
                res = max(res, gold)
                return 0

            for dx, dy in ((-1, 0), (1, 0), (0, -1), (0, 1)):
                goldHere = grid[x][y]
                grid[x][y] = 0
                maxGolds(x+dx, y+dy, gold+goldHere)
                grid[x][y] = goldHere

        for r in range(m):
            for c in range(n):
                maxGolds(r, c, 0)
        return res


if __name__ == "__main__":
    s = Solution()
    A = [[0, 6, 0], [5, 8, 7], [0, 9, 0]]
    print(s.getMaximumGold(A))
