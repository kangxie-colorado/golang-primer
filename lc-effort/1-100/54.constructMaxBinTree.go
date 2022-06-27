// https://leetcode.com/problems/maximum-binary-tree/

package main

/*
analysis:
	build a tree recursived..
	don't seem to hard
	find the max..

	can use the binary search of course.. ah no..

	maybe can use two pass, forward/backware monotonical stack to get
	but ain't so necessary..

	becasue 1 <= nums.length <= 1000
	will just linear search


*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree1(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxN := -1
	maxIdx := 0
	for i := range nums {
		if nums[i] > maxN {
			maxN = nums[i]
			maxIdx = i
		}
	}

	return &TreeNode{
		Val:   maxN,
		Left:  constructMaximumBinaryTree1(nums[:maxIdx]),
		Right: constructMaximumBinaryTree1(nums[maxIdx+1:]),
	}

}

/*
Runtime: 24 ms, faster than 42.34% of Go online submissions for Maximum Binary Tree.
Memory Usage: 7.9 MB, less than 7.21% of Go online submissions for Maximum Binary Tree.

uh.. not to bad??
let me explore that idea..

forward, where is the next bigger value
backward, where is the next bigger value..

where is the O(N) solution?/

damn... I didn't read the problem description accurately
"A maximum binary tree can be built recursively from nums using the following algorithm:" -- can be built this way.. doesn't mean you have to build it this way
*/

func _ugly_but_work_constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	treeNodes := make([]TreeNode, len(nums))
	for i := range treeNodes {
		treeNodes[i] = TreeNode{
			Val:   nums[i],
			Left:  nil,
			Right: nil,
		}
	}

	// store the index in monotonically decreasing order
	stack := []int{}
	for i := range nums {
		if len(stack) > 0 {
			stackTop := stack[len(stack)-1]
			if nums[stackTop] > nums[i] {
				if treeNodes[stackTop].Right == nil {
					treeNodes[stackTop].Right = &treeNodes[i]

				} else {
					if nums[i] > treeNodes[stackTop].Right.Val {

						treeNodes[i].Left = treeNodes[stackTop].Right
						treeNodes[stackTop].Right = &treeNodes[i]
					} else {
						treeNodes[stackTop].Right.Right = &treeNodes[i]
					}

				}
				// mush push onto stack to build the order
				// don't hang tree when popping
				stack = append(stack, i)
			} else {
				for len(stack) > 0 {
					stackTop = stack[len(stack)-1]
					if nums[stackTop] < nums[i] {
						treeNodes[i].Left = &treeNodes[stackTop]
						stack = stack[:len(stack)-1]
					} else {
						treeNodes[stackTop].Right = &treeNodes[i]
						break
					}

				}

				stack = append(stack, i)

			}

		} else {
			stack = append(stack, i)
		}
	}

	for len(stack) > 1 {
		// need to do reblance
		stackTop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nextTop := stack[len(stack)-1]
		if treeNodes[stackTop].Val > treeNodes[nextTop].Right.Val {
			treeNodes[stackTop].Left = treeNodes[nextTop].Right
			treeNodes[nextTop].Right = &treeNodes[stackTop]
		}

	}

	stackTop := stack[0]
	return &treeNodes[stackTop]
}

/*
Runtime: 17 ms, faster than 76.58% of Go online submissions for Maximum Binary Tree.
Memory Usage: 6.8 MB, less than 98.20% of Go online submissions for Maximum Binary Tree.

omg.. ugly but it passed..
the debug technique is adding a big number (or any number) to a working test and see what it triggers
turns out something like

stack: [6,0]
incoming: 5

5.Left -> 0
then I pushed 5 into the stack

however, you need to hang 6.Right -> 5... in the else {}
				for len(stack) > 0 {
					stackTop = stack[len(stack)-1]
					if nums[stackTop] < nums[i] {
						treeNodes[i].Left = &treeNodes[stackTop]
						stack = stack[:len(stack)-1]
					} else {
						treeNodes[stackTop].Right = &treeNodes[i] // need to do this
						break
					}

				}

now it is working.. let me see how to make the code better
*/

