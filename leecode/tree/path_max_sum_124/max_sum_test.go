// Time        : 2019/07/23
// Description :

package path_max_sum_124

import (
	"fmt"
	. "go_learn/leecode/tree/base"
	"testing"
)

func TestMaxPath(t *testing.T) {
	root := NewNode(-10)
	root.Left = NewNode(9)
	root.Right = NewNode(20)
	root.Right.Left = NewNode(15)
	root.Right.Right = NewNode(7)
	fmt.Println(maxPathSum(root))
	root = NewNode(9)
	root.Left = NewNode(6)
	root.Right = NewNode(-3)
	root.Right.Left = NewNode(-6)
	root.Right.Right = NewNode(2)
	root.Right.Right.Left = NewNode(2)
	root.Right.Right.Left.Left = NewNode(-6)
	root.Right.Right.Left.Right = NewNode(-6)
	root.Right.Right.Left.Left.Left = NewNode(-6)
	fmt.Println(maxPathSum(root))
}
