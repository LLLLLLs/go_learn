// Time        : 2019/09/02
// Description :

package mongotest

import (
	"context"
	model2 "golearn/sundry/mongo-test/model"
	"golearn/util"
	"math"
	"strconv"
)

func insertStudent(stu model2.StudentValue) {
	ctx := context.Background()
	collection := client.Database("test").Collection("student")
	_, err := collection.InsertOne(ctx, stu)
	util.OkOrPanic(err)
}

func insertRole(id string, stuNum int) model2.Role {
	ctx := context.Background()
	collection := client.Database("test").Collection("role")
	role := model2.Role{
		RoleId:   id,
		Students: make([]model2.StudentValue, 0),
	}
	for i := 1; i <= stuNum; i++ {
		role.Students = append(role.Students, model2.StudentValue{
			Id:            "studentId" + strconv.Itoa(i),
			Name:          "studentName" + strconv.Itoa(i),
			BeautyNo:      int16(i),
			Sex:           int16(i)%2 + 1,
			Talent:        int16(i),
			Power:         999 * int64(i),
			Prof:          int16(i),
			Status:        int16(i),
			Exp:           10,
			RecoverRemain: int64(util.RandInt(1000, 1800)),
		})
	}
	_, err := collection.InsertOne(ctx, role)
	util.OkOrPanic(err)
	return role
}

func insertTest() {
	ctx := context.Background()
	collection := client.Database("test").Collection("test")
	info := struct {
		Id      int `bson:"_id"`
		F32     float32
		F64     float64
		B       bool
		UI8     uint8
		UI16    uint16
		UI      uint
		UI32    uint32
		UI32Max uint32
		UI64    uint64
		Object  *struct{}
	}{
		Id:      util.RandInt(10000, 99999),
		F32:     100.1,
		F64:     100.2,
		B:       false,
		UI8:     99,
		UI16:    12345,
		UI:      123456789,
		UI32:    123456789,
		UI32Max: math.MaxUint32,
		UI64:    math.MaxUint64 / 2,
		Object:  &struct{}{},
	}
	_, err := collection.InsertOne(ctx, info)
	util.OkOrPanic(err)
}

func insertPhase() {
	ctx := context.Background()
	collection := client.Database("test2").Collection("phase")
	_, err := collection.InsertOne(ctx, model2.Phase{
		Index1: 2,
		Index2: 3,
		Index3: 4,
		Conf:   "phase 2.3.4 config",
	})
	util.OkOrPanic(err)
}
