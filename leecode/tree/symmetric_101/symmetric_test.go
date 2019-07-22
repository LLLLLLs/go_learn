// Time        : 2019/07/19
// Description :

package symmetric_101

import (
	"fmt"
	. "go_learn/leecode/tree/base"
	"testing"
)

func TestSymmetric(t *testing.T) {
	root := NewNode(2)
	root.Left = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right = NewNode(3)
	root.Right.Right = NewNode(4)
	fmt.Println(isSymmetric(root))
}
