"""
https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/

reading lee's code
https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/discuss/252232/JavaC%2B%2BPython-O(N)-Solution

    def bstFromPreorder(self, A):
        if not A: return None
        root = TreeNode(A[0])
        i = bisect.bisect(A, A[0])
        root.left = self.bstFromPreorder(A[1:i])
        root.right = self.bstFromPreorder(A[i:])
        return root

noticet the bisect here..
i = bisect.bisect(A, A[0])

1. the search space is A, not A[1:]; because if you remove one element, the index actually changed
but I think I can do +1 to offset that

2. it is bisect.bisect, not bisect.bisect_left
bisect() will find the last element that is <=, i.e the first element that is strictly >
and return the > index

if I use bisect_left here.. it will TLE
but I think I can compensate for 1 and use A[1:], might work

yes, this works
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def bstFromPreorder(self, preorder: List[int]) -> Optional[TreeNode]:
        def helper(l):
            if len(l) == 0:
                return None
            root = TreeNode(l[0])
            
            right = bisect.bisect_left(l[1:], root.val) + 1
            # right = bisect.bisect(l[1:], root.val) + 1, works too, because no duplicate
            
            root.left = helper(l[1:right])
            root.right = helper(l[right:])
            
            return root
        return helper(preorder)
            
his other solution is liking walking on the rope
I can feel it but cannot clear see thru

    i = 0
    def bstFromPreorder(self, A, bound=float('inf')):
        if self.i == len(A) or A[self.i] > bound:
            return None
        root = TreeNode(A[self.i])
        self.i += 1
        root.left = self.bstFromPreorder(A, root.val)
        root.right = self.bstFromPreorder(A, bound)
        return root


"""


import sys
from typing import List, Optional

from utils import TreeNode


class Solution:
    def bstFromPreorder(self, preorder: List[int]) -> Optional[TreeNode]:
        i = 0

        def helper(divide):
            if len(preorder) == i or preorder[i] >= divide:
                return None

            root = TreeNode(preorder[i])
            i += 1

            root.left = helper(root.val)
            root.right = helper(divide)
            return root

        return helper(sys.maxsize)


"""
Runtime: 86 ms, faster than 5.64% of Python3 online submissions for Construct Binary Search Tree from Preorder Traversal.
Memory Usage: 14 MB, less than 52.84% of Python3 online submissions for Construct Binary Search Tree from Preorder Traversal

so what happened.. 
like in 179.1382.balanceBST.py, in which I construct BST from inorder

this turns element from Array into a node.. 
root = TreeNode(preorder[i])

because the structure of pre-order.. the first element of a whole tree traverse is always the root
then it divides its following into two parts: left, and right
left is smaller than
right is greater than.

if smallers than, turn it into left tree.. it stops on first right tree node (don't i+=1 for this)
then it stamps on a valid left tree... by passing 
            if len(preorder) == i or preorder[i] >= bound:
                return None

then turn it into a node
            root = TreeNode(preorder[i])
            i += 1
i+=1 to consume this node so it move to next element..

after it is done with left tree
            root.left = helper(root.val)

it will stop until it hits first right tree node.. (return None)
it moves onto the right tree
            root.right = helper(bound)

of course this is a recursive structure 
            so sub-tree kind of take care itself...
maybe I'll name the bound theDivide, which shows the intention more clearly

the right...right..right-most branch needs no bound(divide)..

not this inspire me to do that validateBST another way
inorder it... then if it is ever increasing.. then it is fine.. if it brokes the order
then over..
"""
