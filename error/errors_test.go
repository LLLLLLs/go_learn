// Time        : 2019/01/04
// Description :

package error

import (
	"arthur/utils/errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := third()
	fmt.Println(errors.Cause(err))
	fmt.Printf("%+v\n", err)
}
