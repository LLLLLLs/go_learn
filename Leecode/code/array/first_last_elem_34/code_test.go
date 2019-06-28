// Time        : 2019/06/28
// Description :

package first_last_elem_34

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func BenchmarkSearchRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	}
}
