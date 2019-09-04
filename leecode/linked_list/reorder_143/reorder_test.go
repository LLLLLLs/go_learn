// Time        : 2019/07/30
// Description :

package reorder_143

import (
	"golearn/leecode/linked_list/base"
	"testing"
)

func TestReorder(t *testing.T) {
	list := base.NewList(1, 2, 3, 4, 5)
	reorderList2(list)
	list.Print()
}
