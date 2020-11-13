// Time        : 2019/09/26
// Description :

package stupiddi

import (
	"reflect"
)

type DI struct {
	constructor map[reflect.Type]interface{}
	values      map[reflect.Type]interface{}
}

func NewDI() *DI {
	return &DI{
		constructor: make(map[reflect.Type]interface{}),
		values:      make(map[reflect.Type]interface{}),
	}
}

var errType = reflect.TypeOf((*error)(nil)).Elem()

func isErr(t reflect.Type) bool {
	return t.Implements(errType)
}

func (di *DI) Provide(constructor ...interface{}) {
	for _, ctor := range constructor {
		t := reflect.TypeOf(ctor)
		if t.Kind() != reflect.Func {
			panic("constructor must be a func")
		}
		for i := 0; i < t.NumOut(); i++ {
			out := t.Out(i)
			if !isErr(out) {
				di.constructor[out] = ctor
			}
		}
	}
}

func (di *DI) Get(model interface{}) interface{} {
	value := reflect.ValueOf(model)
	if value.Type().Kind() == reflect.Ptr {
		value = value.Elem()
	}
	typ := value.Type()
	v, ok := di.values[typ]
	if ok {
		return v
	}
	return di.create(typ)
}

func (di *DI) create(typ reflect.Type) interface{} {
	ctor, ok := di.constructor[typ]
	if !ok {
		panic("no such constructor")
	}
	cType := reflect.TypeOf(ctor)
	var params = make([]reflect.Value, cType.NumIn())
	for i := range params {
		params[i] = reflect.ValueOf(di.Get(reflect.New(cType.In(i)).Interface()))
	}
	returned := reflect.ValueOf(ctor).Call(params)
	for _, rv := range returned {
		if isErr(rv.Type()) {
			if err := rv.Interface().(error); err != nil {
				panic(err)
			}
		}
		di.values[rv.Type()] = rv.Interface()
		if rv.Type() == typ {
			return rv.Interface()
		}
	}
	panic("no value")
}
