// Time        : 2019/07/22
// Description :

package restore_from_pre_inorder_105

import (
	"fmt"
	"testing"
)

func TestRestore(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root.PreorderTraversal())
	fmt.Println(root.InorderTraversal())
}
