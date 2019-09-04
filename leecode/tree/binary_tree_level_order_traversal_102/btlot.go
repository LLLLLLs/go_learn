// Time        : 2019/07/19
// Description :

package binary_tree_level_order_traversal_102

import . "golearn/leecode/tree/base"

// Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
//
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its level order traversal as:
// [
//   [3],
//   [9,20],
//   [15,7]
// ]

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	preorder(root, 0, &result)
	return result
}

func preorder(root *TreeNode, index int, result *[][]int) {
	if root == nil {
		return
	}
	if index > len(*result)-1 {
		*result = append(*result, []int{})
	}
	(*result)[index] = append((*result)[index], root.Val)
	preorder(root.Left, index+1, result)
	preorder(root.Right, index+1, result)
}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	levels := [2][]*TreeNode{{root}}
	for len(levels[0]) != 0 {
		result = append(result, []int{})
		for i := range levels[0] {
			result[len(result)-1] = append(result[len(result)-1], levels[0][i].Val)
			if levels[0][i].Left != nil {
				levels[1] = append(levels[1], levels[0][i].Left)
			}
			if levels[0][i].Right != nil {
				levels[1] = append(levels[1], levels[0][i].Right)
			}
		}
		levels[0], levels[1] = levels[1], []*TreeNode{}
	}
	return result
}
