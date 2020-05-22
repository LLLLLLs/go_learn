//@author: lls
//@time: 2020/05/12
//@desc:

package example

import (
	"golearn/pkgtest/go-memdb/example/po"
	"golearn/pkgtest/go-memdb/memdb"
	"golearn/util"
	"testing"
)

var n = 99901

func TestMain(t *testing.M) {
	insertN(n)
	t.Run()
}

// 1k3
func BenchmarkGet(b *testing.B) {
	tx := memdb.Tx{Txn: memdb.DB().Txn(false)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rt := newModel(i % n)
		has, err := tx.Get(&rt)
		if !has {
			panic("1324")
		}
		util.MustNil(err)
	}
}

// 2k
func BenchmarkFindByPK(b *testing.B) {
	tx := memdb.Tx{Txn: memdb.DB().Txn(false)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rt := newModel(i % n)
		result := make([]po.RunningTotal, 0)
		err := tx.Find(po.RunningTotal{Id: rt.Id}, &result)
		util.MustNil(err)
	}
}

// 6.7w
func BenchmarkFindByFK(b *testing.B) {
	tx := memdb.Tx{Txn: memdb.DB().Txn(false)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rt := newModel(99900)
		result := make([]po.RunningTotal, 0)
		err := tx.Find(po.RunningTotal{BelongTo: rt.BelongTo}, &result)
		util.MustNil(err)
	}
}

// 2.5w
func BenchmarkFindByIndex(b *testing.B) {
	tx := memdb.Tx{Txn: memdb.DB().Txn(false)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rt := newModel(i % n)
		result := make([]po.RunningTotal, 0)
		err := tx.Find(po.RunningTotal{Type: rt.Type, BelongTo: rt.BelongTo}, &result)
		util.MustNil(err)
	}
}
