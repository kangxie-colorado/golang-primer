"""
https://leetcode.com/problems/balance-a-binary-search-tree/

this should be a classical problem
but I kind of stuck which way to go

top-down or bottom-up
top-down.. faces duplicated calculation

bottom-up.. looks like the subtree can be changed by parent level re-balances..

seems like I have to top-down first..
or re-construct the tree

let me try a few

"""
from turtle import left
from utils import TreeNode


class Solution:
    def balanceBST(self, root: TreeNode) -> TreeNode:
        def height(root):
            if root is None:
                return 0

            return max(height(left), height(root))+1

        def helper(root):
            leftH = height(root.left)
            rightH = height(root.right)

            if leftH - rightH >= 2:
                ...
                # hit a nail...
                # diff 2: look for first left (then right)
                # diff 3: look for 2nd left(then its right...)
                # this complicated things..
                # so yeah.. top-down not going to work


"""
then thinking bottom-up.. but still it has that problem
so looking at the hints

Convert the tree to a sorted array using an in-order traversal.
Construct a new balanced tree from the sorted array recursively.

haha.. so maybe this is the way to go
"""


class Solution:
    def balanceBST(self, root: TreeNode) -> TreeNode:
        inOrder = []

        def inorder(root):
            if root is None:
                return

            inorder(root.left)
            inOrder.append(root.val)
            inorder(root.right)

        inorder()

        def construct(A):
            if len(A) == 0:
                return None
            mid = len(A)//2
            root = TreeNode(A[mid],
                            construct(A[:mid]),
                            construct(A[mid+1:])
                            )
            return root

        return construct(inOrder)


"""
Runtime: 880 ms, faster than 34.45% of Python3 online submissions for Balance a Binary Search Tree.
Memory Usage: 21.7 MB, less than 20.14% of Python3 online submissions for Balance a Binary Search Tree.

so this pass.. taste this a bit longer
what happened?
            mid = len(A)//2
            root = TreeNode(A[mid],
                            construct(A[:mid]),
                            construct(A[mid+1:])
                            )
so basically this is divide the array by 2, by 2..
and simple turn every node into a node.. the order in the array dictates the order in the tree natually..

this should be O(N)

after doing this.. that construct BST from pre-order becomes clearer to me..
the only variation is it needs a bound

every element is also turned into a node... the order in the array and condition > < bound decides its 
position in the tree...

though not 100% but let me re-do that

"""
