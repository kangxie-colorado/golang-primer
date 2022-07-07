// https://leetcode.com/problems/subtree-of-another-tree/

package main

import "fmt"

/*
not super hard, but be sure to get all the combinations

they can be compared on the
root vs subroot level
or root.right vs subroot level
or root.left vs subroot level
*/
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var isEqual func(tree1, tree2 *TreeNode) bool
	isEqual = func(tree1, tree2 *TreeNode) bool {
		if tree1 == nil && tree2 == nil {
			return true
		}

		if tree1 == nil || tree2 == nil || tree1.Val != tree2.Val {
			return false
		}

		return isEqual(tree1.Left, tree2.Left) && isEqual(tree1.Right, tree2.Right)
	}

	return isEqual(root, subRoot) ||
		(root.Left != nil && isSubtree(root.Left, subRoot)) ||
		(root.Right != nil && isSubtree(root.Right, subRoot))

}

/*
failed
[3,4,5,1,2]
[4,1,2,1]

hmm...
fuck it a typo -- two uses of tree1
		if tree1 == nil && tree1 == nil {

failed
[1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,null,1,2]
[1,null,1,null,1,null,1,null,1,null,1,2]

okay let me actually build this tree to test it out

oh I see, I only called upon isSubTree once..
it has to be recursive as well

also has to gurad Left/Right, on nil
	return isEqual(root, subRoot) ||
		(root.Left != nil && isSubtree(root.Left, subRoot)) ||
		(root.Right != nil && isSubtree(root.Right, subRoot))

pass

Runtime: 18 ms, faster than 65.54% of Go online submissions for Subtree of Another Tree.
Memory Usage: 7.2 MB, less than 21.56% of Go online submissions for Subtree of Another Tree.

*/

func testIsSubTree() {

	tree1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		},
	}

	tree2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
	}

	fmt.Println(isSubtree(tree1, tree2))
}
