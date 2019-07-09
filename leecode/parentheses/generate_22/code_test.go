// Time        : 2019/06/26
// Description :

package generate_22

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	fmt.Println(generateParenthesis1(3))
}

func TestGenerate2(t *testing.T) {
	fmt.Print(generateParenthesis2(3))
}

func BenchmarkGenerate1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generateParenthesis1(10)
	}
}

func BenchmarkGenerate2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generateParenthesis2(10)
	}
}
