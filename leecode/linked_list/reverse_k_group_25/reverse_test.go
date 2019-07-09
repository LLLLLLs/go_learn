// Time        : 2019/01/23
// Description :

package reverse_k_group_25

import (
	"go_learn/leecode/linked_list/base"
	"testing"
)

func TestReverse1(t *testing.T) {
	l := &base.ListNode{}
	l.Add(1, 2, 3, 4, 5)
	l = l.Next
	l.Print()
	l = reverseKGroup1(l, 3)
	l.Print()
}

func TestReverse2(t *testing.T) {
	l := &base.ListNode{}
	l.Add(1, 2, 3, 4, 5)
	l = l.Next
	l.Print()
	l = reverseKGroup2(l, 3)
	l.Print()
}

func TestReverse(t *testing.T) {
	l := base.NewList(1, 2)
	reverseKGroup(l, 2).Print()
}
