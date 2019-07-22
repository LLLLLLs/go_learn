// Time        : 2019/07/22
// Description :

package binary_tree_level_order_traversal_107

import . "go_learn/leecode/tree/base"

// Given a binary tree, return the bottom-up level order traversal of its nodes' values.
//  (ie, from left to right, level by level from leaf to root).
//
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its bottom-up level order traversal as:
// [
//   [15,7],
//   [9,20],
//   [3]
// ]

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	traversal(root, 0, &result)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}

func traversal(root *TreeNode, index int, result *[][]int) {
	if root == nil {
		return
	}
	traversal(root.Left, index+1, result)
	for len(*result) < index+1 {
		*result = append(*result, []int{})
	}
	(*result)[index] = append((*result)[index], root.Val)
	traversal(root.Right, index+1, result)
}
