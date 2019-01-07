package query

import (
	"bytes"
	"errors"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"gitlab.dianchu.cc/go_chaos/fort/factory"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
	"gitlab.dianchu.cc/go_chaos/fort/utils"
)

type MySQLQueryDriver interface {
	getSchema() *factory.SQLFactory
	getCondition(ConditionType int) *bytes.Buffer
	updateCondition(ConditionType int, Condition string) (int, error)
	setArgs(cover bool, args ...interface{})
	updateDML(dml string) (int, error)
	check() error
}

const (
	//getCondition updateCondition
	where    = iota
	order
	limit
	offset
	complete  //拼接后的完整条件(getCondition)
)

type MySQLQuery struct {
	schema    *factory.SQLFactory
	where     bytes.Buffer
	order     bytes.Buffer
	limit     bytes.Buffer
	offset    bytes.Buffer
	dml       bytes.Buffer
	openCache bool
	args      []interface{}
}

func (m *MySQLQuery) reset() {
	m.where.Reset()
	m.dml.Reset()
	m.order.Reset()
	m.limit.Reset()
	m.offset.Reset()
	m.args = make([]interface{}, 0)
	m.openCache = false
}

func (m *MySQLQuery) check()error{
	var err error
	if m.dml.Len() < 1 {
		errStr := "No query strng! "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		err = errors.New(errStr)
	}
	return err
}

func (m *MySQLQuery) getSchema() *factory.SQLFactory {
	return m.schema
}

func (m *MySQLQuery) getCondition(conditionType int) *bytes.Buffer {
	switch conditionType {
	case where:
		return &m.where
	case order:
		return &m.order
	case limit:
		return &m.limit
	case offset:
		return &m.offset
	case complete:
		var (
			condition bytes.Buffer
		)
		whereCondition := m.where.String()
		if whereCondition == "" {
			break
		}
		for _, v := range []string{whereCondition, m.order.String(), m.limit.String(), m.offset.String()} {
			if v != "" {
				condition.WriteString(v)
				condition.WriteString(" ")
			}
		}
		return &condition
	}
	return &bytes.Buffer{}
}

func (m *MySQLQuery) updateCondition(conditionType int, condition string) (int, error) {
	switch conditionType {
	case where:
		return m.where.WriteString(condition)
	case order:
		return m.order.WriteString(condition)
	case limit:
		m.limit.Reset()
		return m.limit.WriteString(condition)
	case offset:
		m.offset.Reset()
		return m.offset.WriteString(condition)
	}
	return -1, errors.New("Type Error ")
}

func (m *MySQLQuery) setArgs(cover bool, args ...interface{}) {
	if cover { // 替换
		m.args = args
	} else {
		if len(args) == 1 && args[0] == nil {
			return
		}
		m.args = append(m.args, args...)
	}
}

func (m *MySQLQuery) updateDML(where string) (int, error) {
	return m.dml.WriteString(where)
}

//根据schema获得columns列表
func (m *MySQLQuery) schemaColumns(schema *factory.Schema) []string {
	var (
		fieldNumber int
		i           int
		columns     []string
	)
	fieldNumber = len(schema.FieldSchema)
	columns = make([]string, fieldNumber)

	i = 0
	for field := range schema.FieldSchema {
		columns[i] = field
		i++
	}
	sort.Strings(columns) // 确保产生顺序一致，防止Row Scan时出错
	return columns
}

//根据column列表生成select子语句
func (m *MySQLQuery) selectColumns(columns []string) string {
	var columnsStr bytes.Buffer

	if len(columns) == 0 {
		return "*"
	}

	for i, v := range columns {
		columnsStr.WriteString("`")
		columnsStr.WriteString(v)
		columnsStr.WriteString("`")
		if i+1 < len(columns) {
			columnsStr.WriteString(",")
		}
	}
	return columnsStr.String()

}

//模型自定义解析为Where语句
type handleFunc func(q MySQLQueryDriver, schema *factory.Schema, tableField string, fieldVal interface{})

func (m *MySQLQuery) modleToWhere(modle interface{}, handle handleFunc) error {
	schema, sv, err := m.getSchema().ParseValue(modle)
	if err != nil {
		return err
	}
	for tableField, field := range schema.FieldSchema {
		fieldVal := sv.FieldByName(field.Name)
		if utils.IsZero(fieldVal) {
			continue
		}

		val, err := utils.AsKind(fieldVal)
		if err != nil {
			return err
		}
		handle(m, schema, tableField, val)
	}
	return nil
}


