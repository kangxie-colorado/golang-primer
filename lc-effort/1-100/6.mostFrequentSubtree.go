// https://leetcode.com/problems/most-frequent-subtree-sum/

/*
	analysis: subtree sum, most frequent ones
	so naturally, find out all the sub-tree sum and put in a map...
	then sort it

	let me do this first, then think of optimization

   	* Definition for a binary tree node.
	* type TreeNode struct {
	*     Val int
	*     Left *TreeNode
	*     Right *TreeNode
	* }

subtree_sum_map := [int] *Treenode {}
findFrequentTreeSum(root *TreeNode)
		if root == nil
			return 0

		leftSum := findFrequentTreeSum(root.left)
		rightSum :=  findFrequentTreeSum(root.right)

		subtree_sum_map[leftSum+rightSum+root.Val]++

		return leftSum+rightSum+root.Val

*/

package main

import "fmt"

// the times a subtree sum appears
var subtreeSums map[int]int

func findAllSubtreeSums(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSum := findAllSubtreeSums(root.Left)
	rightSum := findAllSubtreeSums(root.Right)

	treeSum := leftSum + rightSum + root.Val
	if _, found := subtreeSums[treeSum]; !found {
		subtreeSums[treeSum] = 1
	} else {
		subtreeSums[treeSum]++
	}

	return treeSum
}

func findFrequentTreeSum(root *TreeNode) []int {
	subtreeSums = make(map[int]int)
	findAllSubtreeSums(root)

	max := -1
	sums := []int{}
	for k, v := range subtreeSums {
		if v > max {
			sums = []int{k}
			max = v
		} else if v == max {
			sums = append(sums, k)
		}
	}

	return sums
}

/*
Success
Details
Runtime: 28 ms, faster than 8.57% of Go online submissions for Most Frequent Subtree Sum.
Memory Usage: 6.4 MB, less than 82.86% of Go online submissions for Most Frequent Subtree Sum.

I didn't use as much as memory, so it suggests I could use some more memory?

Now let me try to think of something quicker
take a break first

go watering some flowers

is there really a better algorithm?
by definition, you need to traverse the tree

the accounting part which going thru the map is not the dominate force or at least O(n)
yes!! there is no method faster than traversing the tree

check the example solution, it is 7x faster
so possible my other parts are slower but not theoritically

so let it go

	func findFrequentTreeSum(root *TreeNode) []int {
		res := []int{}
		hash_map := make(map[int]int)
		sum(root, &hash_map)
		mp := make(map[int][]int)
		l := 0
		for k, v := range hash_map {
			mp[v] = append(mp[v], k)
			if l <= v {
				l = v
				res = mp[v]
			}
		}
		return res
	}

	func sum(root *TreeNode, hash_map *map[int]int) int {
		if root == nil {
			return 0
		}
		root_sum := root.Val + sum(root.Left, hash_map) + sum(root.Right, hash_map)
		(*hash_map)[root_sum]++
		return root_sum
	}


*/

func testFreqSubtreeSums() {
	root := &TreeNode{
		5,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			-5,
			nil,
			nil,
		},
	}

	fmt.Println(findFrequentTreeSum(root))
}
