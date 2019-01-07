/*
Author : Haoyuan Liu
Time   : 2018/6/25
*/
package dcapi

import (
	"fmt"
)

type Action interface {
	ID() string
	Params([]byte) (interface{}, error) //将传入bytes转为参数结构体
	Do(Context)                         //执行动作
}

type BaseAct struct {
	id string
}

func (act *BaseAct) ID() string {
	return act.id
}

func NewBaseAction(id string) BaseAct {
	return BaseAct{id}
}

//Manager 用于管理Context要执行的handler和Action的映射
type Manager struct {
	handlers  HandlersChain
	actionMap map[string]Action
}

func (m *Manager) Handlers() HandlersChain {
	return m.handlers
}

func (m *Manager) Register(act Action) {
	actId := act.ID()
	_, ok := m.actionMap[actId]
	if ok {
		panic(fmt.Sprintf("action: %s has been registered", actId))
	}
	m.actionMap[actId] = act
}

func (m *Manager) GetAction(actId string) (act Action, ok bool) {
	act, ok = m.actionMap[actId]
	return act, ok
}

func (m *Manager) Use(middleware ...HandlerFunc) {
	m.handlers = append(m.handlers, middleware...)
}

func New() *Manager {
	return &Manager{
		handlers:  make(HandlersChain, 0),
		actionMap: make(map[string]Action),
	}
}
