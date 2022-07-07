// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

package main

/*
I should have done this before... but why I didn't have a submission
hmm...

navie solution is to do a inorder traverse and keep an count..
hard to keep a count..

just inorder and then use kth
should be O(n)

let me just do that
*/

func _inorder_kthSmallest(root *TreeNode, k int) int {
	var inorder func(root *TreeNode) []int

	inorder = func(root *TreeNode) []int {
		if root == nil {
			return nil
		}
		subNums := []int{}
		subNums = append(subNums, inorder(root.Left)...)
		subNums = append(subNums, root.Val)
		subNums = append(subNums, inorder(root.Right)...)

		return subNums
	}

	nums := inorder(root)
	res := 0
	for i := 0; i < len(nums); i++ {
		if i == k-1 {
			res = nums[i]
		}
	}
	return res
}

/*
Runtime: 18 ms, faster than 30.51% of Go online submissions for Kth Smallest Element in a BST.
Memory Usage: 8.2 MB, less than 5.05% of Go online submissions for Kth Smallest Element in a BST.

now this
Follow up: If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?

think we can count the prefix number
left tree is root's prefix
left+root is right's prefix

or we can actually modify above algorithm to return num and build it up..
the problem may be converted as follow
1. root, k
2. if left tree >= k, the root.left, k
3. if left tree <k, the root.right k - (left-tree's node no)
*/

func kthSmallest(root *TreeNode, k int) int {
	var countNode func(root *TreeNode) int
	countNode = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftCount := countNode(root.Left)
		rightCount := countNode(root.Right)
		return leftCount + 1 + rightCount

	}

	leftCount := countNode(root.Left)
	if leftCount == k-1 {
		// root will be kth
		return root.Val
	} else if leftCount >= k {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-leftCount-1)
	}

}

/*
Runtime: 10 ms, faster than 79.19% of Go online submissions for Kth Smallest Element in a BST.
Memory Usage: 6.5 MB, less than 69.49% of Go online submissions for Kth Smallest Element in a BST.

fast enough?

this countNode is actually traverse all the tree (half to whole)
*/
