"""
https://leetcode.com/problems/swap-nodes-in-pairs/

Sounds like we cannot change the next pointer..
then it sounds like we just need to swap the values..

too poiners naturally
"""
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
from tkinter.messagebox import NO
from typing import List, Optional
import ListNode


class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return head

        dummy = ListNode(-1, head)

        oddNode = head
        evenNode = oddNode.next

        while oddNode is not None and evenNode is not None:
            oddNode.val, evenNode.val = evenNode.val, oddNode.val
            evenNode = oddNode.next if oddNode is not None else None
            evenNode = oddNode.next

        return dummy.next


"""
Runtime: 52 ms, faster than 39.25% of Python3 online submissions for Swap Nodes in Pairs.
Memory Usage: 14 MB, less than 17.97% of Python3 online submissions for Swap Nodes in Pairs.
"""
