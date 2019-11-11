// Time        : 2019/01/08
// Description : 跳跃表

package skip_list

import (
	"fmt"
	"math/rand"
)

type node struct {
	key   int
	value int
	level []*node
}

func (n node) String() string {
	return fmt.Sprintf("key:%d,value:%d", n.key, n.value)
}

func newNode(k, v int) *node {
	return &node{
		key:   k,
		value: v,
		level: []*node{nil},
	}
}

type list struct {
	level int
	head  *node
}

func newList() *list {
	return &list{
		level: 1,
		head: &node{
			level: []*node{nil},
		},
	}
}

// 查找
func (l *list) search(k int) *node {
	x := l.head
	for i := l.level - 1; i >= 0; i-- {
		for x.level[i] != nil && x.level[i].key < k {
			x = x.level[i]
		}
	}
	if x.level[0].key == k {
		return x.level[0]
	}
	return nil
}

// 插入新节点
func (l *list) insert(k, v int) {
	x := l.head
	for i := l.level - 1; i >= 0; i-- {
		for x.level[i] != nil && x.level[i].key < k {
			x = x.level[i]
		}
	}
	// 最底层
	n := newNode(k, v)
	next := x.level[0]
	x.level[0] = n
	n.level[0] = next
	// 新节点层数
	level := 1
	// 每次都有50%概率高一层
	for rand.Intn(2) == 0 {
		level++
		// 超过头的层数
		if l.level < level {
			l.level++
			l.head.level = append(l.head.level, n)
			n.level = append(n.level, nil)
			continue
		}
		// 已有该层数节点
		x := l.head
		for i := l.level - 1; i >= level-1; i-- {
			for x.level[i] != nil && x.level[i].key < n.key {
				x = x.level[i]
			}
		}
		next = x.level[level-1]
		x.level[level-1] = n
		n.level = append(n.level, next)
	}
}

func (l *list) delete(k int) {
	x := l.head
	for i := l.level - 1; i >= 0; i-- {
		for x.level[i] != nil && x.level[i].key < k {
			x = x.level[i]
		}
		if x.level[i] == nil || x.level[i].key != k {
			continue
		}
		x.level[i] = x.level[i].level[i]
	}
}

func (l *list) print() {
	for i := l.level - 1; i >= 0; i-- {
		x := l.head
		info := fmt.Sprintf("level.%d:head", i)
		for x.level[i] != nil {
			x = x.level[i]
			info = fmt.Sprintf("%s->%d", info, x.key)
		}
		info = fmt.Sprintf("%s->nil", info)
		fmt.Println(info)
	}
}
