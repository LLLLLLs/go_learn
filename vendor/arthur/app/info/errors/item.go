/*
Author : Haoyuan Liu
Time   : 2018/7/9
*/
package errors

var (
	ErrNoItem      = New("no_item")      //道具不足
	ErrNotCompound = New("not_compound") //无合成配置
)
