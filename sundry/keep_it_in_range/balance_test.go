//@author: lls
//@date: 2020/6/10
//@desc:

package keepitinrange

import (
	"context"
	"testing"
	"time"
)

func TestBalance(t *testing.T) {
	b := NewBalance(5, 10, 20)
	ctx := context.Background()
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Second*20))
	b.Balance(ctx)
}
