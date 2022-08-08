"""
https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

for a moment, no idea
then I think.. traverse and still hashmap

before that I was think traverse and convert to a list then hashmap
then I think why just traverse with a hashmap...
"""

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


from pyparsing import Optional

from utils import TreeNode


class Solution:
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        m = set()

        def helper(root):
            if root is None:
                return False
            if k-root.val in m:
                return True

            m.add(root.val)
            return helper(root.left) or helper(root.right)

        return helper(root)


"""
Runtime: 129 ms, faster than 52.73% of Python3 online submissions for Two Sum IV - Input is a BST.
Memory Usage: 18.3 MB, less than 46.88% of Python3 online submissions for Two Sum IV - Input is a BST.

passed... but I still produced some confusion.. I was thinking 
why looking for left or right, while the value could be one in left, one in right..

well... I beat myself.. I was not doing binary search.. I was doing full tree traverse and use hashmap to pinpoint
the target-val (the diff)... just like the list version

some be flexiable..
"""
