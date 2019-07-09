// Time        : 2019/07/09
// Description : 瞎比实现lru内存缓存

package lru

type DoubleNode struct {
	key   interface{}
	value interface{}
	pre   *DoubleNode
	next  *DoubleNode
}

type Cache struct {
	cache map[interface{}]*DoubleNode
	head  *DoubleNode
	tail  *DoubleNode
	cap   int
}

func NewLRUCache(n int) *Cache {
	return &Cache{
		cache: make(map[interface{}]*DoubleNode),
		head:  nil,
		tail:  nil,
		cap:   n,
	}
}

func (c *Cache) Get(key interface{}) interface{} {
	node, ok := c.cache[key]
	if !ok {
		return nil
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

func (c *Cache) Put(key, value interface{}) {
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