func (m *MySQLQuery) genGetStatement(schema *factory.Schema, sv *reflect.Value, columns *[]string, arg ...interface{}) error {
	m.updateDML("SELECT ")
	m.updateDML(m.selectColumns(*columns))
	m.updateDML(" FROM ")
	m.updateDML(schema.Name)
	m.updateDML(" WHERE ")
	if len(arg) == 1 {
		m.updateDML("`")
		m.updateDML(schema.PrimaryKey)
		m.updateDML("` =?")
		m.setArgs(true, arg...)
		schema.PrimaryKeyValue=arg[0]
	} else {
		pkVal := sv.FieldByName(utils.ToCamelName(schema.PrimaryKey))
		if utils.IsZero(pkVal) {
			errStr := "PrimaryKey Value is nil. "
			syslog.FortLog.ShowLog(syslog.ERROR, errStr)
			return errors.New(errStr)
		}
		m.updateDML(" `")
		m.updateDML(utils.ToUnderscoreName(schema.PrimaryKey))
		m.updateDML("` = ?")
		val, err := utils.AsKind(pkVal)
		if err != nil {
			return err
		}
		schema.PrimaryKeyValue = val
		m.setArgs(false, val)
	}
	return nil
}

func (m *MySQLQuery) genFindStatement(schema *factory.Schema, columns *[]string,whereCondition ...interface{}) error {
	var (
		err       error
		condition string
		ok        bool
	)
	if len(whereCondition) > 1 {
		condition, ok = whereCondition[0].(string)
		args := whereCondition[1:]
		condition, args = inlineConditionParse(condition, args)
		if !ok {
			errStr := "Query condition is not string type! "
			syslog.FortLog.ShowLog(syslog.ERROR, errStr)
			return errors.New(errStr)
		}
		m.setArgs(true, args...)
	} else {
		condition = m.getCondition(complete).String()
		if len(condition) < 1 {
			errStr := "No where clause! "
			syslog.FortLog.ShowLog(syslog.ERROR, errStr)
			return  errors.New(errStr)
		}
	}
	if err != nil {
		return  err
	}
	m.dml.Reset()
	m.updateDML("SELECT ")
	m.updateDML(m.selectColumns(*columns))
	m.updateDML(" FROM `")
	m.updateDML(schema.Name)
	m.updateDML("` ")
	if condition != "" {
		m.updateDML(" Where ")
		m.updateDML(condition)
	}
	return  nil
}

func (m *MySQLQuery) genAllStatement(schema *factory.Schema, columns *[]string) error{
	m.dml.Reset()
	m.updateDML("SELECT ")
	m.updateDML(m.selectColumns(*columns))
	m.updateDML(" FROM `")
	m.updateDML(schema.Name)
	m.updateDML("` ")
	if order := m.order.String();order != ""{
		m.updateDML(order)
	}
	return nil
}



func (m *MySQLQuery) genFirstStatement(schema *factory.Schema, columns *[]string, whereCondition ...interface{}) error {
	m.updateDML("SELECT ")
	m.updateDML(m.selectColumns(*columns))
	m.updateDML(" FROM ")
	m.updateDML(schema.Name)
	switch len(whereCondition) {
	case 0:
		sqlWhere := m.getCondition(where)
		condition := sqlWhere.String()
		if len(condition) > 0 {
			m.updateDML(" WHERE ")
			m.updateDML(condition)
		} else { // 获取第一条记录，按主键排序
			m.updateDML(" ORDER BY `")
			m.updateDML(schema.PrimaryKey)
			m.updateDML("`")
			m.setArgs(true, make([]interface{}, 0)...)
		}
		m.updateDML(" LIMIT 1")
	case 1: // 使用主键获取记录
		m.updateDML(" WHERE `")
		m.updateDML(schema.PrimaryKey)
		m.updateDML("` =? LIMIT 1")
		m.setArgs(true, whereCondition...)
	default:
		m.updateDML(" WHERE ")
		condition, ok := whereCondition[0].(string)
		args := whereCondition[1:]
		if !ok {
			errStr := "Query condition is not string type! "
			syslog.FortLog.ShowLog(syslog.ERROR, errStr)
			return errors.New(errStr)
		}
		condition, args = inlineConditionParse(condition, args)
		m.updateDML(condition)
		m.updateDML(" LIMIT 1")
		m.setArgs(true, args...)
	}
	return nil
}

