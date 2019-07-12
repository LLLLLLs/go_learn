// Time        : 2019/07/12
// Description :

package postorder_traversal

import (
	"fmt"
	. "go_learn/leecode/tree/base"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	root := NewNode(1)
	root.Right = NewNode(2)
	root.Right.Left = NewNode(3)
	fmt.Println(postorderTraversal(root))
	fmt.Println(root.PostorderTraversal())
}
