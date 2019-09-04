// Time        : 2019/07/31
// Description :

package intersection_of_two_list_160

import (
	"golearn/leecode/linked_list/base"
	"testing"
)

func TestIntersection(t *testing.T) {
	headA := base.NewList(4, 1, 8, 4, 5)
	headB := base.NewList(5, 0, 1)
	headB.Next.Next.Next = headA.Next.Next
	getIntersectionNode(headA, headB).Print()
}
