"""
https://leetcode.com/problems/n-ary-tree-postorder-traversal/

I am thinking pre-order reverse
"""


from inspect import stack
import re
from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""


class Solution:
    def postorder(self, root: 'Node') -> List[int]:
        def preorder(root):
            stack = [root]
            res = []
            while len(stack):
                node = stack.pop()
                if node is None:
                    continue

                res.append(node.val)
                for i in range(len(node.children)-1, -1, -1):
                    stack.append(node.children[i])
            return res

        return preorder(root)[::-1]


"""
failed..
so only binary tree has that preorder <-> post-order

okay... even binary tree it is not reverse
preorder: 1 l:2 r:3, will be 1 2 3
postorder will be 2 3 1..

so what I am thinking?

was thinking this 

class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        stack = []
        curr = root
        res = []
        while curr or stack:
            # the write order is mid-right-left.. just the reverse
            while curr:
                res.append(curr.val)
                stack.append(curr)
                curr = curr.right

            node = stack.pop()
            curr = node.left

        return res[::-1]

so this is mid-right-left... it is reflected pre-order..
let me try that

just change 
                for i in range(len(node.children)-1, -1, -1):
                    stack.append(node.children[i])
                
                to

                for i in range(len(node.children)):
                    stack.append(node.children[i])
"""


"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""


class Solution:
    def postorder(self, root: 'Node') -> List[int]:
        def pre_mid_right_left(root):
            stack = [root]
            res = []
            while len(stack):
                node = stack.pop()
                if node is None:
                    continue

                res.append(node.val)
                for i in range(len(node.children)):
                    stack.append(node.children[i])
            return res

        return pre_mid_right_left(root)[::-1]


"""
and cool

Runtime: 66 ms, faster than 72.96% of Python3 online submissions for N-ary Tree Postorder Traversal.
Memory Usage: 16 MB, less than 98.02% of Python3 online submissions for N-ary Tree Postorder Traversal.

now let do another way..
I think to cut the link would be really hard here?????
because so many children here.. you need to keep track of every children... would be really hard

so go right... and look at left child one at a time
still feeling it different..

that is super hard...

so checked a few others' solutions, actually same to my above one..
"""
