// Time        : 2019/07/30
// Description :

package postorder_traversal_iteratively_145

import . "go_learn/leecode/tree/base"

// Given a binary tree, return the postorder traversal of its nodes' values.
//
// Example:
//
// Input: [1,null,2,3]
//    1
//     \
//      2
//     /
//    3
//
// Output: [3,2,1]
// Follow up: Recursive solution is trivial, could you do it iteratively?

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || root != nil {
		for root == nil {
			if len(stack) == 0 {
				goto outLoop
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, root.Left)
		stack = append(stack, root.Right)
		res = append(res, root.Val)
		root = nil
	}
outLoop:
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
