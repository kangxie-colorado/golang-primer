"""
https://leetcode.com/problems/network-delay-time/

so use a heap to maintain next to visited
but keep the res incremented each time by the next cost

until it reaches n nodes.. or not able to
"""


from cmath import cos
from collections import defaultdict
from email.policy import default
import heapq
from time import time
from typing import List


class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        res = 0

        # first convert to adjacent lists
        M = defaultdict(list)
        for edge in times:
            u, v, w = edge
            M[u].append((v, w))

        h = []  # (cost, vertex)
        heapq.heappush(h, (0, k))

        while h and n > 0:
            cost, u = heapq.heappop(h)
            n -= 1
            res = max(res, cost)

            for v, w in M[u]:
                heapq.heappush(h, (w+cost, v))

        if n > 0:
            return -1
        return res


"""
[[1,2,1],[2,3,7],[1,3,4],[2,1,2]]
3
1

failed here.. there is a cycle here
I thought there is no cycle.. let me see
"""


class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        res = 0

        # first convert to adjacent lists
        M = defaultdict(list)
        for edge in times:
            u, v, w = edge
            M[u].append((v, w))

        h = []  # (cost, vertex)
        heapq.heappush(h, (0, k))
        visited = set()

        while h and n > 0:
            cost, u = heapq.heappop(h)
            if u in visited:
                continue
            visited.add(u)

            n -= 1
            res = max(res, cost)

            for v, w in M[u]:
                heapq.heappush(h, (w+cost, v))

        if n > 0:
            return -1
        return res


"""
Runtime: 656 ms, faster than 61.35% of Python3 online submissions for Network Delay Time.
Memory Usage: 16 MB, less than 71.19% of Python3 online submissions for Network Delay Time.

Runtime: 528 ms, faster than 82.24% of Python3 online submissions for Network Delay Time.
Memory Usage: 16.1 MB, less than 55.64% of Python3 online submissions for Network Delay Time.
"""
if __name__ == '__main__':
    s = Solution()
    assert 2 == s.networkDelayTime([[2, 1, 1], [2, 3, 1], [3, 4, 1]], n=4, k=2)
    assert 1 == s.networkDelayTime([[2, 1, 1], [2, 3, 1], [3, 4, 1]], n=3, k=2)
    assert 1 == s.networkDelayTime(times=[[1, 2, 1]], n=2, k=1)
    assert -1 == s.networkDelayTime(times=[[1, 2, 1]], n=2, k=2)
    assert 4 == s.networkDelayTime(
        [[1, 2, 1], [2, 3, 7], [1, 3, 4], [2, 1, 2]], 3, 1)
