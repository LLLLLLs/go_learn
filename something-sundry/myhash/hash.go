// Time        : 2019/01/21
// Description :

package myhash

const size_default = 2

// 键值对
type slot struct {
	key   int
	value interface{}
	next  *slot
}

// 哈希Map
type hash struct {
	ht     [2][]*slot    // 哈希表
	count  int           // 键值对数量
	size   int           // 哈希表大小
	expend int           // 扩展标识
	f      func(int) int // 哈希函数
}

func newHash() *hash {
	h := &hash{
		ht:     [2][]*slot{make([]*slot, size_default), nil},
		count:  0,
		size:   size_default,
		expend: -1,
	}
	h.f = func(i int) int {
		return i % h.size
	}
	return h
}

// 计算扩展因子
func (h *hash) check() bool {
	return h.count == h.size
}

// 扩展开始
func (h *hash) expendBegin() {
	h.size *= 2
	h.ht[1] = make([]*slot, h.size)
	h.expend = 0
	h.f = func(i int) int {
		return i % h.size
	}
}

// 扩展一次
func (h *hash) expendOnce() {
	s := h.ht[0][h.expend]
	h.ht[0][h.expend] = nil
	for s != nil {
		next := s.next
		index := h.f(s.key)
		s.next = h.ht[1][index]
		h.ht[1][index] = s
		s = next
	}
	h.expend++
	if h.expend >= h.size/2 {
		h.expend = -1
		h.ht[0] = h.ht[1]
		h.ht[1] = nil
	}
}

// 设置键值对
func (h *hash) set(k int, v interface{}) {
	index := h.f(k)
	ht := 0
	if h.expend != -1 {
		h.expendOnce()
		if h.expend != -1 {
			ht = 1
		}
	}
	slot := &slot{k, v, nil}
	slot.next = h.ht[ht][index]
	h.ht[ht][index] = slot
	h.count++
	if h.check() {
		h.expendBegin()
	}
}

// 获取值
func (h *hash) get(k int) interface{} {
	index := h.f(k)
	ht := 0
	if h.expend != -1 {
		h.expendOnce()
		if h.expend > index {
			ht = 1
		}
	}
	s := h.ht[ht][index]
	for s != nil {
		if s.key == k {
			break
		}
		s = s.next
	}
	if s != nil {
		return s.value
	}
	return nil
}

// 删除键值对
func (h *hash) remove(k int) {
	index := h.f(k)
	ht := 0
	if h.expend != -1 {
		h.expendOnce()
		if h.expend > index {
			ht = 1
		}
	}
	s := h.ht[ht][index]
	var prev *slot
	for s != nil {
		if s.key == k {
			if prev == nil {
				h.ht[ht][index] = prev
			} else {
				prev.next = s.next
			}
			h.count--
			break
		}
		prev = s
		s = s.next
	}
}
