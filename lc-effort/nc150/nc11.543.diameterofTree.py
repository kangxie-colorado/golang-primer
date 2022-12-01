"""
https://leetcode.com/problems/diameter-of-binary-tree/


"""


from typing import Optional
from utils import TreeNode


class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        res = 0

        def helper(root):
            nonlocal res
            if root is None:
                return 0

            left = helper(root.left)
            right = helper(root.right)

            res = max(res, left + right)

            return max(left, right)+1

        helper(root)
        return res
