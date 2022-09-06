"""
https://leetcode.com/problems/stone-game-vi/

there are 10 such problems.. great serie
I don't know other version but this one I might see the solution already

alice: alice-bob.. maintain the max-heap
bob: bob-alice.. maintain the map-heap

also maintain a taken set.. if the ball is already taken then keep pop until not
let me try
"""


from curses.ascii import SO
import enum
import heapq
from typing import List


class Solution:
    def stoneGameVI(self, aliceValues: List[int], bobValues: List[int]) -> int:
        # from alice, should be aliceValue-bobValue but the maxHeap has the negative -
        # so reverse it
        aliceHeap = [(t[1]-t[0], i, t[0])
                     for i, t in enumerate(zip(aliceValues, bobValues))]
        bobHeap = [(t[0]-t[1], i, t[1])
                   for i, t in enumerate(zip(aliceValues, bobValues))]

        heapq.heapify(aliceHeap)
        heapq.heapify(bobHeap)

        heaps = [aliceHeap, bobHeap]
        turn = 0  # alice's turn
        currHeap = heaps[turn]
        taken = set()
        rawValues = [aliceValues, bobValues]
        values = [0, 0]

        def returnValue():
            if values[0] > values[1]:
                return 1
            if values[0] < values[1]:
                return -1
            return 0

        while currHeap:
            val, i, w = heapq.heappop(currHeap)
            while i in taken:
                if currHeap:
                    val, i, w = heapq.heappop(currHeap)
                else:
                    return returnValue()

            val = -val

            # need to also consider how many I lose by leaving other choices
            other = turn ^ 1
            otherVal = 0
            otherW = 0
            j = -1
            otherHeap = heaps[other]
            if otherHeap:
                otherVal, j, otherW = heapq.heappop(otherHeap)
                while j in taken:
                    if otherHeap:
                        otherVal, j, otherW = heapq.heappop(otherHeap)
                    else:
                        break
                otherVal = -otherVal

            if j != -1:
                if val > otherVal:
                    taken.add(i)
                    values[turn] += rawValues[turn][i]
                    heapq.heappush(otherHeap, (otherVal, j, otherW))
                else:
                    taken.add(j)
                    values[turn] += rawValues[turn][j]
                    heapq.heappush(currHeap, (val, i, w))
            else:
                taken.add(i)
                values[turn] += w

            turn ^= 1
            currHeap = heaps[turn]

        return returnValue()


"""
I complicated it
maybe I should maintain one heap with the abs(a-b).. max heap of course
because alice and bob both wants to max their advantage so always take the biggest difference one

and then keep the index
it does check out this one 
[6,5,1,2,10,6]
[7,7,7,7,3,7]
"""


class Solution:
    def stoneGameVI(self, aliceValues: List[int], bobValues: List[int]) -> int:
        absDiffs = [(-abs(t[0]-t[1]), i)
                    for i, t in enumerate(zip(aliceValues, bobValues))]
        absDiffs.sort()

        evenSum = sum(aliceValues[t[1]]
                      for i, t in enumerate(absDiffs) if i % 2 == 0)
        oddSum = sum(bobValues[t[1]]
                     for i, t in enumerate(absDiffs) if i % 2 == 1)

        if evenSum > oddSum:
            return 1
        if evenSum < oddSum:
            return -1

        return 0


""""
okay.. maybe I over-simply things...

such problem
you must analyze enough cases

hint: Greedily choose the stone with the maximum aliceValues[i] + bobValues[i].

huh.. 
yes.. it makes sense now

my gain, your pain
[6,5,1,2,10,6]
[7,7,7,7,3,7]

if I take 10, I gain 10, your pain is 3, gain+pain=12
if I take 6, I gain  6, your pain is 7.. gain+pain=13
ah.. so
[6,1]
[7,8]

I need to take 6, and leave 8 to you, the net loss is 2
if I take 1, leave 7 to you, the net loss is 6..

holy s...
"""


class Solution:
    def stoneGameVI(self, aliceValues: List[int], bobValues: List[int]) -> int:
        gainAndPain = [(-t[0]-t[1], i)
                       for i, t in enumerate(zip(aliceValues, bobValues))]
        gainAndPain.sort()

        evenSum = sum(aliceValues[t[1]]
                      for i, t in enumerate(gainAndPain) if i % 2 == 0)
        oddSum = sum(bobValues[t[1]]
                     for i, t in enumerate(gainAndPain) if i % 2 == 1)

        if evenSum > oddSum:
            return 1
        if evenSum < oddSum:
            return -1

        return 0


if __name__ == '__main__':
    s = Solution()
    A = [6, 1]
    B = [7, 8]
    assert -1 == s.stoneGameVI(A, B)

    A = [6, 5, 1, 2, 10, 6]
    B = [7, 7, 7, 7, 3, 7]
    assert -1 == s.stoneGameVI(A, B)

    assert 1 == s.stoneGameVI(aliceValues=[1, 3], bobValues=[2, 1])
    aliceValues = [1, 2]
    bobValues = [3, 1]
    assert 0 == s.stoneGameVI(aliceValues, bobValues)
    aliceValues = [2, 4, 3]
    bobValues = [1, 6, 7]
    assert -1 == s.stoneGameVI(aliceValues, bobValues)
