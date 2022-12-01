"""
https://leetcode.com/problems/max-points-on-a-line/

because 1 <= points.length <= 300
at most 300*300 = 90000

then I think I can do all the lines 
and then go thru the point... see which line it falls on and count
but that is complicated

then I think I can still do the lines... build on the go
then use the lines to go thru points... 

if it go thru the point then this is on the line
this way, 90000 lines * 300 points.. 27,000,000  < 1,000,000,000, might pass for brute force
"""


from collections import defaultdict
from typing import List


class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        res = 1

        for i in range(len(points)):
            for j in range(i+1, len(points)):
                p1, p2 = points[i], points[j]
                # p1,p2 is a line
                lineRes = 2
                x1, y1 = p1
                x2, y2 = p2
                for x, y in points:
                    if [x, y] != [x1, y1] and [x, y] != [x2, y2] and (y-y1)*(x-x2) == (y-y2) * (x-x1):
                        lineRes += 1

                res = max(res, lineRes)
        return res


"""
32 / 35 test cases passed.

not bad first try and failed 
[[0,0]]

maybe the default res is 1, at least one point then res at least 1
last case TLE 34 / 35 test cases passed.
[[-27,-31],[-81,44],[-101,-138],[-120,-46],[-143,-71],[109,25],[-66,93],[-106,143],[75,-148],[-136,-25],[123,118],[37,-133],[78,22],[16,121],[-16,96],[-137,-116],[118,68],[12,10],[122,24],[-114,33],[-111,-9],[88,82],[-116,-72],[114,-98],[-103,33],[59,9],[-144,2],[105,18],[-111,-40],[55,60],[-20,-3],[108,-30],[71,78],[32,-1],[-10,-29],[59,33],[-130,3],[86,-22],[-138,-11],[-145,-70],[58,-26],[-23,42],[64,-53],[141,47],[12,111],[60,23],[109,-82],[-43,-117],[97,126],[-83,74],[-136,91],[53,-71],[-56,-2],[-61,106],[128,-29],[128,-39],[70,87],[-92,-58],[74,117],[148,-44],[116,27],[89,48],[66,8],[-85,-58],[-27,-96],[118,-16],[-115,-141],[-150,-58],[81,106],[138,-36],[-149,-65],[3,10],[-133,9],[6,125],[105,-139],[117,90],[104,108],[144,69],[-95,-129],[-113,109],[-130,-85],[-29,69],[3,30],[-77,-130],[-69,9],[124,119],[34,-29],[41,-100],[-49,-32],[-106,-64],[-136,-124],[-124,-29],[86,141],[80,-34],[-24,-127],[10,-52],[-43,-16],[-90,149],[42,18],[79,-123],[-2,-92],[-2,88],[67,-111],[113,-11],[56,68],[71,87],[-49,-61],[120,-125],[-57,19],[20,120],[-38,148],[131,-148],[101,-23],[-126,-149],[-78,-88],[144,-117],[83,128],[117,102],[23,-29],[123,-62],[53,86],[-102,7],[-91,139],[-65,-126],[127,82],[-89,-83],[142,137],[-122,83],[82,-115],[46,-82],[-73,67],[-78,110],[-62,131],[-85,-15],[-117,-50],[-85,-138],[-24,-114],[93,-146],[-28,71],[-17,124],[-31,36],[-101,41],[-52,40],[80,-114],[-100,29],[-108,-137],[-73,-67],[100,-71],[-100,84],[3,106],[10,-47],[-4,135],[-107,35],[-12,136],[69,35],[106,114],[-84,-98],[143,28],[-43,97],[-17,108],[5,16],[134,-106],[139,-135],[24,-31],[134,20],[-105,61],[46,-31],[-67,-21],[87,56],[-88,-10],[120,-108],[-125,59],[73,133],[-62,-93],[-31,-28],[16,-100],[115,-79],[15,87],[-44,63],[20,122],[145,54],[-89,22],[-39,88],[-35,95],[-108,-138],[124,-22],[-13,-72],[-141,66],[-22,-137],[-109,93],[-76,-118],[43,-79],[61,-79],[-97,-61],[102,96],[-6,-97],[-82,-1],[-86,109],[124,149],[-1,-108],[135,41],[77,-62],[122,-22],[145,-74],[77,33],[97,-16],[-94,15],[-124,-86],[111,58],[88,-46],[-136,63],[-138,124],[111,39],[117,37],[-74,-127],[15,-53],[35,-82],[-68,148],[65,-48],[-127,-4],[-56,139],[-129,-107],[-17,-98],[72,106],[30,-54],[117,24],[43,-74],[-127,8],[-19,-108],[-81,88],[-76,-94],[-52,-30],[107,-139],[4,-51],[43,127],[29,-3],[-16,-143],[78,-110],[34,-26],[-115,-40],[17,111],[20,134],[85,-113],[62,120],[70,-63],[25,-47],[142,35],[92,-101],[119,-122],[57,-19],[117,133],[40,102],[-107,1],[123,-112],[13,73],[6,-140],[72,-20],[-23,-39],[35,138],[-120,140],[-92,-70],[-114,97],[86,121],[76,92],[63,49],[123,76],[30,112],[41,-52],[4,-58],[-64,-78],[-97,149],[61,148],[51,-108],[70,-95],[-35,149],[-28,57],[-108,-20],[-79,95],[20,9],[-143,-134],[71,-56],[-13,-135],[-98,-28],[-64,-74],[-30,74],[-88,-113],[104,12],[-4,89],[-62,71],[-14,105],[26,39],[-75,96],[-20,22],[-33,102],[55,50],[-110,-86],[43,-110],[-117,143],[129,-93],[-67,108]]

where to optimize?
hmm... 

when I look at the len(A) = 300
I am thinking, might be I can divide and conqure 

and combine...
"""


