// Time        : 2019/07/04
// Description :

package rotate_61

import (
	"go_learn/leecode/code/linked_list/base"
	"testing"
)

func TestRotate(t *testing.T) {
	list := base.NewList(1, 2, 3, 4, 5)
	list = rotateRight(list, 2)
	list.Print()
	list = base.NewList(1)
	list = rotateRight(list, 0)
	list.Print()
}
