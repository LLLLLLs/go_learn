// Time        : 2019/11/08
// Description :

package runtime

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	Stack()
}

func TestGetGoroutineId(t *testing.T) {
	fmt.Println(GetGoroutineId())
}
