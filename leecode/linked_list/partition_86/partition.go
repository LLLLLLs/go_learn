// Time        : 2019/07/11
// Description :

package partition_86

import . "golearn/leecode/linked_list/base"

// Given a linked list and a value x,
// partition it such that all nodes less than x come before nodes greater than or equal to x.
//
// You should preserve the original relative order of the nodes in each of the two partitions.
//
// Example:
//
// Input: head = 1->4->3->2->5->2, x = 3
// Output: 1->2->2->4->3->5

func partition(head *ListNode, x int) *ListNode {
	less, cur := &ListNode{}, head
	less.Next = head
	result := less
	prev := less
	continuous := true
	for cur != nil {
		if cur.Val < x {
			if continuous {
				less.Next = cur
				less = less.Next
				prev = cur
				cur = cur.Next
			} else {
				prev.Next = cur.Next
				cur.Next = less.Next
				less.Next = cur
				cur = prev.Next
				less = less.Next
			}
		} else {
			prev = cur
			cur = cur.Next
			continuous = false
		}
	}
	return result.Next
}
