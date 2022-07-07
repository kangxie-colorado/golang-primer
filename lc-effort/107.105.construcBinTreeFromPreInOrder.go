// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package main

/*
use pre-order to pinpoint the root
use the in-order to divide the left and right

and notice left/root always appear before right so the composite will be like
pre-order root 			left...left 	right...right
in-order  left...left 	root			right...right

notice how they match up
left-part and root are just swapped and right-part stays in the same range
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	// find left, right part
	for i := range inorder {
		if inorder[i] == root.Val {
			leftInorder := inorder[:i]
			rightInorder := inorder[i+1:]

			leftPreorder := preorder[1 : i+1]
			rightPreorder := preorder[i+1:]

			root.Left = buildTree(leftPreorder, leftInorder)
			root.Right = buildTree(rightPreorder, rightInorder)

			return root
		}
	}

	return nil
}

/*
Runtime: 6 ms, faster than 66.50% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
Memory Usage: 4.2 MB, less than 51.86% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
*/
