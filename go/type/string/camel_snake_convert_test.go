// Time        : 2019/11/29
// Description :

package string

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(CamelToSnake("RoleName"))
	fmt.Println(CamelToSnake("roleName"))
	fmt.Println(CamelToSnake("role_name"))

	fmt.Println(SnakeToCamel("role_name"))
	fmt.Println(SnakeToCamel("_role_name"))
	fmt.Println(SnakeToCamel("role_Name"))
}
