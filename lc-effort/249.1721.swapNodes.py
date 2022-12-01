"""
https://leetcode.com/problems/swapping-nodes-in-a-linked-list/

careful state tracking..
ah.. just need to swap values.. 
no need to swap nodes at all
"""


from multiprocessing import dummy
from typing import Optional
from utils import ListNode


class Solution:
    def swapNodes(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        kthNode = head
        while k > 1 and kthNode:
            kthNode = kthNode.next
            k -= 1

        kthNodeRev = head
        tail = kthNode
        while tail.next:
            kthNodeRev, tail = kthNodeRev.next, tail.next

        kthNode.val, kthNodeRev.val = kthNodeRev.val, kthNode.val
        return dummy.next


"""
Runtime: 1016 ms, faster than 97.11% of Python3 online submissions for Swapping Nodes in a Linked List.
Memory Usage: 48.4 MB, less than 50.89% of Python3 online submissions for Swapping Nodes in a Linked List.
"""
