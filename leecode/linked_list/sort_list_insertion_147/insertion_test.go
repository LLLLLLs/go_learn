// Time        : 2019/07/30
// Description :

package sort_list_insertion_147

import (
	"golearn/leecode/linked_list/base"
	"testing"
)

func TestInsertion(t *testing.T) {
	head := base.NewList(-1, 5, 3, 4, 0)
	insertionSortList(head)
	head.Print()
}
