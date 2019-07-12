// Time        : 2019/07/12
// Description :

package reverse_II_92

import . "go_learn/leecode/linked_list/base"

// Reverse a linked list from position m to n. Do it in one-pass.
//
// Note: 1 ≤ m ≤ n ≤ length of list.
//
// Example:
//
// Input: 1->2->3->4->5->NULL, m = 2, n = 4
// Output: 1->4->3->2->5->NULL

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	mid := &ListNode{}
	tail := mid
	result := &ListNode{}
	result.Next = head
	resultBegin := result
	cur := head
	for index := 1; index <= n; index++ {
		if index < m {
			resultBegin.Next = cur
			resultBegin = resultBegin.Next
			cur = cur.Next
			continue
		}
		if index == m {
			tail = cur
		}
		tmp := cur.Next
		cur.Next = mid
		mid = cur
		cur = tmp
	}
	resultBegin.Next = mid
	tail.Next = cur
	return result.Next
}
