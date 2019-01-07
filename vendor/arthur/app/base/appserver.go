/*
Author : Haoyuan Liu
Time   : 2018/4/19 9:38
*/
package base

import (
	"arthur/utils/panicutils"
	"strconv"
	"strings"
	"arthur/env"
)

var sep = env.REDIS_KEY_SEP

// AppServerId
type ASID string

func (asid ASID) ToString() string {
	return string(asid)
}

//区服信息
//
//	由于 AppId 和 ServerId 有着强关联的关系，不少地方需要同时传入这两个编号
//	因此声明 AppServer 结构体，携带区服信息（包括RegionId）。
type AppServer struct {
	AppId       int
	ServerId    int
	appServerId ASID
}

func NewAppServer(appId, serverId int) AppServer {
	as := AppServer{}
	as.AppId = appId
	as.ServerId = serverId
	as.appServerId = CombineASID(appId, serverId)
	return as
}

// ASID
func (as AppServer) ID() ASID {
	return as.appServerId
}

func (as AppServer) Equal(other AppServer) bool {
	return as.appServerId == other.appServerId
}

// 大区Id
func (as AppServer) RegionId() string {
	return ""
}

// 将appId和serverId合并
func CombineASID(appId, serverId int) ASID {
	return ASID(strings.Join([]string{strconv.Itoa(appId), strconv.Itoa(serverId)}, sep))
}

// 将appServerId拆解为appId和serverId
func separate(appServerId string) (appId, serverId int) {
	asiSlice := strings.Split(appServerId, sep)
	if len(asiSlice) != 2 {
		panic("wrong appServerId")
	}
	var err1, err2 error
	appId, err1 = strconv.Atoi(asiSlice[0])
	serverId, err2 = strconv.Atoi(asiSlice[1])
	panicutils.OkOrPanic(err1)
	panicutils.OkOrPanic(err2)
	return appId, serverId
}
