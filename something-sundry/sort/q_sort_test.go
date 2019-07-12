// Time        : 2019/07/12
// Description :

package sort

import (
	"fmt"
	"go_learn/utils"
	"testing"
)

var nums []int

func init() {
	nums = make([]int, 10)
	for i := range nums {
		nums[i] = utils.RandInt(0, 10)
	}
}

func TestQSort(t *testing.T) {
	qSort(nums)
	fmt.Println(nums)
}

func TestQSortGo(t *testing.T) {
	qSortGo(nums)
	fmt.Println(nums)
}

// 使用协程反而更慢
// 使用wg length = 1000：
// BenchmarkQSort-12           2000            612874 ns/op               0 B/op          0 allocs/op
// BenchmarkQSortGo-12         1000           1669565 ns/op              19 B/op          1 allocs/op
// 使用channel length = 10：
// BenchmarkQSort-12       20000000                95.3 ns/op             0 B/op          0 allocs/op
// BenchmarkQSortGo-12       100000             16196 ns/op             962 B/op         10 allocs/op
func BenchmarkQSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qSort(nums)
	}
}

func BenchmarkQSortGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		qSortGo(nums)
	}
}
