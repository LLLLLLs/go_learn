/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type MailAward struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"varchar(36)"` //奖励id
	StartTime   int64  `db:"bigint(20)"`              //奖励领取开始时间
	EndTime     int64  `db:"bigint(20)"`              //奖励领取结束时间
	CreateTime  int64  `db:"bigint(20)"`              //创建时间
	IsFull      int16  `db:"smallint(255)"`           //发送状态 (1: 全服 2:不是全服 )
	IsShield    int16  `db:"smallint(255)"`           //是否屏蔽状态(1:否 ， 2：是 )
	FailureId   string `db:"text"`                    //系统邮件未能成功发送的玩家show_id
	TemplateId  string `db:"varchar(32)"`             //模板id
	AwardItem   string `db:"varchar(255)"`            //奖励
	BlackList   string `db:"text"`                    //黑名单(全服邮件屏蔽人员)
	ToRoles     string `db:"text"`                    //非全服邮件接收者Id
	Comment     string `db:"varchar(255)"`            //备注
}
