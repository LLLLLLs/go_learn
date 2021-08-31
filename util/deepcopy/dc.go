// @author: lls
// @time: 2021/07/02
// @desc: 支持非导出字段的拷贝

package deepcopy

import (
	"reflect"
	"time"
	"unsafe"
)

// Copy creates a deep copy of whatever is passed to it and returns the copy
// in an interface{}.  The returned value will need to be asserted to the
// correct type.
func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	// Make the interface a reflect.Value
	original := reflect.ValueOf(src)

	// Make a copy of the same type as the original.
	// Recursively copy the original.
	// Return the copy as an interface.
	return copyRecursive(original).Interface()
}

// copyRecursive does the actual copying of the interface. It currently has
// limited support for what it can handle. Add as needed.
func copyRecursive(original reflect.Value) reflect.Value {
	var notPtr bool
	if original.Kind() != reflect.Ptr && !original.CanAddr() {
		notPtr = true
		val := original
		original = reflect.New(original.Type())
		original.Elem().Set(val)
	}
	cpy := reflect.New(original.Type()).Elem()
	// handle according to original's Kind
	switch original.Kind() {
	case reflect.Ptr:
		// Get the actual value being pointed to.
		originalValue := original.Elem()

		// if  it isn't valid, return.
		if !originalValue.IsValid() {
			return cpy
		}
		cpy = reflect.New(originalValue.Type())
		cpy.Elem().Set(copyRecursive(originalValue))

	case reflect.Interface:
		// If this is a nil, don't do anything
		if original.IsNil() {
			return cpy
		}
		// Get the value for the interface, not the pointer.
		originalValue := original.Elem()

		// Get the value by calling Elem().
		copyValue := copyRecursive(originalValue)
		cpy.Set(copyValue)

	case reflect.Struct:
		t, ok := original.Interface().(time.Time)
		if ok {
			cpy.Set(reflect.ValueOf(t))
			return cpy
		}
		// Go through each field of the struct and copy it.
		for i := 0; i < original.NumField(); i++ {
			oriField, cpyField := original.Field(i), cpy.Field(i)
			// The Type's StructField for a given field is checked to see if StructField.PkgPath
			// is set to determine if the field is exported or not because CanSet() returns false
			// for settable fields.  I'm not sure why.  -mohae
			if original.Type().Field(i).PkgPath != "" {
				oriField = reflect.NewAt(oriField.Type(), unsafe.Pointer(oriField.UnsafeAddr())).Elem()
				cpyField = reflect.NewAt(cpyField.Type(), unsafe.Pointer(cpyField.UnsafeAddr())).Elem()
			}
			cpyField.Set(copyRecursive(oriField))
		}

	case reflect.Slice:
		if original.IsNil() {
			return cpy
		}
		// Make a new slice and copy each element.
		cpy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i++ {
			cpy.Index(i).Set(copyRecursive(original.Index(i)))
		}

	case reflect.Map:
		if original.IsNil() {
			return cpy
		}
		cpy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			copyValue := copyRecursive(originalValue)
			copyKey := Copy(key.Interface())
			cpy.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	default:
		cpy.Set(original)
	}
	if notPtr {
		cpy = cpy.Elem()
	}
	return cpy
}