"""
class Solution:
    def maxPoints(self, P: List[List[int]]) -> int:
        res = 1

        def helper(points):
            m = {}
            for i in range(len(points)):
                for j in range(i+1, len(points)):
                    p1, p2 = points[i], points[j]
                    # p1,p2 is a line
                    lineRes = 2
                    x1, y1 = p1
                    x2, y2 = p2
                    for x, y in points:
                        if [x, y] != [x1, y1] and [x, y] != [x2, y2] and (y-y1)*(x-x2) == (y-y2) * (x-x1):
                            lineRes += 1

                    detlaX, deltaY = x2-x1, y2-y1
                    m[(detlaX, deltaY)] = lineRes
            return m

        def divideAndConq(points):
            if len(points) <= 4:
                return helper(points)
            else:
                lMap = helper(points[:len(points)//2])
                rMap = helper(points[len(points)//2:])

                m = defaultdict(int)
                rUsedKey = set()
                for lKey, lVal in lMap.items():
                    m[lKey] = lVal
                    for rKey, rVal in rMap.items():
                        lDeltaX, lDeltaY = lKey
                        rDeltaX, rDeltaY = rKey
                        # same line?
                        if lDeltaX and lDeltaY and (rDeltaX*lDeltaY == rDeltaY*lDeltaX):
                            # non-zero
                            m[lKey] += rVal
                            rUsedKey.add(rKey)
                        elif lDeltaX == rDeltaX == 0 or lDeltaY == rDeltaY == 0:
                            m[lKey] += rVal
                            rUsedKey.add(rKey)
                ...

        return res

no this is not gonna work
maybe I think I can keep all the lines in a set
using the deltaX, deltaY
for example: (1,1) (3,3).. detlas=2,2
then try to get the gcd of 2,2 which is 2

if delta=2,4 then the minimum increament will be 1,2?
y=2*x
y=2*x+1, is the same delta
so thesis is not fully correct

but basically any line could be expressed with y=a*x+b
coefficients (a,b) will identify the line

y1 = a*x1+b
y2 = a*x2+b
a*(x2-x1)=(y2-y1)
a = rational(y2-y1, x2-x1)
b = rational(y1*x2-y2*x1, x2-x1)

maybe this can id a line?
and specialize X-axis and Y-axis?

nah... 
actually I think my solution of 90000*300 should just work...
"""


if __name__ == '__main__':
    ...
