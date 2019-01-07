/*
Author      : lls
Time        : 2018/08/20
Description : 嵌入uow的行为日志
*/

package behaviorloguow

import (
	"arthur/utils/timeutils"
	"gitlab.dianchu.cc/chaos_go_sdk/flume_client_sdk_go/flumesdk"
	log2 "log"
	"os"
)

var (
	eventSender *flumesdk.EventSender
	Toggle      bool
)

func Init(zkServers []string, flumePath, confPath string, debug bool) {

	var logging *log2.Logger
	logging = log2.New(os.Stderr, "Flume_SDK", log2.Ldate|log2.Ltime|log2.Lshortfile)

	// 获取日志发送器,日志发送器为全局单例
	eventSender = flumesdk.NewEventSender(zkServers, flumePath, confPath, logging, debug)
}

func EventSender() *flumesdk.EventSender {
	return eventSender
}

func Close() {
	if Toggle {
		eventSender.Close()
	}
}

/*
	虚拟币（元宝）日志
		roleId			角色rid
		actionId		scene数值变化场景, 在此为接口号
		old				原始值
		new				变化后的值
		appId			应用id
		serverId		区服id
		remark			备注
		reserve			保留字段
*/
func CoinLog(fieldName, roleId, actionId, remark, reserve string, lifeTime, old, new, appId, serverId int) flumesdk.EventLog {
	newLifeTime := int16(lifeTime)
	newOld := int64(old)
	newNew := int64(new)
	newServerId := int32(serverId)
	newAppId := int32(appId)
	log := flumesdk.NewCoinLog(fieldName, newOld, newNew, newLifeTime, actionId, remark, roleId, reserve, newServerId, newAppId, timeutils.Now())
	return log
}

/*
	行为日志
		roleId			角色id
		actId			接口标识
		req				请求内容
		res				响应内容
		reserve			保留字段
		actStat			访问状态
		procTime		处理时间
		appId			应用id
		serverId		区服id
*/
func ActionLog(roleId, actId, req, res, reserve string, actStat, procTime, appId, serverId int) flumesdk.EventLog {
	newActStat := int32(actStat)
	newProcTime := int32(procTime)
	newAppId := int32(appId)
	newServerId := int32(serverId)
	log := flumesdk.NewActLog(actId, newActStat, req, res, newProcTime, newServerId, newAppId, roleId, reserve, timeutils.Now())
	return log
}

/*
	登录日志
		roleId			角色rid
		userId			用户标识id
		deviceId		设备标识id
		deviceType		设备类型
		deviceOs		设备系统
		reserve			保留字段
		ip				登入地址
		retailId		渠道标识id
		level			角色等级
		upTime			在线时长
		appId			应用id
		serverId		区服id
*/
func LoginLog(roleId, userId, deviceId, deviceType, deviceOs, reserve, ip string, retailId, level, upTime, appId, serverId int) flumesdk.EventLog {
	newRetailId := int32(retailId)
	newLv := int32(level)
	newUpTime := int32(upTime)
	newAppId := int32(appId)
	newServerId := int32(serverId)
	log := flumesdk.NewLoginLog(userId, deviceId, deviceType, deviceOs, newRetailId, newLv, newUpTime, newServerId, newAppId, ip, roleId, reserve, timeutils.Now())
	return log
}

/*
	注册日志
		roleId			角色rid
		userId			用户标识id
		deviceId		设备标识id
		deviceType		设备类型
		deviceOs		设备系统
		reserve			保留字段
		retailId		渠道标识id
		appId			应用id
		serverId		区服id
*/
func RegisterLog(roleId, userId, deviceId, deviceType, deviceOs, reserve string, retailId, appId, serverId int, ip string) flumesdk.EventLog {
	newRetailId := int32(retailId)
	newAppId := int32(appId)
	newServerId := int32(serverId)
	log := flumesdk.NewRegLog(roleId, userId, deviceId, deviceType, deviceOs, newRetailId, newServerId, newAppId, reserve, timeutils.Now(), ip)
	return log
}

/*
	数值日志
		roleId				角色id
		fieldName			重要数值字段名称
		actId				scene数值变化场景, 在此为接口号
		remark				备注
		reserve				保留字段
		old					原始值
		new					变化后的值
		appId				应用id
		serverId			区服id
*/
func VarLog(roleId, fieldName, actId, remark, reserve string, old, new, appId, serverId int) flumesdk.EventLog {
	newOld := int64(old)
	newNew := int64(new)
	newAppId := int32(appId)
	newServerId := int32(serverId)
	log := flumesdk.NewVarLog(fieldName, actId, roleId, newOld, newNew, newServerId, newAppId, remark, reserve, timeutils.Now())
	return log
}
