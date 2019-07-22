// Time        : 2019/07/22
// Description :

package binary_tree_level_order_traversal_107

import (
	"fmt"
	. "go_learn/leecode/tree/base"
	"testing"
)

func TestTraversal(t *testing.T) {
	root := NewNode(3)
	root.Left = NewNode(9)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	fmt.Println(levelOrderBottom(root))
}
