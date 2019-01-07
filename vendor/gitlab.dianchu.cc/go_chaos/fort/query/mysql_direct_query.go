package query

import (
	"context"
	"database/sql"
	"errors"
	"reflect"

	"gitlab.dianchu.cc/go_chaos/fort/factory"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
	"gitlab.dianchu.cc/go_chaos/fort/utils"

	"github.com/go-sql-driver/mysql"
	"gitlab.dianchu.cc/go_chaos/fort/engine"
)

type DirectSQLQuery struct {
	MySQLQuery
	db interface{}
}

func NewDirectSQLQuery(sqlEngine engine.ISQLEngine, schema *factory.SQLFactory) *DirectSQLQuery {
	var (
		q *DirectSQLQuery
	)
	q = new(DirectSQLQuery)
	q.db = sqlEngine.DB()
	q.schema = schema
	return q
}

func (q *DirectSQLQuery) query(ctx context.Context) (*sql.Rows, error) {
	err := q.check()
	if err != nil {
		return nil, err
	}
	syslog.FortLog.ShowLog(syslog.DEBUG, q.dml.String(), q.args)
	switch q.db.(type) {
	case *sql.DB:
		return q.db.(*sql.DB).QueryContext(ctx, q.dml.String(), q.args...)
	case *sql.Tx:
		return q.db.(*sql.Tx).QueryContext(ctx, q.dml.String(), q.args...)
	default:
		errStr := "Not sql.DB or sql.Tx type. "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, errors.New(errStr)
	}
}

func (q *DirectSQLQuery) queryRow(ctx context.Context) (*sql.Row, error) {
	err := q.check()
	if err != nil {
		return nil, err
	}
	syslog.FortLog.ShowLog(syslog.DEBUG, q.dml.String(), q.args)
	switch q.db.(type) {
	case *sql.DB:
		return q.db.(*sql.DB).QueryRowContext(ctx, q.dml.String(), q.args...), nil
	case *sql.Tx:
		return q.db.(*sql.Tx).QueryRowContext(ctx, q.dml.String(), q.args...), nil
	default:
		errStr := "No db object! "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, errors.New(errStr)
	}
}

func (q *DirectSQLQuery) queryRowScan(ctx context.Context, model interface{}) (bool, error) {
	var (
		err   error
		count int64
	)
	if count, err = q.directQuery(ctx, model); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return count > 0, nil
}

func (q *DirectSQLQuery) directQuery(ctx context.Context, model interface{}) (int64, error) {
	var (
		modelValue  reflect.Value
		modelKind   reflect.Kind
		err         error
		schema      *factory.Schema
		columns     []string
		fieldNumber int
		rows        *sql.Rows
		row         *sql.Row
	)
	defer func() {
		q.reset()
		if rows != nil {
			rows.Close()
		}
	}()

	if err = q.check(); err != nil {
		return 0, err
	}

	modelValue = reflect.ValueOf(model)
	modelKind = modelValue.Elem().Kind()
	if schema, err = q.schema.ParseModel(model); err != nil {
		return 0, err
	}
	columns = q.schemaColumns(schema)
	fieldNumber = len(columns)

	if modelKind != reflect.Struct && modelKind != reflect.Slice {
		errStr := "No matching type "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
	}
	if modelValue.Kind() != reflect.Ptr {
		errStr := "Needs a pointer to a value "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return 0, errors.New(errStr)
	}
	// 多结果  reflect.Slice
	if modelKind == reflect.Slice {
		if rows, err = q.query(ctx); err != nil {
			return 0, err
		}
		results := reflect.ValueOf(model).Elem()
		bucket := reflect.MakeSlice(results.Type(), utils.ROW_CAP, utils.ROW_CAP)
		i := 0
		for rows.Next() {
			elem := reflect.New(schema.ModelType).Elem()
			err = q.scan(rows, fieldNumber, &columns, schema, elem)
			if err != nil {
				return 0, err
			}
			bucket.Index(i).Set(elem)
			i++
			if i == utils.ROW_CAP {
				results.Set(reflect.AppendSlice(results, bucket))
				i = 0
			}
		}
		if i > 0 {
			results.Set(reflect.AppendSlice(results, bucket.Slice(0, i)))
		}
		return int64(i), err
	}
	// 单结果 reflect.Struct
	if row, err = q.queryRow(ctx); err != nil {
		return 0, err
	}
	if err = q.scan(row, fieldNumber, &columns, schema, modelValue.Elem()); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return 1, nil
}

func (q *DirectSQLQuery) scan(cursor interface{}, fieldNumber int, columns *[]string, schema *factory.Schema, result reflect.Value) error {
	var (
		values []interface{}
		i      int
		column string
		elem   reflect.Value
	)
	values = make([]interface{}, fieldNumber)
	for i, column = range *columns {
		if result.FieldByName(column).Kind() == reflect.Ptr {
			values[i] = result.FieldByName(column).Addr().Interface()
		} else {
			elem = reflect.New(reflect.PtrTo(schema.FieldSchema[column].Type)).Elem()
			elem.Set(result.FieldByName(schema.FieldSchema[column].Name).Addr())
			values[i] = elem.Interface()
		}
	}
	return scanValues(cursor, values)
}

