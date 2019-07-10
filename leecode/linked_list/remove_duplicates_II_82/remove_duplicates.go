// Time        : 2019/07/10
// Description :

package remove_duplicates_II_82

import (
	. "go_learn/leecode/linked_list/base"
)

// Given a sorted linked list, delete all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list.
//
// Example 1:
//
// Input: 1->2->3->3->4->4->5
// Output: 1->2->5
// Example 2:
//
// Input: 1->1->1->2->3
// Output: 2->3

func deleteDuplicates(head *ListNode) *ListNode {
	h := &ListNode{}
	cur := h
	var prev *ListNode
	for head != nil {
		if (head.Next == nil || head.Val != head.Next.Val) && (prev == nil || head.Val != prev.Val) {
			cur.Next = head
			cur = cur.Next
		}
		prev = head
		head = head.Next
	}
	cur.Next = nil
	return h.Next
}
