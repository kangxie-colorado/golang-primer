"""
https://leetcode.com/problems/add-two-numbers/

should be straightforward now
"""


from multiprocessing import dummy
from typing import Optional
from utils import ListNode


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode()
        tail = dummy
        carry = 0

        while l1 and l2:
            val = (l1.val+l2.val+carry) % 10
            carry = (l1.val+l2.val+carry) // 10
            tail.next = ListNode(val=val)
            tail = tail.next
            l1, l2 = l1.next, l2.next

        rest = l1 if l1 else l2
        while rest:
            val = (rest.val+carry) % 10
            carry = (rest.val+carry) // 10
            tail.next = ListNode(val=val)
            tail = tail.next
            rest = rest.next

        if carry:
            tail.next = ListNode(val=carry)

        return dummy.next
