package query

import (
	"context"
	"errors"

	"gitlab.dianchu.cc/go_chaos/fort/factory"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
)

// ColumnValue 记录column及其对应的值，该interface{}的type源自数据库字段定义
// 可能的类型详见 github.com/go-sql-driver/mysql/fields.go 中的scanType<t>
type ColumnValue map[string]interface{}

type SQLQueryDriver interface {
	// 根据主键获取model的单条记录
	Get(ctx context.Context, model interface{}, arg ...interface{}) (bool, error)
	// 根据where条件获取多条记录
	Find(ctx context.Context, modelSet interface{}, where ...interface{}) error
	// 获取表内所有记录
	All(ctx context.Context, modelSet interface{}) error
	Where(where interface{}, args ...interface{}) SQLQueryDriver
	Or(where interface{}, arg ...interface{}) SQLQueryDriver
	Order(where string, asc ...bool) SQLQueryDriver
	Limit(arg int) SQLQueryDriver
	Offset(arg int) SQLQueryDriver
	Count(ctx context.Context, model interface{}, n *int) error
	// 执行query语句，返回多条查询结果，ColumnValue中的Value类型如下，与SQL映射关系见github.com/go-sql-driver/mysql/driver_go18_test.go
	// float32
	// float64
	// int8
	// int16
	// int32
	// int64
	// time.Time NULL时值为0001-01-01 00:00:00 +0000 UTC]
	// uint8
	// uint16
	// uint32
	// uint64
	// string
	// interface{} 类型未知
	Query(ctx context.Context, query string, args ...interface{}) ([]ColumnValue, error)
	First(ctx context.Context, model interface{}, where ...interface{}) (bool, error)
	// 非查询方法
	Cache() SQLQueryDriver // 开启缓存
	canCache() bool
}

type MockSQLQuery struct {
	MockResult []interface{}
	QuerySQL   *factory.SQLInfo
}

func (q *MockSQLQuery) putQuery(sql string, args ...interface{}) {
	q.QuerySQL = new(factory.SQLInfo)
	q.QuerySQL.Sql = sql
	q.QuerySQL.Args = args
}

func (q *MockSQLQuery) Get(ctx context.Context, model interface{}) (bool, error) {
	if len(q.MockResult) < 1 {
		errStr := "No mock result! "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return false, errors.New(errStr)
	}
	return false, nil
}

func (q *MockSQLQuery) Find(ctx context.Context, modelSet interface{}, where string, args ...interface{}) error {
	return nil
}

func (q *MockSQLQuery) Query(ctx context.Context, query string, args ...interface{}) ([]ColumnValue, error) {
	return nil, nil
}

//func NewSQLQuery(sqlEngine interface{}, schema *factory.SQLFactory) (SQLQueryDriver, error) {
//	engineValue := reflect.ValueOf(sqlEngine)
//	if engineValue.Kind() != reflect.Ptr {
//		errStr := engineValue.Kind().String() + " is not ptr type!"
//		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
//		return nil, errors.New(errStr)
//	}
//	switch t := engineValue.Type(); t {
//	case reflect.TypeOf(&engine.SQLEngine{}):
//		q := NewDirectSQLQuery(sqlEngine.(*engine.SQLEngine), schema)
//		return q, nil
//	case reflect.TypeOf(&engine.DALEngine{}):
//		q := NewDALSQLQuery(sqlEngine.(*engine.DALEngine), schema)
//		return q, nil
//	default:
//		errStr := t.String() + " is not engine !"
//		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
//		return nil, errors.New(errStr)
//	}
//}
