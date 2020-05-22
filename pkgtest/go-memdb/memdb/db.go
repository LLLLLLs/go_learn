//@author: lls
//@time: 2020/05/07
//@desc: 内存数据库

package memdb

import (
	"github.com/hashicorp/go-memdb"
	"golearn/util"
	"reflect"
	"strings"
)

var (
	db     *memdb.MemDB
	schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{},
	}
	indexCols = make(indexesCols) // table --> indexes --> cols
)

type indexesCols map[string]map[string][]string

func (ic indexesCols) cols(table, index string) []string {
	return ic[table][index]
}

func RegisterSchema(table interface{}) {
	t := reflect.TypeOf(table)
	ig := newIndexGenerator()
	for i := 0; i < t.NumField(); i++ {
		switch t.Field(i).Tag.Get("model") {
		case "pk":
			ig.addPk(t.Field(i))
		case "fk":
			ig.addFk(t.Field(i))
		}
		index := t.Field(i).Tag.Get("index")
		if index == "" {
			continue
		}
		indexes := strings.Split(index, ",")
		for j := range indexes {
			ig.addIndex(indexes[j], t.Field(i))
		}
	}
	indexSchema, indexCol := ig.result()
	schema.Tables[t.Name()] = &memdb.TableSchema{
		Name:    t.Name(),
		Indexes: indexSchema,
	}
	indexCols[t.Name()] = indexCol
}

func DB() *memdb.MemDB {
	if db == nil {
		var err error
		db, err = memdb.NewMemDB(schema)
		util.MustNil(err)
	}
	return db
}
