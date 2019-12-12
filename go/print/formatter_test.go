// Time        : 2019/11/12
// Description :

package print

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	foo := Foo{
		A: 1,
		B: "hello world",
		C: true,
	}
	fmt.Println(foo)
}
