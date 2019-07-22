// Time        : 2019/07/22
// Description :

package path_sum_II_113

import . "go_learn/leecode/tree/base"

// Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
//
// Note: A leaf is a node with no children.
//
// Example:
//
// Given the below binary tree and sum = 22,
//
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \    / \
// 7    2  5   1
// Return:
//
// [
//    [5,4,11,2],
//    [5,8,4,5]
// ]

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	mid := make([]int, 0, 100)
	_pathSum(root, sum, mid, &result)
	return result
}

func _pathSum(root *TreeNode, sum int, mid []int, result *[][]int) {
	if root == nil {
		return
	}
	mid = append(mid, root.Val)
	length := len(mid)
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		*result = append(*result, append([]int{}, mid...))
	}
	_pathSum(root.Left, sum, mid, result)
	mid = mid[:length]
	_pathSum(root.Right, sum, mid, result)
	mid = mid[:length]
}
