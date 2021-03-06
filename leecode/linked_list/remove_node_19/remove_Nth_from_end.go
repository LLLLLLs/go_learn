// Time        : 2019/01/22
// Description :

package remove_node_19

import (
	"golearn/leecode/linked_list/base"
)

// Given linked list: 1->2->3->4->5, and n = 2.
//
// After removing the second node from the end, the linked list becomes 1->2->3->5.
//
// do this in one pass

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func RemoveNthFromEnd(head *base.ListNode, n int) *base.ListNode {
	lm := make(map[int]*base.ListNode)
	index := 0
	ll := head
	for ll != nil {
		lm[index] = ll
		index++
		ll = ll.Next
	}
	if n == len(lm) {
		head = head.Next
	} else {
		lm[len(lm)-n-1].Next = lm[len(lm)-n].Next
	}
	return head
}
