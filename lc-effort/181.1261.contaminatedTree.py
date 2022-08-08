"""
https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/

I did that hashset.. most people can think of
then I see that use binary to represent

think that is very smart
but the example code I didn't figure out the logic completely 
let me code it up

"""


from typing import Optional

from utils import TreeNode


class FindElements:

    def __init__(self, root: Optional[TreeNode]):
        self.root = root

    def find(self, target: int) -> bool:
        path = bin(target+1)[3:]
        root = self.root
        index = 0
        while index < len(path):
            root = root.left if path[index] == '0' else root.right
            if root is None:
                break
            index += 1

        return index == len(path)


"""
Runtime: 79 ms, faster than 99.11% of Python3 online submissions for Find Elements in a Contaminated Binary Tree.
Memory Usage: 17.5 MB, less than 97.99% of Python3 online submissions for Find Elements in a Contaminated Binary Tree.

so after +1
numbers on the same level will have equal length of binary representation 
1, 2 is 1 10... hard to reason around
+1, 2,3 is 10 11... 

3 4 5 6 -> 11 100 101 110... hard to reason around
+1 4 5 6 7.. -> 100 101 110 111.. 

first 1 has no value of routing.. so get rid of it
then on 0, go left; on 1, go right.. 

if at some point, the node doesn't exist.. then no there
otherwise, it has walked the whole path.. it is there... 

this might help another problem I guess.
"""
