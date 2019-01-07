package query

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"gitlab.dianchu.cc/go_chaos/fort/engine"
	"gitlab.dianchu.cc/go_chaos/fort/factory"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
	"gitlab.dianchu.cc/go_chaos/fort/utils"
)

type DALSQLQuery struct {
	MySQLQuery
	db engine.IDALEngine
}

func NewDALSQLQuery(dalEngine engine.IDALEngine, schema *factory.SQLFactory) *DALSQLQuery {
	var (
		q *DALSQLQuery
	)
	q = new(DALSQLQuery)
	q.db = dalEngine
	q.schema = schema
	return q
}

func (q *DALSQLQuery) Cache() SQLQueryDriver {
	q.openCache = true
	return q
}

func (q *DALSQLQuery) canCache() bool {
	if q.db.MySQLCacher() == nil {
		return false
	}
	return q.openCache && q.db.MySQLCacher().CacherCheck()
}

func (q *DALSQLQuery) Get(ctx context.Context, model interface{}, arg ...interface{}) (bool, error) {
	var (
		schema  *factory.Schema
		sv      reflect.Value
		err     error
		columns []string
	)
	defer q.reset()

	schema, sv, err = q.schema.ParseValue(model)
	if err != nil {
		return false, err
	}
	columns = q.schemaColumns(schema)
	err = q.genGetStatement(schema, &sv, &columns, arg...)
	if err != nil {
		return false, err
	}

	return q.querySingle(ctx, model, schema)
}

func (q *DALSQLQuery) Find(ctx context.Context, modelSet interface{}, where ...interface{}) error {
	var (
		schema  *factory.Schema
		err     error
		columns []string
	)
	defer q.reset()
	if schema, err = q.schema.ParseModel(modelSet); err != nil {
		return err
	}
	columns = q.schemaColumns(schema)
	if err = q.genFindStatement(schema, &columns, where...); err != nil {
		return err
	}
	_, err = q.query(ctx, modelSet, schema)
	return err
}

func (q *DALSQLQuery) First(ctx context.Context, model interface{}, where ...interface{}) (bool, error) {
	var (
		schema  *factory.Schema
		err     error
		columns []string
	)
	defer q.reset()
	schema, _, err = q.schema.ParseValue(model)
	if err != nil {
		return false, err
	}

	columns = q.schemaColumns(schema)
	err = q.genFirstStatement(schema, &columns, where...)
	if err != nil {
		return false, err
	}
	return q.querySingle(ctx, model, schema)
}

func (q *DALSQLQuery) Query(ctx context.Context, query string, args ...interface{}) ([]ColumnValue, error) {
	defer q.reset()
	q.dml.WriteString(query)
	q.args = args
	var result []ColumnValue

	_, err := q.query(ctx, &result, nil)
	return result, err
}

func (q *DALSQLQuery) All(ctx context.Context, modelSet interface{}) error {
	var (
		schema  *factory.Schema
		err     error
		columns []string
	)
	defer q.reset()
	schema, err = q.schema.ParseModel(modelSet)
	if err != nil {
		return err
	}
	columns = q.schemaColumns(schema)
	if err = q.genAllStatement(schema, &columns); err != nil {
		return err
	}
	_, err = q.query(ctx, modelSet, schema)
	return err
}

func (q *DALSQLQuery) Where(where interface{}, args ...interface{}) SQLQueryDriver {
	q.genWhereCondition(where, args...)
	return q //执行查询动作时会检查where语句是否为空
}

func (q *DALSQLQuery) Or(where interface{}, arg ...interface{}) SQLQueryDriver {
	q.genOrCondition(where, arg...)
	return q
}

func (q *DALSQLQuery) Order(where string, asc ...bool) SQLQueryDriver {
	q.genOrderCondition(where, asc...)
	return q
}

func (q *DALSQLQuery) Limit(arg int) SQLQueryDriver {
	q.genLimitCondition(arg)
	return q
}

