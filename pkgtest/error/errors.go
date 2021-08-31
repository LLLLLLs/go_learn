// Time        : 2019/01/04
// Description :

package error

import (
	"fmt"
	"github.com/pkg/errors"
)

func first() error {
	return errors.New("first")
}

func second() error {
	return errors.Wrap(first(), "second")
}

//func third() error {
//	return errors.Wrap(second(), "third")
//}

type Error struct {
	errCode uint8
}

func (e *Error) Error() string {
	if e == nil {
		return "nil"
	}
	switch e.errCode {
	case 1:
		return "file not found"
	case 2:
		return "time out"
	case 3:
		return "permission denied"
	default:
		return "unknown error"
	}
}

func checkError(err error) {
	fmt.Println(err == nil)
	if err != nil {
		panic(err)
	}
}
