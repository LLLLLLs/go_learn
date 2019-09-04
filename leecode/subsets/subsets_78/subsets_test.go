// Time        : 2019/07/10
// Description :

package subsets_78

import (
	"golearn/utils"
	"testing"
)

func TestSubSets(t *testing.T) {
	utils.Print2DimensionList(subsets([]int{1, 2, 3}))
}

func TestSubSetsBinary(t *testing.T) {
	utils.Print2DimensionList(subsetsBinary([]int{1, 2, 3}))
}

func BenchmarkSubsetsBacktrack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subsets([]int{1, 2, 3, 4, 5, 6, 7})
	}
}

func BenchmarkSubsetsBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subsetsBinary([]int{1, 2, 3, 4, 5, 6, 7})
	}
}
