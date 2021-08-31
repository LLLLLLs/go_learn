// @author: lls
// @time: 2020/05/11
// @desc:

package time

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPrintTime(t *testing.T) {
	fmt.Println(time.Unix(1589163920, 0))
}

type timer struct {
	timer *time.Timer
	stop  chan struct{}
}

func (t timer) countdown() {
	for {
		select {
		case <-t.timer.C:
			fmt.Println("finish")
			t.stop <- struct{}{}
			return
		}
	}
}

func (t timer) resetTimer() {
	for {
		randSec := rand.Intn(6) + 1
		fmt.Println(randSec, "秒")
		t.timer.Reset(time.Duration(randSec) * time.Second)
		time.Sleep(time.Second * 2)
	}
}

func TestTimer(t *testing.T) {
	rand.Seed(time.Now().Unix())
	tm := timer{
		timer: time.NewTimer(time.Second * 4),
		stop:  make(chan struct{}, 1),
	}
	go func() {
		tk := time.NewTicker(time.Second)
		for {
			select {
			case <-tk.C:
				fmt.Println("ticker")
			}
		}
	}()
	go tm.countdown()
	go tm.resetTimer()
	<-tm.stop
}

func TestTimer2(t *testing.T) {
	tm := time.NewTimer(time.Second)
	for {
		select {
		case <-tm.C:
			fmt.Println("123")
			// tm.Reset(time.Second)
		}
	}
}
