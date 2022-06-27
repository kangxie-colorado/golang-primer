// https://leetcode.com/problems/merge-k-sorted-lists/

package main

/*
I am not sure why this is a hard question
just merge them one by one
*/

func _linear_mergeKLists(lists []*ListNode) *ListNode {
	mergeTwo := func(l1, l2 *ListNode) *ListNode {
		var prev *ListNode
		dummy := ListNode{Next: nil}
		prev = &dummy

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				prev.Next = l1
				prev = l1
				l1 = l1.Next
			} else {
				prev.Next = l2
				prev = l2
				l2 = l2.Next
			}
		}

		if l1 != nil {
			prev.Next = l1
		}

		if l2 != nil {
			prev.Next = l2
		}

		return dummy.Next
	}

	var res *ListNode
	res = nil

	for _, l := range lists {
		res = mergeTwo(res, l)
	}

	return res
}

/*
Runtime: 110 ms, faster than 36.55% of Go online submissions for Merge k Sorted Lists.
Memory Usage: 5.3 MB, less than 71.72% of Go online submissions for Merge k Sorted Lists.
hmm.. obviously I am at the 2nd tier..

there is one that is more efficient.
is it divide and merge?
*/

func mergeKLists(lists []*ListNode) *ListNode {
	mergeTwo := func(l1, l2 *ListNode) *ListNode {
		var prev *ListNode
		dummy := ListNode{Next: nil}
		prev = &dummy

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				prev.Next = l1
				prev = l1
				l1 = l1.Next
			} else {
				prev.Next = l2
				prev = l2
				l2 = l2.Next
			}
		}
		if l1 != nil {
			prev.Next = l1
		}
		if l2 != nil {
			prev.Next = l2
		}
		return dummy.Next
	}

	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	firstHalf := mergeKLists(lists[:len(lists)/2])
	secondHalf := mergeKLists(lists[len(lists)/2:])

	return mergeTwo(firstHalf, secondHalf)
}

/*
Runtime: 11 ms, faster than 77.82% of Go online submissions for Merge k Sorted Lists.
Memory Usage: 5.3 MB, less than 71.72% of Go online submissions for Merge k Sorted Lists.

probably this is the other tier
*/
