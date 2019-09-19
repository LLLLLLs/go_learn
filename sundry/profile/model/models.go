// Time        : 2019/09/05
// Description :

package model

import (
	"fmt"
	"reflect"
)

var models = make(map[string]reflect.Type)

func RegisterModel(model interface{}) {
	t := reflect.ValueOf(model).Type()
	name := t.Name()
	if _, ok := models[name]; ok {
		panic(fmt.Sprintf("repeat register: %s", name))
	}
	models[name] = t
}

func AllModels() map[string]reflect.Type {
	return models
}
