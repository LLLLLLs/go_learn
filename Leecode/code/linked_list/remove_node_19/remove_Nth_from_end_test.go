// Time        : 2019/01/22
// Description :

package remove_node_19

import (
	"go_learn_test/Leecode/code/linked_list/base"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := base.NewList(1)
	head.Add(2, 3, 4, 5)
	head.Print()
	head = RemoveNthFromEnd(head, 2)
	if head != nil {
		head.Print()
	}
}
