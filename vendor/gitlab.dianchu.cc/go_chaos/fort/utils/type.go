package utils

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

type zeroable interface {
	IsZero() bool
}

// 自定义类型转原始类型，参考XORM convert.go asKind
func AsKind(val reflect.Value) (interface{}, error) {
	//直接返回原始类型
	switch val.Interface().(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64, string, bool, time.Time:
		return val.Interface(), nil
	}
	// 转换自定义类型
	switch val.Kind() {
	case reflect.Int:
		return int(val.Int()), nil
	case reflect.Int8:
		return int8(val.Int()), nil
	case reflect.Int16:
		return int16(val.Int()), nil
	case reflect.Int32:
		return int32(val.Int()), nil
	case reflect.Int64:
		return int64(val.Int()), nil
	case reflect.Uint:
		return int64(val.Int()), nil
	case reflect.Uint8:
		return uint8(val.Int()), nil
	case reflect.Uint16:
		return uint16(val.Int()), nil
	case reflect.Uint32:
		return uint32(val.Int()), nil
	case reflect.Uint64:
		return int64(val.Int()), nil
	case reflect.Float32:
		return float32(val.Int()), nil
	case reflect.Float64:
		return float64(val.Int()), nil
	case reflect.Bool:
		return val.Bool(), nil
	case reflect.String:
		return val.String(), nil
	case reflect.Struct:
		if val.Type() == reflect.TypeOf(time.Time{}) {
			return val.Interface().(time.Time).Format(time.RFC3339Nano), nil
		}
	}
	return nil, errors.New("not support condition type")
}

// Go的每个类型都有对应的零值,所以无法判断是否赋值过。所以等于零值的字段默认判断为未赋值
func IsZero(val reflect.Value) bool {
	var (
		interfaceVal = val.Interface()
	)
	//这里使用switch比直接使用反射判断零值更快
	switch v := interfaceVal.(type) {
	case int:
		return v == 0
	case int8:
		return v == 0
	case int16:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	case uint:
		return v == 0
	case uint8:
		return v == 0
	case uint16:
		return v == 0
	case uint32:
		return v == 0
	case uint64:
		return v == 0
	case float32:
		return v == 0
	case float64:
		return v == 0
	case bool:
		return v == false
	case string:
		return v == ""
	case zeroable: // 比如time.Time
		return v.IsZero()
	default: //Struc,List在前面已经被转换了,而且sql/database不支持Map,Slice Func,使用这里只判断普通类型
		z := reflect.Zero(val.Type())
		return val.Interface() == z.Interface()
	}
	return true
}

//interface to string
func InterfaceToString(data interface{}) (string, error) {
	var (
		dataValue reflect.Value
		dataKind  interface{}
		err       error
	)
	dataValue = reflect.ValueOf(data)
	if dataKind, err = AsKind(dataValue); err != nil {
		return "", err
	}

	//这里使用switch比直接使用反射判断零值更快
	switch v := dataKind.(type) {
	case int:
		return strconv.Itoa(v), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(int64(v), 10), nil
	case uint:
		return strconv.FormatInt(int64(v), 10), nil
	case uint8:
		return strconv.FormatInt(int64(v), 10), nil
	case uint16:
		return strconv.FormatInt(int64(v), 10), nil
	case uint32:
		return strconv.FormatInt(int64(v), 10), nil
	case uint64:
		return strconv.FormatInt(int64(v), 10), nil
	case float32:
		return strconv.FormatInt(int64(v), 10), nil
	case float64:
		return strconv.FormatInt(int64(v), 10), nil
	case string:
		return v, nil
	default:
		return "", errors.New("type error ")
	}
}
