"""
https://leetcode.com/problems/cousins-in-binary-tree/



"""


from typing import Optional
from utils import TreeNode


class Solution:
    def isCousins(self, root: Optional[TreeNode], x: int, y: int) -> bool:
        level = [root]
        while True:
            newLevel = []
            vals = set()
            for node in level:
                if node:
                    vals.add(node.val)
                    newLevel.append(node.left)
                    newLevel.append(node.right)

            xIn = x in vals
            yIn = y in vals
            if xIn ^ yIn:
                return False
            if xIn & yIn:
                return True
            level = newLevel


"""
okay.. siblings vs cousins
"""


class Solution:
    def isCousins(self, root: Optional[TreeNode], x: int, y: int) -> bool:
        encodings = {}

        def encode(root, pEncoding, left):
            if not root:
                return

            if left:
                thisNode = pEncoding+"0"
            else:
                thisNode = pEncoding+"1"

            encodings[root.val] = thisNode
            encode(root.left, thisNode, True)
            encode(root.right, thisNode, False)
        encode(root, "", False)
        return len(encodings[x]) == len(encodings[y]) and encodings[x][:-1] != encodings[y][:-1]


if __name__ == '__main__':
    s = Solution()

    print(s.isCousins(TreeNode(1,
                               TreeNode(2,
                                        None,
                                        TreeNode(4)),
                               TreeNode(3)), 2, 3))
