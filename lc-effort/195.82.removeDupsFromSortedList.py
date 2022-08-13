"""
https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/


I was thinking look backward.. and found it kind of uneasy and unnatural
then I think I can look forward to get rid of the dups..

yEAH... list is natrually built to look forward.. not backword..
"""


from typing import List, Optional
from utils import ListNode


class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(-1, head)
        tail = dummy

        while head:
            if head.next and head.next.val == head.val:
                # duplicates
                dupTail = head
                while dupTail and dupTail.next and dupTail.val == dupTail.next.val:
                    dupTail = dupTail.next
                tail.next = dupTail.next
                head = dupTail.next
            else:
                tail.next = head
                tail = head
                head = head.next

        return dummy.next


"""
wonderful... passed on first time

Runtime: 49 ms, faster than 81.92% of Python3 online submissions for Remove Duplicates from Sorted List II.
Memory Usage: 13.9 MB, less than 74.26% of Python3 online submissions for Remove Duplicates from Sorted List II.
"""
