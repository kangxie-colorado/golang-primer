"""
https://leetcode.com/problems/binary-tree-postorder-traversal/

this is not trivial at all
also the morris..

yet to figure out
now it is a good time to go off the track a bit and learn something extra

the tree traversal
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
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        stack = []
        curr = root
        res = []

        while curr or stack:
            while curr:
                stack.append(curr)
                curr = curr.left

            if stack:  # this check is not necessary
                if not stack[-1].right:
                    node = stack.pop()
                    res.append(node.val)
                    # now I am popping up what is stored in stack
                    # not the time to process previously precossed
                    # keep the curr to None
                    curr = None
                else:
                    curr = stack[-1].right
                    # break cycle here..?
                    stack[-1].right = None

        return res


"""
Runtime: 40 ms, faster than 72.29% of Python3 online submissions for Binary Tree Postorder Traversal.
Memory Usage: 13.9 MB, less than 13.25% of Python3 online submissions for Binary Tree Postorder Traversal.

okay.. cool

so the key is to break the cycle..
"""


class Solution:
    def postorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        stack = []
        curr = root
        res = []

        while curr or stack:
            while curr:
                stack.append(curr)
                curr = curr.left

            if not stack[-1].right:
                node = stack.pop()
                res.append(node.val)
                curr = None
            else:
                curr = stack[-1].right
                # break cycle here..?
                stack[-1].right = None

        return res


"""
except to serve the cycle link
post-order can be achieved by walking to right then left first.. 
but the write order is also reversed (or reverse at the end)
"""


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


"""
Runtime: 39 ms, faster than 75.03% of Python3 online submissions for Binary Tree Postorder Traversal.
Memory Usage: 13.9 MB, less than 60.12% of Python3 online submissions for Binary Tree Postorder Traversal.

of course I can pre-pend.. but no big diff
"""

if __name__ == '__main__':
    s = Solution()

    print(s.postorderTraversal(TreeNode(1, None, TreeNode(2, TreeNode(3)))))
    print(s.postorderTraversal(TreeNode(1, TreeNode(2), TreeNode(3))))
