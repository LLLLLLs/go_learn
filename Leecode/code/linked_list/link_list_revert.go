/*
Author      : lls
Time        : 2018/10/29
Description : 链表翻转
*/

package linked_list

// 给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。
// 示例:
//
// 给定这个链表：1->2->3->4->5
//
// 当 k = 2 时，应当返回: 2->1->4->3->5
//
// 当 k = 3 时，应当返回: 3->2->1->4->5
// 说明:
//
// 你的算法只能使用常数的额外空间。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

func Revert(node *Node, k int) *Node {
	list := make([]*Node, 0)
	var next *Node = nil
	for i := 0; i < k; i++ {
		if node == nil {
			break
		}
		list = append(list, node)
		node = node.next
	}
	if node != nil {
		next = node
	}
	head := list[len(list)-1]
	node = head
	for i := len(list) - 1; i > 0; i-- {
		node.next = list[i-1]
		node = node.next
	}
	node.next = nil
	if next != nil {
		node.next = Revert(next, k)
	}
	return head
}

func RevertAll(pre, current *Node) (head *Node) {
	if current.next == nil {
		current.next = pre
		return current
	}
	head = RevertAll(current, current.next)
	current.next = pre
	return head
}
