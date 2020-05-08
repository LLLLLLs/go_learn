//@time:2020/01/04
//@desc:

package strconv

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestParseInt(t *testing.T) {
	str := "032768"
	sleep()
	i, err := strconv.ParseInt(str, 10, 64)
	ast := assert.New(t)
	ast.Nil(err)
	t.Log(i)
}

func sleep() {
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
	}
}

func TestParseFloat(t *testing.T) {
	str := "123.45"
	f, err := strconv.ParseFloat(str, 64)
	ast := assert.New(t)
	ast.Nil(err)
	t.Log(f)
}
