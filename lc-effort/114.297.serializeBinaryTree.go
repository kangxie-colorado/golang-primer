// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
I did that construct tree from preorder and inorder so I can utilize that
*/

/*
version1

type Codec struct {
	seq1 string
	seq2 string
}

func Constructor() Codec {
	return Codec{
		" ",
		",",
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var preorder func(root *TreeNode) string
	var inorder func(root *TreeNode) string

	preorder = func(root *TreeNode) string {
		if root == nil {
			return ""
		}

		return strconv.Itoa(root.Val) + this.seq1 +
			preorder(root.Left) +
			preorder(root.Right)
	}
	inorder = func(root *TreeNode) string {
		if root == nil {
			return ""
		}

		return inorder(root.Left) +
			strconv.Itoa(root.Val) + this.seq1 +
			inorder(root.Right)
	}

	return preorder(root) + this.seq2 + inorder(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, this.seq2)
	preStr, inStr := strings.Trim(parts[0], this.seq1), strings.Trim(parts[1], this.seq1)
	if len(preStr) == 0 {
		return nil
	}

	preOrder := []int{}
	inOrder := []int{}
	for _, n := range strings.Split(preStr, this.seq1) {
		num, _ := strconv.Atoi(n)
		preOrder = append(preOrder, num)
	}
	for _, n := range strings.Split(inStr, this.seq1) {
		num, _ := strconv.Atoi(n)
		inOrder = append(inOrder, num)
	}

	var buildTree func(preorder []int, inorder []int) *TreeNode
	buildTree = func(preorder, inorder []int) *TreeNode {
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

	return buildTree(preOrder, inOrder)
}
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//// version 2
// didn't notice it is two objects
// but it actually crossed my mind...
// but it was not loud enough to stop me doing the map
// anyway, good exercise to put an idea into code anyway
/*
type Codec struct {
	seq1       string
	seq2       string
	nodeSeqMap map[int]int
}

func Constructor() Codec {
	nodeSeqMap := make(map[int]int)
	return Codec{
		" ",
		",",
		nodeSeqMap,
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var preorder func(root *TreeNode) string
	var inorder func(root *TreeNode) string
	var mark func(root *TreeNode)

	marker := 0
	mark = func(root *TreeNode) {
		if root == nil {
			return
		}

		this.nodeSeqMap[marker] = root.Val
		root.Val = marker
		marker++
		mark(root.Left)
		mark(root.Right)
	}

	preorder = func(root *TreeNode) string {
		if root == nil {
			return ""
		}

		return strconv.Itoa(root.Val) + this.seq1 +
			preorder(root.Left) +
			preorder(root.Right)
	}
	inorder = func(root *TreeNode) string {
		if root == nil {
			return ""
		}

		return inorder(root.Left) +
			strconv.Itoa(root.Val) + this.seq1 +
			inorder(root.Right)
	}
	mark(root)
	return preorder(root) + this.seq2 + inorder(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, this.seq2)
	preStr, inStr := strings.Trim(parts[0], this.seq1), strings.Trim(parts[1], this.seq1)

	if len(preStr) == 0 {
		return nil
	}

	preOrder := []int{}
	inOrder := []int{}
	for _, n := range strings.Split(preStr, this.seq1) {
		num, _ := strconv.Atoi(n)
		preOrder = append(preOrder, num)
	}
	for _, n := range strings.Split(inStr, this.seq1) {
		num, _ := strconv.Atoi(n)
		inOrder = append(inOrder, num)
	}

	var buildTree func(preorder []int, inorder []int) *TreeNode
	buildTree = func(preorder, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}

		root := &TreeNode{
			Val: this.nodeSeqMap[preorder[0]],
		}

		// find left, right part
		for i := range inorder {
			if inorder[i] == preorder[0] {
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

	return buildTree(preOrder, inOrder)
}
*/

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

/*
 [3,2,4,3]
 failed here? humm
 mine: [3,null,2,null,3,null,4]
 expect: [3,2,4,3]

 oh.. I see..
the number are not unique two 3s messed up my tree build code
I need to use something else to mark the tree nodes
*/

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);

 was wondering why
[1,2,3,null,null,4,5]
[]
[3,2,4,3]
give
[0,0,0,null,null,0,0]
[]
[0,0,0,0]

yeah, you see it is two objects ser and deser
has to keep this somewhere else or some other way

maybe, turn the marker into decimal number
thus two 3 will be marked as
3.1
3.2
*/

