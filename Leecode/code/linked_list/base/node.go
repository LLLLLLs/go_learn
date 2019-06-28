// Time        : 2019/01/22
// Description :

package base

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(values ...int) *ListNode {
	head := &ListNode{}
	cur := head
	if len(values) == 0 {
		return head
	}
	for _, v := range values {
		cur.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		cur = cur.Next
	}
	return head.Next
}

func (h *ListNode) Print() {
	ll := h
	for ll != nil {
		fmt.Printf("%d-->", ll.Val)
		ll = ll.Next
	}
	fmt.Printf("nil\n")
}

func (h *ListNode) Add(vals ...int) {
	ll := h
	for ll.Next != nil {
		ll = ll.Next
	}
	for _, val := range vals {
		ll.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		ll = ll.Next
	}
}
