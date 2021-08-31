// @author: lls
// @date: 2021/8/31
// @desc:

package main

import (
	"context"
	"fmt"
	"golearn/pkgtest/redis"
	"golearn/util"
	"time"
)

func main() {
	getLoop()
	getBatch()
}

func getLoop() {
	client := redis.Client()
	start := time.Now()
	ctx := context.Background()
	for i := 0; i < 200; i++ {
		res, err := client.Get(ctx, key(i)).Result()
		util.MustNil(err)
		util.MustTrue(res == value(i))
	}
	fmt.Println("遍历耗时:", time.Now().Sub(start).Microseconds(), "us")
}

func getBatch() {
	client := redis.Client()
	start := time.Now()
	ctx := context.Background()
	keys := make([]string, 0)
	for i := 0; i < 200; i++ {
		keys = append(keys, key(i))
	}
	_, err := client.MGet(ctx, keys...).Result()
	util.MustNil(err)
	// fmt.Println(res)
	fmt.Println("批量耗时:", time.Now().Sub(start).Microseconds(), "us")
}

func key(i int) string {
	return fmt.Sprintf("test_key%d", i)
}

func value(i int) string {
	return fmt.Sprintf("test_value%d", i)
}
