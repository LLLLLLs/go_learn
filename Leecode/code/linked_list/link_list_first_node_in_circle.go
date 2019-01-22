/*
Author      : lls
Time        : 2018/10/29
Description : 入环的第一个节点
*/

package linked_list

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

func FirstNodeInCircle(head *Node) *Node {
	current, slow, fast := head.next, head.next, head.next.next
	for {
		for { // 每次循环slow走1步，fast走2步
			var fastStep = 2
			for i := 0; i < fastStep; i++ {
				fast = fast.next
				if fast == slow { // 快慢相遇
					break
				}
				if fast == nil { // 没有环
					return nil
				}
			}
			if fast == slow { // 快慢相遇
				break
			}
			slow = slow.next
		}
		if slow == current { // 三点相遇 --> 环入口
			return current
		}
		current = current.next // 否则 current 走一步
	}
}

// 改良版
//链表头是X，环的第一个节点是Y，slow和fast第一次的交点是Z。各段的长度分别是a,b,c，如图所示
//第一次相遇时slow走过的距离：a+b，fast走过的距离：a+b+c+b
//因为fast的速度是slow的两倍，所以fast走的距离是slow的两倍，有 2(a+b) = a+b+c+b，可以得到a=c（这个结论很重要！）
//这时候，slow从X出发，fast从Z出发，以相同速度走，同时到达Y，Y就是环的入口，即第一个节点

func FirstNodeInCircleImprove(head *Node) *Node {
	slow, fast := head, head
	for {
		var fastStep = 2
		for i := 0; i < fastStep; i++ {
			fast = fast.next
			if fast == slow { // 快慢相遇
				break
			}
			if fast == nil { // 没有环
				return nil
			}
		}
		if fast == slow { // 快慢相遇
			break
		}
		slow = slow.next
	}
	slow = head
	for {
		slow = slow.next
		if fast == slow { // 快慢相遇
			break
		}
		fast = fast.next
		if fast == slow { // 快慢相遇
			break
		}
	}
	return fast
}
