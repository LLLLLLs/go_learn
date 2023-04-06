package patchtime

import (
	"bou.ke/monkey"
	"reflect"
	"sync"
	"time"
)

func withOffset(sec int64) {
	offset.Store(sec)
	manager.refreshWithOffsetChanged()
}

var onceConsume = sync.Once{}

func Patch() {
	monkey.Patch(time.NewTimer, NewTimer)
	replaceTimer()
	monkey.Patch(time.NewTicker, NewTicker)
	replaceTicker()
	monkey.Patch(time.Now, Now)
	onceConsume.Do(func() {
		go manager.consume()
	})
}

func replaceTimer() {
	t := new(time.Timer)
	rt := reflect.TypeOf(t)
	monkey.PatchInstanceMethod(rt, "Reset", func(t *time.Timer, d time.Duration) bool {
		mt, has := manager.ttMapper.Load(t)
		if !has {
			return false
		}
		return manager.reset(mt.(*timer).id, d, 0)
	})
	monkey.PatchInstanceMethod(rt, "Stop", func(t *time.Timer) bool {
		mt, has := manager.ttMapper.Load(t)
		if !has {
			return false
		}
		return manager.stop(mt.(*timer).id)
	})
}
func replaceTicker() {
	t := new(time.Ticker)
	rt := reflect.TypeOf(t)
	monkey.PatchInstanceMethod(rt, "Reset", func(t *time.Ticker, d time.Duration) {
		mt, has := manager.ttMapper.Load(t)
		if !has {
			return
		}
		manager.reset(mt.(*timer).id, d, d)
		return
	})
	monkey.PatchInstanceMethod(rt, "Stop", func(t *time.Ticker) {
		mt, has := manager.ttMapper.Load(t)
		if !has {
			return
		}
		manager.stop(mt.(*timer).id)
		return
	})
}

func Now() time.Time {
	return patchNow()
}

func NewTimer(d time.Duration) *time.Timer {
	return manager.newTimer(d)
}

func NewTicker(d time.Duration) *time.Ticker {
	return manager.newTicker(d)
}
