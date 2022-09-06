"""
https://leetcode.com/problems/swim-in-rising-water/

binary search + dfs(to both ends)
"""


from curses.ascii import SO
import heapq
from typing import List


class Solution:
    def swimInWater(self, grid: List[List[int]]) -> int:
        def reachStartAndEnd(m):
            if m < grid[n-1][n-1]:
                return False

            visited = [[]]*n
            for i in range(n):
                visited[i] = [0]*n
            # start to end (if it connects thru a smaller value, then it is still a yes)
            stack = [(0, 0)]
            while len(stack):
                x, y = stack.pop()
                if x == n-1 and y == n-1:
                    return True

                if x > 0 and visited[x-1][y] != -1 and grid[x-1][y] <= m:
                    stack.append((x-1, y))
                if x < n-1 and visited[x+1][y] != -1 and grid[x+1][y] <= m:
                    stack.append((x+1, y))
                if y > 0 and visited[x][y-1] != -1 and grid[x][y-1] <= m:
                    stack.append((x, y-1))
                if y < n-1 and visited[x][y+1] != -1 and grid[x][y+1] <= m:
                    stack.append((x, y+1))

                visited[x][y] = -1
            return False

        n = len(grid)
        l, r = grid[0][0], n**2-1

        while l < r:
            m = l+(r-l)//2
            if reachStartAndEnd(m):
                r = m
            else:
                l = m+1

        return l


"""
Runtime: 218 ms, faster than 36.54% of Python3 online submissions for Swim in Rising Water.
Memory Usage: 14.5 MB, less than 90.66% of Python3 online submissions for Swim in Rising Water.

Runtime: 172 ms, faster than 60.44% of Python3 online submissions for Swim in Rising Water.
Memory Usage: 14.4 MB, less than 90.66% of Python3 online submissions for Swim in Rising Water.
"""

"""
interesting, this can also be solved by priority queue
    def swimInWater(self, grid):
        N, pq, seen, res = len(grid), [(grid[0][0], 0, 0)], set([(0, 0)]), 0
        while True:
            T, x, y = heapq.heappop(pq)
            res = max(res, T)
            if x == y == N - 1:
                return res
            for i, j in [(x + 1, y), (x, y + 1), (x - 1, y), (x, y - 1)]:
                if 0 <= i < N and 0 <= j < N and (i, j) not in seen:
                    seen.add((i, j))
                    heapq.heappush(pq, (grid[i][j], i, j))

the min-heap, makes sure next elevation that can be seen pops out
the res = max(res, T) makes sure the biggest elevation end up res on the path..

kind of philosiphcal
let me just code it for practice

"""


class Solution:
    def swimInWater(self, grid: List[List[int]]) -> int:
        n = len(grid)
        res = grid[0][0]
        h = [(res, 0, 0)]
        heapq.heapify(h)
        visited = [[]]*n
        for i in range(n):
            visited[i] = [0]*n

        while len(h):
            T, x, y = heapq.heappop(h)
            res = max(res, T)
            if x == y == n-1:
                return res
            if x > 0 and visited[x-1][y] != -1:
                heapq.heappush(h, (grid[x-1][y], x-1, y))
            if x < n-1 and visited[x+1][y] != -1:
                heapq.heappush(h, (grid[x+1][y], x+1, y))
            if y > 0 and visited[x][y-1] != -1:
                heapq.heappush(h, (grid[x][y-1], x, y-1))
            if y < n-1 and visited[x][y+1] != -1:
                heapq.heappush(h, (grid[x][y+1], x, y+1))

            visited[x][y] = -1


"""
Runtime: 219 ms, faster than 36.16% of Python3 online submissions for Swim in Rising Water.
Memory Usage: 14.5 MB, less than 90.66% of Python3 online submissions for Swim in Rising Water.
"""

if __name__ == "__main__":
    s = Solution()

    grid = [[0, 1, 2, 3, 4], [24, 23, 22, 21, 5], [12, 13, 14, 15, 16],
            [11, 17, 18, 19, 20], [10, 9, 8, 7, 6]]  # 16
    print(s.swimInWater(grid))

    grid = [[0, 2], [1, 3]]  # 3
    print(s.swimInWater(grid))

    grid = [[1, 3], [2, 0]]  # 2
    print(s.swimInWater(grid))

    grid = [[2, 3], [1, 0]]  # 2
    print(s.swimInWater(grid))
