"""
https://leetcode.com/problems/all-possible-full-binary-trees/


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
    def allPossibleFBT(self, N: int) -> List[Optional[TreeNode]]:
        if N == 1:
            return [TreeNode(0)]

        def helper(n):
            if n % 2 == 0:
                return None

            if n == 1:
                return [TreeNode(0)]
            if n == 3:
                return [TreeNode(0, TreeNode(0), TreeNode(0))]

            return self.allPossibleFBT(n)

        res = []
        for i in range(1, N+1):
            # use node i as root, 1-based
            # e.g. node1 as root, then left tree is left with 0 ndoes
            # node2 as root, then left tree is left with 1 nodes, right tree left with N-2 nodes
            left = helper(i-1)
            right = helper(N-i)
            if left is None or right is None:
                continue

            for l in left:
                for r in right:
                    root = TreeNode(0)
                    root.left = l
                    root.right = r
                    res.append(root)
        return res


"""
Runtime: 662 ms, faster than 6.48% of Python3 online submissions for All Possible Full Binary Trees.
Memory Usage: 23.8 MB, less than 21.80% of Python3 online submissions for All Possible Full Binary Trees.

ugly..
performance wise, I have lot of duplicate computation

maybe it can be DP'ed
"""


class Solution:
    def allPossibleFBT(self, N: int) -> List[Optional[TreeNode]]:
        # know n%2==0, None
        # know 1->[ [0] ]
        # know 3->[ [0,0,0]]
        # can I dp?

        dpMap = {}
        dpMap[1] = [TreeNode(0)]
        dpMap[3] = [TreeNode(0, TreeNode(0), TreeNode(0))]
        for i in range(0, 21):
            if i % 2 == 0:
                dpMap[i] = []

        for i in range(5, N+1, 2):
            # left number range from 0 to i-1
            dpMap[i] = []
            for k in range(i-1):
                if len(dpMap[k]) == 0 or len(dpMap[i-k-1]) == 0:
                    continue

                for l in dpMap[k]:
                    for r in dpMap[i-k-1]:
                        root = TreeNode(0)
                        root.left = l
                        root.right = r
                        dpMap[i].append(root)

        return dpMap[N]


"""

Runtime: 258 ms, faster than 65.10% of Python3 online submissions for All Possible Full Binary Trees.
Memory Usage: 17.4 MB, less than 65.18% of Python3 online submissions for All Possible Full Binary Trees.
"""