type Codec struct {
	seq1 string
	seq2 string
}

func Constructor() Codec {
	return Codec{
		" ",
		",",
	}
}

type TreeNodeF struct {
	Val   float32
	Left  *TreeNodeF
	Right *TreeNodeF
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var mark func(root *TreeNode) *TreeNodeF
	marker := 1
	decimalize := func(marker int, value int) float32 {
		res := float32(0.0)
		if marker < 10 {
			res = float32(marker) / 10
		} else if marker < 100 {
			res = float32(marker) / 100
		} else if marker < 1000 {
			res = float32(marker) / 1000
		} else if marker < 10000 {
			res = float32(marker) / 10000
		} else if marker < 100000 {
			res = float32(marker) / 100000
		}

		if value < 0 {
			res = -res
		}

		return res
	}
	nextMarker := func(marker int) int {
		marker++
		if marker%10 == 0 {
			marker++
		}

		return marker
	}

	mark = func(root *TreeNode) *TreeNodeF {
		if root == nil {
			return nil
		}

		rootF := &TreeNodeF{
			Val: float32(root.Val) + decimalize(marker, root.Val),
		}
		marker = nextMarker(marker)

		rootF.Left = mark(root.Left)
		rootF.Right = mark(root.Right)

		return rootF
	}

	var preorder func(root *TreeNodeF) string
	var inorder func(root *TreeNodeF) string
	preorder = func(root *TreeNodeF) string {
		if root == nil {
			return ""
		}

		return fmt.Sprintf("%.5f", root.Val) + this.seq1 +
			preorder(root.Left) +
			preorder(root.Right)
	}
	inorder = func(root *TreeNodeF) string {
		if root == nil {
			return ""
		}

		return inorder(root.Left) +
			fmt.Sprintf("%.5f", root.Val) + this.seq1 +
			inorder(root.Right)
	}
	rootF := mark(root)
	return preorder(rootF) + this.seq2 + inorder(rootF)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, this.seq2)
	preStr, inStr := strings.Trim(parts[0], this.seq1), strings.Trim(parts[1], this.seq1)

	if len(preStr) == 0 {
		return nil
	}

	preOrder := []float32{}
	inOrder := []float32{}
	for _, n := range strings.Split(preStr, this.seq1) {
		value, _ := strconv.ParseFloat(n, 32)
		f := float32(value)
		preOrder = append(preOrder, f)
	}
	for _, n := range strings.Split(inStr, this.seq1) {
		value, _ := strconv.ParseFloat(n, 32)
		f := float32(value)
		inOrder = append(inOrder, f)
	}

	var buildTree func(preorder []float32, inorder []float32) *TreeNode
	buildTree = func(preorder, inorder []float32) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}

		root := &TreeNode{
			Val: int(preorder[0]),
		}

		// find left, right part
		for i := range inorder {
			if inorder[i] == preorder[0] {
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

	return buildTree(preOrder, inOrder)
}

/*
Runtime: 34 ms, faster than 29.97% of Go online submissions for Serialize and Deserialize Binary Tree.
Memory Usage: 8.6 MB, less than 25.69% of Go online submissions for Serialize and Deserialize Binary Tree.
*/

/*
ouch
    void serialize(TreeNode* root, ostringstream& out) {
        if(root == nullptr)  {
             out << "#" << " ";
             return;
        }


        out << root->val << " ";
        serialize(root->left, out);
        serialize(root->right, out);
    }

    TreeNode* deserialize(istringstream& in) {
        string data;
        TreeNode *root;

        in>>data;
        {
            if(data == "#")
                return nullptr;
            root = new TreeNode(stoi(data));
            root->left = deserialize(in);
            root->right =  deserialize(in);
        }

        return root;

    }

	yeah.. actually inorder or any order alone is enough as long as you can mark null ptr
	but I found my solution fun too
*/

func testSerialize() {
	codec1 := Constructor()
	root1 := &TreeNode{
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
	str1 := codec1.serialize(root1)
	fmt.Println(str1)
	droot1 := codec1.deserialize(str1)
	fmt.Println(codec1.serialize(droot1))
	fmt.Println("---")

	codec2 := Constructor()
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	str2 := codec2.serialize(root2)
	fmt.Println(str2)
	droot2 := codec2.deserialize(str2)
	fmt.Println(codec2.serialize(droot2))
	fmt.Println("---")
}
