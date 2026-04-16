package script

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

func TestB_Nil(t *testing.T) {
	a := new(A)
	a.B.Func()
	a.B.Empty()
}

type AA struct{}

func (*AA) Print() {
	fmt.Println("aa")
}

type BB struct{}

func (*BB) Print() {
	fmt.Println("bb")
}

func TestNil(t *testing.T) {
	aa := new(AA)
	bb := new(BB)
	cc := new(AA)
	fmt.Println(unsafe.Pointer(aa) == unsafe.Pointer(bb))
	fmt.Println(unsafe.Pointer(cc) == unsafe.Pointer(bb))
	reflect.NewAt(reflect.TypeOf(&AA{}), unsafe.Pointer(bb)).Elem().Interface().(*AA).Print()
	reflect.NewAt(reflect.TypeOf(&BB{}), unsafe.Pointer(aa)).Elem().Interface().(*BB).Print()
	fmt.Println(unsafe.Pointer(aa))

	list := []int32{1, 2, 3}
	list = list[:len(list)-1]
	fmt.Println(list)

	for i := 0; i < 1; i++ {
		fmt.Println(i)
	}
}

func TestDiv(t *testing.T) {
	fmt.Println(float64(0)/float64(100000) >= float64(0))
}

func TestFloat64(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(int64(float64(0) * (0.01 * float64(i))))
	}
}

type C struct {
	a int
}

func (i *C) A() {
	i.a = 1
}

func (i C) AA() {
	fmt.Println(i.a)
}

func TestCFunc(t *testing.T) {
	c := C{a: 100}
	c.AA()
	c.A()
	c.AA()
}

func TestMap(t *testing.T) {
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		m[i] = i
	}
	count := 0
	for k, v := range m {
		fmt.Println(k, v)
		delete(m, k*2)
		count++
	}
	fmt.Println(m)
	fmt.Println(count)
}

// 将 x 左移 16 位，然后与 y 组合
func XYToIndex(x, y int16) int32 {
	return int32(x)<<16 | int32(y)&0xFFFF
}

func IndexToXY(index int32) (x, y int16) {
	x = int16(index >> 16)
	y = int16(index & 0xFFFF)
	return
}

func TestXYToIndex(t *testing.T) {
	fmt.Println(XYToIndex(193, 83))
}

func TestString(t *testing.T) {
	fmt.Println(string(getUint32()))
}

func getUint32() int32 {
	return 57
}

func TestNum(t *testing.T) {
	fmt.Println(math.MaxInt64 - 9223372036841959643)
	fmt.Println(7950734 + 20000000 - 9223372036841959643 + math.MaxInt64)
}

func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	for i := range 10 {
		m.Store(i, i)
	}
	m.Range(func(k, v interface{}) bool {
		m.Delete(k)
		fmt.Println(k, v)
		return true
	})
	fmt.Println("Delete")
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}

func stringToUTF8OctalEscape(s string) string {
	var result strings.Builder
	for _, r := range s {
		// 将rune转换为UTF-8字节序列
		bytes := []byte(string(r))

		// 将每个字节转换为八进制转义序列
		for _, b := range bytes {
			// 使用八进制格式，确保有3位数字，前面补0
			octal := fmt.Sprintf("\\%03o", b)
			result.WriteString(octal)
		}
	}

	return result.String()
}

func TestStringToUTF8(t *testing.T) {
	fmt.Println(stringToUTF8OctalEscape("你好"))
}

func TestInt64(t *testing.T) {
	fmt.Println(int(math.MaxInt64))
}

func TestDefer(t *testing.T) {
	a := 10
	if a == 1 {
		defer fmt.Println(123)
	}
	fmt.Println(321)
}

type TestA struct {
	A int
	B string
}

func TestReflect(t *testing.T) {
	typ := reflect.TypeOf(&TestA{})
	ta := &TestA{
		A: 1,
		B: "hello",
	}
	data, _ := json.Marshal(ta)
	d2 := reflect.New(typ.Elem()).Interface().(*TestA)
	fmt.Println(json.Unmarshal(data, d2))
	d2.A = 100
	fmt.Println(ta, d2)

}

func TestChan(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)

	// 方法1：读取时检查
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed!")
			break
		}
		fmt.Println("Received:", val)
	}
}

func TestMarshal(t *testing.T) {
	var ta *TestA
	data, err := json.Marshal(ta)
	fmt.Println(data, err)

	ta2 := new(TestA)
	err = json.Unmarshal(data, ta2)
	fmt.Println(ta2, err)
}

func TestMapConflict(t *testing.T) {
	m := make(map[int]int)
	m[1] = 1
	res := make(chan int)
	go func() {
		for i := 0; i < 10000; i++ {
			m[1]++
		}
	}()
	go func() {
		a := 0
		for i := 0; i < 10000; i++ {
			a = m[1]
		}
		res <- a
	}()
	fmt.Println(<-res)
}

func TestMapNil(t *testing.T) {
	m := make(map[int]map[int]int)
	fmt.Println(m[1], m[1] == nil)
	fmt.Println(m[1][2])
	sub := m[1]
	fmt.Println(sub[2])
	sub = nil
	fmt.Println(sub[2])
	var sm map[int]int
	fmt.Println(sm[2])
	sm[2] = 1
}

func TestSlice(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(list[len(list)-1:])
	var list1 []int = nil
	fmt.Println(len(list1))
	list1 = append(list1, 1, 2, 3)
	fmt.Println(list1)
}

func TestFormat(t *testing.T) {
	param := []any{int16(1), int16(2)}
	param = append(param, "hello")
	content := "您设置于【%d, %d】的自动续矿，因为【%s】未成功续矿"
	str := fmt.Sprintf(content, param...)
	fmt.Println(str)
}
