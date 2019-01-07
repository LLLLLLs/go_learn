package uow

import (
	"bytes"
	"context"
	"errors"
	"reflect"
	"sort"

	"gitlab.dianchu.cc/chaos_go_sdk/flume_client_sdk_go/flumesdk"
	"gitlab.dianchu.cc/go_chaos/fort/event"
	"gitlab.dianchu.cc/go_chaos/fort/factory"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
	"gitlab.dianchu.cc/go_chaos/fort/utils"
)

type UoWSQL struct {
	transaction *factory.SQLAtom
	schema      *factory.SQLFactory
	eventBus    event.EventAgent
	logAtom     *flumesdk.TransAction
}

func NewUoWSQL(ctx context.Context, f *factory.SQLFactory, agent event.EventAgent, logger *flumesdk.EventSender) (UoW, error) {
	if f == nil {
		errStr := "No *factory.SQLFactory argument! "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, errors.New(errStr)
	}
	if agent == nil {
		errStr := "No event.EventAgent argument! "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, errors.New(errStr)
	}
	uow := new(UoWSQL)
	uow.schema = f
	uow.eventBus = agent
	if logger != nil {
		uow.logAtom = logger.TransAction()
	}
	return uow, nil
}

func holdField(b *bytes.Buffer, f string, i int) {
	if i > 0 {
		b.WriteString(", ")
	}
	if f == "?" {
		b.WriteString(f)
		return
	}
	b.WriteString("`")
	b.WriteString(f)
	b.WriteString("`")
}

func setField(b *bytes.Buffer, schema *factory.Schema, f string, args *[]interface{}, v interface{}, i *int) {
	if f == schema.PrimaryKey {
		return
	}
	holdField(b, f, *i)
	b.WriteString("=?")
	(*args)[*i] = v
	*i++
}

func (uow *UoWSQL) Submit(ctx context.Context, sqlInfo *factory.SQLInfo) error {
	return uow.eventBus.Submit(ctx, sqlInfo)
}

func (uow *UoWSQL) Add(ctx context.Context, model interface{}) error {
	var (
		err        error
		sv         reflect.Value
		schema     *factory.Schema
		sqlBuf     bytes.Buffer
		fieldCount int
		i          int
		attribute  reflect.StructField
		field      string
		args       []interface{}
		modelValue reflect.Value
		sqlHandle  func()
		argsHandle func() error
		pkVal      interface{}
	)

	modelValue = reflect.ValueOf(model).Elem()
	if modelValue.Kind() == reflect.Slice {
		if schema, sv, err = uow.schema.ParseValue(modelValue.Index(0).Addr().Interface()); err != nil {
			syslog.FortLog.ShowLog(syslog.ERROR, err.Error())
			return err
		}
		fieldCount = len(schema.FieldSchema) * modelValue.Len()
		sqlHandle = func() {
			sqlBuf.WriteString(") VALUES ")
			for mIndex := 0; mIndex < modelValue.Len(); mIndex++ {
				if mIndex != 0 {
					sqlBuf.WriteString(",(")
				} else {
					sqlBuf.WriteString(" (")
				}
				for i = 0; i < fieldCount/modelValue.Len(); i++ {
					holdField(&sqlBuf, "?", i)
				}
				sqlBuf.WriteString(")")
			}
		}
		argsHandle = func() error {
			var pkvals []interface{}
			i = 0
			for mIndex := 0; mIndex < modelValue.Len(); mIndex++ {
				var fields []string
				schema, sv, err = uow.schema.ParseValue(modelValue.Index(mIndex).Addr().Interface())
				pkvals = append(pkvals,sv.FieldByName(utils.ToCamelName(schema.PrimaryKey)).Interface())
				//golang map 不保证有序,需要统一排序.防止字段值不对应
				for field := range schema.FieldSchema {
					fields = append(fields, field)
				}
				sort.Strings(fields)
				for _, field := range fields {
					attribute = schema.FieldSchema[field]
					if mIndex == 0 {
						holdField(&sqlBuf, field, i)
					}
					val, err := utils.AsKind(sv.FieldByName(attribute.Name))
					if err != nil {
						return err
					}
					args[i] = val
					i++
				}
			}
			pkVal = pkvals
			return nil
		}
	} else {
		if schema, sv, err = uow.schema.ParseValue(model); err != nil {
			syslog.FortLog.ShowLog(syslog.ERROR, err.Error())
			return err
		}
		fieldCount = len(schema.FieldSchema)
		sqlHandle = func() {
			sqlBuf.WriteString(") VALUES (")
			for i = 0; i < fieldCount; i++ {
				holdField(&sqlBuf, "?", i)
			}
			sqlBuf.WriteString(")")
		}
		argsHandle = func() error {
			i = 0
			for field, attribute = range schema.FieldSchema {
				holdField(&sqlBuf, field, i)
				val, err := utils.AsKind(sv.FieldByName(attribute.Name))
				if err != nil {
					return err
				}
				args[i] = val
				i++
			}
			pkVal = sv.FieldByName(utils.ToCamelName(schema.PrimaryKey)).Interface()
			return nil
		}
	}
	args = make([]interface{}, fieldCount)
	sqlBuf.WriteString("INSERT INTO `")
	sqlBuf.WriteString(schema.Name)
	sqlBuf.WriteString("` (")
	if err := argsHandle(); err != nil {
		syslog.FortLog.ShowLog(syslog.ERROR, err.Error())
		return err
	}
	sqlHandle()
	return uow.Submit(ctx, &factory.SQLInfo{Sql: sqlBuf.String(), Args: args, TableName: schema.Name, PrimaryKeyValue: pkVal}) //todo 涉及缓存
}

