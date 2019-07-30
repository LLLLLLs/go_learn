// Time        : 2019/07/30
// Description :

package reorder_143

import . "go_learn/leecode/linked_list/base"

// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
//
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
// Example 1:
// Given 1->2->3->4, reorder it to 1->4->2->3.
//
// Example 2:
// Given 1->2->3->4->5, reorder it to 1->5->2->4->3.

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	length := 0
	nodes := make([]*ListNode, 0)
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
		length++
	}
	for i := 0; i < length/2; i++ {
		nodes[length-i-1].Next = nodes[i].Next
		nodes[i].Next = nodes[length-i-1]
	}
	nodes[length/2].Next = nil
}

func reorderList2(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	l1, l2 := head, slow.Next
	slow.Next = nil
	// reverse l2
	pre := l2
	cur := l2.Next
	pre.Next = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	l2 = pre
	// merge
	for l2 != nil {
		tmp := l2.Next
		l2.Next = l1.Next
		l1.Next = l2
		l2 = tmp
		l1 = l1.Next.Next
	}
}
