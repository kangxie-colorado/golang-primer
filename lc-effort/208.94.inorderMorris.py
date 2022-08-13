"""
https://leetcode.com/problems/binary-tree-inorder-traversal/

now I know 
- dfs
- stack-based
now there is this 
- morris inorder traversal
    curr = root
    while curr
        if curr.left == null:
            # reached the leftmost of this subtree..
            output curr
            # what if curr.right is none?
            # ahah.. when travelling down.. the parent is aleady on the right's hook
            curr = curr.right
        else
            # curr is stil curr
            pre=curr.left
            # fine the right most child and make curr its right child
            while pre.right and pre.right != curr:
                # goal is not to make pre to null
                # goal is to make pre.right to null
                # or end on the circle...
                pre = pre.right
            
            if pre.right == null:
                # create the cycle
                pre.right = curr
                # and move on the deal with left 
                curr = curr.left
            else:
                # so this is coming back to the cycle
                pre.right = nul
                output curr
                curr = curr.right
"""


from typing import List, Optional
from utils import TreeNode


class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        curr = root
        res = []
        while curr:
            if curr.left is None:
                res.append(curr.val)
                curr = curr.right
            else:
                prev = curr.left
                while prev.right and prev.right != curr:
                    prev = prev.right

                if prev.right is None:
                    prev.right = curr
                    curr = curr.left
                else:
                    prev.right = None
                    res.append(curr.val)
                    curr = curr.right

        return res


"""
Runtime: 52 ms, faster than 38.36% of Python3 online submissions for Binary Tree Inorder Traversal.
Memory Usage: 13.9 MB, less than 13.18% of Python3 online submissions for Binary Tree Inorder Traversal.

cool... let me use this to solve that validate BST
"""
