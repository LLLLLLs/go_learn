// Time        : 2019/10/15
// Description :

package context

import (
	"context"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	<-ctx.Done()
}
