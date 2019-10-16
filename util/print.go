// Time        : 2019/07/03
// Description :

package util

import (
	"fmt"
	"reflect"
)

func Print2DimensionList(list interface{}) {
	v := reflect.ValueOf(list)
	for i := 0; i < v.Len(); i++ {
		fmt.Println(v.Index(i).Interface())
	}
	fmt.Println()
}
