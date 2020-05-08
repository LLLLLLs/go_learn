//@author: lls
//@time: 2020/05/07
//@desc:

package memdb

import (
	"github.com/hashicorp/go-memdb"
	"reflect"
)

type IndexGenerator struct {
	pk      *reflect.StructField
	fk      []reflect.StructField
	indexes map[string][]reflect.StructField
}

func newIndexGenerator() *IndexGenerator {
	return &IndexGenerator{
		pk:      nil,
		fk:      make([]reflect.StructField, 0),
		indexes: make(map[string][]reflect.StructField),
	}
}

func (ig *IndexGenerator) addPk(t reflect.StructField) {
	ig.pk = &t
}

func (ig *IndexGenerator) addFk(t reflect.StructField) {
	ig.fk = append(ig.fk, t)
}

func (ig *IndexGenerator) addIndex(name string, t reflect.StructField) {
	list, ok := ig.indexes[name]
	if !ok {
		list = make([]reflect.StructField, 0)
	}
	list = append(list, t)
	ig.indexes[name] = list
}

func (ig IndexGenerator) result() map[string]*memdb.IndexSchema {
	indexSchemas := make(map[string]*memdb.IndexSchema)
	if ig.pk == nil {
		panic("table must provide pk")
	}
	indexSchemas["id"] = &memdb.IndexSchema{
		Name:         "id",
		AllowMissing: false,
		Unique:       true,
		Indexer:      indexSchema(*ig.pk),
	}
	for i := range ig.fk {
		indexSchemas[ig.fkName(ig.fk[i])] = &memdb.IndexSchema{
			Name:         ig.fkName(ig.fk[i]),
			AllowMissing: false,
			Unique:       false,
			Indexer:      indexSchema(ig.fk[i]),
		}
	}
	for name, list := range ig.indexes {
		indexes := make([]memdb.Indexer, len(list))
		for i := range list {
			indexes[i] = indexSchema(list[i])
		}
		indexSchemas[name] = &memdb.IndexSchema{
			Name:         name,
			AllowMissing: false,
			Unique:       false,
			Indexer: &memdb.CompoundMultiIndex{
				Indexes:      indexes,
				AllowMissing: false,
			},
		}
	}
	return indexSchemas
}

func (ig IndexGenerator) fkName(t reflect.StructField) string {
	return "fk" + t.Name
}

func indexSchema(t reflect.StructField) memdb.Indexer {
	switch t.Type.Kind() {
	case reflect.String:
		return &memdb.StringFieldIndex{
			Field:     t.Name,
			Lowercase: false,
		}
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		return &memdb.IntFieldIndex{Field: t.Name}
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		return &memdb.UintFieldIndex{Field: t.Name}
	default:
		panic("unsupported index typ")
	}
}
