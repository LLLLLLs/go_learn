// Time        : 2019/10/16
// Description :

package error

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var errTest = errors.New("test")

func testReturnErr() (err error) {
	a, err := 1, errTest
	_ = a
	return
}

func TestReturnErr(t *testing.T) {
	err := testReturnErr()
	assert.New(t).Equal(errTest, err)
}
