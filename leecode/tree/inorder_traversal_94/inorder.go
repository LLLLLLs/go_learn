// Time        : 2019/07/12
// Description : 二叉树中序遍历 左->中->右

package inorder_traversal_94

import . "golearn/leecode/tree/base"

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}
