/*
Created on 2018-09-17 17:58:09
author: Auto Generate
*/
package center

type ServerInfo struct {
	Id            string `model:"pk"` //
	AppId         int16               //大区
	ServerId      int16               //服务器
	Stat          int16               //服务器状态:1、推荐；2、火爆；3、维护
	ServerName    string              //服务器名称
	DbName    	  string              //数据库名称
	ServerUrl     string          	  //区服地址
	SocketIp      string              //区服ip
	SocketPort    string              //区服端口
	OpenTime      int64               //开区时间
	AnnounceState bool                //公告状态
	AnnounceUrl   string              //公告地址
	ConnType      string              //链接类型
	AuditService  bool                //审核服
	Version       string              //版本
}
