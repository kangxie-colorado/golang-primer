"""
https://leetcode.com/problems/unique-binary-search-trees-ii/


"""


from typing import List, Optional

from utils import TreeNode, print_tree


class Solution:
    def generateTrees(self, n: int) -> List[Optional[TreeNode]]:
        inorder = [i for i in range(1, n+1)]

        def buildTrees(A):
            if len(A) == 0:
                return [None]

            if len(A) == 1:
                return [TreeNode(A[0])]

            trees = []
            for i, n in enumerate(A):
                left = buildTrees(A[:i])
                right = buildTrees(A[i+1:])
                for l in left:
                    for r in right:
                        # important to create a new root everytime
                        root = TreeNode(n)
                        root.left = l
                        root.right = r
                        trees.append(root)
            return trees

        return buildTrees(inorder)


"""
Runtime: 86 ms, faster than 57.27% of Python3 online submissions for Unique Binary Search Trees II.
Memory Usage: 15.8 MB, less than 42.38% of Python3 online submissions for Unique Binary Search Trees II.
"""

if __name__ == '__main__':
    for tree in Solution().generateTrees(3):
        print_tree(tree)
        print()
