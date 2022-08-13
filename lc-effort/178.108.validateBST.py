"""
https://leetcode.com/problems/validate-binary-search-tree/

let me try usig pre-order in-order to validate 
"""


from bisect import bisect
import math
import sys
from typing import Optional

from utils import TreeNode


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        preOrder = []

        def preorder(root):
            if root is None:
                return

            preOrder.append(root.val)
            preorder(root.left)
            preorder(root.right)

        inOrder = []

        def inorder(root):
            if root is None:
                return

            inorder(root.left)
            inOrder.append(root.val)
            inorder(root.right)

        preorder(root)
        inorder(root)

        def validate(p, i):
            if len(p) <= 1:
                return True

            rootVal = p[0]
            rightStart = bisect(i, p[0])
            leftVal = rightVal = 0

            if rightStart == len(p):
                # right empty
                leftVal = i[-2]
                rightVal = sys.maxsize
            elif rightStart == 1:
                # left empty
                leftVal = -sys.maxsize
                rightVal = i[1]
            else:
                leftVal = i[rightStart-2]
                rightVal = i[rightStart]

            return leftVal < rootVal < rightVal and \
                validate(p[1:rightStart], i[:rightStart-1]) and \
                validate(p[rightStart:], i[rightStart:])

        return validate(preOrder, inOrder)


"""
Runtime: 102 ms, faster than 6.98% of Python3 online submissions for Validate Binary Search Tree.
Memory Usage: 17.6 MB, less than 14.71% of Python3 online submissions for Validate Binary Search Tree.

cool.. as tedious as it is but it passes

"""


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        A = []

        def inorder(root):
            if root is None:
                return

            inorder(root.left)
            A.append(root.val)
            inorder(root.right)
        inorder(root)

        last = A[0]
        for i in range(1, len(A)):
            if A[i] <= last:
                return False
            last = A[i]

        return True


"""
Runtime: 83 ms, faster than 27.44% of Python3 online submissions for Validate Binary Search Tree.
Memory Usage: 17 MB, less than 25.98% of Python3 online submissions for Validate Binary Search Tree.
"""


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        prevNum = None

        curr = root
        while curr:
            if curr.left is None:
                if prevNum is not None and curr.val <= prevNum:
                    return False
                prevNum = curr.val
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
                    if prevNum is not None and curr.val <= prevNum:
                        return False
                    prevNum = curr.val
                    curr = curr.right
        return True


"""
Runtime: 51 ms, faster than 85.88% of Python3 online submissions for Validate Binary Search Tree.
Memory Usage: 16.2 MB, less than 99.49% of Python3 online submissions for Validate Binary Search Tree.

cool cool
"""

if __name__ == '__main__':
    s = Solution()
    root = TreeNode(2,
                    TreeNode(1),
                    TreeNode(3),
                    )
    print(s.isValidBST(root))
    root = TreeNode(2,
                    TreeNode(1),
                    TreeNode(1),
                    )
    print(s.isValidBST(root))
