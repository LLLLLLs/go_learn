/**
**队列模块。使用channel作为存放数据的存储
 */
package flumesdk

import (
	"time"
)

type queue struct {
	value   chan interface{}
	maxSize int
	timeout time.Duration
}

func (q *queue) Get() (interface{}, bool) {
	select {
	case v, ok := <-q.value:
		if !ok {
			return nil, false
		}
		return v, true

	case <-time.After(q.timeout):
		return nil, false
	}
}

//var Lock = new(sync.Mutex)
//func (q *queue) GetNoWait() (interface{}, bool) {
//	Lock.Lock()
//	defer Lock.Unlock()
//	if len(q.value) == 0 {
//		return nil, false
//	}
//	v, ok := <-q.value
//	if !ok {
//		return nil, false
//	}
//	return v, true
//}

func (q *queue) GetNoWait() (interface{}, bool) {
	select {
	case v, ok := <-q.value:
		if !ok {
			return nil, false
		}
		return v, true

	case <-time.After(q.timeout / 10):
		return nil, false
	}
}

func (q *queue) Put(v interface{}) bool {
	select {
	case q.value <- v:
		return true

	case <-time.After(q.timeout):
		return false
	}
}

func (q *queue) Size() int {
	return len(q.value)
}

func (q *queue) Empty() bool {
	return len(q.value) == 0
}

func (q *queue) Full() bool {
	return len(q.value) == cap(q.value)
}

func (q *queue) Close() {
	close(q.value)
}

func newQueue(maxSize int, timeout time.Duration) *queue {
	queue := queue{
		value:   make(chan interface{}, maxSize),
		maxSize: maxSize,
		timeout: timeout,
	}
	return &queue
}
