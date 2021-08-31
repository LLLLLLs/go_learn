//@author: lls
//@time: 2020/12/23
//@desc:

package regexp

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	reg := regexp.MustCompile(`^((\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.){3}(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])$`)
	fmt.Println(reg.MatchString("10.255.255.225"))
}
