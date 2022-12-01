"""
https://leetcode.com/problems/car-fleet/


feel like it is sort then combine?
I can combine (startPos, speed) together for a sort then calculate the time to get to the end

then combine the same time... or set() to uniq it... 
let me see if this is right after all..

the simple examples checks out fine

"""


from time import time
from typing import List


class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        posAndSpeed = []
        for p, s in zip(position, speed):
            posAndSpeed.append((-p, s))

        posAndSpeed.sort()
        timeNeeded = []
        for p, s in posAndSpeed:
            timeNeeded.append((target+p-1)//s+1)

        i = 0
        res = 0
        while i < len(timeNeeded):
            res += 1
            while i+1 < len(timeNeeded) and timeNeeded[i] >= timeNeeded[i+1]:
                i += 1
            i += 1

        return res


""" ^ wrong
10
[0,4,2]
[2,1,3]

wrong here.. 
okay.. 
(4,1) (2,3) (0,1)
translate to 
(6,3,5)

okay.. the last car will catch up to the first one
so you have to eat up all the quicker cars


"""


class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        posAndSpeed = []
        for p, s in zip(position, speed):
            posAndSpeed.append((-p, s))

        posAndSpeed.sort()
        timeNeeded = []
        for p, s in posAndSpeed:
            timeNeeded.append((target+p-1)//s+1)

        i = 0
        res = 0
        while i < len(timeNeeded):
            res += 1
            j = i+1
            while j < len(timeNeeded) and timeNeeded[i] >= timeNeeded[j]:
                j += 1
            i = j

        return res


"""
10
[8,3,7,4,6,5]
[4,4,4,4,4,4]

output 2 
expect 6

okay.. wrong again.. 
okay.. same speed, the 2nd car cannot catch up the first
although they only take the same time... 

so maybe I should use fraction time?
"""


class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        posAndSpeed = []
        for p, s in zip(position, speed):
            posAndSpeed.append((-p, s))

        posAndSpeed.sort()
        timeNeeded = []
        for p, s in posAndSpeed:
            timeNeeded.append((target+p)/s)

        i = 0
        res = 0
        while i < len(timeNeeded):
            res += 1
            j = i+1
            while j < len(timeNeeded) and timeNeeded[i] >= timeNeeded[j]:
                j += 1
            i = j

        return res


"""
Runtime: 2469 ms, faster than 22.15% of Python3 online submissions for Car Fleet.
Memory Usage: 37 MB, less than 19.18% of Python3 online submissions for Car Fleet.

okay.. fraction time it passes
"""
