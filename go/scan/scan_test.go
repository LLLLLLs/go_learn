//@time:2020/04/02
//@desc:

package scan

import (
	"fmt"
	"testing"
)

func TestScanln(t *testing.T) {
	var input string
	e, err := fmt.Scanln(&input)
	fmt.Println(err)
	fmt.Println(e)
	fmt.Println(input)
}
