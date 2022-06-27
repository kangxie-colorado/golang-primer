// https://leetcode.com/problems/maximum-binary-tree-ii/

package main

/*
actually the max binary tree I in  https://leetcode.com/problems/maximum-binary-tree/ made this piece of case
if I didn't go thru the monotonical exercise, hmm... it would be hard

I may think I need to restore the tree to its array form and then re-form the tree, which isn't particularly hard actually
it is just left-root-right traverse...

but here is the smart way

because you are appending..
so only root-right-.. leg is concerned

you push all the right leg onto a queue
then examine the incoming number, if the stacktop is smaller than it... hang stacktop to its left.. then pop
if it hits a bigger stack top... then hang incoming to its right

if the incoming ends up in the stack.. meaning it is the biggest one by result.. then it will be hanging previous root as its left

*/

func _iter_insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	stack := []*TreeNode{}

	treeNode := root
	for treeNode != nil {
		stack = append(stack, treeNode)
		treeNode = treeNode.Right
	}

	newTreeNode := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}

	for len(stack) > 0 {
		// pop all the smaller ones and hang my left to them one by one,
		// so it rests on the largest smaller one
		stackTop := stack[len(stack)-1]
		if stackTop.Val < val {
			newTreeNode.Left = stackTop
			stack = stack[:len(stack)-1]
		} else {
			// if it runs into a bigger one
			// it needs to be hung onto its right
			// this actually takes care the rebalance
			// so yeah, when I was doing evolution
			// I wasn't generic enough...
			stackTop.Right = newTreeNode
			break
		}
	}

	stack = append(stack, newTreeNode)

	stackTop := stack[0]
	return stackTop
}

/*
Runtime: 2 ms, faster than 28.57% of Go online submissions for Maximum Binary Tree II.
Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Maximum Binary Tree II.

cool.. this is utilize previous solution
and it is an iterative procedure

let me change this into a recursive procedure
*/

func insertAt(root, new *TreeNode) *TreeNode {
	if root == nil {
		return new
	}

	if root.Val < new.Val {
		new.Left = root
		return new
	}

	root.Right = insertAt(root.Right, new)
	return root
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {

	newTreeNode := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}

	return insertAt(root, newTreeNode)
}

/*
Runtime: 4 ms, faster than 14.29% of Go online submissions for Maximum Binary Tree II.
Memory Usage: 2.7 MB, less than 42.86% of Go online submissions for Maximum Binary Tree II.

need some thinking
and my first version went into a mind coma..

so the logical is
at a certain root... if it is bigger, then hang the root to its left
since we are walking top-down.. so a bigger value will become root's parent and root becomes its left child
and the new will become new root --- thus return new

if we have walk down the right leg, and bigger value is smaller than any of the node.. then it will just end up to the right-most child
it is important you alwasy hang the root.right to new right (which could be old right of course, but chain them)
so here we return root

and the rebalance is taken care by this two steps above automatically


but the recursive procedure looks prettier
*/
