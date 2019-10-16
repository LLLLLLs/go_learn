// Time        : 2019/07/09
// Description :

package util

import (
	"fmt"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 0)
	DeepCopy(&dst, &src)
	dst[0] = 10
	fmt.Println(dst)
	fmt.Println(src)
}
