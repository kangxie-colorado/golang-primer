"""
https://leetcode.com/problems/delete-nodes-and-return-forest/

think the difficulty is to maintain the pointers..
and de-duplicate..


"""
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right


from typing import List, Optional

from utils import TreeNode, print_tree


class Solution:
    def delNodes(self, root: Optional[TreeNode], to_delete: List[int]) -> List[TreeNode]:
        res = set()

        def helper(node, parent, ancestor, left):
            if node is None:
                return

            if node.val in to_delete:
                if node != ancestor:
                    if left:
                        parent.left = None
                    else:
                        parent.right = None

                    res.add(ancestor)
                if node.left:
                    res.add(node.left)
                if node.right:
                    res.add(node.right)

                if node in res:
                    res.remove(node)

                helper(node.left, None, node.left, True)
                helper(node.right, None, node.right, False)

            else:
                helper(node.left, node, ancestor, True)
                helper(node.right, node, ancestor, False)
        helper(root, None, root, True)
        return res


"""
Runtime: 139 ms, faster than 15.84% of Python3 online submissions for Delete Nodes And Return Forest.
Memory Usage: 14.5 MB, less than 34.07% of Python3 online submissions for Delete Nodes And Return Forest.

"""


class Solution:
    def delNodes(self, root: Optional[TreeNode], to_delete: List[int]) -> List[TreeNode]:
        res = set()

        def helper(node, parent, ancestor, left):
            if node is None:
                return

            helper(node.left, node, root, True)
            helper(node.right, node, root, False)

            if node.val in to_delete:
                if node != ancestor:
                    if left:
                        parent.left = None
                    else:
                        parent.right = None
                    res.add(ancestor)
                if node.left:
                    res.add(node.left)
                if node.right:
                    res.add(node.right)

                if node in res:
                    res.remove(node)

        helper(root, None, root, True)
        return res


"""
Runtime: 75 ms, faster than 86.43% of Python3 online submissions for Delete Nodes And Return Forest.
Memory Usage: 14.4 MB, less than 98.13% of Python3 online submissions for Delete Nodes And Return Forest.

hmm.. this solutin is better?
how?
I visit first.. then when I back up.. I do the deletion
post-order..

then I look at Lee's code
https://leetcode.com/problems/delete-nodes-and-return-forest/discuss/328853/JavaC%2B%2BPython-Recursion-Solution

well.. once again.. educated
I have been looking at left/right children of node and I had to worry about duplicates..
instead he looks at per-node itsef and simply the idea

the idea is if is_root and not-deleted, it should be in the final results..
when a parent got deleted, its child becomes root and if the child is not deleted.. then add to res

but to make this work, need to pay attention to still delete the root when it should be
so helper() should have a return value.. return the tree (root itself or None if deleted) to parent

res = [] works as the side effect
"""


class Solution:
    def delNodes(self, root: Optional[TreeNode], to_delete: List[int]) -> List[TreeNode]:
        res = []

        def helper(root, isRoot):
            if root is None:
                return None

            root_deleted = root.val in to_delete
            if isRoot and not root_deleted:
                res.append(root)

            root.left = helper(root.left,  root_deleted)
            root.right = helper(root.right,  root_deleted)
            return None if root_deleted else root
        helper(root, True)
        return res


"""
Runtime: 143 ms, faster than 13.69% of Python3 online submissions for Delete Nodes And Return Forest.
Memory Usage: 14.6 MB, less than 34.07% of Python3 online submissions for Delete Nodes And Return Forest.
"""

if __name__ == '__main__':
    s = Solution()
    for t in s.delNodes(TreeNode(1, None,
                                 TreeNode(2, None,
                                          TreeNode(3, None,
                                                   TreeNode(4)))), [3]):
        print_tree(t)

    root = TreeNode(1, TreeNode(2, TreeNode(4), TreeNode(3)))
    for t in s.delNodes(root, [2, 3]):
        print_tree(t)
