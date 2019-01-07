/*
Created on 2018-11-06 16:58:42
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type StudentNew struct {
	infra.Model     `model:"-"`                        //
	Id              string `model:"pk"  db:"char(36)"` //
	RoleId          string `db:"char(36)"`             //玩家id
	MotherNo        int16  `db:"smallint(5)"`          //母亲编号
	SeatNo          int16  `db:"smallint(5)"`          //座位号
	Talent          int16  `db:"smallint(5)"`          //学员天赋
	Name            string `db:"char(10)"`             //学员名称
	Sex             bool   `db:"tinyint(1)"`           //学员性别 男=true，女=false
	Avatar          int16  `db:"smallint(5)"`          //学员形象id
	Wise            int    `db:"int(10)"`              //
	Diligent        int    `db:"int(10)"`              //
	Loyalty         int    `db:"int(10)"`              //
	Heroic          int    `db:"int(10)"`              //
	GrowthPoint     int    `db:"int(10)"`              //学员成长点
	GrowthSpeed     int    `db:"int(10)"`              //成长速度
	GrowTimes       int16  `db:"smallint(5)"`          //上次衰减后成长次数
	Stat            int16  `db:"smallint(5)"`          //学员状态1=正常，2=期中待考，3=期末待考，4=待毕业
	Accepted        bool   `db:"tinyint(1)"`           //是否接受
	AcceptTime      int64  `db:"bigint(20)"`           //学员接受时间
	SpecPeriodTime  int64  `db:"bigint(20)"`           //特殊时期时间（在该时间之后才成长，比如考试）
	MarriageTime    int64  `db:"bigint(20)"`           //联姻时间
	MarriageStudent string `db:"char(36)"`             //联姻学员
	MarriageRole    string `db:"char(36)"`             //联姻学员所属玩家
}
