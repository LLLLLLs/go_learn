/*
Author : Haoyuan Liu
Time   : 2018/6/22
*/
package errors

var (
	ErrNoRole               = New("no_role")                //无该角色
	ErrHasBeenInit          = New("has_been_init")          //角色已被初始化
	ErrNameExist            = New("name_exist")             //名字已被占用
	ErrName                 = New("name_illegal")           //名字包含非法字符
	ErrNoGold               = New("no_gold")                //无元宝
	ErrNoRes                = New("no_res")                 //无资源
	ErrMaxRoleLv            = New("max_role_lv")            //角色等级已升至最大
	ErrInsuffExp            = New("insuff_exp")             //经验不足
	ErrNoRoleTitle          = New("no_role_title")          //无角色称号
	ErrReceiveWithoutFinish = New("receive_without_finish") //任务尚未完成，无法领取奖励
	ErrNoRoleValue          = New("no_role_value")          //角色数值不足
)
