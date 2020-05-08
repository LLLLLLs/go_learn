//@time:2020/04/02
//@desc:

package _interface

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIOCopy(t *testing.T) {
	w, err := io.Copy(WriteF(MyWriteFunction), os.Stdin)
	fmt.Println(w, err)
}

func TestCopy(t *testing.T) {
	Proxies := make(map[int]ProxyWrite, 2)
	Proxies[1] = new(Implementer)
	Proxies[2] = new(Implementer)

	/* runs and uses the Write method normally */
	io.Copy(Proxies[1], strings.NewReader("Hello world"))
	fmt.Println(Proxies[1].Result())
	/* gets ./main.go:45: method Proxies[1].SpecialWrite is not an expression, must be called */
	io.Copy(WriteFunc(Proxies[2].SpecialWrite), strings.NewReader("Hello world 123"))
	fmt.Println(Proxies[2].Result())
}