func (m *MySQLQuery) genWhereCondition(whereCondition interface{}, args ...interface{}) {
	if _, ok := whereCondition.(string); ok {
		if w := m.getCondition(where); len(w.String()) != 0 {
			m.updateCondition(where, " AND ")
		} else {
			m.updateCondition(where, " ")
		}
		condition := whereCondition.(string)
		condition, args = inlineConditionParse(condition, args)
		m.updateCondition(where, condition)
		m.setArgs(false, args...)
	} else {
		handle := func(q MySQLQueryDriver, schema *factory.Schema, tableField string, fieldVal interface{}) {
			title := " `"
			if w := m.getCondition(where); len(w.String()) != 0 {
				title = " AND `"
			}
			m.updateCondition(where, title)
			m.updateCondition(where, tableField)
			m.updateCondition(where, "` = ?")
			m.setArgs(false, fieldVal)
		}
		m.modleToWhere(whereCondition, handle)
	}
}

func (m *MySQLQuery) genOrCondition(whereCondition interface{}, args ...interface{}) {
	if _, ok := whereCondition.(string); ok {
		condition := whereCondition.(string)
		condition, args = inlineConditionParse(condition, args)
		m.updateCondition(where, " OR ")
		m.updateCondition(where, condition)
		m.setArgs(false, args...)
	} else {
		handle := func(q MySQLQueryDriver, schema *factory.Schema, tableField string, fieldVal interface{}) {
			q.updateCondition(where, " OR `")
			q.updateCondition(where, tableField)
			q.updateCondition(where, "` = ? ")
			q.setArgs(false, fieldVal)
		}
		m.modleToWhere(whereCondition, handle)
	}
}

func (m *MySQLQuery) genOrderCondition(condition string, asc ...bool) {
	var flag string
	if len(asc) == 1 && !asc[0] {
		flag = " DESC"
	}
	if m.getCondition(order).String() == "" {
		m.updateCondition(order, " ORDER BY `")
		m.updateCondition(order, condition)
		m.updateCondition(order, "`")
	} else {
		m.updateCondition(order, ",`")
		m.updateCondition(order, condition)
		m.updateCondition(order, "`")
	}
	m.updateCondition(order, flag)
}

func (m *MySQLQuery) genLimitCondition(arg int) {
	var temp bytes.Buffer
	temp.WriteString(" LIMIT ")
	temp.WriteString(strconv.Itoa(arg))
	m.updateCondition(limit, temp.String())
}

func (m *MySQLQuery) genOffsetCondition(arg int) {
	var temp bytes.Buffer
	temp.WriteString(" OFFSET ")
	temp.WriteString(strconv.Itoa(arg))
	m.updateCondition(offset, temp.String())
}

func (m *MySQLQuery) genCountStatement(schema *factory.Schema) error {
	var (
		condition string
	)
	sqlWhere := m.getCondition(where)
	condition = sqlWhere.String()
	if len(condition) < 1 {
		errStr := "Query condition is not string type! "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return errors.New(errStr)
	}
	m.updateDML("SELECT COUNT(*) FROM ")
	m.updateDML(schema.Name)
	m.updateDML(" WHERE ")
	m.updateDML(condition)
	return nil
}

func inlineConditionParse(condition string, args []interface{}) (string, []interface{}) {
	var ValSubMap = make(map[int]bool)
	for k, v := range args {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			ValSubMap[k] = true
		}
	}
	if len(ValSubMap) != 0 {
		conditionList := strings.Split(condition, "?")
		var newCondition bytes.Buffer
		for k, condition := range conditionList {
			if _, ok := ValSubMap[k]; ok {
				var w bytes.Buffer
				w.WriteString(conditionList[k])
				w.WriteString(" (")
				for i := 0; i < reflect.ValueOf(args[k]).Len(); i++ {
					if i > 0 {
						w.WriteString(", ")
					}
					w.WriteString("?")
				}
				w.WriteString(") ")
				newCondition.WriteString(w.String())
				continue
			}
			newCondition.WriteString(condition)
			if k < len(conditionList)-1 {
				newCondition.WriteString("?")
			}
		}
		var newArgs []interface{}
		for k, v := range args {
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				arg := reflect.ValueOf(args[k])
				for i := 0; i < arg.Len(); i++ {
					newArgs = append(newArgs, arg.Index(i).Interface())
				}
				continue
			}
			newArgs = append(newArgs, v)
		}
		return newCondition.String(), newArgs
	} else {
		return condition, args
	}
}
