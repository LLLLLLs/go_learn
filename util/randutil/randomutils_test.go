// Time        : 2019/09/19
// Description :

package randutil

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandList(t *testing.T) {
	list := make([]int, 100)
	for i := range list {
		list[i] = i + 1
	}
	rand5 := RandList(list, 5)
	fmt.Println(rand5)
}

func BenchmarkRandList(b *testing.B) {
	list := make([]int, 10000)
	for i := range list {
		list[i] = i + 1
	}
	for i := 0; i < b.N; i++ {
		listCopy := make([]int, len(list))
		copy(listCopy, list)
		RandList(list, 99)
	}
}

func TestRandMap(t *testing.T) {
	dataMap := make(map[int]struct{})
	dataList := make([]int, 10)
	for i := 0; i < 10; i++ {
		dataMap[i+1] = struct{}{}
		dataList[i] = i + 1
	}
	start := time.Now()
	resultMap := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		j := 0
		for key := range dataMap {
			resultMap[key]++
			j++
			if j == 5 {
				break
			}
		}

	}
	fmt.Println("map耗时:", time.Now().Sub(start).Nanoseconds()/int64(time.Millisecond))
	start = time.Now()
	resultList := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		dataCopy := make([]int, len(dataList))
		copy(dataCopy, dataList)
		choice := RandList(dataCopy, 5)
		for _, v := range choice.([]int) {
			resultList[v]++
		}
	}
	fmt.Println("list耗时:", time.Now().Sub(start).Nanoseconds()/int64(time.Millisecond))
	fmt.Printf("%+v\n%+v\n", resultMap, resultList)
}

func TestInitRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float32())
	fmt.Println(rand.ExpFloat64())
	fmt.Println(rand.NormFloat64())
}

func TestCryptoRand(t *testing.T) {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err) // out of randomness, should never happen
	}
	fmt.Printf("%x\n", buf)
}
