// Time        : 2019/07/04
// Description :

package rotate_61

import (
	"golearn/leecode/linked_list/base"
)

// Given a linked list, rotate the list to the right by k places, where k is non-negative.
//
// Example 1:
//
// Input: 1->2->3->4->5->NULL, k = 2
// Output: 4->5->1->2->3->NULL
// Explanation:
// rotate 1 steps to the right: 5->1->2->3->4->NULL
// rotate 2 steps to the right: 4->5->1->2->3->NULL
// Example 2:
//
// Input: 0->1->2->NULL, k = 4
// Output: 2->0->1->NULL
// Explanation:
// rotate 1 steps to the right: 2->0->1->NULL
// rotate 2 steps to the right: 1->2->0->NULL
// rotate 3 steps to the right: 0->1->2->NULL
// rotate 4 steps to the right: 2->0->1->NULL

func rotateRight(head *base.ListNode, k int) *base.ListNode {
	if head == nil {
		return head
	}
	tail := head
	length := 1
	for tail.Next != nil {
		length++
		tail = tail.Next
	}
	if k%length == 0 {
		return head
	}
	tmp := head
	var pre *base.ListNode
	for i := 0; i < length-k%length; i++ {
		pre = tmp
		tmp = tmp.Next
	}
	tail.Next = head
	pre.Next = nil
	return tmp
}
