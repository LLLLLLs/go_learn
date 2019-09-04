// Time        : 2019/06/28
// Description :

package first_last_elem_34

import (
	"fmt"
	"golearn/utils"
	"sort"
	"testing"
)

var list []int
var target int

func init() {
	list = make([]int, 200)
	for i := range list {
		list[i] = utils.RandInt(1, 200)
	}
	sort.Ints(list)
	target = utils.RandInt(1, 200)
}

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func TestSearchRange2(t *testing.T) {
	fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange2([]int{5, 7, 7, 8, 8, 10}, 6))
}

// 大概以200的长度作区分，长度短时从头遍历效率高，长度长时二分效率高
func BenchmarkSearchRang(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchRange(list, target)
	}
}

func BenchmarkSearchRange2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchRange2(list, target)
	}
}
