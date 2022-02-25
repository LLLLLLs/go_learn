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
	s1 := append([]int(nil), 1)
	fmt.Println(s1)
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

var ri RankInfo

func TestMakeList(t *testing.T) {
	elem := reflect.ValueOf(ri)
	typ := elem.Type()
	ss := reflect.MakeSlice(reflect.SliceOf(typ), 10, 20)
	fmt.Println(ss.Type().Name())
	fmt.Println(ss.Interface())
}

func TestSameSlice(t *testing.T) {
	fmt.Println([3]int{1, 2, 3} == [3]int{1, 2, 3})
}

func TestModifySlice(t *testing.T) {
	nums := [20]int{}
	for i := range nums {
		nums[i] = i
	}
	slice := nums[5:10]
	s2 := make([]int, 5)
	slice = append(slice, s2...)
	fmt.Println(slice)
	fmt.Println(nums)
}

func insertList(l []int, n int) []int {
	if len(l) == 0 {
		return []int{n}
	}
	if n < l[0] {
		return append([]int{n}, l...)
	} else if n > l[len(l)-1] {
		return append(l, n)
	}
	left, right := 0, len(l)-1
	insertIndex := func(index int) []int {
		first := append(l[:index], n)
		return append(first, l[index:]...)
		// return append(append(append([]int{}, l[:index]...), n), l[index:]...)
	}
	for left < right {
		mid := (left + right) / 2
		if l[mid] == n {
			return insertIndex(mid)
		} else if l[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return insertIndex(left)
}

func TestInsert(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9}
	fmt.Println(slice, "insert 5 ==>", insertList(slice, 5))
	fmt.Println(slice, "insert 6 ==>", insertList(slice, 6))
	fmt.Println(slice, "insert 8 ==>", insertList(slice, 8))
	fmt.Println(slice, "insert 10 ==>", insertList(slice, 10))
}

func Test_Append(t *testing.T) {
	ss := make([]int, 0, 4)
	ss2 := append(ss, 1)
	fmt.Printf("%p:%v\n", ss, ss)
	fmt.Printf("%p:%v\n", ss2, ss2)

	_ = append(ss[:], 1)
	fmt.Println(ss)

	ss = make([]int, 2, 4)
	ss[0] = 1
	ss[1] = 2
	ss2 = append(ss[:1], 1)
	fmt.Println(ss2)
	fmt.Println(ss)
}

func TestAppendWhenRange(t *testing.T) {
	var list = make([]int, 5, 10)
	var app = false
	for _, j := range list {
		fmt.Println(j)
		if app {
			list = append(list, list[len(list)-1])
		}
		app = !app
	}
}

func TestInitial(t *testing.T) {
	slice := []int{
		0:  1,
		2:  3,
		1:  2,
		10: 11,
	}
	fmt.Println(slice)
}

func TestSliceCopy(t *testing.T) {
	a := make([]int, 5, 10)
	fmt.Println(a)
	b := a[1:]
	fmt.Printf("value=%v,len=%d,cap=%d\n", b, len(b), cap(b))
	b[0] = 2
	fmt.Println(a)
}

func TestSliceIter(t *testing.T) {
	var slice = []int{1, 2, 3, 4, 5}
	var toSlice = make([]*int, len(slice))
	for i, v := range slice { // i:int; v:*int
		toSlice[i] = &v
	}
	fmt.Println(slice)
	fmt.Println(toSlice)
}

func TestTrimNum(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	newSlice1 := trimNum(slice, 3)
	newSlice2 := trimNum(slice, 5)
	newSlice3 := trimNum(newSlice1, 4)
	fmt.Println(slice)
	fmt.Println(newSlice1)
	fmt.Println(newSlice2)
	fmt.Println(newSlice3)
}

func trimNum(list []int, num int) []int {
	for i := range list {
		if list[i] == num {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}
	return list
}

func TestTrimNum1(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)
}

func TestSliceNil(t *testing.T) {
	var slice []int
	fmt.Println(isNil(slice))
	slice = make([]int, 0)
	fmt.Println(isNil(slice))
}

func isNil(s []int) bool {
	return s == nil
}

func TestEmptyLen(t *testing.T) {
	slice := make([]int, 0)
	slice10 := make([]int, 10)
	fmt.Println(cap(slice))
	fmt.Println(cap(slice10))
}
