"""
https://leetcode.com/problems/car-pooling/

feeling very werid 
so I just go thru each time...

if it is more than capacity then false
otherwise true
"""


from code import interact
from collections import defaultdict
from multiprocessing.resource_sharer import stop
from typing import List


class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        m = defaultdict(int)
        for p, s, d in trips:
            for i in range(s, d):
                m[i] += p
                if m[i] > capacity:
                    return False
        return True


"""
Runtime: 731 ms, faster than 5.01% of Python3 online submissions for Car Pooling.
Memory Usage: 14.4 MB, less than 87.05% of Python3 online submissions for Car Pooling.

I don't know what is the time complecity.. 
"""


class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        trips.sort(key=lambda x: (x[1], x[2]))

        p1, int1 = 0, trips[0][1:]
        for p, s, d in trips:
            interval = [s, d]
            if int1[1] > interval[0]:
                # that is overlap and for the overlap I need
                if p1+p > capacity:
                    return False
                # after the overlap
                int1 = [int1[1], interval[1]]
            else:
                p1, int1 = p, interval

        return True


"""
totally messed up... 
I treated this as an interval problem and it quickly messed up my thinking
reading 
https://leetcode.com/problems/car-pooling/discuss/317611/C%2B%2BJava-O(n)-Thousand-and-One-Stops
https://leetcode.com/problems/car-pooling/discuss/317610/JavaC%2B%2BPython-Meeting-Rooms-III

just at every start, -capacity by passenger
at every end, replenish capacity by passenger
see at any point.. if it is negative... this is after sorting
genius!

then the V guy focus on how much capacity is needed, 
at every start, + passenger needed for capacity
at every end, - passenger needed for capcity
what a genius way!!
even no need to sort...
"""


class Solution:
    def carPooling(self, trips, capacity):
        stops = [0]*1001
        for p, s, d in trips:
            stops[s] += p
            stops[d] -= p
        for s in stops:
            capacity -= s
            if capacity < 0:
                return False
        return True


"""
Runtime: 97 ms, faster than 71.42% of Python3 online submissions for Car Pooling.
Memory Usage: 14.3 MB, less than 87.05% of Python3 online submissions for Car Pooling.

this is focusing on how many capacity is needed for each stop and as a whole...
what a good thinking process...

also that sorting then process is pretty good too

this is Lee's code
    def carPooling(self, trips, capacity):
        for i, v in sorted(x for n, i, j in trips for x in [[i, n], [j, - n]]):
            capacity -= v
            if capacity < 0:
                return False
        return True

but maybe no need to sort actually
hmmm need to sort indeed!!!
[[2,1,5],[3,3,7]]
4

otherwise, for example above, it will process +2 at 1 then -2 at 5... it will always to zero..
you do need to process +2 at 1, +3 at 3..

what are the subtle difference between the two??


ah.. okayokay.. the 1001 stops array is an implicit sort
so no matter what, the stops must be processed sorted order
"""

if __name__ == '__main__':
    s = Solution()
    #assert s.carPooling([[3, 2, 7], [3, 7, 9], [8, 3, 9]], 11)
    assert not s.carPooling(
        trips=[[3, 3, 7], [2, 1, 5], [2, 1, 4]], capacity=4)
