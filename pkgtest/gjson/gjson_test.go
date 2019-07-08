// Time        : 2019/06/25
// Description :

package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

const jsonStr = `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44},
    {"first": "Roger", "last": "Craig", "age": 68},
    {"first": "Jane", "last": "Murphy", "age": 47}
  ]
}`

const jsonLine = `
	{"first": "Dale", "last": "Murphy", "age": 44}
    {"first": "Roger", "last": "Craig", "age": 68}
    {"first": "Jane", "last": "Murphy", "age": 47}
`

func TestGJson(t *testing.T) {
	value := gjson.Get(jsonStr, "name.first")
	fmt.Println(value.String())

	value = gjson.Get(jsonStr, "friends.#")
	fmt.Println(value.Int())

	value = gjson.Get(jsonStr, "friends.1.first")
	fmt.Println(value.Array())

	value = gjson.Get(jsonStr, `friends`)
	fmt.Println(value)

	value = gjson.Get(jsonStr, `friends`)
	fmt.Println(value.Array()[0])

	value = gjson.Get(jsonStr, `friends.0`)
	fmt.Println(value)

	value = gjson.Get(jsonLine, "..1")
	fmt.Println(value)

	value = gjson.Get(jsonLine, `..#[first!%"D*"]#.last`)
	fmt.Println(value)
}
