// Time        : 2019/09/03
// Description :

package typeconvert

import (
	"reflect"
	"strconv"
)

func NumberToString(i interface{}) string {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64)
	default:
		panic("not a number")
	}
}
