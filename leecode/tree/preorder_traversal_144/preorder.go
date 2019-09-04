// Time        : 2019/07/12
// Description : 先序遍历 中左右

package preorder_traversal_144

import . "golearn/leecode/tree/base"

func preorderTraversal(root *TreeNode) []int {
	io := make([]int, 0)
	preorder(root, &io)
	return io
}

func preorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	preorder(root.Left, result)
	preorder(root.Right, result)
}
