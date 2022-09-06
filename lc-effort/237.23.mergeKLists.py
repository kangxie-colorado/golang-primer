"""
https://leetcode.com/problems/merge-k-sorted-lists/

I did this before using divide and conquer 
half then half.. 

eventally merge two lists

but I also see a heap solution.. let me see
"""


import heapq
from typing import List, Optional
from utils import ComparbleListNode, ListNode


class ComparbleListNode(ListNode):
    def __init__(self, val=0, next=None):
        super().__init__(val, next)

    def __lt__(self, other):
        return self.val < other.val


class Solution:

    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        dummy = ListNode(0)
        tail = dummy
        h = []  # (val, node)

        for l in lists:
            if l:
                l = ComparbleListNode(l.val, l.next)
                heapq.heappush(h, (l.val, l))

        while h:
            _, node = heapq.heappop(h)
            tail.next = node
            tail = tail.next
            nextNode = node.next
            if nextNode:
                nextNode = ComparbleListNode(nextNode.val, nextNode.next)
                heapq.heappush(h, (nextNode.val, nextNode))

        return dummy.next
