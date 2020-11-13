// Time        : 2019/08/07
// Description :

package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	str := "123"
	([]byte)(str)[1] = 'b'
	fmt.Printf("%p\n", &str)
	fmt.Printf("%p\n", ([]byte)(str))
}

func TestSplit(t *testing.T) {
	str := "hello;"
	list := strings.Split(str, ";")
	fmt.Println(list, len(list))
}
