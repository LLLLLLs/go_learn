/*
Author      : lls
Time        : 2018/10/30
Description :
*/

package code

type DoubleNode struct {
	key   int
	value int
	pre   *DoubleNode
	next  *DoubleNode
}

type LRUCache struct {
	cache map[int]*DoubleNode
	head  *DoubleNode
	tail  *DoubleNode
	cap   int
}

func NewLRUCache(n int) *LRUCache {
	return &LRUCache{
		cache: make(map[int]*DoubleNode),
		head:  nil,
		tail:  nil,
		cap:   n,
	}
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1
	}
	if node != c.head { // node 不是 head
		node.pre.next = node.next
		node.next.pre = node.pre
		node.next = c.head
		c.head.pre = node
		node.pre = nil
		c.head = node
		if node == c.tail {
			c.tail = node.pre
		}
	}
	return node.value
}

func (c *LRUCache) Put(key, value int) {
	node, ok := c.cache[key]
	if ok { // 该node已存在
		if node != c.head {
			node.pre.next = node.next
			node.next = c.head
			node.pre = nil
		}
		node.value = value
		return
	}

	node = &DoubleNode{
		key:   key,
		value: value,
		pre:   nil,
		next:  nil,
	}

	c.cache[key] = node
	if len(c.cache) == 1 { // 第一个node
		c.head = node
		c.tail = node
	}
	c.head.pre = node
	node.next = c.head
	c.head = node
	if len(c.cache) > c.cap { // 超过容量
		delete(c.cache, c.tail.key)
		c.tail = c.tail.pre
		c.tail.next = nil
		return
	}
}
