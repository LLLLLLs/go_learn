// Time        : 2019/06/26
// Description :

package merge_two_sorted_list_21

import (
	"go_learn/leecode/linked_list/base"
)

// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
// Example:
//
// Input: 1->2->4, 1->3->4
//
// Output: 1->1->2->3->4->4

// AC
func merge(l1, l2 *base.ListNode) *base.ListNode {
	var head = base.NewList(0)
	var cur = head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return head.Next
}
