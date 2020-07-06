//@author: lls
//@date: 2020/6/10
//@desc: 保证在最大值和最小值之间，期间增减是随机的

package keepitinrange

import (
	"context"
	"fmt"
	"time"
)

type Balancer interface {
	Balance(ctx context.Context)
}

type balance struct {
	num, min, max              int
	biggerThanMin, lessThanMax chan struct{}
}

func NewBalance(num, min, max int) Balancer {
	return &balance{
		num:           num,
		min:           min,
		max:           max,
		biggerThanMin: make(chan struct{}, 1),
		lessThanMax:   make(chan struct{}, 1),
	}
}

func (b *balance) Balance(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond * 200)
	for {
		b.produce()
		fmt.Println(b.num)
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			select {
			case <-b.biggerThanMin:
				b.num--
			case <-b.lessThanMax:
				b.num++
			}
			b.consume()
		}
	}
}

func (b balance) produce() {
	if b.num < b.max {
		b.lessThanMax <- struct{}{}
	}
	if b.num > b.min {
		b.biggerThanMin <- struct{}{}
	}
}

func (b balance) consume() {
	for {
		select {
		case <-b.biggerThanMin:
		case <-b.lessThanMax:
		default:
			return
		}
	}
}
