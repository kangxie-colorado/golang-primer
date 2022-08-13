"""
https://leetcode.com/problems/path-sum-ii/


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
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> List[List[int]]:
        res = []

        def helper(root, run, S):
            if root is None:
                return

            if root.left is None and root.right is None and S+root.val == targetSum:
                res.append(run+[root.val])
                return

            helper(root.left, run+[root.val], S+root.val)
            helper(root.right, run+[root.val], S+root.val)

        helper(root, [], 0)
        return res


"""
Runtime: 51 ms, faster than 85.61% of Python3 online submissions for Path Sum II.
Memory Usage: 19.1 MB, less than 29.36% of Python3 online submissions for Path Sum II.
"""
if __name__ == '__main__':
    s = Solution()
    tree = TreeNode(1,
                    TreeNode(2),
                    TreeNode(3)
                    )

    print(s.pathSum(tree, 3))
