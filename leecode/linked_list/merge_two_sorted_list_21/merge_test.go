// Time        : 2019/06/26
// Description :

package merge_two_sorted_list_21

import (
	"go_learn/leecode/linked_list/base"
	"testing"
)

func TestMerge(t *testing.T) {
	l1 := base.NewList(1)
	l1.Add(2, 4)
	l2 := base.NewList(1)
	l2.Add(3, 4)
	merge(l1, l2).Print()
}
