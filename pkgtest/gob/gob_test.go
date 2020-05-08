//@author: lls
//@time: 2020/04/10
//@desc:

package gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGob(t *testing.T) {
	ast := assert.New(t)

	gob.Register(A{})
	gob.Register(A{})

	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(B{AA: A{Y: 1}})
	ast.Nil(err)
	var data B
	decoder := gob.NewDecoder(buf)
	err = decoder.Decode(&data)
	ast.Nil(err)
	fmt.Println(data)

	var ia IA
	fmt.Println(reflect.TypeOf(ia) == reflect.TypeOf(interface{}(nil)))
	fmt.Println(reflect.TypeOf(A{}) == reflect.TypeOf(interface{}(nil)))
	var b B
	fmt.Println(reflect.TypeOf(b.AA) == reflect.TypeOf(interface{}(nil)))

	fmt.Println((*A)(nil) == nil)
	fmt.Println(interface{}((*A)(nil)) == nil)
}
