// https://leetcode.com/problems/validate-binary-search-tree/

package main

import "math"

/*
I thought this is not a difficult problem
but the pass rate is really low

got me doubted

so what should it be?
maybe many edge cases???

let me see
*/

func _too_long_isValidBST(root *TreeNode) bool {

	type SubTreeMinMax struct {
		leftMin int // left subTree's min.. the leftMost
		leftMax int // left subtree's max.. to left, then right

		rightMax int // right subtree's max.. the rightmost
		rightMin int // right subtree's min.. to right, then left
	}

	var validAndMinMax func(root *TreeNode) (bool, SubTreeMinMax)
	validAndMinMax = func(root *TreeNode) (bool, SubTreeMinMax) {
		if root.Left == nil && root.Right == nil {
			return true, SubTreeMinMax{
				root.Val,
				root.Val,
				root.Val,
				root.Val,
			}
		}

		validLeft := true
		validRight := true
		var leftTreeMinMax SubTreeMinMax
		var rightTreeMinMax SubTreeMinMax
		if root.Left != nil && root.Right != nil {
			validLeft, leftTreeMinMax = validAndMinMax(root.Left)
			validRight, rightTreeMinMax = validAndMinMax(root.Right)
			if !validLeft || !validRight || leftTreeMinMax.rightMax >= root.Val || rightTreeMinMax.leftMin <= root.Val {
				return false, SubTreeMinMax{
					leftMin:  leftTreeMinMax.leftMin,
					leftMax:  leftTreeMinMax.rightMax,
					rightMax: rightTreeMinMax.rightMax,
					rightMin: rightTreeMinMax.leftMin,
				}
			}
			return true, SubTreeMinMax{
				leftMin:  leftTreeMinMax.leftMin,
				leftMax:  leftTreeMinMax.rightMax,
				rightMax: rightTreeMinMax.rightMax,
				rightMin: rightTreeMinMax.leftMin,
			}
		} else if root.Left == nil {
			// has only right child
			validRight, rightTreeMinMax = validAndMinMax(root.Right)
			if !validRight || rightTreeMinMax.leftMin <= root.Val {
				return false, SubTreeMinMax{
					leftMin:  root.Val,
					leftMax:  root.Val,
					rightMax: rightTreeMinMax.rightMax,
					rightMin: rightTreeMinMax.leftMin,
				}
			}

			return true, SubTreeMinMax{
				leftMin:  root.Val,
				leftMax:  root.Val,
				rightMax: rightTreeMinMax.rightMax,
				rightMin: rightTreeMinMax.leftMin,
			}

		} else {
			// only left child
			validLeft, leftTreeMinMax = validAndMinMax(root.Left)
			if !validLeft || leftTreeMinMax.rightMax >= root.Val {
				return false, SubTreeMinMax{
					leftMin:  leftTreeMinMax.leftMin,
					leftMax:  leftTreeMinMax.rightMax,
					rightMax: root.Val,
					rightMin: root.Val,
				}
			}
			return true, SubTreeMinMax{
				leftMin:  leftTreeMinMax.leftMin,
				leftMax:  leftTreeMinMax.rightMax,
				rightMax: root.Val,
				rightMin: root.Val,
			}
		}

	}

	res, _ := validAndMinMax(root)
	return res
}

/*
[2,2,2] failed
okay.. this is an edge case

[0,-1]
hmm...

ah..
		if !validLeft || !validRight || leftTreeMinMax.rightMax >= root.Val || rightTreeMinMax.leftMin <= root.Val {
			return false, SubTreeMinMax{
				leftMin:  leftTreeMinMax.leftMin,
				leftMax:  leftTreeMinMax.rightMax,
				rightMax: rightTreeMinMax.rightMax,
				rightMin: rightTreeMinMax.leftMin,
			}
		}
the structure has a default of zero and will fail the test...

wow...

Runtime: 8 ms, faster than 60.65% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.1 MB, less than 83.35% of Go online submissions for Validate Binary Search Tree.

Runtime: 4 ms, faster than 87.71% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.4 MB, less than 17.34% of Go online submissions for Validate Binary Search Tree.
*/

