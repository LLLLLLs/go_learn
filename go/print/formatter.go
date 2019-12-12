// Time        : 2019/11/12
// Description :

package print

import (
	"fmt"
	"strconv"
)

type Foo struct {
	A int
	B string
	C bool
}

func (f2 Foo) Format(f fmt.State, c rune) {
	mustWriteOk(f.Write([]byte("exec Foo.Format():")))
	mustWriteOk(f.Write([]byte(f2.B)))
	mustWriteOk(f.Write([]byte(",")))
	mustWriteOk(f.Write([]byte(intToString(f2.A))))
	mustWriteOk(f.Write([]byte(",")))
	mustWriteOk(f.Write([]byte(boolToString(f2.C))))
}

func mustWriteOk(_ int, err error) {
	if err != nil {
		panic(err)
	}
}

func intToString(i int) string {
	return strconv.Itoa(i)
}

func boolToString(b bool) string {
	var boolMapper = map[bool]string{
		true:  "true",
		false: "false",
	}
	return boolMapper[b]
}

func (f2 Foo) String() string {
	return "exec Foo.String()"
}
