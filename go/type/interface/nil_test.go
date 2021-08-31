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
	fmt.Println(models[0] == nil)
}

type AB interface {
	A
	B
}

type A interface {
	A()
}

type B interface {
	B()
}

type a struct{}

func (a a) A() {}

type b struct{}

func (b b) B() {}

func GetAB() AB {
	return nil
}

func GetIA() A {
	return nil
}

func GetA() *a {
	return nil
}

func TestGetNil(t *testing.T) {
	x := GetIA()
	fmt.Println(x == nil)
	x = GetAB()
	fmt.Println(x == nil)
	x = GetA()
	fmt.Println(x == nil)
}

func TestInterfaceAst(t *testing.T) {
	var num interface{}
	num = uint(123)
	fmt.Println(num.(int))
}
