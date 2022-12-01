"""
https://leetcode.com/problems/distribute-coins-in-binary-tree/


if there is one coin on a node, it would make no sense to move it to another and get a re-suppliment 
that is just shift the load... 

so the problem is for a zero coin node to reach to a >1 coin node... the cumulitive smallest sum
but there is also decision to make

2-0-2-0
see this branch.. the first 0 has two choices of one
but if it take the right coin, then the 2nd 0 would need 3 steps... 

another idea struck me.. 
if I can focus on a single tree


"""


from typing import Optional

from utils import TreeNode


class Solution:
    def distributeCoins(self, root: Optional[TreeNode]) -> int:
        res = 0

        def helper(root):
            nonlocal res
            if root is None:
                return 0, 0

            node = 1
            coin = root.val

            leftNodes, leftCoins = helper(root.left)
            rightNodes, rightCoins = helper(root.right)

            res += abs(leftNodes - leftCoins)
            res += abs(rightNodes - rightCoins)

            return node+leftNodes+rightNodes, coin+leftCoins+rightCoins
        helper(root)
        return res


if __name__ == '__main__':
    s = Solution()
    print(s.distributeCoins(TreeNode(3,
                                     TreeNode(0),
                                     TreeNode(0))))
