"""
https://leetcode.com/problems/delete-leaves-with-a-given-value/

post-order 
then pay attention to return None or root depending on if root should be deleted

kind of delicate but beautiful
"""

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


from typing import Optional

from utils import TreeNode


class Solution:
    def removeLeafNodes(self, root: Optional[TreeNode], target: int) -> Optional[TreeNode]:
        def helper(root):
            if root is None:
                return None

            root.left = helper(root.left)
            root.right = helper(root.right)

            if root.left is None and root.right is None and root.val == target:
                return None

            return root
        return helper(root)
