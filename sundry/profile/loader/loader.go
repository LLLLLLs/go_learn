// Time        : 2019/09/09
// Description :

package loader

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var Separator = ":"

//数据加载器
type Loader interface {
	Load() (dataMap map[string]map[string]interface{}, dataList map[string]interface{}, err error)
}

func IndexToString(i interface{}) string {
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

func generateIndexKey(value reflect.Value) string {
	typ := value.Type()
	indexes := make([]string, 0)
	for i := 0; i < typ.NumField(); i++ {
		index := typ.Field(i).Tag.Get("index")
		if index != "" {
			indexes = append(indexes, IndexToString(value.Field(i).Interface()))
		}
	}
	key := strings.Join(indexes, Separator)
	return key
}
