//@time:2020/01/04
//@desc:

package strconv

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestParseInt(t *testing.T) {
	str := "032768"
	i, err := strconv.ParseInt(str, 10, 64)
	ast := assert.New(t)
	ast.Nil(err)
	t.Log(i)
}