func (q *DALSQLQuery) Offset(arg int) SQLQueryDriver {
	q.genOffsetCondition(arg)
	return q
}

func (q *DALSQLQuery) Count(ctx context.Context, model interface{}, n *int) error {
	var (
		err    error
		r      ColumnValue
		schema *factory.Schema
	)
	defer q.reset()

	schema, err = q.getSchema().ParseModel(model)
	if err != nil {
		return err
	}

	err = q.genCountStatement(schema)
	if err != nil {
		return err
	}

	if _, err = q.query(ctx, &r, schema); err != nil {
		return err
	}
	if i, ok := r["COUNT(*)"].(int64); ok {
		*n = int(i)
		return nil
	}
	errStr := "Query error! [COUNT(*)] fields do not exist "
	syslog.FortLog.ShowLog(syslog.ERROR, errStr)
	return errors.New(errStr)
}

func (q *DALSQLQuery) query(ctx context.Context, model interface{}, schema *factory.Schema) (int64, error) {
	var (
		fieldSchema    map[string]reflect.StructField
		res            []map[string]interface{}
		err            error
		isSupportCache bool
		canGenVersion  = false // 是否允许更新缓存版本
		canUpdateCache = true  // 是否允许更新缓存
		resCount       int64
	)
	defer func() {
		q.reset()
	}()
	if err = q.check(); err != nil {
		return 0, err
	}

	// 先查询缓存
	if q.canCache() { //todo 日志输出
		//todo 判断PrimaryKeyValue 是否正确
		exist, rows, err := q.db.MySQLCacher().GetCache(q.db.GetDBName(), schema.Name, []interface{}{schema.PrimaryKeyValue}, model)
		if exist && err == nil {
			return rows, err
		} else if err != nil {
			switch err.Error() {
			case "The version does not exist ": // 有缓存,无缓存版本号的情况下,自动生成缓存版本号.
				canGenVersion = true
			case "The cache is locked ": // 如果缓存版本号被锁定,则不对缓存进行更新
				canUpdateCache = false
			}
		}
	}

	// 数据访问服务查询
	cmd := engine.CmdData{
		Statement: q.dml.String(),
		Args:      q.args,
	}
	syslog.FortLog.ShowLog(syslog.DEBUG, cmd)
	if resCount, res, err = q.db.DALQuery(ctx, cmd); err != nil {
		return resCount, err
	}

	if resCount == 0 {
		return resCount, nil
	}

	if schema != nil {
		fieldSchema = schema.FieldSchema
	}

	// 处理返回的查询结果
	isSupportCache, err = resultHandle(res, model, fieldSchema)

	// 更新缓存:返回结果处理无误 && 支持缓存 && 返回结果支持存入缓存 && 允许更新缓存
	if err == nil && q.canCache() && isSupportCache && canUpdateCache {
		if data, err := utils.Marshal(model); err == nil {
			cacheErr := q.db.MySQLCacher().UpdateCache(q.db.GetDBName(), schema.Name, []interface{}{schema.PrimaryKeyValue}, data)
			if cacheErr == nil && canGenVersion {
				q.db.MySQLCacher().UpdateCacheVersion(q.db.GetDBName(), schema.Name, schema.PrimaryKeyValue, false)
			}
		}
	}

	return resCount, err
}

// 检查单条结果是否为空 或 查询是否错误
func (q *DALSQLQuery) querySingle(ctx context.Context, model interface{}, schema *factory.Schema) (bool, error) {
	// 数据库查询
	if resCount, err := q.query(ctx, model, schema); err != nil || resCount == 0 {
		return false, err
	}
	return true, nil
}

