"""
https://leetcode.com/problems/reverse-nodes-in-k-group/

so okay I know I did this before 
idea is not super complicated

a prev(dummy) will be used to point to the head
a tail will be maintained and initialized to None

then move kAdv by k steps from head
reverse the list between prev.next and kAdv

when prev.next catches up with kAdv, move kAdv by k further steps
if able to move kAdv by k steps successfully then continue to reverse 

maybe to deal with just multiple times of k length, need a sentinel, huh... 
yeah the tricky part is how to deal with the last group

if I move k times, and kAdv just becomes None, then it is a full sub-grp as well
if kAdv beomces None in less than k times, then it is a incomplete sub-grp

incomplete subgrp, the old head will be connecting to it


okay... I made a mistake
it is not a full reverse..

the last k-grp will reverse but its tail will connect the current k-grp

okay.. I did walk the wrong way, spending 40 minutes..
gosh... future is so doom..

now I see a recursive structure 
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


from typing import Optional
from utils import ListNode, print_list


class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummy = ListNode(-1, head)

        def reverseK(prev):
            kAdv = prev.next
            i = 0
            while kAdv and i < k:
                i += 1
                kAdv = kAdv.next

            if i < k:
                return prev.next

            newHead = prev.next
            newTail = prev.next
            tail = None

            while prev.next != kAdv:
                newHead = prev.next
                prev.next = prev.next.next
                newHead.next = tail
                tail = newHead

            newTail.next = reverseK(prev)
            return newHead

        return reverseK(dummy)


"""
Runtime: 55 ms, faster than 88.82% of Python3 online submissions for Reverse Nodes in k-Group.
Memory Usage: 15.3 MB, less than 40.03% of Python3 online submissions for Reverse Nodes in k-Group.
"""

if __name__ == "__main__":
    s = Solution()
    n1 = ListNode(1, None)
    n2 = ListNode(2, None)
    n3 = ListNode(3, None)
    n1.next = n2
    n2.next = n3
    print_list(n1)
    print_list(s.reverseKGroup(n1, 2))
