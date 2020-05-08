//@time:2020/04/02
//@desc:

package context

import (
	"context"
	"fmt"
	"time"
)

func newContext() context.Context {
	ctx := context.WithValue(context.Background(), "key", 123)
	return ctx
}

// 如果parent ctx中已经有了deadline,并且该deadline早于本次定义的deadline,
// 那么本次定义的deadline会以parent的为主(包括done和cancel)
func timeOutNest() {
	ctx := context.Background()
	go ticker()
	timeOut5Sec(ctx)
	return
}

func ticker() {
	var sec = 1
	tk := time.NewTicker(time.Second)
	for {
		select {
		case <-tk.C:
			fmt.Printf("di da...%d\n", sec)
			sec++
		}
	}
}

func timeOut5Sec(c context.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()
	go timeOut4Sec(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("5 second done!")
	}
}

func timeOut4Sec(c context.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*4)
	defer cancel()
	go timeOut6Sec(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("4 second done!")
		//case <-time.After(time.Second * 3): // 可注释该case来获取done的结果
		//	fmt.Println("exit without 4s ctx done")
	}
}

func timeOut6Sec(c context.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*6)
	defer cancel()
	go timeOutAno5Sec(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("6 second done!")
	}
}

func timeOutAno5Sec(c context.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("another 5 second done!")
	}
}
