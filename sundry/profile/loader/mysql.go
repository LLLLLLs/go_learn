// Time        : 2019/09/09
// Description :

package loader

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"golearn/sundry/profile/model"
	"golearn/util"
	"reflect"
	"sync"
	"xorm.io/core"
)

var conn *xorm.Engine
var once sync.Once
var driverName = "mysql"

func getConn() *xorm.Engine {
	once.Do(func() {
		var err error
		conn, err = xorm.NewEngine(driverName, getUri())
		util.OkOrPanic(err)
		conn.SetTableMapper(core.SnakeMapper{})
		conn.SetColumnMapper(core.SnakeMapper{})
	})
	return conn
}

type MysqlLoader struct{}

func (m MysqlLoader) Load() (dataMap map[string]map[string]interface{}, dataList map[string]interface{}, err error) {
	conn := getConn()
	dataMap = make(map[string]map[string]interface{})
	dataList = make(map[string]interface{})
	for name, typ := range model.AllModels() {
		slice := reflect.New(reflect.SliceOf(typ))
		err := conn.Find(slice.Interface())
		util.OkOrPanic(err)
		var data = make(map[string]interface{})
		slice = slice.Elem()
		for i := 0; i < slice.Len(); i++ {
			row := slice.Index(i)
			key := generateIndexKey(row)
			data[key] = row.Interface()
		}
		dataMap[name] = data
		dataList[name] = slice.Interface()
	}
	return
}

func getUri() string {
	s := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		"root",
		"123456",
		"localhost",
		3306,
		"mytest",
		"utf8",
	)
	return s
}
