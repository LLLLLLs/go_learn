/*
Author : Haoyuan Liu
Time   : 2018/6/19
*/
package infra

import (
	"arthur/app/base"
	goctx "context"
	fortq "gitlab.dianchu.cc/go_chaos/fort/query"
	"reflect"
)

type Queryer interface {
	Get(ctx goctx.Context, model interface{}, arg ...interface{}) (bool, error)
	Find(ctx goctx.Context, modelSet interface{}, where ...interface{}) error
	All(ctx goctx.Context, modelSet interface{}) error
	Where(where interface{}, args ...interface{}) Queryer
	Or(where interface{}, arg ...interface{}) Queryer
	Order(where string, asc ...bool) Queryer
	Limit(arg int) Queryer
	Offset(arg int) Queryer
	Count(ctx goctx.Context, model interface{}, n *int) error
	Query(ctx goctx.Context, query string, args ...interface{}) ([]fortq.ColumnValue, error)
	First(ctx goctx.Context, model interface{}, where ...interface{}) (bool, error)
}

func Query(as base.AppServer) Queryer {
	var q fortq.SQLQueryDriver
	switch getDBConn() {
	case dalConn:
		if dalConnMap.connMap == nil {
			panic("must initDalConn() first ")
		}
		q = fortq.NewDALSQLQuery(dalConnMap.connMap[as.ID()], factory)
	case directConn:
		if sqlConnMap.connMap == nil {
			panic("must initSQLConn() first ")
		}
		q = fortq.NewDirectSQLQuery(sqlConnMap.connMap[as.ID()], factory)
	default:
		panic("wrong connection method")
	}
	if q == nil {
		panic("cannot find appserver")
	}
	return asQuery{q, as}
}

func QueryCenter() Queryer {
	var q fortq.SQLQueryDriver

	switch getDBConn() {
	case dalConn:
		q = fortq.NewDALSQLQuery(dalCenterConn, factory)
	case directConn:
		q = fortq.NewDirectSQLQuery(sqlCenterConn, factory)
	default:
		panic("wrong connection method")
	}
	if q == nil {
		panic("wrong center conn")
	}
	return queryCenter{q}
}

type asQuery struct {
	raw fortq.SQLQueryDriver
	as base.AppServer
}

func (q asQuery) Where(where interface{}, args ...interface{}) Queryer {
	q.raw.Where(where, args...)
	return q
}

func (q asQuery) Or(where interface{}, arg ...interface{}) Queryer {
	q.raw.Or(where, arg...)
	return q
}

func (q asQuery) Order(where string, asc ...bool) Queryer {
	q.raw.Order(where, asc...)
	return q
}

func (q asQuery) Limit(arg int) Queryer {
	q.raw.Limit(arg)
	return q
}

func (q asQuery) Offset(arg int) Queryer {
	q.raw.Offset(arg)
	return q
}

func (q asQuery) Count(ctx goctx.Context, model interface{}, n *int) error {
	return q.raw.Count(ctx, model, n)
}

func (q asQuery) Query(ctx goctx.Context, query string, args ...interface{}) ([]fortq.ColumnValue, error) {
	return q.raw.Query(ctx, query, args...)
}

func (q asQuery) Get(ctx goctx.Context, model interface{}, arg ...interface{}) (bool, error){
	ok, err := q.raw.Get(ctx, model, arg...)
	if err == nil && ok {
		err = setModel(model, q.as)
		if err != nil {
			return false, err
		}
	}
	return ok, err
}

func (q asQuery) First(ctx goctx.Context, model interface{}, where ...interface{}) (bool, error){
	ok, err := q.raw.First(ctx, model, where...)
	if err == nil && ok {
		err = setModel(model, q.as)
		if err != nil {
			return false, err
		}
	}
	return ok, err
}

func (q asQuery) All(ctx goctx.Context, modelSet interface{}) error {
	err := q.raw.All(ctx, modelSet)
	if err == nil {
		err = setModelSet(modelSet, q.as)
		if err != nil {
			return err
		}
	}
	return err
}

func (q asQuery) Find(ctx goctx.Context, modelSet interface{}, where ...interface{}) error {
	err := q.raw.Find(ctx, modelSet, where...)
	if err == nil {
		err = setModelSet(modelSet, q.as)
		if err != nil {
			return err
		}
	}
	return err
}

type queryCenter struct {
	raw fortq.SQLQueryDriver
}
func (q queryCenter) Where(where interface{}, args ...interface{}) Queryer {
	q.raw.Where(where, args...)
	return q
}

func (q queryCenter) Or(where interface{}, arg ...interface{}) Queryer {
	q.raw.Or(where, arg...)
	return q
}

func (q queryCenter) Order(where string, asc ...bool) Queryer {
	q.raw.Order(where, asc...)
	return q
}

func (q queryCenter) Limit(arg int) Queryer {
	q.raw.Limit(arg)
	return q
}

func (q queryCenter) Offset(arg int) Queryer {
	q.raw.Offset(arg)
	return q
}

func (q queryCenter) Count(ctx goctx.Context, model interface{}, n *int) error {
	return q.raw.Count(ctx, model, n)
}

func (q queryCenter) Query(ctx goctx.Context, query string, args ...interface{}) ([]fortq.ColumnValue, error) {
	return q.raw.Query(ctx, query, args...)
}

func (q queryCenter) Get(ctx goctx.Context, model interface{}, arg ...interface{}) (bool, error){
	return q.raw.Get(ctx, model, arg...)
}

func (q queryCenter) First(ctx goctx.Context, model interface{}, where ...interface{}) (bool, error){
	return q.raw.First(ctx, model, where...)
}

func (q queryCenter) All(ctx goctx.Context, modelSet interface{}) error {
	return q.raw.All(ctx, modelSet)
}

func (q queryCenter) Find(ctx goctx.Context, modelSet interface{}, where ...interface{}) error {
	return q.raw.Find(ctx, modelSet, where...)
}

func setModel(model interface{}, as base.AppServer) error{
	schema, rv, err := factory.ParseValue(model)
	if err != nil {
		return err
	}
	id := rv.FieldByName(schema.FieldSchema[schema.PrimaryKey].Name).String()
	rv.FieldByName("Model").Set(reflect.ValueOf(NewModel(id, as)))
	return nil
}

func setModelSet(modelSet interface{}, as base.AppServer) error{
	schema, err := factory.ParseModel(modelSet)
	if err != nil {
		return err
	}
	rv := reflect.ValueOf(modelSet)
	count := rv.Elem().Len()
	pkName := schema.FieldSchema[schema.PrimaryKey].Name
	for i := 0; i < count; i++ {
		id := rv.Elem().Index(i).FieldByName(pkName).String()
		rv.Elem().Index(i).FieldByName("Model").Set(reflect.ValueOf(NewModel(id, as)))
	}
	return nil
}
