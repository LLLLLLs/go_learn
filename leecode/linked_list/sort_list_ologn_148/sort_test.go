// Time        : 2019/07/30
// Description :

package sort_list_ologn_148

import (
	"golearn/leecode/linked_list/base"
	"testing"
)

func TestSort(t *testing.T) {
	head := base.NewList(-1, 5, 3, 4, 0)
	sortList(head)
	head.Print()
}
