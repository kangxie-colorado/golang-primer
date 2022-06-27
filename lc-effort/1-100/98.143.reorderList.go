// https://leetcode.com/problems/reorder-list/
package main

/*
not hard to see this is
1. find the half, and reverse 2nd half; this can be done using a slow/fast pointer
2. rever the 2nd half
3. merge two lists together
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	dummy := ListNode{Next: head}
	slow, fast := &dummy, &dummy

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reverseList := func(head *ListNode) *ListNode {
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

	leftHalf := dummy.Next
	rightHalf := reverseList(slow.Next)
	slow.Next = nil

	mergeTwo := func(l1, l2 *ListNode) *ListNode {
		var prev *ListNode
		dummy := ListNode{Next: nil}
		prev = &dummy

		choices := []*ListNode{l1, l2}
		which := 0

		for choices[which] != nil {
			prev.Next = choices[which]
			prev = prev.Next
			choices[which] = choices[which].Next
			which ^= 1
		}
		which ^= 1
		prev.Next = choices[which]

		return dummy.Next
	}

	dummy.Next = mergeTwo(leftHalf, rightHalf)
}

/*
Runtime: 10 ms, faster than 76.92% of Go online submissions for Reorder List.
Memory Usage: 6 MB, less than 40.72% of Go online submissions for Reorder List.

Runtime: 8 ms, faster than 85.52% of Go online submissions for Reorder List.
Memory Usage: 6.1 MB, less than 32.35% of Go online submissions for Reorder List.
*/

func testReorderList() {

	l5 := ListNode{5, nil}
	l4 := ListNode{4, &l5}
	l3 := ListNode{3, &l4}
	l2 := ListNode{2, &l3}
	l1 := ListNode{1, &l2}

	reorderList(&l1)

}

func main() {
	testReorderList()
}
