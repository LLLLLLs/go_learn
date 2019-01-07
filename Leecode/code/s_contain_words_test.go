// Time        : 2019/01/03
// Description :

package code

import (
	"testing"
)

func TestWordsContainer(t *testing.T) {
	WordsContainerWithTree("barfoothefoobarman", []string{"foo", "bar"})
}

func BenchmarkWordsContainer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WordsContainerWithTree("barfoothefoobarman", []string{"foo", "bar"})
	}
}
