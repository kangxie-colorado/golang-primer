"""
https://leetcode.com/problems/binary-tree-preorder-traversal/

hmm.. not that simple

need to use stack and append right before left
this way, it can do the right order..
"""

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


from typing import List, Optional

from utils import TreeNode


class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        s = [root]
        res = []
        while len(s):
            node = s[-1]
            s = s[:len(s)-1]
            if node is None:
                continue
            res.append(node.val)

            s.append(node.right)
            s.append(node.left)
        return res


"""
Runtime: 52 ms, faster than 38.95% of Python3 online submissions for Binary Tree Preorder Traversal.
Memory Usage: 13.9 MB, less than 60.05% of Python3 online submissions for Binary Tree Preorder Traversal.
"""


class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        stack = []
        curr = root
        res = []
        while curr or stack:
            while curr:
                res.append(curr.val)
                stack.append(curr)
                curr = curr.left

            node = stack.pop()
            curr = node.right

        return res


"""
Runtime: 33 ms, faster than 90.92% of Python3 online submissions for Binary Tree Preorder Traversal.
Memory Usage: 14 MB, less than 13.06% of Python3 online submissions for Binary Tree Preorder Traversal.
"""
