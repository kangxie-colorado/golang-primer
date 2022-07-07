// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

package main

/*
I did this before and still have some memory but don't trust it

so I think it depends on these scenarios
1. one to the root's left, one to root's right -- then root, no other possibility
2. both to the root's left,
	then use root.Left as new root and check,
3. both to the root's right.
4. one is root and one is the child.. then the root one

let me try a bit, if not working, leave to tomorrow afternoon with whiteboard

*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	if root.Val > p.Val && root.Val < q.Val {
		return root
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// pay attention to left right
	// root < min(p,q), then should search right tree
	// root > max(p,q), then should search left tree
	if root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return nil
}

/*
Runtime: 31 ms, faster than 33.33% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.
Memory Usage: 7.3 MB, less than 56.45% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.

Runtime: 23 ms, faster than 65.96% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.
Memory Usage: 7.4 MB, less than 44.11% of Go online submissions for Lowest Common Ancestor of a Binary Search Tree.

*/
