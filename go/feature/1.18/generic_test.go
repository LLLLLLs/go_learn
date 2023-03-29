package __18

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	uintM := map[string]uint{
		"1": 1,
		"2": 2,
	}
	uint8M := map[string]uint8{
		"3": 3,
		"4": 4,
	}
	fmt.Println(sum(uintM))
	fmt.Println(sum(uint8M))
}

// func TestStruct(t *testing.T) {
// 	a := &addA{num: 0}
// 	fmt.Println(add(a, 10))
// }
type Bucket interface {
}

func TestComparable(t *testing.T) {
	m := make(map[Bucket]interface{})
	m[1] = "1"
	m["2"] = "2"
	m["1"] = "3"
	fmt.Println(m)
}

func TestToUint(t *testing.T) {
	res := toUint("123", uint(0))
	fmt.Printf("%d,%T\n", res, res)
}

type d1 struct {
	a int
}

type d2 struct {
	b int
}

func TestValue(t *testing.T) {
	var (
		i1  interface{} = d1{a: 100}
		i1p interface{} = &d1{a: 100}
		i2  interface{} = d2{b: 1000}
		i2p interface{} = &d2{b: 100}
	)
	fmt.Println(value[d1](i1))
	fmt.Println(value[d1](i1p))
	fmt.Println(value[d1](i2))
	fmt.Println(value[d2](i2p))
}
