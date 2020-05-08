//@author: lls
//@time: 2020/05/07
//@desc:

package example

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golearn/pkgtest/go-memdb/example/po"
	"golearn/pkgtest/go-memdb/memdb"
	"golearn/util"
	"testing"
)

func TestMemDb(t *testing.T) {
	ast := assert.New(t)
	db := memdb.DB()

	tx := db.Txn(true)
	err := tx.Insert("RunningTotal", po.RunningTotal{
		Id:         "abcd",
		ServerId:   1,
		Type:       1,
		BelongTo:   "my",
		Value:      100,
		UpdateTime: 123,
		Extend:     nil,
	})
	ast.Nil(err)
	raw, err := tx.Get("RunningTotal", "id", "abcd")
	ast.Nil(err)
	data := raw.Next().(po.RunningTotal)
	fmt.Println(data)

	d, err := tx.First("RunningTotal", "type_belong", uint16(1), "my")
	ast.Nil(err)
	fmt.Println(d.(po.RunningTotal))

	d, err = tx.First("RunningTotal", "type_belong", uint16(1), "my")
	ast.Nil(err)
	fmt.Println(d.(po.RunningTotal))

	tx.Abort()

	tx = db.Txn(true)
	raw, err = tx.Get("RunningTotal", "id", "abcd")
	ast.Nil(err)
	ast.Nil(raw.Next())
}

func newModel(i int) po.RunningTotal {
	return po.RunningTotal{
		Id:         fmt.Sprintf("%s%d", "rt", i),
		ServerId:   1,
		Type:       uint16(i % 100),
		BelongTo:   fmt.Sprintf("%s%d", "belong", i),
		Value:      int64(i),
		UpdateTime: int64(i),
		Extend:     nil,
	}
}

func insertN(count int) {
	db := memdb.DB()
	tx := db.Txn(true)
	for i := 0; i < count; i++ {
		util.MustNil(tx.Insert("RunningTotal", newModel(i)))
	}
	tx.Commit()
}

func BenchmarkGet(b *testing.B) {
	n := 500000
	insertN(n)
	b.ResetTimer()
	tx := memdb.DB().Txn(false)
	for i := 0; i < b.N; i++ {
		rt := newModel(i % n)
		data, err := tx.First("RunningTotal", "type_belong", rt.Type, rt.BelongTo)
		util.MustNil(err)
		_ = data.(po.RunningTotal)
	}
}
