// Time        : 2019/07/12
// Description :

package preorder_traversal_144

import (
	"fmt"
	. "go_learn/leecode/tree/base"
	"testing"
)

func TestPreorder(t *testing.T) {
	root := NewNode(1)
	root.Right = NewNode(2)
	root.Right.Left = NewNode(3)
	fmt.Println(preorderTraversal(root))
	fmt.Println(root.PreorderTraversal())
}
