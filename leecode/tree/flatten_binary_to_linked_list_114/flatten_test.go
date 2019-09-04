// Time        : 2019/07/22
// Description :

package flatten_binary_to_linked_list_114

import (
	"fmt"
	. "golearn/leecode/tree/base"
	"testing"
)

//     1
//    / \
//   2   5
//  / \   \
// 3   4   6
func TestFlatten(t *testing.T) {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Left.Left = NewNode(3)
	root.Left.Right = NewNode(4)
	root.Right = NewNode(5)
	root.Right.Right = NewNode(6)
	flatten(root)
	fmt.Println(root.PreorderTraversal())
	fmt.Println(root.InorderTraversal())
}
