// Time        : 2019/07/12
// Description : 后序遍历

package postorder_traversal

import . "go_learn/leecode/tree/base"

// 后序遍历 左右中
func postorderTraversal(root *TreeNode) []int {
	io := make([]int, 0)
	postorder(root, &io)
	return io
}

func postorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, result)
	postorder(root.Right, result)
	*result = append(*result, root.Val)
}
