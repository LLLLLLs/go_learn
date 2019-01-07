/*
Created on 2018-11-06 16:58:41
author: Auto Generate
*/
package dyn

import "arthur/app/infra"

type MailTemplate struct {
	infra.Model `model:"-"`
	Id          string `model:"pk"  db:"varchar(64)"` //邮件模板id
	Lang        string `db:"varchar(255)"`            //语言标识
	Typ         string `db:"varchar(255)"`            //模板类型
	Title       string `db:"varchar(255)"`            //模板标题
	Content     string `db:"text"`                    //模板内容
}
