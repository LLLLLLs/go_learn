// Time        : 2019/07/08
// Description :

package _struct

import (
	"fmt"
	"go_learn/leecode/linked_list/base"
	"reflect"
	"testing"
)

func TestEmpty(t *testing.T) {
	ea, eb := new(empty), new(empty)
	fmt.Printf("ea=%p,eb=%p\n", ea, eb)
	ea.hello()
	eb.world()
	fmt.Println(reflect.TypeOf(base.ListNode{}).Name())
}
