// Time        : 2019/11/08
// Description :

package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestFields(t *testing.T) {
	str := "hello world"
	fmt.Println(strings.Fields(str))
	str = "hello      wor ld"
	fmt.Println(strings.Fields(str))
}
