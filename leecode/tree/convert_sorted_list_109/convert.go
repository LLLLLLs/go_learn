// Time        : 2019/07/22
// Description :

package convert_sorted_list_109

import (
	. "go_learn/leecode/linked_list/base"
	. "go_learn/leecode/tree/base"
)

// Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
//
// For this problem, a height-balanced binary tree is defined as a binary tree
// in which the depth of the two subtrees of every node never differ by more than 1.
//
// Example:
//
// Given the sorted linked list: [-10,-3,0,5,9],
//
// One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
//
//       0
//      / \
//    -3   9
//    /   /
//  -10  5

func sortedListToBST(head *ListNode) *TreeNode {
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	return _sortedListToBST(head, length)
}

func _sortedListToBST(head *ListNode, length int) *TreeNode {
	if length == 0 {
		return nil
	}
	cur := head
	for i := 0; i < length/2; i++ {
		cur = cur.Next
	}
	root := &TreeNode{
		Val:   cur.Val,
		Left:  _sortedListToBST(head, length/2),
		Right: _sortedListToBST(cur.Next, length/2-(length+1)%2),
	}
	return root
}