func (q *DirectSQLQuery) getColumnValue(cursor interface{}, fieldNumber int, columns *[]string, columnTypes []*sql.ColumnType) (ColumnValue, error) {
	var (
		values []interface{}
		column string
		i      int
		err    error
		result ColumnValue
	)

	values = make([]interface{}, fieldNumber)
	result = make(ColumnValue, fieldNumber)
	for i, column = range *columns {
		typ := columnTypes[i].ScanType()
		elem := reflect.New(reflect.PtrTo(typ).Elem())
		values[i] = elem.Interface()
	}

	scanValues(cursor, values)

	//converted value map
	for i, column = range *columns {
		var (
			v   reflect.Value //actual value
			vi  interface{}   //actual value interface
			cvi interface{}   //converted value interface
		)
		v = reflect.ValueOf(values[i]).Elem()
		vi = v.Interface()

		switch r := v.Interface().(type) {
		case mysql.NullTime:
			cvi = r.Time
		case sql.NullFloat64:
			cvi = r.Float64
		case sql.NullInt64:
			cvi = r.Int64
		case sql.RawBytes:
			cvi = string(r)
		default:
			cvi = vi
		}
		result[column] = cvi
	}
	return result, err
}

func (q *DirectSQLQuery) Get(ctx context.Context, model interface{}, arg ...interface{}) (bool, error) {
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
	return q.queryRowScan(ctx, model)
}

func (q *DirectSQLQuery) First(ctx context.Context, model interface{}, where ...interface{}) (bool, error) {
	var (
		schema  *factory.Schema
		err     error
		columns []string
	)
	defer q.reset()
	schema, err = q.schema.ParseModel(model)
	if err != nil {
		return false, err
	}

	columns = q.schemaColumns(schema)

	err = q.genFirstStatement(schema, &columns, where...)
	if err != nil {
		return false, err
	}

	return q.queryRowScan(ctx, model)
}

func (q *DirectSQLQuery) Find(ctx context.Context, modelSet interface{}, where ...interface{}) error {
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
	_, err = q.directQuery(ctx, modelSet)
	return err
}

func (q *DirectSQLQuery) All(ctx context.Context, modelSet interface{}) error {
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
	q.genAllStatement(schema,&columns)
	_, err = q.directQuery(ctx, modelSet)
	return err
}

func (q *DirectSQLQuery) Where(where interface{}, args ...interface{}) SQLQueryDriver {
	q.genWhereCondition(where, args...)
	return q //执行查询动作时会检查where语句是否为空
}

func (q *DirectSQLQuery) Or(where interface{}, arg ...interface{}) SQLQueryDriver {
	q.genOrCondition(where, arg...)
	return q
}

func (q *DirectSQLQuery) Order(where string, asc ...bool) SQLQueryDriver {
	q.genOrderCondition(where, asc...)
	return q
}

func (q *DirectSQLQuery) Limit(arg int) SQLQueryDriver {
	q.genLimitCondition(arg)
	return q
}

func (q *DirectSQLQuery) Offset(arg int) SQLQueryDriver {
	q.genLimitCondition(arg)
	return q
}

func (q *DirectSQLQuery) Count(ctx context.Context, model interface{}, n *int) error {
	var (
		err    error
		row    *sql.Row
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

	row, err = q.queryRow(ctx)
	if err != nil {
		return err
	}

	err = row.Scan(n)
	if err != nil {
		return err
	}
	return nil
}

func (q *DirectSQLQuery) Query(ctx context.Context, query string, args ...interface{}) ([]ColumnValue, error) {
	var (
		err         error
		columns     []string
		columnTypes []*sql.ColumnType
		fieldNumber int
		bucket      []ColumnValue
		rows        *sql.Rows
		i           int
		results     []ColumnValue
	)
	defer q.reset()
	q.dml.WriteString(query)
	q.args = args
	rows, err = q.query(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns = make([]string, fieldNumber)
	columns, err = rows.Columns()
	if err != nil {
		return nil, err
	}
	columnTypes = make([]*sql.ColumnType, fieldNumber)
	columnTypes, err = rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	fieldNumber = len(columns)
	bucket = make([]ColumnValue, utils.ROW_CAP, utils.ROW_CAP)
	i = 0
	for rows.Next() {
		columnValue, err := q.getColumnValue(rows, fieldNumber, &columns, columnTypes)
		if err != nil {
			return nil, err
		}
		bucket[i] = columnValue
		i++
		if i == utils.ROW_CAP {
			results = append(results, bucket...)
			i = 0
		}
	}
	if i > 0 {
		results = append(results, bucket[0:i]...)
	}

	return results, nil
}

func (q *DirectSQLQuery) Cache() SQLQueryDriver {
	q.openCache = true
	return q
}

func (q *DirectSQLQuery) canCache() bool {
	return q.openCache
}

func scanValues(cursor interface{}, values []interface{}) error {
	if cursor == nil {
		errStr := "Cursor is nil. "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return errors.New(errStr)
	}
	switch cursor.(type) {
	case *sql.Row:
		return cursor.(*sql.Row).Scan(values...)
	case *sql.Rows:
		return cursor.(*sql.Rows).Scan(values...)
	default:
		errStr := "Not *sql.Row or *sql.Rows type. "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return errors.New(errStr)
	}
}
