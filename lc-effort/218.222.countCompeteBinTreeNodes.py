"""
https://leetcode.com/problems/count-complete-tree-nodes/

I am not thinking best solutions yet
just have some fun
"""

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


from tkinter.messagebox import RETRY
from tkinter.tix import Tree
from typing import Optional

from utils import TreeNode


class Solution:
    def countNodes(self, root: Optional[TreeNode]) -> int:
        def helper(root):
            curr = root
            rightDepth = 0
            while curr:
                rightDepth += 1
                curr = curr.right

            curr = root
            leftDepth = 0
            while curr:
                leftDepth += 1
                curr = curr.left

            if leftDepth == rightDepth:
                return 2**leftDepth-1
            else:
                return 1 + helper(root.left) + helper(root.right)
        return helper(root)


"""
Runtime: 121 ms, faster than 56.44% of Python3 online submissions for Count Complete Tree Nodes.
Memory Usage: 21.4 MB, less than 47.73% of Python3 online submissions for Count Complete Tree Nodes.
is this it?
Runtime: 74 ms, faster than 98.27% of Python3 online submissions for Count Complete Tree Nodes.
Memory Usage: 21.5 MB, less than 47.73% of Python3 online submissions for Count Complete Tree Nodes.

maybe... but how is this binary search?
oh I think I see the binary search 

if the guess is 7, then 1 1 1... walk right three times.. if the nodes is there.. 
then this is possible, but could be more so move l=m
if not there... then this is not possible... so move r=m-1

this is right converging to left.. ??

"""


class Solution:
    def countNodes(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0

        def isNodeThere(m):
            # decide the highest bit first
            n = 0
            while n < 16:  # 2**16 = 65536
                if 2**n <= m < 2**(n+1):
                    break
                n += 1

            if n == 0:
                return root is not None
            mask = 1 << (n-1)
            curr = root
            while curr and n > 0:
                curr = curr.right if m & mask else curr.left
                n -= 1
                mask >>= 1

            return curr is not None

        l, r = 0, 50000
        while l < r:
            m = r-(r-l)//2
            if isNodeThere(m):
                l = m
            else:
                r = m-1
        return l


"""
Runtime: 82 ms, faster than 93.62% of Python3 online submissions for Count Complete Tree Nodes.
Memory Usage: 21.5 MB, less than 47.73% of Python3 online submissions for Count Complete Tree Nodes.

that walk left/right can be improved because left is guranteed to be >= right depth

"""


class Solution:
    def countNodes(self, root: Optional[TreeNode]) -> int:
        def helper(root):

            left = root
            right = root
            depth = 0
            while right:
                right = right.right
                left = left.left
                depth += 1

            if left is None:
                return 2**depth - 1
            else:
                return 1 + helper(root.left) + helper(root.right)
        return helper(root)


if __name__ == '__main__':
    s = Solution()

    print(s.countNodes(None))

    root = TreeNode(1, TreeNode(2))
    print(s.countNodes(root))

    root = TreeNode(1,
                    TreeNode(2,
                             TreeNode(4)),
                    TreeNode(3))
    print(s.countNodes(root))

    root = TreeNode(1,
                    TreeNode(2,
                             TreeNode(4),
                             TreeNode(5)),
                    TreeNode(3,
                             TreeNode(6)))
    print(s.countNodes(root))
