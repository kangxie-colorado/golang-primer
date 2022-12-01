# https://leetcode.com/problems/max-points-on-a-line/description/


"""
so okay
use the slope and a fix point to identy a line

three situations -- maybe 2 actually
1. deltaX != 0 
c1 = (deltaY/gcd, deltaX/gcd)
c2 can be calculated by y1=c1*x1+c2 => c2=y1-c1*x1

2. deltaX == 0
c1 = (inf, 0)
c2 = x1 or x2 (the same)

in the middle of coding, i found it still had to deal with map/set/map-of-set sort of things

so again noticet at most 90K lines
each line has a slope and a start point.. so why don't I just deal with 90K tuples...


"""


from collections import defaultdict
from dataclasses import dataclass
from email.policy import default
import math
from typing import List
from utils import gcd


class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        if len(points) == 1:
            return 1

        lines = defaultdict(int)
        for i in range(len(points)):
            for j in range(i+1, len(points)):
                deltaX = points[i][0] - points[j][0]
                deltaY = points[i][1] - points[j][1]

                if deltaX == 0:
                    lines[(1, 0, points[i][0], 0)] += 1
                elif deltaY == 0:
                    lines[(0, 1, 0, points[i][1])] += 1
                else:
                    sign = 1
                    if deltaX * deltaY < 0:
                        sign = -1
                    gcdDelta = gcd(abs(deltaX), abs(deltaY))
                    deltaX = abs(deltaX) // gcdDelta
                    deltaY = abs(deltaY) // gcdDelta
                    x, y = points[i]
                    #c2 = (y*deltaX - x*deltaY, deltaX)
                    lines[(sign*abs(deltaY), abs(deltaX),
                           y*deltaX - sign*x*deltaY, deltaX)] += 1

        maxLine = max(lines.values())
        # from number of same lines (L) to the number of points(P)
        # because L = P*(P-1)/2
        # so P = int(sqrt[2L]+1)
        return int(math.sqrt(maxLine*2)+1)


"""
Runtime 209 ms Beats 49.19%

"""


class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        n = len(points)
        res = 0

        for i in range(n-1):
            hashmap = defaultdict(int)
            x1, y1 = points[i]
            for j in range(i+1, n):
                x2, y2 = points[j]

                slope = float("+inf") if x2 - \
                    x1 == 0 else (y2-y1)*1.0 / (x2-x1)*1.0

                hashmap[slope] += 1

            res = max(res, max(hashmap.values()))

        return res + 1


"""
I was kind of confusing why this worked?
so it is to fix a point (x1,y1)
then pass this point, how many slope are there...

and notice hashmap is created every time for each fixed point..
ah.. this is indeed simple and good

fix a point, also removes the worry for parallel lines...
so yeah... not smart enough and blind to see this but I should feel good enough to come up with two solutions
"""


if __name__ == '__main__':
    s = Solution()
    print(s.maxPoints([[1, 1], [2, 2], [1, 2], [2, 3]]))
    print(s.maxPoints([[0, 0], [4, 5], [7, 8], [8, 9], [5, 6], [3, 4], [1, 1]]
                      ))
