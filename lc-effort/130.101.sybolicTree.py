# https://leetcode.com/problems/symmetric-tree/

# do some easy to warm up..
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

from typing import Optional

from utils import TreeNode

"""
okay.. it is mirroring effect
I went for equal left/right.. and knocked down on first example...


"""


class Solution:
    def _recursive_isSymmetric(self, root: Optional[TreeNode]) -> bool:
        def symmetric_tree(root1, root2) -> bool:
            if root1 is None and root2 is None:
                return True
            if root1 is None or root2 is None:
                return False
            return root1.val == root2.val \
                   and symmetric_tree(root1.left, root2.right) \
                   and symmetric_tree(root1.right, root2.left)

        return symmetric_tree(root.left, root.right)

        '''
        Runtime: 34 ms, faster than 94.30% of Python3 online submissions for Symmetric Tree.
        Memory Usage: 14 MB, less than 60.42% of Python3 online submissions for Symmetric Tree.
        '''

    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        level = [root.left, root.right]
        while len(level):
            nextLevFirstHalf = []
            nextLevSecHalf = []
            l, r = 0, len(level) - 1
            while l < r:
                if level[l] is None and level[r] is None:
                    l += 1
                    r -= 1
                    continue
                if level[l] is None or level[r] is None or level[l].val != level[r].val:
                    return False

                nextLevFirstHalf.extend([level[l].left, level[l].right])
                nextLevSecHalf.extend([level[r].right, level[r].left])
                l += 1
                r -= 1

            if len(nextLevFirstHalf) != len(nextLevSecHalf):
                return False

            level = nextLevFirstHalf
            level.extend(nextLevSecHalf[::-1])

        return True
