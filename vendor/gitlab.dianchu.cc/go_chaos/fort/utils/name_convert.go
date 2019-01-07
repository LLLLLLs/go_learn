/*
Author: Minsi Ruan
Data: 2018/6/5 9:16
*/

package utils

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

func WriteBuffer(b *bytes.Buffer, c interface{}) {
	switch val := c.(type) {
	case int:
		b.WriteString(strconv.Itoa(val))
	case int64:
		b.WriteString(strconv.FormatInt(val, 10))
	case uint:
		b.WriteString(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.WriteString(strconv.FormatUint(val, 10))
	case string:
		b.WriteString(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
}

// 驼峰式写法转为下划线写法
func ToUnderscoreName(name string) string {
	b := new(bytes.Buffer)
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				WriteBuffer(b, '_')
			}
			WriteBuffer(b, unicode.ToLower(r))
		} else {
			WriteBuffer(b, r)
		}
	}
	return b.String()
}

// 下划线转为驼峰
func ToCamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
