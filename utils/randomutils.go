// Time        : 2019/06/25
// Description :

package utils

import (
	"math/rand"
	"reflect"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// 从切片中随机获取N个元素
// 函数内用反射拷贝与外部copy后传入性能对比（10000个元素中取5个）,因此函数默认接受copy后的切片
// reflect: BenchmarkRandList-12    	    4292	    254447 ns/op
// copy:    BenchmarkRandList-12    	  136270	      8746 ns/op
// result:  请传入copy后的切片,list内的元素顺序将被修改
func RandList(list interface{}, n int) interface{} {
	value := reflect.ValueOf(list)
	if value.Type().Kind() != reflect.Slice {
		panic("input must be a slice")
	}
	length := value.Len()
	if length < n {
		n = length
	}
	for i := 0; i < n; i++ {
		index := RandInt(0, length-i-1)
		tmp := reflect.ValueOf(value.Index(index).Interface())
		value.Index(index).Set(value.Index(length - i - 1))
		value.Index(length - i - 1).Set(tmp)
	}
	return value.Slice(length-n, length).Interface()
}
