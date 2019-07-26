// Time        : 2019/07/24
// Description :

package sum_root_to_leaf_129

import . "go_learn/leecode/tree/base"

// Given a binary tree containing digits from 0-9 only,
// each root-to-leaf path could represent a number.
//
// An example is the root-to-leaf path 1->2->3 which represents the number 123.
//
// Find the total sum of all root-to-leaf numbers.
//
// Note: A leaf is a node with no children.
//
// Example:
//
// Input: [1,2,3]
//     1
//    / \
//   2   3
// Output: 25
// Explanation:
// The root-to-leaf path 1->2 represents the number 12.
// The root-to-leaf path 1->3 represents the number 13.
// Therefore, sum = 12 + 13 = 25.
// Example 2:
//
// Input: [4,9,0,5,1]
//     4
//    / \
//   9   0
//  / \
// 5   1
// Output: 1026
// Explanation:
// The root-to-leaf path 4->9->5 represents the number 495.
// The root-to-leaf path 4->9->1 represents the number 491.
// The root-to-leaf path 4->0 represents the number 40.
// Therefore, sum = 495 + 491 + 40 = 1026.

func sumNumbers(root *TreeNode) int {
	sum := 0
	getNums(root, 0, &sum)
	return sum
}

func getNums(root *TreeNode, mid int, sum *int) {
	if root == nil {
		return
	}
	mid = mid*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += mid
		return
	}
	getNums(root.Left, mid, sum)
	getNums(root.Right, mid, sum)
}
