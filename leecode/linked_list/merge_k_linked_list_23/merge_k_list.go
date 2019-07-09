// Time        : 2019/01/22
// Description :

package merge_k_linked_list_23

import (
	"go_learn/leecode/linked_list/base"
	"sort"
)

//Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
//Example:
//
//Input:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//Output: 1->1->2->3->4->4->5->6

func mergeKLists1(lists []*base.ListNode) *base.ListNode {
	var head = &base.ListNode{Val: 1, Next: nil}
	var now = head
	var clear = false
	for !clear {
		var n *base.ListNode
		var index = 0
		clear = true
		for i, node := range lists {
			if node == nil {
				continue
			}
			clear = false
			if n == nil || node.Val < n.Val {
				n = node
				index = i
			}
		}
		if clear {
			break
		}
		now.Next = n
		now = n
		lists[index] = lists[index].Next
	}
	return head.Next
}

// 使用协程每次两两分组合并
func mergeKLists2_1(lists []*base.ListNode) *base.ListNode {
	if len(lists) == 0 {
		return nil
	}
	f := func(l1, l2 *base.ListNode, c chan *base.ListNode) {
		head := &base.ListNode{Val: 1, Next: nil}
		now := head
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				now.Next = l1
				l1 = l1.Next
			} else {
				now.Next = l2
				l2 = l2.Next
			}
			now = now.Next
		}
		if l1 == nil {
			now.Next = l2
		} else {
			now.Next = l1
		}
		c <- head.Next
	}
	for len(lists) > 1 {
		group := len(lists) / 2
		cs := make(chan *base.ListNode, group)
		for i := 0; i < group; i++ {
			go f(lists[i*2], lists[i*2+1], cs)
		}
		lists = lists[2*group:]
		for i := 0; i < group; i++ {
			lists = append(lists, <-cs)
		}
	}
	return lists[0]
}

// 不使用协程
func mergeKLists2_2(lists []*base.ListNode) *base.ListNode {
	if len(lists) == 0 {
		return nil
	}
	f := func(l1, l2 *base.ListNode) *base.ListNode {
		head := &base.ListNode{Val: 1, Next: nil}
		now := head
		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				now.Next = l1
				l1 = l1.Next
			} else {
				now.Next = l2
				l2 = l2.Next
			}
			now = now.Next
		}
		if l1 == nil {
			now.Next = l2
		} else {
			now.Next = l1
		}
		return head.Next
	}
	list := lists[0]
	lists = lists[1:]
	for len(lists) > 0 {
		list = f(list, lists[0])
		lists = lists[1:]
	}
	return list
}

// 第三方算法
func mergeKLists3(lists []*base.ListNode) *base.ListNode {
	heap := NewMinHeap(len(lists))
	for _, v := range lists {
		heap.Insert(v)
	}

	if heap.Size() == 0 {
		return nil
	}

	head := heap.Pop()
	current := head
	heap.Insert(current.Next)

	for heap.Size() > 1 {
		current.Next = heap.Pop()
		current = current.Next
		if current != nil {
			heap.Insert(current.Next)
		}
	}

	// link the rest of the last remaining list
	current.Next = heap.Pop()

	return head
}

type MinHeap struct {
	buffer []*base.ListNode
}

func NewMinHeap(capacity int) *MinHeap {
	return &MinHeap{buffer: make([]*base.ListNode, 0, capacity)}
}

// to get index of parent of node at index i
func parent(i int) int {
	return (i - 1) / 2
}

// to get index of left child of node at index i
func left(i int) int {
	return 2*i + 1
}

// to get index of right child of node at index i
func right(i int) int {
	return 2*i + 2
}

// A recursive method to heapify a subtree with the root at given index
// This method assumes that the subtrees are already heapified
func (h *MinHeap) heapify(i int) {
	l, r, smallest := left(i), right(i), i
	if l < len(h.buffer) && h.buffer[l].Val < h.buffer[i].Val {
		smallest = l
	}
	if r < len(h.buffer) && h.buffer[r].Val < h.buffer[smallest].Val {
		smallest = r
	}
	if smallest != i {
		h.buffer[i], h.buffer[smallest] = h.buffer[smallest], h.buffer[i]
		h.heapify(smallest)
	}
}

func (h *MinHeap) Insert(k *base.ListNode) {
	if k == nil {
		return
	}
	h.buffer = append(h.buffer, k)
	i := len(h.buffer) - 1
	for i != 0 && h.buffer[parent(i)].Val > h.buffer[i].Val {
		h.buffer[i], h.buffer[parent(i)] = h.buffer[parent(i)], h.buffer[i]
		i = parent(i)
	}
}

func (h *MinHeap) Pop() *base.ListNode {
	heapSize := h.Size()
	if heapSize == 0 {
		return nil
	}

	item := h.buffer[0]
	h.buffer[0] = h.buffer[heapSize-1]
	heapSize--
	h.buffer = h.buffer[:heapSize]
	if heapSize > 0 {
		h.heapify(0)
	}
	return item
}

func (h *MinHeap) Size() int {
	return len(h.buffer)
}

// 算法4
func mergeKLists4(lists []*base.ListNode) *base.ListNode {
	if len(lists) == 0 {
		return nil
	}

	q := NewQueue(lists)
	for q.Length() > 1 {
		q.Push(mergeTwoLists(q.Pop(), q.Pop()))
	}

	return q.Pop()
}

// You don't need to read on.

// Queue implementation
type Queue interface {
	Length() int
	Push(node *base.ListNode)
	Pop() *base.ListNode
}

type que struct {
	elements []*base.ListNode
}

func NewQueue(lists []*base.ListNode) Queue {
	if lists == nil {
		lists = make([]*base.ListNode, 0)
	}

	return &que{
		elements: lists,
	}
}

func (q *que) Length() int {
	return len(q.elements)
}

func (q *que) Push(node *base.ListNode) {
	q.elements = append(q.elements, node)
}

func (q *que) Pop() *base.ListNode {
	length := len(q.elements)
	if length == 0 {
		panic("queue is empty.")
	}
	n := q.elements[0]
	q.elements = q.elements[1:]
	return n
}

// Merge 2 lists
func mergeTwoLists(l1 *base.ListNode, l2 *base.ListNode) *base.ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := base.ListNode{}
	prev := &head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	}

	if l2 != nil {
		prev.Next = l2
	}

	return head.Next
}

// 算法5 用slice排序
func mergeKLists5(lists []*base.ListNode) *base.ListNode {
	l := make([]int, 0)
	for _, node := range lists {
		for node != nil {
			l = append(l, node.Val)
			node = node.Next
		}
	}
	sort.Ints(l)
	head := &base.ListNode{}
	now := head
	for _, val := range l {
		now.Next = &base.ListNode{Val: val}
		now = now.Next
	}
	return head.Next
}
