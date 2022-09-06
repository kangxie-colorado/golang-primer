"""
https://leetcode.com/problems/sliding-window-maximum/

python will TLE???

"""


from collections import deque
import heapq
from multiprocessing import heap
from tracemalloc import start
from sortedcontainers import SortedList
from typing import Deque, List


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        def search(dq, num):
            l, r = 0, len(dq)
            while l < r:
                m = l+(r-l)//2
                if dq[m][1] > num:
                    l = m+1
                else:
                    r = m
            return l

        dq = []  # (index, number)
        res = []

        for i in range(len(nums)):
            if i >= k:
                if dq[0][0] == i-k:
                    dq = dq[1:]

            idx = search(dq, nums[i])
            dq = dq[:idx]

            dq.append((i, nums[i]))
            if i >= k-1:
                res.append(dq[0][1])

        return res


"""
oh I think when dq[-1][1] == nums[i]
I can kick the previous one out.. since they are in the same window and the later one will outlive the first one
hmm.. it will back fire and fails at first case!!!

so what about the TLE...
maybe binary search


did that and also get rid of equal precedsors, 
bearly passes in python.. should there be something sooner in python?
"""


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        def search(dq, num):
            l, r = 0, len(dq)
            while l < r:
                m = l+(r-l)//2
                if dq[m][1] > num:
                    l = m+1
                else:
                    r = m
            return l

        dq = []  # (index, number)
        res = []

        for i in range(len(nums)):
            if i >= k:
                if dq[0][0] == i-k:
                    dq = dq[1:]

            idx = search(dq, nums[i])
            dq = dq[:idx]

            dq.append((i, nums[i]))
            if i >= k-1:
                res.append(dq[0][1])

        return res


"""
this barely passes..
maybe that SortedList?
"""


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:

        sl = SortedList()
        res = []

        for i in range(len(nums)):
            if i >= k:
                sl.remove(nums[i-k])

            sl.add(nums[i])
            if i >= k-1:
                res.append(sl[-1])

        return res


"""
Runtime: 5343 ms, faster than 5.01% of Python3 online submissions for Sliding Window Maximum.
Memory Usage: 30.9 MB, less than 11.53% of Python3 online submissions for Sliding Window Maximum.

hmm.. so this one is better.

maybe I can use a maxHeap but tracking the idx.. if it is out of window, keep popping..
but I need to put it back for a valid one
"""


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:

        h = []  # (-val, index)
        res = []

        for i in range(len(nums)):
            heapq.heappush(h, (-nums[i], i))

            if i >= k-1:
                val, idx = heapq.heappop(h)
                while idx <= i-k:
                    val, idx = heapq.heappop(h)
                heapq.heappush(h, (val, idx))
                res.append(-val)
        return res


"""
Runtime: 3580 ms, faster than 16.02% of Python3 online submissions for Sliding Window Maximum.
Memory Usage: 38.7 MB, less than 6.42% of Python3 online submissions for Sliding Window Maximum.

hmm.. interesting..

let me try using a python dequeue instead of the plain list
"""


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        dq = deque()  # (index, number)
        res = []

        for i in range(len(nums)):
            if i >= k:
                if dq[0][0] == i-k:
                    dq.popleft()

            while len(dq) and dq[-1][1] <= nums[i]:
                dq.pop()

            dq.append((i, nums[i]))
            if i >= k-1:
                res.append(dq[0][1])

        return res


"""
Runtime: 2469 ms, faster than 68.52% of Python3 online submissions for Sliding Window Maximum.
Memory Usage: 30 MB, less than 66.50% of Python3 online submissions for Sliding Window Maximum.

well.. well.

Runtime: 1881 ms, faster than 88.79% of Python3 online submissions for Sliding Window Maximum.
Memory Usage: 29.9 MB, less than 80.34% of Python3 online submissions for Sliding Window Maximum.
"""


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        dq = []  # (index, number)
        res = []
        readIdx = 0
        for i in range(len(nums)):
            if i >= k:
                if dq[readIdx][0] == i-k:
                    readIdx += 1

            while len(dq) > readIdx and dq[-1][1] < nums[i]:
                dq.pop()

            dq.append((i, nums[i]))
            if i >= k-1:
                res.append(dq[readIdx][1])

        return res


"""
Runtime: 2430 ms, faster than 70.27% of Python3 online submissions for Sliding Window Maximum.
Memory Usage: 34.5 MB, less than 9.61% of Python3 online submissions for Sliding Window Maximum.
"""


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        dq = []  # (index, number)
        res = []
        readIdx = 0
        for i in range(len(nums)):
            if i >= k:
                if dq[readIdx][0] == i-k:
                    readIdx += 1

            l, r = readIdx, len(dq)
            while l < r:
                m = l+(r-l)//2
                if dq[m][1] >= nums[i]:
                    l = m+1
                else:
                    r = m

            dq = dq[:l]

            dq.append((i, nums[i]))
            if i >= k-1:
                res.append(dq[readIdx][1])

        return res


if __name__ == '__main__':
    s = Solution()
    print(s.maxSlidingWindow([1, 3, -1, -3, 5, 3, 6, 7], 3))
    print(s.maxSlidingWindow([9, 10, 9, -7, -4, -8, 2, -6], 5))
