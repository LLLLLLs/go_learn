/*
Author      : lls
Time        : 2018/10/29
Description : 链表基础
*/

package linked_list

import "fmt"

type Node struct {
	num  int
	next *Node
}

func GetInitLinkList(n int) *Node {
	head := &Node{num: 0, next: nil}
	current := head
	for i := 1; i < n; i++ {
		node := &Node{num: i, next: nil}
		current.next = node
		current = node
	}
	return head
}

func GetCircleLinkList(n, c int) *Node {
	head := &Node{num: 1, next: nil}
	current := head
	circle := head
	for i := 2; i <= n; i++ {
		node := &Node{num: i, next: nil}
		current.next = node
		current = node
		if c == i {
			circle = current
		}
	}
	current.next = circle
	return head
}

func PrintLinkList(node *Node) {
	for node != nil {
		fmt.Printf("%d->", node.num)
		node = node.next
	}
	fmt.Printf("nil\n")
}
