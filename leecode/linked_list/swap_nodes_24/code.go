// Time        : 2019/06/26
// Description :

package swap_nodes_24

import (
	"go_learn/leecode/linked_list/base"
)

// Given a linked list, swap every two adjacent nodes and return its head.
//
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
// Example:
//
// Given 1->2->3->4, you should return the list as 2->1->4->3.

// AC
func swapPairs(head *base.ListNode) *base.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head.Next
	head.Next = swapPairs(tmp.Next)
	tmp.Next = head
	return tmp
}
