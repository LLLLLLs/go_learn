// Time        : 2019/07/22
// Description :

package flatten_binary_to_linked_list_114

import . "go_learn/leecode/tree/base"

// Given a binary tree, flatten it to a linked list in-place.
//
// For example, given the following tree:
//
//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
// The flattened tree should look like:
//
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		left := root.Left
		for left.Right != nil {
			left = left.Right
		}
		left.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}
