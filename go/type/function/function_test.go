// Time        : 2019/08/22
// Description :

package function

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	f1 := NewFunc()
	fmt.Println(f1())
	f2 := f1
	fmt.Println(f2())
	fmt.Println(f1())
	printFunc(f1)
	printFunc(f1)
	printFunc(f1)
	cf := ContentFunc{f: f1}
	cf.Do()
	cf.Do()
	printFunc(f1)
}

func printFunc(f func() int) {
	fmt.Println(f())
}
