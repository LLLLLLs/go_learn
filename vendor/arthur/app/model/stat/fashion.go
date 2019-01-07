/*
Created on 2018-08-21 15:46:26
author: Auto Generate
*/
package stat

type Fashion struct {
	No            int16  `model:"pk"  db:"smallint(5)"` //时装编号
	Name          string `db:"char(36)"`                //时装名称
	TimeLimit     int    `db:"int(10)"`                 //时装期限（秒）
	CanEdit       bool   `db:"tinyint(2)"`              //是否可编辑宣言
	Remark        string `db:"char(255)"`               //说明、备注
	Level         int16  `db:"smallint(5)"`             //皮肤等级
	IsDefault     bool   `db:"tinyint(2)"`              //是否为初始时装
	TimeLimitDesc string `db:"char(255)"`               //时效期限描述
}
