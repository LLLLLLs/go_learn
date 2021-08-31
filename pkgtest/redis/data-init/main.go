// @author: lls
// @date: 2021/8/31
// @desc:

package main

import (
	"context"
	"fmt"
	"golearn/pkgtest/redis"
	"golearn/util"
)

func main() {
	client := redis.Client()
	ctx := context.Background()
	// client.FlushDB(ctx)
	for i := 0; i < 200; i++ {
		util.MustNil(client.Set(ctx, key(i), value(i), -1).Err())
		util.MustNil(client.HSet(ctx, "test_hash", key(i), value(i)).Err())
	}
}

func key(i int) string {
	return fmt.Sprintf("test_key%d", i)
}

func value(i int) string {
	return fmt.Sprintf("test_value%d", i)
}