func _still_long_isValidBST(root *TreeNode) bool {

	type SubTreeMinMax struct {
		min int
		max int
	}

	var validAndMinMax func(root *TreeNode) (bool, SubTreeMinMax)
	validAndMinMax = func(root *TreeNode) (bool, SubTreeMinMax) {
		if root.Left == nil && root.Right == nil {
			return true, SubTreeMinMax{
				root.Val,
				root.Val,
			}
		}

		validLeft := true
		validRight := true
		var leftTreeMinMax SubTreeMinMax
		var rightTreeMinMax SubTreeMinMax
		if root.Left != nil && root.Right != nil {
			validLeft, leftTreeMinMax = validAndMinMax(root.Left)
			validRight, rightTreeMinMax = validAndMinMax(root.Right)
			if !validLeft || !validRight || leftTreeMinMax.max >= root.Val || rightTreeMinMax.min <= root.Val {
				return false, SubTreeMinMax{}
			}
			return true, SubTreeMinMax{
				min: leftTreeMinMax.min,
				max: rightTreeMinMax.max,
			}
		} else if root.Left == nil {
			// has only right child
			validRight, rightTreeMinMax = validAndMinMax(root.Right)
			if !validRight || rightTreeMinMax.min <= root.Val {
				return false, SubTreeMinMax{}
			}

			return true, SubTreeMinMax{
				min: root.Val,
				max: rightTreeMinMax.max,
			}

		} else {
			// only left child
			validLeft, leftTreeMinMax = validAndMinMax(root.Left)
			if !validLeft || leftTreeMinMax.max >= root.Val {
				return false, SubTreeMinMax{}
			}
			return true, SubTreeMinMax{
				min: leftTreeMinMax.min,
				max: root.Val,
			}
		}

	}

	res, _ := validAndMinMax(root)
	return res
}

/*
simplify this a bit to save only a min/max for any subtree

Runtime: 6 ms, faster than 72.84% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.3 MB, less than 32.41% of Go online submissions for Validate Binary Search Tree.


yeah now I see, this is working but ugly
it keeps doing some unnecessary information

also the leaf node, nil, is hard to process, really? maybe I can get away

this is to pull subtree to parent.. like most tree algorithms do
but there is a solution that push parent to subtree

use root.val as leftTree's max, and rightTree's min
make sure it is indeed the max in the left tree and min the right tree
*/

func _concise_isValidBST(root *TreeNode) bool {

	type SubTreeMinMax struct {
		min int
		max int
	}

	var validAndMinMax func(root *TreeNode) (bool, *SubTreeMinMax)
	validAndMinMax = func(root *TreeNode) (bool, *SubTreeMinMax) {
		if root == nil {
			return true, &SubTreeMinMax{
				min: math.MaxInt,
				max: math.MinInt,
			}
		}

		validLeft, left := validAndMinMax(root.Left)
		validRight, right := validAndMinMax(root.Right)

		if !validLeft || !validRight || left.max >= root.Val || right.min <= root.Val {
			return false, nil
		}

		return true, &SubTreeMinMax{
			min: min(root.Val, left.min),
			max: max(root.Val, right.max),
		}

	}

	res, _ := validAndMinMax(root)
	return res
}

/*
Runtime: 7 ms, faster than 68.19% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.9 MB, less than 12.78% of Go online submissions for Validate Binary Search Tree.

so it can be pretty concise as well

let me now push the root.val (min/max down )
I was a liitle held back by this -2^31 <= Node.val <= 2^31 - 1


*/

func isValidBST(root *TreeNode) bool {
	var verifyValidBST func(root *TreeNode, min, max int) bool

	verifyValidBST = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}

		return root.Val > min && root.Val < max &&
			verifyValidBST(root.Left, min, root.Val) &&
			verifyValidBST(root.Right, root.Val, max)

	}

	return verifyValidBST(root, math.MinInt, math.MaxInt)
}

/*
Runtime: 3 ms, faster than 93.86% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.2 MB, less than 83.35% of Go online submissions for Validate Binary Search Tree.
*/
