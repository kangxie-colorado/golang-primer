"""
https://leetcode.com/problems/path-with-minimum-effort/

just seems like a heap+dfs 
well heap based bfs/dfs.. 

using a queue, bfs
using a stack, dfs
using a heap... dijstra
"""


import heapq
from typing import List


class Solution:
    def minimumEffortPath(self, heights: List[List[int]]) -> int:
        res = 0
        m, n = len(heights), len(heights[0])
        h = [(0, 0, 0)]  # (effort, x,y)
        heapq.heapify(h)
        while len(h):
            node = heapq.heappop(h)
            res = max(res, node[0])
            x, y = node[1], node[2]
            if heights[x][y] == -1:
                continue
            if x == m-1 and y == n-1:
                return res

            if x > 0 and heights[x-1][y] != -1:
                heapq.heappush(
                    h, (abs(heights[x][y] - heights[x-1][y]), x-1, y))
            if x < m-1 and heights[x+1][y] != -1:
                heapq.heappush(
                    h, (abs(heights[x][y] - heights[x+1][y]), x+1, y))
            if y > 0 and heights[x][y-1] != -1:
                heapq.heappush(
                    h, (abs(heights[x][y] - heights[x][y-1]), x, y-1))
            if y < n-1 and heights[x][y+1] != -1:
                heapq.heappush(
                    h, (abs(heights[x][y] - heights[x][y+1]), x, y+1))
            heights[x][y] = -1


"""
Runtime: 984 ms, faster than 81.21% of Python3 online submissions for Path With Minimum Effort.
Memory Usage: 15.3 MB, less than 78.42% of Python3 online submissions for Path With Minimum Effort.
"""

if __name__ == '__main__':
    s = Solution()
    m = [[10, 8], [10, 8], [1, 2], [10, 3], [1, 3], [6, 3], [5, 2]]
    print(s.minimumEffortPath(m))
    print(s.minimumEffortPath([[1, 2, 2], [3, 8, 2], [5, 3, 5]]))
    print(s.minimumEffortPath([[1, 2, 3], [3, 8, 4], [5, 3, 5]]))
