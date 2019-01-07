/*
Created on 2018-09-17 17:58:39
author: Auto Generate
*/
package center

type LoginRecords struct {
	Id        string `model:"pk"` //
	UserId    string              //用户id
	AppId     int16               //应用号
	ServerId  int16               //区服号
	LoginTime int64               //登录时间
}
