/*
@Time : 2018/11/26 9:27
@Author : linfeng
@File : querywapper
@Desc:
*/

package infra

import (
	"arthur/app/base"
	"arthur/utils/struitls"
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type QueryWapper interface {
	Params(params interface{}) QueryWapper
	OrderField(field string) QueryWapper
	OrderBy(order orderby) QueryWapper
	SetPage(pageNo, pageSize int) QueryWapper
	Limit(start, offset int) QueryWapper
	QueryPage(ctx Context, server base.AppServer, model interface{}, total *int, data interface{}) error
	Find(ctx Context, server base.AppServer, data interface{}) error
}

func NewQueryWapper() QueryWapper {
	return &queryWapper{order: OrderNone, start: -1, offset: -1}
}

type orderby int8

const (
	OrderNone orderby = iota
	OrderAsc
	OrderDesc
)

//传0返回ASC，传其他返回DESC
func GetOrderBy(order int16) orderby {
	switch order {
	case 0:
		return OrderAsc
	default:
		return OrderDesc

	}
}

const (
	eq       = "eq"
	eq_      = "="
	noteq    = "noteq"
	noteq_   = "!="
	like     = "like"
	gt       = "gt"
	gt_      = ">"
	lt       = "lt"
	lt_      = "<"
	ge       = "ge"
	ge_      = ">="
	le       = "le"
	le_      = "<="
	in       = "in"
	notin    = "notin"
	notin_   = "not in"
	notlike  = "notlike"
	notlike_ = "not like"
)

type queryWapper struct {
	params        interface{}
	order         orderby
	orderField    string
	basicSQL      string
	basicArgs     []interface{}
	initReflect   bool
	typeof        reflect.Type
	valueof       reflect.Value
	start, offset int
}

func (w *queryWapper) Find(ctx Context, server base.AppServer, data interface{}) error {
	sql, args := w.GetSQL()
	return Query(server).Where(sql, args...).Find(ctx, data)
}

func (w *queryWapper) QueryPage(ctx Context, server base.AppServer, model interface{}, total *int, data interface{}) error {
	basicSQL, basicArgs := w.BasicSQL()
	err := Query(server).Where(basicSQL, basicArgs...).Count(ctx, model, total)
	if err != nil {
		return err
	}
	if *total == 0 {
		return nil
	}
	sql, args := w.GetSQL()
	err = Query(server).Where(sql, args...).Find(ctx, data)
	return err
}

func (w *queryWapper) Params(params interface{}) QueryWapper {
	w.params = params
	w.basicSQL = ""
	w.basicArgs = nil
	w.initReflect = false
	return w
}

func (w *queryWapper) OrderField(field string) QueryWapper {
	w.orderField = field
	return w
}

func (w *queryWapper) OrderBy(order orderby) QueryWapper {
	w.order = order
	return w
}

func (w *queryWapper) SetPage(pageNo, pageSize int) QueryWapper {
	if pageNo <= 0 {
		panic("ErrorPageNo")
	}
	if pageSize <= 0 {
		panic("ErrorPageSize")
	}
	start := (pageNo - 1) * pageSize
	return w.Limit(start, pageSize)
}

func (w *queryWapper) Limit(start, offset int) QueryWapper {
	if start < 0 {
		panic("ErrorStart")
	}
	if offset <= 0 {
		panic("ErrorOffset")
	}
	w.start, w.offset = start, offset
	return w
}

func (w *queryWapper) Parse() (sql string, args []interface{}) {
	return
}

func (w *queryWapper) BasicSQL() (sql string, args []interface{}) {
	if w.params == nil {
		panic("ParamsIsNil")
	}
	if w.basicSQL != "" {
		sql = w.basicSQL
		args = w.basicArgs
		return
	}
	var whereStr bytes.Buffer
	whereStr.WriteString(" 1=1 ")
	args = make([]interface{}, 0)
	for i := 0; i < w.ParamsNumField(); i++ {
		value, valid := w.GetValue(i)
		if !valid {
			continue
		}
		if whereStr.String() != "" {
			whereStr.WriteString(" and ")
		}
		fieldName, queryType := w.MetaField(i)
		str := fmt.Sprintf(" %s %s ? ", fieldName, queryType)
		whereStr.WriteString(str)
		args = append(args, value)
	}
	//将SQL中两个空格替换成一个空格
	sql = struitls.DeepReplace(whereStr.String(), "  ", " ")
	w.basicSQL = sql
	w.basicArgs = args
	args = w.basicArgs
	return
}

func (w *queryWapper) ParamsNumField() int {
	w.InitReflect()
	return w.typeof.NumField()
}
func (w *queryWapper) InitReflect() {
	if w.params == nil {
		panic("ParamsIsNil")
	}
	if w.initReflect {
		return
	}
	typeof := reflect.TypeOf(w.params)
	w.typeof = typeof
	valueof := reflect.ValueOf(w.params)
	w.valueof = valueof
	w.initReflect = true
}

func (w *queryWapper) MetaField(index int) (fieldName string, queryType string) {
	w.InitReflect()
	structField := w.typeof.Field(index)
	fieldName = structField.Tag.Get("field")
	if fieldName == "" {
		fieldName = struitls.SnakeString(structField.Name)
	}
	qtype := structField.Tag.Get("type")
	queryType = GetQueryType(qtype)
	return
}

func GetQueryType(qtype string) string {
	qtype = struitls.RemoveBlank(qtype)
	switch strings.ToLower(qtype) {
	case eq:
		fallthrough
	case eq_:
		fallthrough
	default:
		return eq_
	case noteq_:
		fallthrough
	case noteq:
		return noteq_
	case like:
		return like
	case gt_:
		fallthrough
	case gt:
		return gt_
	case lt_:
		fallthrough
	case lt:
		return lt_
	case ge_:
		fallthrough
	case ge:
		return ge_
	case le_:
		fallthrough
	case le:
		return le_
	case in:
		return in
	case notin:
		return notin_
	case notlike:
		return notlike_
	}
	panic("ErrorQueryType")
}

func (w *queryWapper) GetSQL() (sql string, args []interface{}) {
	sql, args = w.BasicSQL()
	var whereStr bytes.Buffer
	whereStr.WriteString(sql)
	field, order := w.GetOrderField()
	var orderstr string
	if field != "" && (order == OrderAsc || order == OrderNone) {
		orderstr = fmt.Sprintf(" order by %s %s ", field, " asc ")
	} else if field != "" && order == OrderDesc {
		orderstr = fmt.Sprintf(" order by %s %s ", field, " desc ")
	}
	whereStr.WriteString(orderstr)
	sql = whereStr.String()
	if w.start < 0 || w.offset <= 0 {
		return
	}
	whereStr.WriteString(" limit ?,?")
	sql = struitls.DeepReplace(whereStr.String(), "  ", " ")
	args = append(args, w.start)
	args = append(args, w.offset)
	return
}

func (w *queryWapper) GetOrderField() (field string, order orderby) {
	if w.orderField == "" {
		return "", OrderNone
	}
	w.InitReflect()
	if w.order == OrderNone { //如果未设置排序顺序，则默认使用升序
		w.order = OrderAsc
	}
	order = w.order
	fieldNew, ok := w.typeof.FieldByName(w.orderField)
	if !ok {
		return w.orderField, order
	}
	field = fieldNew.Tag.Get("field")
	if field != "" {
		return
	}
	field = struitls.SnakeString(fieldNew.Name)
	return
}

func (w *queryWapper) GetValue(index int) (value interface{}, valid bool) {
	w.InitReflect()
	valueField := w.valueof.Field(index)
	kind := valueField.Kind()
	structField := w.typeof.Field(index)
	validTag := structField.Tag.Get("valid")
	validTag = struitls.RemoveBlank(validTag)
	if validTag == "require" {
		valid = true
	} else if validTag == "" || validTag == "option" {
		valid = false
	} else {
		panic("ErrorValid")
	}
	qtype := structField.Tag.Get("type")
	qtype = strings.ToLower(qtype)
	switch kind {
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		valueNew := valueField.Int()
		if valueNew != 0 {
			valid = true
		}
		value = valueNew
		return
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		valueNew := valueField.Float()
		if valueNew != 0 {
			valid = true
		}
		value = valueNew
		return
	case reflect.String:
		valueNew := valueField.String()
		if valueNew != "" {
			valid = true
		}
		if qtype == like || qtype == notlike {
			value = fmt.Sprintf("%%%s%%", valueNew)
		} else {
			value = valueNew
		}
		return
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return valueField.Interface(), !valueField.IsNil()

	}
	return nil, false
}
