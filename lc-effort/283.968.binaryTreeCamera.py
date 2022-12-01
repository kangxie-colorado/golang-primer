"""
https://leetcode.com/problems/binary-tree-cameras/

I read Lee's post
genius...

https://leetcode.com/problems/binary-tree-cameras/discuss/211180/JavaC%2B%2BPython-Greedy-DFS
"""


from typing import Optional
from utils import TreeNode


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def minCameraCover(self, root: Optional[TreeNode]) -> int:
        res = 0

        def helper(root):
            nonlocal res
            if root is None:
                return 0, 0  # nodeCount, marked

            leftNode, leftMark = helper(root.left)
            rightNode, rightMark = helper(root.right)

            if leftNode == rightNode == 0:
                # leaf
                return 1, 0

            if leftNode == 1 or rightNode == 1:
                # leaf's parent
                res += 1
                return leftNode+rightNode+1, 1

            if leftMark or rightMark:
                # this is equivilent to removing this node
                return 0, 0

        # the root node needs a special treatement
        # if it is just the leaf's parent's parent, then no need to consider - which is marked
        #   the child level returns a mark
        # if is just the leaf's parent, then node will be "removed" so returned 0 no need
        # if it is not leaf parent, but remained a leaf... then + 1
        # e.g. O-O-O-O, the root node will be left alone...

        node, mark = helper(root)
        if node and not mark:
            res += 1
        return res
