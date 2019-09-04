// Time        : 2019/07/22
// Description :

package restore_from_pre_inorder_105

import . "golearn/leecode/tree/base"

// Given preorder and inorder traversal of a tree, construct the binary tree.
//
// Note:
// You may assume that duplicates do not exist in the tree.
//
// For example, given
//
// preorder = [3,9,20,15,7]
// inorder = [9,3,15,20,7]
// Return the following binary tree:
//
//     3
//    / \
//   9  20
//     /  \
//    15   7

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	for i := range inorder {
		if inorder[i] == root.Val {
			root.Left = buildTree(preorder[1:i+1], inorder[:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}
	return root
}
