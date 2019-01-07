/*
Created on 2018-11-08 15:26:40
author: Auto Generate
*/
package stat

type DiplomacyAchievement struct {
	No           int     `model:"pk"  db:"int(11)"` //
	Name         string  `db:"varchar(255)"`        //每日外交成就名称
	Require      int     `db:"int(255)"`            //连胜需求
	Rate         float64 `db:"float(255,2)"`        //奖励系数
	Award        string  `db:"varchar(255)"`        //奖励组
	Resource     float64 `db:"float(255,2)"`        //资源系数
	ResourceRate int     `db:"int(11)"`             //资源奖励出现的概率
}
