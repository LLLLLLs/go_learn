// Time        : 2019/07/19
// Description :

package recover_binary_search_99

import (
	. "go_learn/leecode/tree/base"
	"math"
)

//Two elements of a binary search tree (BST) are swapped by mistake.
//
//Recover the tree without changing its structure.
//
//Example 1:
//
//Input: [1,3,null,null,2]
//
//   1
//  /
// 3
//  \
//   2
//
//Output: [3,1,null,null,2]
//
//   3
//  /
// 1
//  \
//   2
//Example 2:
//
//Input: [3,1,4,null,null,2]
//
//  3
// / \
//1   4
//   /
//  2
//
//Output: [2,1,4,null,null,3]
//
//  2
// / \
//1   4
//   /
//  3

func recoverTree(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	pop := func() *TreeNode {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return node
	}
	prev := &TreeNode{}
	var maxN, minN *TreeNode
	min := int64(math.MinInt64)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = pop()
		if int64(root.Val) <= min {
			if maxN == nil {
				maxN = prev
				minN = root
				min = int64(root.Val)
			} else {
				minN = root
				maxN.Val, root.Val = root.Val, maxN.Val
				return
			}
		} else {
			prev = root
			min = int64(root.Val)
		}
		root = root.Right
	}
	maxN.Val, minN.Val = minN.Val, maxN.Val
}
