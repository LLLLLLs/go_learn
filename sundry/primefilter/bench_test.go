// Time        : 2019/07/08
// Description :

package primefilter

import (
	"golearn/util/randutil"
	"testing"
)

func BenchmarkNthPrimeWithList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := randutil.RandInt(5, 10)
		NthPrimeWithList(n)
	}
}

func BenchmarkNthPrimeWithFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := randutil.RandInt(5, 10)
		NthPrimeWithFilter(n)
	}
}
