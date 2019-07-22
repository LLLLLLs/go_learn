// Time        : 2019/07/22
// Description :

package restore_from_in_postorder_106

import (
	"fmt"
	"testing"
)

func TestRestore(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	fmt.Println(root.InorderTraversal())
	fmt.Println(root.PostorderTraversal())
	inorder = []int{2, 1}
	postorder = []int{2, 1}
	root = buildTree(inorder, postorder)
	fmt.Println(root.InorderTraversal())
	fmt.Println(root.PostorderTraversal())
}
