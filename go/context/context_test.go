// Time        : 2019/10/15
// Description :

package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	<-ctx.Done()
}

func TestCtxWithTimeOutNest(t *testing.T) {
	timeOutNest()
}

func TestWithValue(t *testing.T) {
	ctx := context.Background()
	setCtx(ctx, "my context")
	v := ctx.Value("value")
	fmt.Println(v)
}

func setCtx(ctx context.Context, value interface{}) {
	ctx = context.WithValue(ctx, "value", value)
}

func TestNewContext(t *testing.T) {
	ctx := newContext()
	_ = ctx
}
