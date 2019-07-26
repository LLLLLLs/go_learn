// Time        : 2019/07/23
// Description :

package path_max_sum_124

import (
	. "go_learn/leecode/tree/base"
	"math"
)

// Given a non-empty binary tree, find the maximum path sum.
//
// For this problem, a path is defined as any sequence of nodes
// from some starting node to any node in the tree along the parent-child connections.
// The path must contain at least one node and does not need to go through the root.
//
// Example 1:
//
// Input: [1,2,3]
//
//        1
//       / \
//      2   3
//
// Output: 6
// Example 2:
//
// Input: [-10,9,20,null,null,15,7]
//
//    -10
//    / \
//   9  20
//     /  \
//    15   7
//
// Output: 42

func maxPathSum(root *TreeNode) int {
	_, max := _maxPathSum(root)
	return max
}

func _maxPathSum(root *TreeNode) (int, int) {
	if root == nil {
		return 0, math.MinInt32
	}
	left, lMax := _maxPathSum(root.Left)
	right, rMax := _maxPathSum(root.Right)
	val := root.Val
	if left > 0 {
		val += left
	}
	left += root.Val
	if right > 0 {
		val += right
	}
	right += root.Val
	return max(root.Val, max(left, right)), max(val, max(lMax, rMax))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
