// Time        : 2019/07/19
// Description :

package symmetric_101

import . "go_learn/leecode/tree/base"

//Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
//For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//
//
//But the following [1,2,2,null,3,null,3] is not:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3

func isSymmetric(root *TreeNode) bool {
	queue := [][2]*TreeNode{{root, root}}
	poll := func() [2]*TreeNode {
		nodes := queue[0]
		for i := 0; i < len(queue)-1; i++ {
			queue[i], queue[i+1] = queue[i+1], queue[i]
		}
		queue = queue[:len(queue)-1]
		return nodes
	}
	for ; len(queue) > 0; queue = queue[1:] {
		nodes := poll()
		t1 := nodes[0]
		t2 := nodes[1]
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil || t1.Val != t2.Val {
			return false
		}
		queue = append(queue, [2]*TreeNode{t1.Left, t2.Right})
		queue = append(queue, [2]*TreeNode{t1.Right, t2.Left})
	}
	return true
}
