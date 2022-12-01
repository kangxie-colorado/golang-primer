"""
https://leetcode.com/problems/min-cost-to-connect-all-points/

didn't do this before 
at a glance, this needs mini span tree knowledge probably I don't master
let me see

a heap problem?
but where to start?

might just be wrong.. but let me try heap... 
not natural... 

okay... don't waste my brain cells and time
just go reading some MST background theory

that PRIM algorithm is exactly like what I am thinking
"""


from dis import dis
from typing import List
import heapq


class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        h = []  # (dist, x,y)
        heapq.heappush(h, (0, *points[0]))
        connected = set()
        res = 0

        while h:
            dist, x, y = heapq.heappop(h)
            if (x, y) in connected:
                continue
            connected.add((x, y))
            res += dist
            for pX, pY in points:
                if (pX, pY) in connected:
                    continue
                heapq.heappush(h, (abs(pX-x)+abs(pY-y), pX, pY))

        return res


"""
at least verified the solution 71 / 72 test cases passed.

TLE at last test

maybe there is another twist?
so maybe I should use the Krushal's algorithm


func union(x, y int, parents []int) {
	parents[find(x, parents)] = find(y, parents)
}

func find(x int, parents []int) int {
	// parents[x] will be initialized to x
	// so when x != parents[x], it has been unioned into another set

	if x != parents[x] {
		// follow the parentage link to the source
		//
		parents[x] = find(parents[x], parents)
	}

	return parents[x]
}

for ref. I kind of forgot how to write the union/find 
"""


class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        parents = [i for i in range(len(points))]

        def find(x):
            if parents[x] != x:
                parents[x] = find(parents[x])
            return parents[x]

        def union(x, y):
            parents[find(x)] = find(y)

        edges = []  # (len, x, y)
        for i in range(len(points)):
            for j in range(i+1, len(points)):
                p1, p2 = points[i], points[j]
                edges.append((abs(p1[0]-p2[0])+abs(p1[1]-p2[1]), i, j))

        res = 0
        edges.sort()
        for cost, p1, p2 in edges:
            if find(p1) != find(p2):
                res += cost
                union(p2, p1)
        return res


"""
Runtime: 6512 ms, faster than 15.20% of Python3 online submissions for Min Cost to Connect All Points.
Memory Usage: 81.7 MB, less than 68.80% of Python3 online submissions for Min Cost to Connect All Points.

okay.. this can pass

the prim needs some optimization 
I see the issue is I added multiple copies to the heap..

maybe i can remember all the min distance like this
so only add new minimum if there is any... 
otherwise, no touch heap so that logV don't grow too much
            # 3. update min distance for neighbors in graph if not in mst and add to heap
            for v in range(n):
                if v not in mst:
                    d = abs(points[v][0]-points[min_idx][0]) + abs(points[v][1]-points[min_idx][1])
                    if d < dist[v]:
                        dist[v] = d
                        heapq.heappush(q, (d, v))


"""


class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        h = []  # (dist, x,y)
        heapq.heappush(h, (0, *points[0]))
        connected = set()
        res = 0
        dists = [10**7] * len(points)

        while h:
            dist, x, y = heapq.heappop(h)
            if (x, y) in connected:
                continue
            connected.add((x, y))
            res += dist
            for i in range(len(points)):
                pX, pY = points[i]
                d = abs(pX-x)+abs(pY-y)
                if d < dists[i]:
                    dists[i] = d
                    heapq.heappush(h, (d, pX, pY))

        return res


"""
Runtime: 2139 ms, faster than 76.56% of Python3 online submissions for Min Cost to Connect All Points.
Memory Usage: 17.2 MB, less than 92.31% of Python3 online submissions for Min Cost to Connect All Points.

oh, there might be more than n elements so while h as my loop condition is wasting
it has to deal with more than n loops... 

so keeping a count to break may also help
"""


class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        h = []  # (dist, x,y)
        heapq.heappush(h, (0, *points[0]))
        connected = set()
        res = 0
        cnt = 0

        while h:
            dist, x, y = heapq.heappop(h)
            if (x, y) in connected:
                continue
            connected.add((x, y))
            res += dist
            cnt += 1
            if cnt >= len(points):
                break
            for pX, pY in points:
                if (pX, pY) in connected:
                    continue
                heapq.heappush(h, (abs(pX-x)+abs(pY-y), pX, pY))

        return res


"""
Runtime: 1169 ms, faster than 95.19% of Python3 online submissions for Min Cost to Connect All Points.
Memory Usage: 66 MB, less than 89.49% of Python3 online submissions for Min Cost to Connect All Points.

so this is the major break/make twist
"""
