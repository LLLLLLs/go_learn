// Time        : 2019/07/16
// Description :

package file

import (
	"fmt"
	"go_learn/utils"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, err := os.Open("testfile\\type.go")
	utils.OkOrPanic(err)
	defer file.Close()
	body := make([]byte, 100)
	n, err := file.Read(body)
	utils.OkOrPanic(err)
	fmt.Println(string(body), n)
}
