// Time        : 2019/07/30
// Description :

package sort_list_insertion_147

import . "golearn/leecode/linked_list/base"

// Sort a linked list using insertion sort.
//
// A graphical example of insertion sort.
// The partial sorted list (black) initially contains only the first element in the list.
// With each iteration one element (red) is removed from the input data and inserted in-place into the sorted list
//
// Algorithm of Insertion Sort:
//
// Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
// At each iteration, insertion sort removes one element from the input data,
// finds the location it belongs within the sorted list, and inserts it there.
// It repeats until no input elements remain.
//
// Example 1:
//
// Input: 4->2->1->3
// Output: 1->2->3->4
//
// Example 2:
// Input: -1->5->3->4->0
// Output: -1->0->3->4->5

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var cur, next, cmp, end *ListNode
	end = head
	for cur, head.Next = head.Next, nil; cur != nil; cur = next {
		next = cur.Next
		if cur.Val > end.Val {
			end.Next = cur
			end = cur
			cur.Next = nil
			continue
		}
		if cur.Val < head.Val {
			cur.Next = head
			head = cur
			continue
		}
		cmp = head
		for cmp.Next != nil && cur.Val > cmp.Next.Val {
			cmp = cmp.Next
		}
		cur.Next = cmp.Next
		cmp.Next = cur
	}
	return head
}
