// Time        : 2019/11/08
// Description :

package runtime

import (
	"fmt"
	"golearn/util"
	"runtime"
	"strconv"
	"strings"
)

func Stack() {
	buf := make([]byte, 100)
	n := runtime.Stack(buf, false)
	_ = n
	fmt.Println(string(buf))
}

func GetGoroutineId() int {
	buf := make([]byte, 1000)
	n := runtime.Stack(buf, false)
	idStr := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine"))[0]
	id, err := strconv.Atoi(idStr)
	util.MustNil(err)
	return id
}
