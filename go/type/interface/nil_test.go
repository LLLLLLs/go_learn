// Time        : 2019/08/26
// Description :

package _interface

import (
	"fmt"
	"reflect"
	"testing"
)

type Nil struct {
	A int
}

type Nil2 struct {
	B int
}

func TestNil(t *testing.T) {
	var i interface{}
	fmt.Println(reflect.ValueOf(i).CanSet())
	i = &Nil{A: 1}
	fmt.Println(reflect.ValueOf(i).Elem().CanSet())
	reflect.ValueOf(i).Elem().Set(reflect.ValueOf(&Nil{A: 2}).Elem())
	i = &Nil2{B: 2}
	reflect.ValueOf(i).Elem().Set(reflect.ValueOf(&Nil2{B: 2}).Elem())
}

func TestSliceInterface(t *testing.T) {
	nilTest(nil)
}

func nilTest(models ...interface{}) {
	fmt.Println(models[0] != nil)
}
