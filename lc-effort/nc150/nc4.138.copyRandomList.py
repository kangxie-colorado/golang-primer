"""
https://leetcode.com/problems/copy-list-with-random-pointer/

didn't seem I did this before 
so okay.. after doing this I am also done for the day


so this is just a two pass 
1. one pass to create the nodes, and establish the mapping table
2. connect the random pointers..
"""


from typing import Optional


class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random


class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        dummy = Node(-1, head)
        tail = dummy
        m = {}  # re-directing mapping

        while tail.next:
            newNode = Node(tail.next.val, tail.next.next, tail.next.random)
            m[tail.next] = newNode

            tail.next = newNode
            tail = tail.next

        tail = dummy.next
        while tail:
            if tail.random:
                tail.random = m[tail.random]
            tail = tail.next
        return dummy.next


"""
Runtime: 45 ms, faster than 77.89% of Python3 online submissions for Copy List with Random Pointer.
Memory Usage: 15 MB, less than 18.84% of Python3 online submissions for Copy List with Random Pointer.
"""
