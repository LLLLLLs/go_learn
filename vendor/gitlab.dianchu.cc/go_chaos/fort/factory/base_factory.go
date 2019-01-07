package factory

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
	"gitlab.dianchu.cc/go_chaos/fort/utils"
)

type ModelFactory interface {
	//Register 注册模型，注册后的模型可
	Register(modelSet ...interface{}) error
	ParseValue(model interface{}) (*Schema, reflect.Value, error)
	ParseModel(model interface{}) (*Schema, error)
}

//Schema SQL表Schema
type Schema struct {
	ModelType       reflect.Type
	Name            string
	PrimaryKey      string
	PrimaryKeyValue interface{}
	AttributeField  map[string]string
	FieldSchema     map[string]reflect.StructField
}

type SchemaFactory struct {
	ModelSpec map[string]*Schema //模型Schema映射
}

type TableName interface {
	TableName() string
}

func (mf *SchemaFactory) Register(modelSet ...interface{}) error {
	var (
		model     interface{}
		st        reflect.Type
		sv        reflect.Value
		modelName string
		field     reflect.StructField
	)
	for _, model = range modelSet {
		st = reflect.TypeOf(model).Elem()
		sv = reflect.Indirect(reflect.ValueOf(model))
		modelName = st.Name()
		mf.ModelSpec[modelName] = new(Schema)
		mf.ModelSpec[modelName].AttributeField = make(map[string]string, st.NumField())
		mf.ModelSpec[modelName].FieldSchema = make(map[string]reflect.StructField, st.NumField())
		if tb, ok := sv.Interface().(TableName); ok {
			mf.ModelSpec[modelName].Name = utils.ToUnderscoreName(tb.TableName())
		} else {
			mf.ModelSpec[modelName].Name = utils.ToUnderscoreName(modelName)
		}
		for i := 0; i < st.NumField(); i++ {
			field = st.Field(i)
			if err := mf.parseTag(&field, modelName); err != nil {
				return err
			}
		}
		mf.ModelSpec[modelName].ModelType = st
		if mf.ModelSpec[modelName].PrimaryKey == "" {
			errStr := "primary key is empty"
			syslog.FortLog.ShowLog(syslog.ERROR, errStr)
			return errors.New(errStr)
		}
	}
	return nil
}

func (mf *SchemaFactory) ParseValue(model interface{}) (*Schema, reflect.Value, error) {
	var (
		sv     reflect.Value
		schema *Schema
		ok     bool
	)
	sv = reflect.ValueOf(model)
	if sv.Kind() != reflect.Ptr {
		errStr := sv.Kind().String() + " is not ptr type!"
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, sv, errors.New(errStr)
	}
	sv = sv.Elem()
	if sv.Kind() != reflect.Struct {
		errStr := sv.Kind().String() + " is not struct type!"
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, sv, errors.New(errStr)
	}
	if schema, ok = mf.ModelSpec[sv.Type().Name()]; !ok {
		errStr := "Model " + sv.Type().Name() + " does not exist!"
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, sv, errors.New(errStr)
	}
	return schema, sv, nil
}

func (mf *SchemaFactory) ParseModel(model interface{}) (*Schema, error) {
	var (
		st     reflect.Type
		schema *Schema
		ok     bool
	)
	st = reflect.TypeOf(model)
	if st.Kind() != reflect.Ptr && st.Kind() != reflect.Slice {
		errStr := st.Kind().String() + " is not ptr type or not slice type!"
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, errors.New(errStr)
	}
	st = st.Elem()
	if st.Kind() == reflect.Slice {
		st = st.Elem()
	}
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}
	if st.Kind() != reflect.Struct {
		errStr := st.Kind().String() + " is not struct type!"
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, errors.New(errStr)
	}
	schema, ok = mf.ModelSpec[st.Name()]
	if !ok {
		errStr := "Model " + st.Name() + " does not exist!"
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return nil, errors.New(errStr)
	}
	return schema, nil
}

func (mf *SchemaFactory) clearModel(modelName string) {
	mf.ModelSpec[modelName] = nil
}

func (mf *SchemaFactory) repeatCheck(modelName, fieldName string) error {
	if _, ok := mf.ModelSpec[modelName].AttributeField[fieldName]; ok {
		mf.clearModel(modelName)
		errStr := modelName + ":" + fieldName + " is repeating "
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return errors.New(errStr)
	}
	return nil
}

func splitTag(tag string) map[string]interface{} {
	tag = strings.TrimSpace(tag)
	if len(tag) == 0 {
		return nil
	}
	if tag[len(tag)-1] != ';' {
		tag = tag + ";"
	}
	var (
		tags     = make(map[string]interface{})
		hasValue = false
		lastIdx  = 0
		lastTag  = ""
	)
	get := func(i int) {
		if lastIdx < i {
			lastTag = strings.TrimSpace(tag[lastIdx:i])
			tags[lastTag] = true
			lastIdx = i + 1
		}
	}
	for i, t := range tag {
		switch t {
		case ':':
			hasValue = !hasValue
			get(i)
		case ';':
			if hasValue {
				hasValue = !hasValue
				tags[lastTag] = strings.TrimSpace(tag[lastIdx:i])
				lastIdx = i + 1
			} else {
				get(i)
			}
		}
	}
	return tags
}

func (mf *SchemaFactory) parseTag(field *reflect.StructField, modelName string) error {
	fieldTag := splitTag(field.Tag.Get("model"))
	fieldName := field.Name
	if t, ok := fieldTag["-"].(bool); ok && t {
		return nil
	}

	if field.Type.Kind() == reflect.Struct && field.Type != reflect.TypeOf(time.Time{}) {
		if !field.Anonymous {
			return nil
		}
		mf.Register(reflect.New(field.Type).Interface())
		return mf.mergeModelField(modelName, fieldName)
	}

	if len(fieldTag) > 0 {
		if t, ok := fieldTag["column"].(string); ok {
			fieldName = t
		}
		if t, ok := fieldTag["pk"].(bool); ok && t {
			if mf.ModelSpec[modelName].PrimaryKey != "" {
				errStr := "primary key is repeating"
				syslog.FortLog.ShowLog(syslog.ERROR, errStr)
				return errors.New(errStr)
			}
			mf.ModelSpec[modelName].PrimaryKey = utils.ToUnderscoreName(fieldName)
		}
	}
	if err := mf.repeatCheck(modelName, fieldName); err != nil {
		return err
	}
	tableField := utils.ToUnderscoreName(fieldName)
	mf.ModelSpec[modelName].AttributeField[fieldName] = tableField
	mf.ModelSpec[modelName].FieldSchema[tableField] = *field
	return nil
}

func (mf *SchemaFactory) mergeModelField(masterModelName, slaveModelName string) error {
	Master := mf.ModelSpec[masterModelName]
	Slave := mf.ModelSpec[slaveModelName]
	if Master.PrimaryKey != "" && Slave.PrimaryKey != "" {
		errStr := "primary key is repeating"
		syslog.FortLog.ShowLog(syslog.ERROR, errStr)
		return errors.New(errStr)
	} else if Master.PrimaryKey == "" {
		Master.PrimaryKey = Slave.PrimaryKey
	}

	for fieldName, tableName := range Slave.AttributeField {
		if err := mf.repeatCheck(masterModelName, fieldName); err != nil {
			return err
		}
		Master.AttributeField[fieldName] = tableName
		Master.FieldSchema[tableName] = Slave.FieldSchema[tableName]
	}
	return nil
}
