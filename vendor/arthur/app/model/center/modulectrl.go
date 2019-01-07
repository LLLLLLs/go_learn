/*
Created on 2018-11-22 15:31:20
author: Auto Generate
*/
package center

import "arthur/app/enum/ModuleNo"

type ModuleCtrl struct {
	Id       string `model:"pk"  db:"varchar(36)"`
	No       ModuleNo.Type  `db:"int(11)"`    //模块id
	AppId    int    `db:"int(11)"`    //应用id
	ServerId int    `db:"int(11)"`    //服务id
	IsOpen   bool   `db:"tinyint(1)"` //是否开启
}
