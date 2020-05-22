//@author: lls
//@time: 2020/05/09
//@desc:

package memdb

import (
	"errors"
	"github.com/hashicorp/go-memdb"
	"reflect"
	"strings"
)

type Tx struct {
	zeroMatchedCols []string
	*memdb.Txn
}

func (t Tx) reset() {
	t.zeroMatchedCols = make([]string, 0)
}

func (t Tx) WithCol(cols ...string) Tx {
	t.zeroMatchedCols = append(t.zeroMatchedCols, cols...)
	return t
}

func (t Tx) zeroMatched(col string) bool {
	for i := range t.zeroMatchedCols {
		if t.zeroMatchedCols[i] == col {
			return true
		}
	}
	return false
}

var lackOfPK = errors.New("lack of pk")

// 根据主键查询
func (t Tx) Get(model interface{}) (bool, error) {
	defer t.reset()
	modelV := reflect.ValueOf(model)
	if modelV.Kind() != reflect.Ptr {
		panic("must provide a ptr model")
	}
	return t.WithCol(indexCols.cols(modelV.Elem().Type().Name(), "id")[0]).get(modelV)
}

func (t Tx) get(value reflect.Value) (bool, error) {
	if value.Kind() != reflect.Ptr {
		panic("must provide a ptr model")
	}
	elem := value.Elem()
	typ := elem.Type()
	fieldName := indexCols.cols(typ.Name(), "id")[0]
	field := elem.FieldByName(fieldName)
	if field.IsZero() && !t.zeroMatched(fieldName) {
		return false, lackOfPK
	}
	result, err := t.First(typ.Name(), "id", field.Interface())
	if err != nil || result == nil {
		return false, err
	}
	elem.Set(reflect.ValueOf(result))
	return true, nil
}

type findHelper struct {
	tx          Tx
	cond        interface{}
	condValue   reflect.Value
	condType    reflect.Type
	result      interface{}
	resultPtr   reflect.Value
	resultValue reflect.Value
}

func (t Tx) buildFindHelper(cond interface{}, result interface{}) *findHelper {
	h := findHelper{tx: t}
	h.cond = cond
	h.condValue = reflect.ValueOf(cond)
	if h.condValue.Kind() == reflect.Ptr {
		h.condValue = h.condValue.Elem()
	}
	h.condType = h.condValue.Type()
	h.result = result
	h.resultPtr = reflect.ValueOf(result)
	if h.resultPtr.Kind() != reflect.Ptr {
		panic("result must provide a ptr")
	}
	h.resultValue = reflect.MakeSlice(h.resultPtr.Elem().Type(), 0, 1)
	return &h
}

func (fh *findHelper) tableName() string {
	return fh.condType.Name()
}

func (fh *findHelper) byPK() error {
	pkResult := reflect.New(fh.condType)
	pkResult.Elem().Set(fh.condValue)
	has, err := fh.tx.get(pkResult)
	if err != nil || !has {
		return err
	}
	if fh.tx.equal(fh.cond, pkResult) {
		fh.resultValue = reflect.Append(fh.resultValue, reflect.ValueOf(pkResult))
		fh.resultPtr.Elem().Set(fh.resultValue)
	}
	return nil
}

func (fh *findHelper) byFK(fks []string) (bool, error) {
	for i := range fks {
		fieldName := indexCols.cols(fh.tableName(), fks[i])[0]
		field := fh.condValue.FieldByName(fieldName)
		if field.IsZero() && !fh.tx.zeroMatched(fieldName) {
			continue
		}
		resultIt, err := fh.tx.Txn.Get(fh.tableName(), fks[i], field.Interface())
		if err != nil {
			return false, err
		}
		fh.byIter(resultIt)
		return true, nil
	}
	return false, nil
}

func (fh *findHelper) byMultiIndexes(multi map[string][]memdb.Indexer) (bool, error) {
	for name, indexes := range multi {
		cols := make([]interface{}, 0)
		for i := range indexes {
			fieldName := indexCols.cols(fh.tableName(), name)[i]
			field := fh.condValue.FieldByName(fieldName)
			if field.IsZero() && !fh.tx.zeroMatched(fieldName) {
				break
			}
			cols = append(cols, field.Interface())
		}
		if len(cols) != len(indexes) {
			continue
		}
		resultIt, err := fh.tx.Txn.Get(fh.tableName(), name, cols...)
		if err != nil {
			return false, err
		}
		fh.byIter(resultIt)
		return true, nil
	}
	return false, nil
}

func (fh *findHelper) byIter(it memdb.ResultIterator) {
	data := it.Next()
	for ; data != nil; data = it.Next() {
		if fh.tx.equal(fh.cond, data) {
			fh.resultValue = reflect.Append(fh.resultValue, reflect.ValueOf(data))
		}
	}
	fh.resultPtr.Elem().Set(fh.resultValue)
}

func (t Tx) Find(cond interface{}, result interface{}) error {
	defer t.reset()

	fh := t.buildFindHelper(cond, result)
	mSchema := schema.Tables[fh.tableName()]

	fks := make([]string, 0)
	multi := make(map[string][]memdb.Indexer)
	uniqueMulti := make(map[string][]memdb.Indexer)
	for name := range mSchema.Indexes {
		if strings.HasPrefix(name, "fk") {
			fks = append(fks, name)
			continue
		}
		if name != "id" && strings.HasSuffix(name, "unique") {
			uniqueMulti[name] = mSchema.Indexes[name].Indexer.(*memdb.CompoundMultiIndex).Indexes
		} else if name != "id" {
			multi[name] = mSchema.Indexes[name].Indexer.(*memdb.CompoundMultiIndex).Indexes
		}
	}
	err := fh.byPK()
	if err != lackOfPK {
		return err
	}
	matched, err := fh.byMultiIndexes(uniqueMulti)
	if err != nil || matched {
		return err
	}
	matched, err = fh.byFK(fks)
	if err != nil || matched {
		return err
	}
	_, err = fh.byMultiIndexes(multi)
	return err
}

func (t Tx) equal(target, result interface{}) bool {
	tValue := reflect.ValueOf(target)
	tType := tValue.Type()
	rValue := reflect.ValueOf(result)
	if tValue.Type() != rValue.Type() {
		return false
	}
	for i := 0; i < tValue.NumField(); i++ {
		tField := tValue.Field(i)
		if tField.IsZero() && !t.zeroMatched(tType.Field(i).Name) {
			continue
		}
		if !reflect.DeepEqual(tField.Interface(), rValue.Field(i).Interface()) {
			return false
		}
	}
	return true
}
