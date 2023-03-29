package unexported

import (
	"fmt"
	"golearn/go/reflect/unexported/define"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	typ := reflect.TypeOf(define.Test{})
	for i := 0; i < typ.NumMethod(); i++ {
		fmt.Println(typ.Method(i).Name)
		fmt.Println(typ.Method(i).Type.NumIn())
	}
	fmt.Println(reflect.ValueOf(define.Test{}).Type().Name())
	fmt.Println(typ.Name())
}
