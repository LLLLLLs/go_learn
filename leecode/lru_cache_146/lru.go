// Time        : 2019/07/30
// Description :

package lru_cache_146

// Design and implement a data structure for Least Recently Used (LRU) cache.
// It should support the following operations: get and put.
//
// get(key) - Get the value (will always be positive) of the key
// 	if the key exists in the cache, otherwise return -1.
// put(key, value) - Set or insert the value if the key is not already present.
// 	When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
//
// The cache is initialized with a positive capacity.
//
// Follow up:
// Could you do both operations in O(1) time complexity?
//
// Example:
//
// LRUCache cache = new LRUCache( 2 /* capacity */ );
//
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4

type doubleNode struct {
	key  int
	val  int
	pre  *doubleNode
	next *doubleNode
}

type LRUCache struct {
	values map[int]*doubleNode
	head   *doubleNode
	tail   *doubleNode
	cap    int
}

func Constructor(capacity int) LRUCache {
	head, tail := &doubleNode{}, &doubleNode{}
	head.next, tail.pre = tail, head
	return LRUCache{
		values: make(map[int]*doubleNode, capacity),
		head:   head,
		tail:   tail,
		cap:    capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	node := lru.values[key]
	if node == nil {
		return -1
	}
	lru.moveToHead(node)
	return node.val
}

func (lru *LRUCache) moveToHead(node *doubleNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	node.next = lru.head.next
	node.pre = lru.head
	lru.head.next.pre = node
	lru.head.next = node
}

func (lru *LRUCache) Put(key int, value int) {
	node := lru.values[key]
	if node != nil {
		node.val = value
		lru.moveToHead(node)
	} else {
		node = &doubleNode{
			key:  key,
			val:  value,
			pre:  lru.head,
			next: lru.head.next,
		}
		lru.head.next.pre = node
		lru.head.next = node
		lru.values[key] = node
		if len(lru.values) > lru.cap {
			delete(lru.values, lru.tail.pre.key)
			lru.tail.pre = lru.tail.pre.pre
			lru.tail.pre.next = lru.tail
		}
	}
}
