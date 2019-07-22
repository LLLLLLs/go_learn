// Time        : 2019/07/22
// Description :

package balanced_binary_tree_judge_110

import . "go_learn/leecode/tree/base"

// Given a binary tree, determine if it is height-balanced.
//
// For this problem, a height-balanced binary tree is defined as:
//
// a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
// Example 1:
//
// Given the following tree [3,9,20,null,null,15,7]:
//
//     3
//    / \
//   9  20
//     /  \
//    15   7
// Return true.
//
// Example 2:
//
// Given the following tree [1,2,2,3,3,null,null,4,4]:
//
//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4
// Return false.

func isBalanced(root *TreeNode) bool {
	_, ok := _isBalanced(root)
	return ok
}

func _isBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	lh, lok := _isBalanced(root.Left)
	if !lok {
		return 0, false
	}
	rh, rok := _isBalanced(root.Right)
	if !rok {
		return 0, false
	}
	if abs(lh-rh) > 1 {
		return 0, false
	}
	return max(lh, rh) + 1, true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
