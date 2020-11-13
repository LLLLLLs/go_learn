//@author: lls
//@time: 2020/08/27
//@desc:

package compare_version_165

import (
	"fmt"
	"testing"
)

func TestCompare1(t *testing.T) {
	fmt.Println(compareVersion("0.1", "1.1"))
	fmt.Println(compareVersion("1.0.1", "1"))
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
}
