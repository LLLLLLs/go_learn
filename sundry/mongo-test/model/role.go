// Time        : 2019/09/05
// Description :

package model

import (
	"golearn/sundry/mongo-test/util"
)

type Role struct {
	RoleId   string `bson:"_id"`
	Students interface{}
}

func (r *Role) MarshalStudents(model interface{}) {
	util.MarshalExtend(r.Students, model)
	r.Students = model
}