func (uow *UoWSQL) Save(ctx context.Context, model interface{}) error {
	var (
		err        error
		sv         reflect.Value
		schema     *factory.Schema
		sqlBuf     bytes.Buffer
		fieldCount int
		i          int
		attribute  reflect.StructField
		field      string
		args       []interface{}
	)
	schema, sv, err = uow.schema.ParseValue(model)
	if err != nil {
		return err
	}
	fieldCount = len(schema.FieldSchema) - 1
	sqlBuf.WriteString("UPDATE ")
	sqlBuf.WriteString(schema.Name)
	sqlBuf.WriteString(" SET ")
	args = make([]interface{}, fieldCount+1)
	i = 0
	for field, attribute = range schema.FieldSchema {
		val, err := utils.AsKind(sv.FieldByName(attribute.Name))
		if err != nil {
			syslog.FortLog.ShowLog(syslog.ERROR, err.Error())
			return err
		}
		setField(&sqlBuf, schema, field, &args, val, &i)
	}
	sqlBuf.WriteString(" WHERE ")
	sqlBuf.WriteString(schema.PrimaryKey)
	sqlBuf.WriteString("=?")
	args[fieldCount], err = utils.AsKind(sv.FieldByName(schema.FieldSchema[schema.PrimaryKey].Name))
	if err != nil {
		syslog.FortLog.ShowLog(syslog.ERROR, err.Error())
		return err
	}
	//return uow.Submit(ctx, sqlBuf.String(), args)
	return uow.Submit(ctx, &factory.SQLInfo{Sql: sqlBuf.String(), Args: args, TableName: schema.Name, PrimaryKeyValue: args[fieldCount]})
}

func (uow *UoWSQL) Modify(ctx context.Context, model interface{}, fields map[string]interface{}) error {
	var (
		err         error
		sv          reflect.Value
		schema      *factory.Schema
		sqlBuf      bytes.Buffer
		updateCount int
		args        []interface{}
		f           string
		v           interface{}
		i           int
	)
	updateCount = len(fields)
	if updateCount < 1 {
		errStr := "No fields need to update! "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return errors.New(errStr)
	}
	schema, sv, err = uow.schema.ParseValue(model)
	if err != nil {
		return err
	}
	sqlBuf.WriteString("UPDATE `")
	sqlBuf.WriteString(schema.Name)
	sqlBuf.WriteString("` SET ")
	args = make([]interface{}, updateCount+1)
	i = 0
	for f, v = range fields {
		setField(&sqlBuf, schema, f, &args, v, &i)
	}
	sqlBuf.WriteString(" WHERE ")
	sqlBuf.WriteString(schema.PrimaryKey)
	sqlBuf.WriteString("=?")

	val, err := utils.AsKind(sv.FieldByName(schema.FieldSchema[schema.PrimaryKey].Name))
	if err != nil {
		syslog.FortLog.ShowLog(syslog.ERROR, err.Error())
	}
	args[i] = val

	pkVal := sv.FieldByName(utils.ToCamelName(schema.PrimaryKey)).Interface()
	return uow.Submit(ctx, &factory.SQLInfo{Sql: sqlBuf.String(), Args: args, TableName: schema.Name, PrimaryKeyValue: pkVal})
}

func (uow *UoWSQL) Remove(ctx context.Context, model interface{}) error {
	var (
		err    error
		sv     reflect.Value
		schema *factory.Schema
		sqlBuf bytes.Buffer
		args   []interface{}
	)
	schema, sv, err = uow.schema.ParseValue(model)
	if err != nil {
		return err
	}
	sqlBuf.WriteString("DELETE FROM `")
	sqlBuf.WriteString(schema.Name)
	sqlBuf.WriteString("` WHERE `")
	sqlBuf.WriteString(schema.PrimaryKey)
	sqlBuf.WriteString("` =?")

	val, err := utils.AsKind(sv.FieldByName(schema.FieldSchema[schema.PrimaryKey].Name))
	if err != nil {
		syslog.FortLog.ShowLog(syslog.ERROR, err.Error())
	}

	args = []interface{}{val}

	pkVal := sv.FieldByName(utils.ToCamelName(schema.PrimaryKey)).Interface()
	return uow.Submit(ctx, &factory.SQLInfo{Sql: sqlBuf.String(), Args: args, TableName: schema.Name, PrimaryKeyValue: pkVal})
}

func (uow *UoWSQL) Commit(ctx context.Context) error {
	err := uow.eventBus.Commit(ctx)
	if err != nil {
		return err
	}
	if uow.logAtom != nil {
		go uow.logAtom.Commit()
	}
	return nil
}

func (uow *UoWSQL) Extend(commands *[]*factory.SQLInfo) error {
	return uow.eventBus.Extend(commands)
}

func (uow *UoWSQL) Log(eventLog flumesdk.EventLog) error {
	return uow.logAtom.Add(eventLog)
}

func (uow *UoWSQL) GetAgent() event.EventAgent {
	return uow.eventBus
}

func (uow *UoWSQL) SetSourceName(sourceName string) {
	uow.eventBus.SetSourceName(sourceName)
}
