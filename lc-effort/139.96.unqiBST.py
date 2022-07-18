"""
https://leetcode.com/problems/unique-binary-search-trees/

notice that inorder for 1-n BST is always 1,2,3...,n... simply because it is binary search tree
the mid should be larger than left and smaller than right...

then maybe I can permutate the preorder and combine with inorder to verify if it is a valid tree
e.g. 1/2/3
pre-order [1,2,3] [1,2,3] [2,1,3] [3,2,1] [3,1,2] all can
but [2,3,1] cannot..

it is going to o(n!).. n<=19.. will it pass?

possibly not... 
but let me just code this idea...


now think how to validate a bst.. given pre-oder and in-order(implicitly known as 1..n)

given [2,3,1] as pre-order 2 is root, then because in-order is naturally 1..3
so the in the in-order, 1/2/3... up to the node prior to 2 will be left, so the length plus root is naturally 2 

thus preorder[1:2] naturally is the left and preorder[2:] is naturally the right.. 
pay attention to the 1-based node value.. so be careful here..

to validate a bst, you need to push down the min/max to validate..
pull up the subtree's min/max will be harder to do
"""


from bisect import bisect_left


def validBST(preorder):

    def validHelper(nums, inorder, min, max):
        if len(nums) == 0:
            return True

        root = nums[0]
        split = bisect_left(inorder, root)

        left = nums[1:split+1]  # 1-based node values..
        right = nums[split+1:]

        leftInorder = inorder[:split]
        rightInorder = inorder[split+1:]

        return min < root < max and \
            validHelper(left, leftInorder, min, root) and \
            validHelper(right, rightInorder, root, max)

    inorder = [i for i in range(1, len(preorder)+1)]
    return validHelper(preorder, inorder, 0, 20)  # 1 <= n <= 19


class Solution_bf(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        for i in range(len(nums)-1, -2, -1):
            if i >= 0 and i+1 < len(nums) and nums[i] < nums[i+1]:
                break

        if i == -1:
            return None
        for j in range(len(nums)-1, i, -1):
            if nums[j] > nums[i]:
                nums[j], nums[i] = nums[i], nums[j]
                break
        a = nums[i+1:]
        nums[i+1:] = sorted(a)

        return nums

    def numTrees(self, n):
        """
        :type n: int
        :rtype: int
        """
        start = [i for i in range(1, n+1)]
        count = 0
        while start is not None:
            if validBST(start):
                count += 1

            start = self.nextPermutation(start)
        return count

    def validNums(self, n):
        start = [i for i in range(1, n+1)]

        while start is not None:
            if not validBST(start):
                print(start, validBST(start))
            start = self.nextPermutation(start)


"""
okay.. the algorithm seems correct but it will only be able to cal up to n=9 without TLE
"""


class Solution(object):
    def numTrees(self, n):
        dp = [0]*(n+1)
        dp[0] = dp[1] = 1

        for i in range(2, n+1):
            for j in range(0, i):
                dp[i] += dp[j]*dp[i-1-j]

        return dp[n]


"""
Runtime: 20 ms, faster than 78.66% of Python online submissions for Unique Binary Search Trees.
Memory Usage: 13.2 MB, less than 83.77% of Python online submissions for Unique Binary Search Trees.
"""

if __name__ == "__main__":
    """
    print(validBST([1, 2, 3]))
    print(validBST([1, 3, 2]))
    print(validBST([2, 1, 3]))
    print(validBST([2, 3, 1]))
    print(validBST([3, 1, 2]))
    print(validBST([3, 2, 1]))
        s = Solution()



    """
    s = Solution()
    print(s.numTrees(1))
    print(s.numTrees(2))
    print(s.numTrees(3))
    print(s.numTrees(4))
    print(s.numTrees(5))
    print(s.numTrees(6))
    print(s.numTrees(7))
    print(s.numTrees(8))
    print(s.numTrees(9))
    print(s.numTrees(10))
