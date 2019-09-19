// Time        : 2019/09/03
// Description :

package typeconvert

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToString(i interface{}) string {
	switch result := i.(type) {
	case string:
		return result
	case int8, int16, int32, int, int64:
		return strconv.FormatInt(reflect.ValueOf(i).Int(), 10)
	case uint8, uint16, uint32, uint, uint64:
		return strconv.FormatUint(reflect.ValueOf(i).Uint(), 10)
	case float32:
		return fmt.Sprintf("%g", result)
	case float64:
		return strconv.FormatFloat(result, 'g', -1, 64)
	default:
		return fmt.Sprintf("%v", result)
	}
}
