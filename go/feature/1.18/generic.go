package __18

import (
	"fmt"
	"strconv"
)

type Uint interface {
	uint8 | uint16 | uint32 | uint64 | uint
}

func sum[K comparable, V Uint](m map[K]V) V {
	var result V
	for _, v := range m {
		result += v
	}
	return result
}

func toUint[V Uint](str string, _ V) V {
	ui, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return V(ui)
}

// type IAdd interface {
// 	add(n int)
// }
//
// type addA struct {
// 	num int
// }
//
// func (a *addA) add(n int) {
// 	a.num += n
// }
//
// type addB struct {
// 	num int
// }
//
// func (a *addB) add(n int) {
// 	a.num += n
// }
//
// func add[a *addA | *addB](t a, num int) int {
// 	t.add(num)
// 	return t.num
// }

func value[T any](data interface{}) T {
	d, has := data.(T)
	if has {
		return d
	}
	pd, has := data.(*T)
	if has {
		return *pd
	}
	fmt.Println("error")
	return d
}
