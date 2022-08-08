"""
https://leetcode.com/problems/partition-list/

hmm.. I think just using two queues?
smaller go to one; bigger or equal go to another... then..
or not really queue.. just remember the tail of each partition..

let me do both
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


from tkinter.messagebox import NO
from pyparsing import Optional
from utils import ListNode


class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        smQ = []
        bgQ = []

        while head is not None:
            next = head.next
            head.next = None
            if head.val < x:
                smQ.append(head)
            else:
                bgQ.append(head)

            head = next

        dummy = ListNode()
        tail = dummy
        for n in smQ+bgQ:
            tail.next = n
            tail = n

        return dummy.next


"""
Runtime: 47 ms, faster than 67.21% of Python3 online submissions for Partition List.
Memory Usage: 13.9 MB, less than 75.71% of Python3 online submissions for Partition List.

okay.. now no queue, just remember and maintain a bunch of pointers
"""


class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        dummyS = smTail = ListNode()
        dummyB = bgTail = ListNode()

        while head is not None:
            next = head.next
            head.next = None
            if head.val < x:
                smTail.next = head
                smTail = head
            else:
                bgTail.next = head
                bgTail = head

            head = next

        smTail.next = dummyB.next
        return dummyS.next


"""
Runtime: 38 ms, faster than 90.25% of Python3 online submissions for Partition List.
Memory Usage: 13.9 MB, less than 75.71% of Python3 online submissions for Partition List.
"""
