// Time        : 2019/07/26
// Description :

package cycle_II_142

import . "golearn/leecode/linked_list/base"

// Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
//
// To represent a cycle in the given linked list,
// we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to.
// If pos is -1, then there is no cycle in the linked list.
//
// Note: Do not modify the linked list.
//
// Example 1:
//
// Input: head = [3,2,0,-4], pos = 1
// Output: tail connects to node index 1
// Explanation: There is a cycle in the linked list, where tail connects to the second node.
//
// Example 2:
//
// Input: head = [1,2], pos = 0
// Output: tail connects to node index 0
// Explanation: There is a cycle in the linked list, where tail connects to the first node.
//
// Example 3:
//
// Input: head = [1], pos = -1
// Output: no cycle
// Explanation: There is no cycle in the linked list.
//
// Follow-up:
// Can you solve it without using extra space?

func detectCycle(head *ListNode) *ListNode {
	one, two := head, head
	for one != nil && two != nil && two.Next != nil {
		one = one.Next
		two = two.Next.Next
		if one == two {
			break
		}
	}
	if two == nil || two.Next == nil {
		return nil
	}
	for head != two {
		head = head.Next
		two = two.Next
	}
	return head
}
