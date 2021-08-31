// Time        : 2019/01/04
// Description :

package error

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestError(t *testing.T) {
	err := second()
	fmt.Println(errors.Cause(err))
	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", nil)
}

func TestEqual(t *testing.T) {
	e1 := first()
	e2 := first()
	fmt.Println(e1 == e2)
}

func TestCheckError(t *testing.T) {
	var e *Error = nil
	checkError(nil)
	checkError(e)
}
