// Time        : 2019/10/24
// Description :

package util

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type StudentValue struct {
	Id            string `bson:"_id"` // 唯一ID
	Name          string // 名字
	BeautyNo      int16  // 名媛
	Sex           int16  // 性别
	Talent        int16  // 资质
	Power         int64  // 属性
	Prof          int16  // 职业
	Status        int16  // 状态 1=婴儿 2=幼年 3=成年 4=待授勋
	Exp           int    // 经验
	RecoverRemain int64  // 活力回满剩余时间
}

// 31663354                36.3 ns/op             8 B/op          1 allocs/op
func BenchmarkMarshal_Int64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ms int64
		MarshalExtend(int64(100), &ms)
	}
}

// 23592555                50.5 ns/op            16 B/op          1 allocs/op
func BenchmarkMarshal_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ms string
		MarshalExtend("extend", &ms)
	}
}

// 708322              1747 ns/op             640 B/op         12 allocs/op
func BenchmarkMarshal_1Attr_Doc(b *testing.B) {
	extend := bson.D{
		{Key: "id", Value: "studentId0"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ms StudentValue
		MarshalExtend(extend, &ms)
	}
}

// 708322              1747 ns/op             640 B/op         12 allocs/op
func BenchmarkMarshal_5Attr_Doc(b *testing.B) {
	extend := bson.D{
		{Key: "id", Value: "studentId0"},
		{Key: "name", Value: "studentName0"},
		{Key: "beautyno", Value: 0},
		{Key: "sex", Value: 1},
		{Key: "talent", Value: 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ms StudentValue
		MarshalExtend(extend, &ms)
	}
}

// 429296              2820 ns/op             989 B/op         23 allocs/op
func BenchmarkMarshal_5Attr_Map_Bson(b *testing.B) {
	extend := map[string]interface{}{
		"id":       "studentId0",
		"name":     "studentName0",
		"beautyno": 0,
		"sex":      1,
		"talent":   0,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ms StudentValue
		marshalBson(extend, &ms)
	}
}

// 316402              3864 ns/op            1112 B/op         23 allocs/op
func BenchmarkMarshal_5Attr_Map_Json(b *testing.B) {
	extend := map[string]interface{}{
		"id":       "studentId0",
		"name":     "studentName0",
		"beautyno": 0,
		"sex":      1,
		"talent":   0,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ms StudentValue
		marshalJson(extend, &ms)
	}
}

// 429726              2713 ns/op             672 B/op         17 allocs/op
func BenchmarkMarshal_10Attr_Doc(b *testing.B) {
	extend := bson.D{
		{Key: "id", Value: "studentId0"},
		{Key: "name", Value: "studentName0"},
		{Key: "beautyno", Value: 0},
		{Key: "sex", Value: 1},
		{Key: "talent", Value: 0},
		{Key: "power", Value: 0},
		{Key: "prof", Value: 0},
		{Key: "status", Value: 0},
		{Key: "exp", Value: 10},
		{Key: "recoverremain", Value: 1321},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ms StudentValue
		MarshalExtend(extend, &ms)
	}
}

// 414502              2957 ns/op             817 B/op         20 allocs/op
func BenchmarkMarshal_1Elem_Array(b *testing.B) {
	extend := bson.A{
		bson.D{
			{Key: "id", Value: "studentId0"},
			{Key: "name", Value: "studentName0"},
			{Key: "beautyno", Value: 0},
			{Key: "sex", Value: 1},
			{Key: "talent", Value: 0},
			{Key: "power", Value: 0},
			{Key: "prof", Value: 0},
			{Key: "status", Value: 0},
			{Key: "exp", Value: 10},
			{Key: "recoverremain", Value: 1321},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ms = make([]StudentValue, 0)
		MarshalExtend(extend, &ms)
	}
}

// 106477             11325 ns/op            3043 B/op         71 allocs/op
func BenchmarkMarshal_4Elem_Array(b *testing.B) {
	extend := bson.A{
		bson.D{
			{Key: "id", Value: "studentId0"},
			{Key: "name", Value: "studentName0"},
			{Key: "beautyno", Value: 0},
			{Key: "sex", Value: 1},
			{Key: "talent", Value: 0},
			{Key: "power", Value: 0},
			{Key: "prof", Value: 0},
			{Key: "status", Value: 0},
			{Key: "exp", Value: 10},
			{Key: "recoverremain", Value: 1321},
		},
		bson.D{
			{Key: "id", Value: "studentId1"},
			{Key: "name", Value: "studentName1"},
			{Key: "beautyno", Value: 12},
			{Key: "sex", Value: 1},
			{Key: "talent", Value: 4},
			{Key: "power", Value: 777},
			{Key: "prof", Value: 1},
			{Key: "status", Value: 2},
			{Key: "exp", Value: 10},
			{Key: "recoverremain", Value: 1254},
		},
		bson.D{
			{Key: "id", Value: "studentId2"},
			{Key: "name", Value: "studentName2"},
			{Key: "beautyno", Value: 3},
			{Key: "sex", Value: 2},
			{Key: "talent", Value: 5},
			{Key: "power", Value: 9999},
			{Key: "prof", Value: 3},
			{Key: "status", Value: 3},
			{Key: "exp", Value: 999},
			{Key: "recoverremain", Value: 1452},
		},
		bson.D{
			{Key: "id", Value: "studentId3"},
			{Key: "name", Value: "studentName3"},
			{Key: "beautyno", Value: 10},
			{Key: "sex", Value: 1},
			{Key: "talent", Value: 3},
			{Key: "power", Value: 123456},
			{Key: "prof", Value: 1},
			{Key: "status", Value: 2},
			{Key: "exp", Value: 980},
			{Key: "recoverremain", Value: 1526},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var ms = make([]StudentValue, 0)
		MarshalExtend(extend, &ms)
	}
}
