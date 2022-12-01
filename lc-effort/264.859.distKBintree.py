"""
https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

I think this is a very interesting problem
probably many solutions

but right now I am seeing a binary encoding way..
using the example case 
Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2


I could encoding the tree like 
1 [10 [100 [101 [1010 1011]]]] [11 [110 111]]

then I could analyze the common prefix to deduct the distance
like 
10 vs 1010.. common prefix is 10, and 1010 has two more level... then it would be 2
10 vs 1011.. same story

10 vs 11.. common prefix is 1, the left is one level and the right is one level, so 2

let me verify 
10 vs 110... common prefix is 1.. left 1, right 2... then 3
so more common ways 
the common prefix len, comLen.. res=(leftLen-comLen)+(rightLen-comLen)
"""

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


from typing import List
from utils import TreeNode, print_tree


class Solution:
    def distanceK(self, root: TreeNode, target: TreeNode, k: int) -> List[int]:
        encodings = {}
        targetEncoding = ""

        def encode(root, pEncoding, left):
            nonlocal targetEncoding
            if not root:
                return

            if left:
                thisNode = pEncoding+"0"
            else:
                thisNode = pEncoding+"1"

            if root.val == target.val:
                targetEncoding = thisNode

            encodings[thisNode] = root.val
            encode(root.left, thisNode, True)
            encode(root.right, thisNode, False)

        encode(root, "", False)

        def commonPrefix(s1, s2):
            res = 0
            for c1, c2 in zip(s1, s2):
                if c1 != c2:
                    break
                res += 1
            return res

        res = []
        for e, val in encodings.items():
            cPfix = commonPrefix(e, targetEncoding)
            if len(e)+len(targetEncoding) - 2*cPfix == k:
                res.append(val)

        return res


if __name__ == "__main__":
    root = TreeNode(3,
                    TreeNode(5,
                             TreeNode(6),
                             TreeNode(2,
                                      TreeNode(7),
                                      TreeNode(4))),
                    TreeNode(1,
                             TreeNode(0),
                             TreeNode(8)),

                    )

    s = Solution()
    print(s.distanceK(root, TreeNode(5), 2))
