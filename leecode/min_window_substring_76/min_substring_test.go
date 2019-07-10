// Time        : 2019/07/10
// Description :

package min_window_substring_76

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func TestMinWindow2(t *testing.T) {
	fmt.Println(minWindow2("ADOBECODEBANC", "ABC"))
}

func BenchmarkMinWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minWindow("ADOBECODEBANC", "ABC")
	}
}

func BenchmarkMinWindow2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minWindow2("ADOBECODEBANC", "ABC")
	}
}
