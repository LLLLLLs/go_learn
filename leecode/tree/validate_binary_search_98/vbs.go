// Time        : 2019/07/19
// Description :

package validate_binary_search_98

import (
	. "go_learn/leecode/tree/base"
	"math"
)

//Given a binary tree, determine if it is a valid binary search tree (BST).
//
//Assume a BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than the node's key.
//The right subtree of a node contains only nodes with keys greater than the node's key.
//Both the left and right subtrees must also be binary search trees.
//
//Example 1:
//
//    2
//   / \
//  1   3
//
//Input: [2,1,3]
//Output: true
//Example 2:
//
//    5
//   / \
//  1   4
//     / \
//    3   6
//
//Input: [5,1,4,null,null,3,6]
//Output: false
//Explanation: The root node's value is 5 but its right child's value is 4.

func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	pop := func() *TreeNode {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return node
	}
	min := int64(math.MinInt64)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = pop()
		if int64(root.Val) <= min {
			return false
		}
		min = int64(root.Val)
		root = root.Right
	}
	return true
}
