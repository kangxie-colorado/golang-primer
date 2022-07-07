// https://leetcode.com/problems/binary-tree-maximum-path-sum/

package main

import "fmt"

/*
I see a solution but maybe the timing will suffer

at any node
	N
   / \
  L   R

the max will be one of
max(N+L+R, N+L, N+R, N)

N+L+R will be one locally is used
max(N+L,N+R,N) can be passed to upper level
*/
func maxPathSum(root *TreeNode) int {
	var helper func(root *TreeNode) int
	res := -1001
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		N := root.Val
		lMax := helper(root.Left)
		rMax := helper(root.Right)

		res = max(res, max(N+lMax+rMax, max(N+lMax, max(N+rMax, N))))

		return max(N, max(N+lMax, N+rMax))

	}

	helper(root)
	return res
}

func testMaxPathSum() {

	root := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(maxPathSum(root))
}

/*
Runtime: 14 ms, faster than 91.86% of Go online submissions for Binary Tree Maximum Path Sum.
Memory Usage: 7.5 MB, less than 70.60% of Go online submissions for Binary Tree Maximum Path Sum.

surprise...
*/
