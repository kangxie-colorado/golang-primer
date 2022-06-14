// https://leetcode.com/problems/reverse-linked-list-ii/

package main

/*
one pass

so remember a few position while going thru the list

0. before l, business as usual
1. when num reach l, remember this node as nodeL; it will be used to connect what is after r
	and the stack begin, push the node in
2. between l and r, incoming just hang it next point to stack top
3. at r, prev->next should point to it; also this node should hang its next for stack top
4. after r, nodeL should point to this node and set prev to node
5. then business as usual

problem is l=r... then I should just the list itself
my above procedure cannot hanlde it.. ??

*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if left == right {
		return head
	}

	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	prev := dummy
	node := dummy
	L := dummy
	reverseHead := dummy

	num := 0

	for node != nil {
		theNext := node.Next
		if num < left {

			prev = node
		} else if num == left {
			L = node
			reverseHead = node
			// if not setting to nil, there will be loop and memory error
			L.Next = nil

		} else if num > left && num < right {
			node.Next = reverseHead
			reverseHead = node
		} else if num == right {
			node.Next = reverseHead
			prev.Next = node
			prev = L
		} else if num > right {
			prev.Next = node
			prev = node
		}

		node = theNext
		num++
	}

	return dummy.Next
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List II.
Memory Usage: 2 MB, less than 66.28% of Go online submissions for Reverse Linked List II.
*/

func testReverseBetween() {

}
