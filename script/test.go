// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"strconv"
	"time"
)

type WeightForRand interface {
	WeightNum() int
}

func RandWithWeight(list ...WeightForRand) WeightForRand {
	candidate := make([]WeightForRand, 0, len(list))
	weightTotal := 0
	for i := range list {
		if list[i].WeightNum() > 0 { // 权重大于0的才能加入随机列表
			weightTotal += list[i].WeightNum()
			candidate = append(candidate, list[i])
		}
	}
	// 无随机道具
	if weightTotal == 0 {
		return nil
	}
	// 从随机道具中抽取一种
	// bingo := 1 + rand.Intn(weightTotal) // [1,total]
	// curWeight := 0
	// for i := range candidate {
	// 	curWeight += candidate[i].WeightNum()
	// 	if bingo <= curWeight { // rand返回的是左闭右开随机数 因此右端点属于下一条
	// 		return candidate[i]
	// 	}
	// }
	return nil
}

type ran struct {
	weight int
	data   int
}

func (t ran) WeightNum() int {
	return t.weight
}

func main() {
	fmt.Println(Key())
	return
	list := make([]WeightForRand, 3)
	now := time.Now().UnixNano()
	fmt.Println(now)
	// fmt.Println(rand.Int(100))
	list[0] = ran{weight: int(5), data: 3}
	list[1] = ran{weight: int(1), data: 10}
	list[2] = ran{weight: int(15), data: 20}
	log.Print(RandWithWeight(list...), "=======1")
}

func Key() int64 {
	buf := make([]byte, 4)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err) // out of randomness, should never happen
	}
	i, err := strconv.ParseInt(fmt.Sprintf("%x", buf), 16, 64)
	if err != nil {
		panic(err)
	}
	return i
}
