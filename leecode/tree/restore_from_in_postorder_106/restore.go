// Time        : 2019/07/22
// Description :

package restore_from_in_postorder_106

import . "go_learn/leecode/tree/base"

// Given inorder and postorder traversal of a tree, construct the binary tree.
//
// Note:
// You may assume that duplicates do not exist in the tree.
//
// For example, given
//
// inorder = [9,3,15,20,7]
// postorder = [9,15,7,20,3]
// Return the following binary tree:
//
//     3
//    / \
//   9  20
//     /  \
//    15   7

func buildTree(inorder []int, postorder []int) *TreeNode {
	return _buildTree(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func _buildTree(inorder []int, iBegin, iEnd int, postorder []int, pBegin, pEnd int) *TreeNode {
	if iEnd < iBegin {
		return nil
	}
	root := &TreeNode{Val: postorder[pEnd]}
	for left := 0; left <= iEnd-iBegin; left++ {
		if inorder[iBegin+left] == root.Val {
			root.Left = _buildTree(inorder, iBegin, iBegin+left-1, postorder, pBegin, pBegin+left-1)
			root.Right = _buildTree(inorder, iBegin+left+1, iEnd, postorder, pBegin+left, pEnd-1)
			break
		}
	}
	return root
}
