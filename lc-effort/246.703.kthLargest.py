"""
https://leetcode.com/problems/kth-largest-element-in-a-stream/

maybe two heaps?
a heap k elements.. the smallest on this heap at the top will be kth largest

when there is a bigger element to be added
pop the top and add it to right heap
add the bigger one to left heap

what if duplicates?
doesn't matter
"""


import heapq
from multiprocessing import heap
from typing import List


class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.k = k
        self.kHeap = nums
        heapq.heapify(self.kHeap)
        while len(self.kHeap) > k:
            heapq.heappop(self.kHeap)

    def add(self, val: int) -> int:
        if len(self.kHeap) < self.k:
            heapq.heappush(self.kHeap, val)
        elif val > self.kHeap[0]:
            heapq.heappop(self.kHeap)
            heapq.heappush(self.kHeap, val)

        return self.kHeap[0]


# Your KthLargest object will be instantiated and called as such:
# obj = KthLargest(k, nums)
# param_1 = obj.add(val)


""""
works but kind of clumsy, right?
see someone simplify this

just need to maintain the heap to be k size...
"""


class KthLargest:

    def __init__(self, k: int, nums: List[int]):
        self.k = k
        self.nums = nums
        heapq.heapify(self.nums)

    def add(self, val: int) -> int:
        heapq.heappush(self.nums, val)

        while len(self.nums) > self.k:
            heapq.heappop(self.nums)

        return self.nums[0]
