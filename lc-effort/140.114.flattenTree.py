"""
https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

so it feels like just for every tree, return its head and tail
head is used for hook up parent node.. tail is used for hook up the right tree is there is any
"""


# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution1(object):
    def flatten(self, root):
        """
        :type root: TreeNode
        :rtype: None Do not return anything, modify root in-place instead.
        """
        def helper(root):
            """
            return: subtree's flatten head/tail pair
            """
            if root is None:
                return None, None

            leftHead, leftTail = helper(root.left)
            rightHead, rightTail = helper(root.right)

            root.right = leftHead if leftHead is not None else rightHead
            if leftTail is not None:
                leftTail.right = rightHead

            tail = rightTail or leftTail or root

            root.left = None
            return root, tail

        helper(root)


"""
gosh it passes

Runtime: 59 ms, faster than 5.25% of Python online submissions for Flatten Binary Tree to Linked List.
Memory Usage: 14.1 MB, less than 88.92% of Python online submissions for Flatten Binary Tree to Linked List.

it is very tedious to make it all right
but being super careful and don't miss any link and care the order of hooks up

it is not super hard..
"""


class Solution(object):
    def flatten(self, root):
        if root is None:
            return

        left = root.left
        right = root.right

        self.flatten(left)
        self.flatten(right)

        root.left = None
        root.right = left

        # now walk to the very end
        cur = root
        while cur.right is not None:
            cur = cur.right

        cur.right = right


"""
Runtime: 40 ms, faster than 38.78% of Python online submissions for Flatten Binary Tree to Linked List.
Memory Usage: 14.5 MB, less than 8.60% of Python online submissions for Flatten Binary Tree to Linked List.

this seems faster? why
"""
