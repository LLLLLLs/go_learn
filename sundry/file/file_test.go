// Time        : 2019/07/16
// Description :

package file

import (
	"fmt"
	"golearn/util"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, err := os.Open("testfile\\type.go")
	util.MustNil(err)
	defer file.Close()
	body := make([]byte, 100)
	n, err := file.Read(body)
	util.MustNil(err)
	fmt.Println(string(body), n)
}
