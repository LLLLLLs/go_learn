// Time        : 2019/07/12
// Description :

package inorder_traversal_94

import (
	"fmt"
	. "golearn/leecode/tree/base"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := NewNode(1)
	root.Right = NewNode(2)
	root.Right.Left = NewNode(3)
	fmt.Println(inorderTraversal(root))
	fmt.Println(root.InorderTraversal())
}
