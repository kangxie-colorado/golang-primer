// https://leetcode.com/problems/reverse-linked-list/

package main

/*
iterative or recursive
or just use a stack
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func _iterative_reverseList(head *ListNode) *ListNode {
	dummy := ListNode{
		Next: head,
	}

	var tail *ListNode
	tail = nil

	for head != nil {
		nextHead := head.Next
		dummy.Next = head
		head.Next = tail
		tail = head
		head = nextHead
	}

	return dummy.Next
}

/*
Runtime: 5 ms, faster than 29.30% of Go online submissions for Reverse Linked List.
Memory Usage: 2.6 MB, less than 79.13% of Go online submissions for Reverse Linked List.

Runtime: 3 ms, faster than 64.80% of Go online submissions for Reverse Linked List.
Memory Usage: 2.6 MB, less than 79.13% of Go online submissions for Reverse Linked List.

now recursive
*/

func _recursive_reverseList(head *ListNode) *ListNode {

	tail := ListNode{
		Next: nil,
	}

	var helper func(head *ListNode) *ListNode
	helper = func(head *ListNode) *ListNode {

		if head == nil {
			return &tail
		}
		subList := helper(head.Next) // remember to call helper() not outer function here...
		head.Next = nil
		subList.Next = head

		return head
	}

	helper(head)

	return tail.Next
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
Memory Usage: 2.9 MB, less than 6.45% of Go online submissions for Reverse Linked List.
*/

/*
okay
just using a stack


a little try and it is actually the iterative
*/
