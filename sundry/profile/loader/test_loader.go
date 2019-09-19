// Time        : 2019/09/09
// Description :

package loader

import (
	"golearn/sundry/profile/model"
	"reflect"
	"sync"
)

const TEST_VERSION = "version:test"

type TestLoader struct {
	data map[string]map[string]interface{}
	list map[string]interface{}
}

func NewTestLoader() *TestLoader {
	definer := &TestLoader{
		data: make(map[string]map[string]interface{}),
		list: make(map[string]interface{}),
	}
	definer.AddTable([]model.Version{{Version: TEST_VERSION}})
	return definer
}

func (l *TestLoader) Load() (dataMap map[string]map[string]interface{}, dataList map[string]interface{}, err error) {
	return l.data, l.list, nil
}

// 添加一张表的静态配置数据，应传入一个配置切片
func (l *TestLoader) AddTable(conf interface{}) {
	slice := reflect.ValueOf(conf)
	var data = make(map[string]interface{})
	var name string
	var once sync.Once
	for i := 0; i < slice.Len(); i++ {
		row := slice.Index(i)
		once.Do(func() {
			name = row.Type().Name()
		})
		key := generateIndexKey(row)
		data[key] = row.Interface()
	}
	l.data[name] = data
	l.list[name] = slice.Interface()
}
