// Time        : 2019/11/29
// Description :

package string

import (
	"bytes"
	"strings"
	"unicode"
)

// 驼峰式写法转为下划线写法
func CamelToSnake(name string) string {
	buffer := bytes.Buffer{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}

	return buffer.String()
}

// 下划线写法转为驼峰写法
func SnakeToCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
