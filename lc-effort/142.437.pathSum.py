"""
https://leetcode.com/problems/path-sum-iii/

this should be similar to a previous problem

any root can produce 4 sums for its own
root+left, root+right, root, root+left+right

the first three can be send back to higher level, the last one stays on itself
"""


from collections import defaultdict


class Solution_Premature_Thinking(object):
    def pathSum(self, root, targetSum):
        """
        :type root: TreeNode
        :type targetSum: int
        :rtype: int
        """
        def helper(root):
            """
            return: 
                (root+left, root+right, root)
            """
            if root is None:
                return (0, 0, 0)

            leftT = helper(root.left)
            rightT = helper(root.right)
            ...


"""
hmm... just by examing one bit I see it will fan out more than 3
it will be 3*3 *3 *3... so not managable

then I see this 
The path does not need to start or end at the root or a leaf,
but it must go downwards (i.e., traveling only from parent nodes to child nodes).

so I didn't even finish reading the statement.. shame on me

so if still do the pulling up, it will still face the fanning out dilemma..
I might just do the push down..

hum.. not easy as well..
rethink... 

still pulling up
so one node can push up these values

root root+left root+right
on parent

parent can add itself, if there is a target then target +1
then append itself to the list.. and push up

no other better way I can see

"""


class Solution(object):
    def pathSum(self, root, targetSum):
        count = 0

        def helper(root):
            """
            return: list of sums at root
            """
            if root is None:
                return []

            res = []
            left = helper(root.left)
            right = helper(root.right)

            for i in left+right:
                if root.val+i == targetSum:
                    count += 1
                res.append(root.val+i)

            if root.val == targetSum:
                count += 1
            res.append(root.val)
            return res

        helper(root)
        return count


"""
Runtime: 713 ms, faster than 26.43% of Python3 online submissions for Path Sum III.
Memory Usage: 15.5 MB, less than 31.00% of Python3 online submissions for Path Sum III.

hmm.. same approach let me try push down..
"""


class Solution(object):
    def pathSum(self, root, targetSum):
        count = 0

        def helper(root, sums):
            if root is None:
                return

            nonlocal count
            pushdowns = []
            if root.val == targetSum:
                count += 1
            pushdowns.append(root.val)

            for s in sums:
                if root.val + s == targetSum:
                    count += 1
                pushdowns.append(root.val + s)

            helper(root.left, pushdowns)
            helper(root.right, pushdowns)

        helper(root, [])
        return count


"""
Runtime: 839 ms, faster than 23.52% of Python3 online submissions for Path Sum III.
Memory Usage: 42.7 MB, less than 6.21% of Python3 online submissions for Path Sum III.

okay.. I am O(N*logN) because on every node, I go thru up to the heights of partial sums

how about I use a hashmap to save the O(logN) sum search..

but then you still go thru all of them to add the current root.val to pass down..

hmm... there must be some other trick to keep it together..
let me check other's 

https://leetcode.com/problems/path-sum-iii/discuss/141424/Python-step-by-step-walk-through.-Easy-to-understand.-Two-solutions-comparison.-%3A-)

comparing to the brute force described here
I see he is doing a traversal and from each node, do a dfs.. which is basically same as mine pushing-down method

however I do one traversal but i pushed down all the partial sums.. which is also O(H) so yeah. the same

then the optimization he has is about memorization.. I wonder how
"""

"""
the memorization of path sums.. wow

use curreentPathsum - target ... in a map.. if it is a hit.. then between that path and now, the sum is target
fucking genius.. 
"""


class Solution(object):
    def pathSum(self, root, targetSum):
        sumMap = defaultdict(int)
        count = 0
        sumMap[0] = 1

        def helper(root, sumToRoot):
            nonlocal sumMap
            nonlocal count

            if root is None:
                return

            currSum = sumToRoot+root.val
            count += sumMap[currSum-targetSum]

            sumMap[currSum] += 1
            helper(root.left, currSum)
            helper(root.right, currSum)
            # going up, this path sum should disappear by one times..
            sumMap[currSum] -= 1

        helper(root, 0)
        return count


"""
[1]
1

failed here

hm.. thinking we can get rid of 
            # this will need to be handled as edge case
            # it cannot become a sum itself..
            if root.val == targetSum:
                count += 1

because when a node's val is target, the pathsum will presum+targetsum and prevsum naturally appeared before

Runtime: 63 ms, faster than 86.69% of Python3 online submissions for Path Sum III.
Memory Usage: 15.4 MB, less than 31.18% of Python3 online submissions for Path Sum III.

a very interesting problem, helped to think
"""
