"""
https://leetcode.com/problems/cheapest-flights-within-k-stops/

tricky is the stop must be at most k
so I think still do the usual dijstra

but when stop is used up.. put its next's cost to a super big number to disqualify that...

"""


from collections import defaultdict
import heapq
from typing import List


class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        adjs = defaultdict(list)

        for s, d, p in flights:
            adjs[s].append((d, p))

        h = []  # (price, leftStop, nodeIdx)
        heapq.heappush(h, (0, k+1, src))
        seen = set()

        while h:
            price, leftStop, node = heapq.heappop(h)
            if node in seen or leftStop < 0:
                continue
            if node == dst:
                return price
            seen.add(node)

            for neighbor, neiPrice in adjs[node]:
                if leftStop == 0:
                    neiPrice = 2*100000
                heapq.heappush(h, (price+neiPrice, leftStop-1, neighbor))

        return -1


"""
failed here
4
[[0,1,1],[0,2,5],[1,2,1],[2,3,1]]
0
3
1

got -1, exp 6
okay... because 2 is passed by 0-1-2 and 0-5
and in the processing, I used seen set
but in this case.. it is not necessary to use seen at all
if there is a cycle
it will blow itself out by accumulating once and once again, right? not sure. let me see

"""


class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        adjs = defaultdict(list)

        for s, d, p in flights:
            adjs[s].append((d, p))

        h = []  # (price, leftStop, nodeIdx)
        heapq.heappush(h, (0, k+1, src))

        while h:
            price, leftStop, node = heapq.heappop(h)
            if node == dst:
                return price
            if leftStop <= 0:
                continue

            for neighbor, neiPrice in adjs[node]:
                if leftStop == 0:
                    neiPrice = 2*100000
                heapq.heappush(h, (price+neiPrice, leftStop-1, neighbor))

        return -1


"""
47 / 51 test cases passed.
and TLE

13
[[11,12,74],[1,8,91],[4,6,13],[7,6,39],[5,12,8],[0,12,54],[8,4,32],[0,11,4],[4,0,91],[11,7,64],
[6,3,88],[8,5,80],[11,10,91],[10,0,60],[8,7,92],[12,6,78],[6,2,8],[4,3,54],[3,11,76],[3,12,23],
[11,6,79],[6,12,36],[2,11,100],[2,5,49],[7,0,17],[5,8,95],[3,9,98],[8,10,61],[2,12,38],[5,7,58],
[9,4,37],[8,6,79],[9,0,1],[2,3,12],[7,10,7],[12,10,52],[7,2,68],[12,2,100],[6,9,53],[7,4,90],
[0,5,43],[11,2,52],[11,8,50],[12,4,38],[7,9,94],[2,7,38],[3,7,88],[9,12,20],[12,0,26],[10,5,38],
[12,8,50],[0,2,77],[11,0,13],[9,10,76],[2,6,67],[5,6,34],[9,7,62],[5,3,67]]
10
1
10

result is -1 and correct... so what is the optimization?

shall I add a bfs to detect if it is at all possible before diving in?
"""


class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        adjs = defaultdict(list)

        for s, d, p in flights:
            adjs[s].append((d, p))

        def reachAble(k):
            q = [src]
            while q and k:
                newq = []
                for n in q:
                    if n == dst:
                        return True
                    for nei, _ in adjs[n]:
                        newq.append(nei)
                q = newq
                k -= 1

            return False

        if not reachAble(k+2):
            return -1

        h = []  # (price, leftStop, nodeIdx)
        heapq.heappush(h, (0, k+1, src))

        while h:
            price, leftStop, node = heapq.heappop(h)
            if node == dst:
                return price
            if leftStop <= 0:
                continue

            for neighbor, neiPrice in adjs[node]:
                if leftStop == 0:
                    neiPrice = 2*100000
                heapq.heappush(h, (price+neiPrice, leftStop-1, neighbor))

        return -1


"""
this reduce the time to tell it -1 from 1m1s to 7s
but still not good enough
"""


"""
bellman-ford!
"""


class Solution:
    def findCheapestPrice(self, n: int, flights: List[List[int]], src: int, dst: int, k: int) -> int:
        prices = [200000]*n
        prices[src] = 0

        for i in range(k+1):
            pricesCopy = prices.copy()
            for s, d, p in flights:
                pricesCopy[d] = min(pricesCopy[d], prices[s]+p)
            prices = pricesCopy
        return prices[dst] if prices[dst] != 200000 else -1


if __name__ == '__main__':
    s = Solution()
    print(s.findCheapestPrice(n=4, flights=[[0, 1, 100], [1, 2, 100], [
          2, 0, 100], [1, 3, 600], [2, 3, 200]], src=0, dst=3, k=1))
    print(s.findCheapestPrice(n=3, flights=[[0, 1, 100], [
          1, 2, 100], [0, 2, 500]], src=0, dst=2, k=1))
    print(s.findCheapestPrice(n=3, flights=[[0, 1, 100], [
          1, 2, 100], [0, 2, 500]], src=0, dst=2, k=0))
    print(s.findCheapestPrice(n=13, flights=[[11, 12, 74], [1, 8, 91], [4, 6, 13], [7, 6, 39], [5, 12, 8], [0, 12, 54], [8, 4, 32], [0, 11, 4], [4, 0, 91], [11, 7, 64], [6, 3, 88], [8, 5, 80], [11, 10, 91], [10, 0, 60], [8, 7, 92], [12, 6, 78], [6, 2, 8], [4, 3, 54], [3, 11, 76], [3, 12, 23], [11, 6, 79], [6, 12, 36], [2, 11, 100], [2, 5, 49], [7, 0, 17], [5, 8, 95], [3, 9, 98], [8, 10, 61], [
          2, 12, 38], [5, 7, 58], [9, 4, 37], [8, 6, 79], [9, 0, 1], [2, 3, 12], [7, 10, 7], [12, 10, 52], [7, 2, 68], [12, 2, 100], [6, 9, 53], [7, 4, 90], [0, 5, 43], [11, 2, 52], [11, 8, 50], [12, 4, 38], [7, 9, 94], [2, 7, 38], [3, 7, 88], [9, 12, 20], [12, 0, 26], [10, 5, 38], [12, 8, 50], [0, 2, 77], [11, 0, 13], [9, 10, 76], [2, 6, 67], [5, 6, 34], [9, 7, 62], [5, 3, 67]], src=10, dst=1, k=10))
