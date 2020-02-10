//@time:2019/12/25
//@desc:

package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type RoleId string

func TestReflect(t *testing.T) {
	var roleId = RoleId("roleId")
	value := reflect.ValueOf(roleId)
	fmt.Println(value.String())
}
