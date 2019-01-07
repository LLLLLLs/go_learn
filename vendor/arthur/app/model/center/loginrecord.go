/*
Created on 2018-09-17 17:58:33
author: Auto Generate
*/
package center

type LoginRecord struct {
	UserId        string `model:"pk"` //用户id
	CurrentApp    int16               //当前登录大区
	CurrentServer int16               //当前登录服务器
	LastApp       int16               //上次登录大区
	LastServer    int16               //上次登录服务器
}