//return 查询结果是否支持缓存,错误
func resultHandle(res []map[string]interface{}, model interface{}, fieldSchema map[string]reflect.StructField) (bool, error) {

	isSupportCache := false

	modelValue := reflect.ValueOf(model)
	modelKind := modelValue.Elem().Kind()

	if modelKind == reflect.Map {
		if m, ok := model.(*ColumnValue); ok { // Count查询
			*m = res[0]
			return false, nil
		}
		errStr := "Map is not matching "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return false, errors.New(errStr)
	}

	if modelKind != reflect.Struct && modelKind != reflect.Slice {
		errStr := "No matching type "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return false, errors.New(errStr)
	}

	if modelValue.Kind() != reflect.Ptr {
		errStr := "Needs a pointer to a value "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return false, errors.New(errStr)
	}

	if modelValue.Elem().Kind() == reflect.Slice {
		slice := reflect.MakeSlice(reflect.TypeOf(model).Elem(), len(res), len(res))
		if slice.Index(0).Type() == reflect.TypeOf(ColumnValue{}) { //自定义语句查询:Query
			if m, ok := model.(*[]ColumnValue); ok { // Count
				var temp []ColumnValue
				for _, v := range res {
					temp = append(temp, v)
				}
				*m = temp
				return false, nil
			}
			errStr := "Slice is not matching "
			syslog.FortLog.ShowLog(syslog.ERROR, errStr)
			return false, errors.New(errStr)
		}

		//多结果查询:Find
		isSupportCache = true
		for i, r := range res {
			mValue := reflect.New(slice.Index(i).Type())
			if err := map2struct(r, mValue, fieldSchema); err != nil {
				return false, err
			}
			slice.Index(i).Set(mValue.Elem())
		}
		modelValue.Elem().Set(reflect.AppendSlice(modelValue.Elem(), slice))
	} else { //单结果查询:Get
		isSupportCache = true
		if len(res) > 0 {
			if err := map2struct(res[0], modelValue, fieldSchema); err != nil {
				return false, err
			}
		}
	}
	return isSupportCache, nil
}

func map2struct(resData map[string]interface{}, sv reflect.Value, fieldSchema map[string]reflect.StructField) error {
	t := sv.Elem()
	for resKey, resVal := range resData {
		fName := fieldSchema[resKey].Name
		modelRV := t.FieldByName(fName)  // struct字段Value
		resRV := reflect.ValueOf(resVal) // 返回结果Value
		if !modelRV.CanSet() {
			continue
		}
		switch modelRV.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			s := asString(resVal)
			i64, err := strconv.ParseInt(s, 10, resRV.Type().Bits())
			if err != nil {
				return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", resRV, s, resRV.Kind(), err)
			}
			modelRV.SetInt(i64)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			s := asString(resVal)
			u64, err := strconv.ParseUint(s, 10, resRV.Type().Bits())
			if err != nil {
				return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", resVal, s, resRV.Kind(), err)
			}
			modelRV.SetUint(u64)
		case reflect.Float32, reflect.Float64:
			s := asString(resVal)
			f64, err := strconv.ParseFloat(s, resRV.Type().Bits())
			if err != nil {
				return fmt.Errorf("converting driver.Value type %T (%q) to a %s: %v", resVal, s, resRV.Kind(), err)
			}
			modelRV.SetFloat(f64)
		case reflect.String:
			switch v := resVal.(type) {
			case string:
				modelRV.SetString(v)
			case []byte:
				modelRV.SetString(string(v))
			}
		case reflect.Bool:
			s := asString(resVal)
			if s == "0" {
				modelRV.SetBool(false)
			} else {
				modelRV.SetBool(true)
			}
		case reflect.Struct: //time.Time{}
			var (
				tm  time.Time
				err error
			)
			if len(resVal.(string)) > 10 {
				if tm, err = time.Parse("2006-01-02 15:04:05", resVal.(string)); err != nil {
					return err
				}
			} else {
				if tm, err = time.Parse("2006-01-02", resVal.(string)); err != nil {
					return err
				}
			}
			modelRV.Set(reflect.ValueOf(tm))
		default:
			return fmt.Errorf("no matching type")
		}
	}
	return nil
}

func asString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv := reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
}
