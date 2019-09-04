// Time        : 2019/07/22
// Description :

package binary_zigzag_level_order_traversal_103

import (
	"fmt"
	. "golearn/leecode/tree/base"
	"testing"
)

// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its zigzag level order traversal as:
// [
//   [3],
//   [20,9],
//   [15,7]
// ]
func TestZigzagTraversal(t *testing.T) {
	root := NewNode(3)
	root.Left = NewNode(9)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	fmt.Println(zigzagLevelOrder(root))
}
