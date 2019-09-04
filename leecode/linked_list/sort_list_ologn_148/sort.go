// Time        : 2019/07/30
// Description :

package sort_list_ologn_148

import . "golearn/leecode/linked_list/base"

// Sort a linked list in O(n log n) time using constant space complexity.
//
// Example 1:
//
// Input: 4->2->1->3
// Output: 1->2->3->4
// Example 2:
//
// Input: -1->5->3->4->0
// Output: -1->0->3->4->5

func sortList(head *ListNode) *ListNode {
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	return sort(head, length)
}

func sort(head *ListNode, length int) *ListNode {
	if length <= 1 {
		return head
	}
	l2 := head
	for i := 0; i < length/2-1; i++ {
		l2 = l2.Next
	}
	newL2 := sort(l2.Next, length-length/2)
	l2.Next = nil
	newL1 := sort(head, length/2)
	return merge(newL1, newL2)
}

func merge(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	head.Next = l1
	cur := head
	var next *ListNode
	for ; l2 != nil && cur != nil; l2 = next {
		next = l2.Next
		for cur.Next != nil && cur.Next.Val < l2.Val {
			cur = cur.Next
		}
		if cur.Next == nil {
			cur.Next = l2
			break
		}
		l2.Next = cur.Next
		cur.Next = l2
	}
	return head.Next
}
