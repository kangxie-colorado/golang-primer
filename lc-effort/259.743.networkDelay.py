"""
https://leetcode.com/problems/network-delay-time/

bellman ford obviously 
but I kind of forgot how to do 

so for every node, relax all the edges.
O(V.E)??

okay.. actually I did this before using heap
but now let me use bellman ford
"""


import math
import sys
from typing import List


class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        dist = [sys.maxsize] * (n+1)
        # this is only for conviniece but don't let this spot take any meaning
        dist[0] = 0
        dist[k] = 0

        for i in range(n):
            for s, d, w in times:
                dist[d] = min(dist[d], dist[s]+w)

        res = max(dist)

        return res if res < sys.maxsize else -1


"""
Runtime: 3458 ms, faster than 5.01% of Python3 online submissions for Network Delay Time.
Memory Usage: 15.8 MB, less than 97.09% of Python3 online submissions for Network Delay Time.

"""
