// Time        : 2019/01/18
// Description :

package slice

import (
	"fmt"
	"testing"
)

func TestBase(t *testing.T) {
	sliceBase()
}

func TestAppend(t *testing.T) {
	sliceAppend()
}

func TestSort(t *testing.T) {
	var init = make([]RankInfo, len(list))
	copy(init, list)
	sortSlice()
	for i := range list {
		fmt.Println(init[i], "==>", list[i])
	}
}
