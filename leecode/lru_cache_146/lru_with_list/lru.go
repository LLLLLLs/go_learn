// Time        : 2019/07/30
// Description :

package lru_with_list

import "container/list"

type node struct {
	key   int
	value int
}

type LRUCache struct {
	values map[int]*list.Element
	cap    int
	l      *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		values: make(map[int]*list.Element, capacity),
		l:      new(list.List),
		cap:    capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	elem := lru.values[key]
	if elem == nil {
		return -1
	}
	lru.l.MoveToFront(elem)
	return elem.Value.(*list.Element).Value.(*node).value
}

func (lru *LRUCache) Put(key int, value int) {
	elem := lru.values[key]
	if elem != nil {
		elem.Value.(*list.Element).Value.(*node).value = value
		lru.l.MoveToFront(elem)
	} else {
		elem = &list.Element{Value: &node{key: key, value: value}}
		ptr := lru.l.PushFront(elem)
		lru.values[key] = ptr
		if len(lru.values) > lru.cap {
			tail := lru.l.Back()
			delete(lru.values, tail.Value.(*list.Element).Value.(*node).key)
			lru.l.Remove(tail)
		}
	}
}