func _1_constructMaximumBinaryTree(nums []int) *TreeNode {
	treeNodes := make([]TreeNode, len(nums))
	for i := range treeNodes {
		treeNodes[i] = TreeNode{
			Val:   nums[i],
			Left:  nil,
			Right: nil,
		}
	}

	// store the index in monotonically decreasing order
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		if len(stack) > 0 {
			stackTop := stack[len(stack)-1]
			if nums[stackTop] > nums[i] {
				treeNodes[stackTop].Right = &treeNodes[i]

				// mush push onto stack to build the order
				// don't hang tree when popping
				stack = append(stack, i)
			} else {
				for len(stack) > 0 {
					// pop all the smaller ones and hang my left to them one by one,
					// so it rests on the largest smaller one
					stackTop = stack[len(stack)-1]
					if nums[stackTop] < nums[i] {
						treeNodes[i].Left = &treeNodes[stackTop]
						stack = stack[:len(stack)-1]
					} else {
						// if it runs into a bigger one
						// it needs to be hung onto its right
						// this actually takes care the rebalance
						// so yeah, when I was doing evolution
						// I wasn't generic enough...
						treeNodes[stackTop].Right = &treeNodes[i]
						break
					}

				}

				stack = append(stack, i)

			}

		} else {
			stack = append(stack, i)
		}
	}

	stackTop := stack[0]
	return &treeNodes[stackTop]
}

/*
Runtime: 21 ms, faster than 56.76% of Go online submissions for Maximum Binary Tree.
Memory Usage: 7.6 MB, less than 31.53% of Go online submissions for Maximum Binary Tree.


wow.. so the logic is more concise and pretty now

actually It can be further simplified to
*/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	treeNodes := make([]TreeNode, len(nums))
	for i := range treeNodes {
		treeNodes[i] = TreeNode{
			Val:   nums[i],
			Left:  nil,
			Right: nil,
		}
	}

	// store the index in monotonically decreasing order
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 {
			// pop all the smaller ones and hang my left to them one by one,
			// so it rests on the largest smaller one
			stackTop := stack[len(stack)-1]
			if nums[stackTop] < nums[i] {
				treeNodes[i].Left = &treeNodes[stackTop]
				stack = stack[:len(stack)-1]
			} else {
				// if it runs into a bigger one
				// it needs to be hung onto its right
				// this actually takes care the rebalance
				// so yeah, when I was doing evolution
				// I wasn't generic enough...
				treeNodes[stackTop].Right = &treeNodes[i]
				break
			}
		}

		stack = append(stack, i)

	}

	stackTop := stack[0]
	return &treeNodes[stackTop]
}

/*
Runtime: 18 ms, faster than 72.07% of Go online submissions for Maximum Binary Tree.
Memory Usage: 6.7 MB, less than 98.20% of Go online submissions for Maximum Binary Tree.

probably the best form I can do

but actually it is not O(n) it is O(nlogn)
the previouly is also O(nlogn)
*/

func testConstructMaximumBinaryTree() {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5, 9})
	//constructMaximumBinaryTree([]int{48, 259, 222, 129, 17, 245, 174, 68, 8, 261, 233, 112, 263, 41, 108, 209, 22, 35, 167, 133, 23, 201, 91, 190, 252, 182, 86, 15, 296, 103, 195, 207, 146, 275, 21, 204, 271, 248, 280, 66, 183, 28, 202, 78, 240, 92, 223, 264, 64, 163, 262, 25, 184, 242, 281, 288, 104, 158, 165, 67, 40, 272, 198, 273, 127, 290, 155, 197, 106, 226, 109, 81, 113, 119, 37, 168, 75, 214, 295, 237, 63, 192, 215, 251, 142, 218, 161, 80, 105, 20, 62, 100, 266, 39, 179, 83, 247, 269, 85, 234, 82, 118, 185, 277, 140, 122, 162, 128, 93, 139, 4, 216, 152, 285, 42, 102, 194, 175, 61, 210, 284, 14, 145, 299, 53, 213, 51, 0, 34, 79, 211, 1, 294, 94, 282, 125, 5, 249, 99, 173, 116, 220, 270, 45, 224, 144, 98, 177, 260, 46, 268, 230, 49, 107, 166, 77, 297, 178, 44, 231, 157, 159, 235, 131, 283, 120, 241, 6, 172, 123, 256, 19, 110, 150, 206, 33, 227, 170, 95, 31, 225, 130, 134, 257, 38, 30, 87, 254, 193, 3, 12, 236, 52, 186, 55, 180, 65, 72, 229, 154, 60, 115, 121, 219, 228, 76, 13, 238, 97, 217, 243, 27, 287, 88, 10, 169, 137, 244, 84, 73, 32, 286, 205, 156, 24, 151, 292, 160, 239, 50, 200, 70, 136, 138, 124, 189, 203, 191, 148, 153, 143, 276, 18, 221, 258, 278, 69, 57, 246, 2, 267, 176, 135, 16, 26, 187, 250, 181, 9, 11, 291, 255, 232, 265, 274, 149, 196, 212, 58, 89, 47, 117, 188, 132, 293, 54, 298, 171, 141, 208, 56, 147, 7, 101, 164, 114, 43, 199, 59, 111, 126, 74, 29, 279, 253, 71, 36, 289, 90, 96})
}
