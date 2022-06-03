package main

import "fmt"

// https://leetcode.com/problems/rotate-list/

/** rotate the list to right by K places

	I think, first look at the contraints since the problem is easily understood

	Constraints:
	- The number of nodes in the list is in the range [0, 500].
	- -100 <= Node.val <= 100
	- 0 <= k <= 2 * 10^9

	the k can be really high but the list is not really long
	because of that.. so I can actually walk thru the list to get a length
	then I can module it down to <500 rotations

	and actually I may convert this into an array then copy/paste it out.. that seems super easy
	yeah, let me code this quick solution up

**/

func _rotateRight(head *ListNode, k int) *ListNode {
	listAsSlice := make([]int, 500)

	len := 0
	headCpy := head
	for headCpy != nil {
		listAsSlice[len] = headCpy.Val
		headCpy = headCpy.Next
		len++
	}

	if len == 0 {
		return head
	}

	rotates := k % len

	if rotates == 0 {
		return head
	}

	// re-construct the list
	newList := ListNode{
		-101,
		nil,
	}
	nextNodePtr := &newList
	for i := rotates; i > 0; i-- {
		newNodePtr := &ListNode{
			listAsSlice[len-i],
			nil,
		}

		nextNodePtr.Next = newNodePtr
		nextNodePtr = newNodePtr
	}
	for i := 0; i < len-rotates; i++ {
		newNodePtr := &ListNode{
			listAsSlice[i],
			nil,
		}

		nextNodePtr.Next = newNodePtr
		nextNodePtr = newNodePtr
	}

	return newList.Next
}

/**
	Success
	Details
	Runtime: 3 ms, faster than 53.55% of Go online submissions for Rotate List.
	Memory Usage: 2.7 MB, less than 6.72% of Go online submissions for Rotate List.

	not really bad for first solution
	looking at it again, you don't need to keep the slice around, you just need a len, which is absolutely necessary in this constraints setting

	write the pseudo-code

roateRight(head, k)
	if head == nil:
		return head

	headCpy = head
	tail = head # keep a tail, so I can build up the list very quickly
	len = 0
	for headCpy != nil
		len++
		tail = headCpy
		headCpy = headCpy.Next


	rotates = k%len
	if rotates == 0
		return head

	# fine the roate head
	rotateHead = head
	for len-rotates>0 {
		rotateHead = rotateHead.Next
		len--
	}

	tail.Next = head
	return rotateHead


**/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	headCpy := head
	tail := head
	len := 0
	for headCpy != nil {
		len++
		tail = headCpy
		headCpy = headCpy.Next
	}

	rotates := k % len
	if rotates == 0 {
		return head
	}

	roateHead := head
	newTail := head
	for len > rotates {
		newTail = roateHead
		roateHead = roateHead.Next
		len--
	}

	// rotateHead will still be conncted to left nodes, and it will form a loop
	// it will OOM, so we need to cut the connection

	newTail.Next = nil

	tail.Next = head
	return roateHead

}

/**
Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate List.
Memory Usage: 2.6 MB, less than 13.45% of Go online submissions for Rotate List.
**/

func testRotateRight() {

	list1 := &ListNode{-101, nil}
	tail := list1
	for i := 0; i < 5; i++ {
		tail.Next = &ListNode{
			i,
			nil,
		}
		tail = tail.Next
	}

	fmt.Println(rotateRight(list1.Next, 2))

}
