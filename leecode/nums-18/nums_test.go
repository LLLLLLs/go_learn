// Time        : 2019/01/21
// Description :

package nums_18

import (
	"fmt"
	"golearn/util"
	"testing"
)

var testList = func() []int {
	l := make([]int, 60)
	for i := range l {
		l[i] = util.RandInt(-20, 20)
	}
	return l
}()

func TestFourNums1(t *testing.T) {
	fmt.Println(fourSum1([]int{1, 0, -1, 0, -2, 2}, 0))
}

func TestFourNums2(t *testing.T) {
	fmt.Println(fourSum2([]int{1, 0, -1, 0, -2, 2}, 0))
}

func TestFourNums3(t *testing.T) {
	fmt.Println(fourSum3([]int{-5, -2, -4, -2, -5, -4, 0, 0}, -13))
}

func TestFourNums4(t *testing.T) {
	fmt.Println(fourSum4([]int{-5, -2, -4, -2, -5, -4, 0, 0}, -13))
}

func BenchmarkNums1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fourSum1(testList, 0)
	}
}

func BenchmarkNums2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fourSum2(testList, 0)
	}
}

func BenchmarkNums3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fourSum3(testList, 0)
	}
}

func BenchmarkNums4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fourSum4(testList, 0)
	}
}
