
/*
Created on 2018-12-12 11:06:29
author: Auto Generate
*/
package center

type ServerName struct {
	Id       	string `model:"pk"` //
	AppId    	int16               //大区
	ServerId 	int16               //服务器Id
	Name        string				//服务器名
	Language    string				//语言类型
}
