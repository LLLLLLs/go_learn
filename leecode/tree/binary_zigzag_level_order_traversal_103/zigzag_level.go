// Time        : 2019/07/22
// Description :

package binary_zigzag_level_order_traversal_103

import . "golearn/leecode/tree/base"

// Given a binary tree, return the zigzag level order traversal of its nodes' values.
// (ie, from left to right, then right to left for the next level and alternate between).
//
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its zigzag level order traversal as:
// [
//   [3],
//   [20,9],
//   [15,7]
// ]

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	nodes := [2][]*TreeNode{{root}}
	left := true
	for len(nodes[0]) != 0 {
		list := make([]int, 0)
		for i := len(nodes[0]) - 1; i >= 0; i-- {
			if nodes[0][i] == nil {
				continue
			}
			list = append(list, nodes[0][i].Val)
			if left {
				nodes[1] = append(nodes[1], nodes[0][i].Left, nodes[0][i].Right)
			} else {
				nodes[1] = append(nodes[1], nodes[0][i].Right, nodes[0][i].Left)
			}
		}
		if len(list) != 0 {
			result = append(result, list)
			left = !left
		}
		nodes[0], nodes[1] = nodes[1], []*TreeNode{}
	}
	return result
}
