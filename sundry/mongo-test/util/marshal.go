// Time        : 2019/10/24
// Description :

package util

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"golearn/util"
	"reflect"
)

var errMustPtrType = errors.New("must provide a ptr model")
var errMustSliceType = errors.New("must provide a slice model")

func MarshalExtend(extend interface{}, model interface{}) {
	marshal(extend, model)
}

func marshal(extend interface{}, model interface{}) {
	switch extend.(type) {
	case bson.A: // Array
		marshalA(extend.(bson.A), model)
	case bson.D, map[string]interface{}: // Initial Document or Fort Document
		marshalBson(extend, model)
	default: // basic type:int,string,float...
		marshalDefault(extend, model)
	}
}

func marshalA(extend bson.A, model interface{}) {
	rv := reflect.ValueOf(model)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	} else {
		panic(errMustPtrType)
	}
	if rv.Kind() != reflect.Slice {
		panic(errMustSliceType)
	}
	length := len(extend)
	if length == 0 {
		return
	}
	slice := reflect.MakeSlice(rv.Type(), length, length)
	mt := slice.Index(0).Type()
	for i := range extend {
		m := reflect.New(mt)
		marshal(extend[i], m.Interface())
		slice.Index(i).Set(m.Elem())
	}
	rv.Set(slice)
}

func marshalBson(extend interface{}, model interface{}) {
	data, err := bson.Marshal(extend)
	util.MustNil(err)
	err = bson.Unmarshal(data, model)
	util.MustNil(err)
}

func marshalJson(extend interface{}, model interface{}) {
	data, err := json.Marshal(extend)
	util.MustNil(err)
	err = json.Unmarshal(data, model)
	util.MustNil(err)
}

func marshalDefault(extend interface{}, model interface{}) {
	reflect.ValueOf(model).Elem().Set(reflect.ValueOf(extend))
}
