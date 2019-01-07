/*
Author : Haoyuan Liu
Time   : 2018/7/25
*/
package infra

import (
	"arthur/app/base"
)

//唯一标识符，Id确定唯一对象
type Identifier interface {
	ID() string                //数据库主键
	AppServer() base.AppServer //所属区服
}

//区服Model
type Model interface {
	Identifier
}

type model struct {
	id string
	as base.AppServer
}

func (m model) ID() string {
	return m.id
}

func (m model) AppServer() base.AppServer {
	return m.as
}

func NewModel(id string, as base.AppServer) Model {
	return &model{id: id, as: as}
}

//判断两个Identifier是否相等
func IsModelEqual(my, other Model) bool {
	if my.ID() == other.ID() {
		if my.AppServer().AppId == other.AppServer().AppId && my.AppServer().ServerId == other.AppServer().ServerId {
			return true
		}
	}
	return false
}
