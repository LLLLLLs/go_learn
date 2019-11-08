// Time        : 2019/08/22
// Description :

package function

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func NewFunc() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

type ContentFunc struct {
	f func() int
}

func (cf ContentFunc) Do() {
	fmt.Println(cf.f())
}

func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	totalName := runtime.FuncForPC(pc).Name()
	return extractFuncName(totalName)
}

func GetTargetFuncName(f interface{}) string {
	fv := reflect.ValueOf(f)
	if fv.Type().Kind() != reflect.Func {
		panic("f must be a func")
	}
	totalName := runtime.FuncForPC(fv.Pointer()).Name()
	return extractFuncName(totalName)
}

func extractFuncName(totalName string) string {
	list := strings.Split(totalName, ".")
	return list[len(list)-1]
}

func TraceWithFuncForPC() {
	var pc = make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("%s:%d,%s\n", file, line, f.Name())
	}
}

func TraceWithFrames() {
	var pc = make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])
	for {
		f, more := frames.Next()
		fmt.Printf("%s:%d,%s\n", f.File, f.Line, f.Function)
		if !more {
			break
		}
	}
}
