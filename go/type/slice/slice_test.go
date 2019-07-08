// Time        : 2019/01/18
// Description :

package slice

import (
	"fmt"
	"reflect"
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

type SliceStruct struct {
	list []int
}

func (ss SliceStruct) List() []int {
	return ss.list
}

func TestList(t *testing.T) {
	ss := SliceStruct{list: []int{1, 2, 3}}
	list := ss.List()
	list[1] = 0
	fmt.Println(ss.List())
	fmt.Println(reflect.ValueOf(ss.list).Type())
}
