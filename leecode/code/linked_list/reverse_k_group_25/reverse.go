// Time        : 2019/01/23
// Description :

package reverse_k_group_25

import (
	. "go_learn/leecode/code/linked_list/base"
)

//Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
//k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
//Example:
//
//Given this linked list: 1->2->3->4->5
//
//For k = 2, you should return: 2->1->4->3->5
//
//For k = 3, you should return: 3->2->1->4->5
//
//Note:
//
//Only constant extra memory is allowed.
//You may not alter the values in the list's nodes, only nodes itself may be changed.

func reverseKGroup1(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	now := head
	i := 0
	for ; i < k && now != nil; i++ {
		now = now.Next
	}
	if i != k {
		return head
	} else {
		tail := head
		next := head.Next
		for i = 1; i < k; i++ {
			tmp := next.Next
			next.Next = head
			head = next
			next = tmp
		}
		tail.Next = reverseKGroup1(now, k)
		return head
	}
}

// 递归尾调用
func reverseKGroup2(head *ListNode, k int) *ListNode {
	list := &ListNode{}
	list.Next = head
	return reverseWithTail(list, list, k, true)
}

func reverseWithTail(head, tail *ListNode, k int, first bool) *ListNode {
	mid := tail.Next
	if mid == nil || k <= 1 {
		return head
	}
	now := mid
	i := 0
	for ; i < k && now != nil; i++ {
		now = now.Next
	}
	if i != k {
		return head
	} else {
		thisTail := mid
		next := mid.Next
		tmp := next
		for i = 1; i < k; i++ {
			tmp = tmp.Next
			next.Next = mid
			mid = next
			next = tmp
		}
		if first {
			head.Next = mid
		}
		tail.Next = mid
		return reverseWithTail(head, thisTail, k, false)
	}
}

// AC
func reverseKGroup(head *ListNode, k int) *ListNode {
	// check
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	tail := head
	mid := head
	cur = head.Next
	for i := 0; i < k-1; i++ {
		tmp := cur.Next
		cur.Next = mid
		mid = cur
		cur = tmp
	}
	tail.Next = reverseKGroup(cur, k)
	return mid
}
