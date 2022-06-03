// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func walk(p, c *ListNode, n int, backout *int) {
	if c == nil {
		*backout = 0
		return
	}
	walk(p.Next, c.Next, n, backout)
	*backout++

	if *backout == n {
		p.Next = c.Next
		c = nil
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{-1, head}
	backout := 0
	walk(dummy, head, n, &backout)

	return dummy.Next

}

/*
Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
Memory Usage: 2.2 MB, less than 32.50% of Go online submissions for Remove Nth Node From End of List.
Next challenges:

Apparently the standard one-pass solutin is fast/slow pointer
also open my mind... the faster is not limited to walk 2 over 1(slow)

it can walk n+1 step in advance...
early pointer/late pointer...

open my mind...
*/

func testRemoveNthFromEnd() {
	removeNthFromEnd(&ListNode{1, nil}, 1)
}
