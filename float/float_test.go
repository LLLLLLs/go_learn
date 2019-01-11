// Time        : 2019/01/10
// Description :

package float

import "testing"

// Benchmark test: go test -bench=. -run=none

func BenchmarkAddFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addFloat()
	}
}

func BenchmarkAddZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addZero()
	}
}
