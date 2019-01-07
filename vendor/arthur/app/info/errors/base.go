/*
Author : Haoyuan Liu
Time   : 2018/6/20
*/
package errors

import (
	"arthur/utils/errors"
	"fmt"
)

type GameError interface {
	error
}

func New(info string) GameError {
	return errors.New(info)
}

func Newf(info string, args ...interface{}) GameError {
	return errors.New(fmt.Sprintf(info, args...))
}

var (
	ErrNoAppServer    = New("no_app_server_id") //无应用ID
	ErrNoConf         = New("no_config")        //找不到配置
	ErrType           = New("type_error")       //类型错误
	ErrUser           = New("user_error")       //用户错误
	ErrParamIllegal   = New("param_illegal")    //非法参数
	ErrNoAction       = New("no_action")        //找不到接口
	ErrProtoError     = New("proto_error")      //协议错误
	ErrModuleIsClosed = New("module_is_closed") //模块已关闭
	ErrNoModule       = New("no_module")        //无该模块
)
