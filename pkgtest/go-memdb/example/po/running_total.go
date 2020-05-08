//@author: lls
//@time: 2020/05/07
//@desc:

package po

import "golearn/pkgtest/go-memdb/memdb"

type RunningTotal struct {
	Id         string `model:"pk"`
	ServerId   uint   `index:"server_type"`
	Type       uint16 `index:"type_belong,server_type"`
	BelongTo   string `model:"fk" index:"type_belong"`
	Value      int64
	UpdateTime int64
	Extend     map[string]string
}

func init() {
	memdb.RegisterSchema(RunningTotal{})
}
