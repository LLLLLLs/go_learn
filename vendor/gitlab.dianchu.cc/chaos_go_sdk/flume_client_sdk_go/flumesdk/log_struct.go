package flumesdk

import (
	"strings"
	"time"
)

var timeLocation *time.Location

func init() {
	location, err := time.LoadLocation("PRC")
	if err != nil {
		panic(err)
	}
	timeLocation = location
}

type EventLog interface {
	eventLog() string
}

type coinLog struct {
	FieldName string `json:"field_name"`
	OldVal    int64  `json:"old_val"`
	NewVal    int64  `json:"new_val"`
	LifeTime  int16  `json:"lifetime"`
	Scene     string `json:"scene"`
	Remark    string `json:"remark"`
	LogTime   string `json:"log_time"`
	ServerId  int32  `json:"server_id"`
	LogDate   string `json:"log_date"`
	RoleId    string `json:"role_id"`
	CpAppId   int32  `json:"cp_app_id"`
	Reserve   string `json:"reserve"`
}

func (*coinLog) eventLog() string {
	return "coin_log"
}

func NewCoinLog(fieldName string, oldVal, newVal int64, lifeTime int16, scene, remark, roleId, reserve string, serverId, cpAppId int32, logTime int64) *coinLog {
	logTime_ := time.Unix(logTime, 0).In(timeLocation).Format("2006-01-02 15:04:05")
	logDate := strings.Split(logTime_, " ")[0]
	coin := coinLog{
		FieldName: fieldName,
		OldVal:    oldVal,
		NewVal:    newVal,
		LifeTime:  lifeTime,
		Scene:     scene,
		Remark:    remark,
		LogTime:   logTime_,
		LogDate:   logDate,
		ServerId:  serverId,
		RoleId:    roleId,
		CpAppId:   cpAppId,
		Reserve:   reserve,
	}
	return &coin
}

type actLog struct {
	ActId    string `json:"act_id"`
	ActStat  int32  `json:"act_stat"`
	Req      string `json:"req"`
	Res      string `json:"res"`
	ProcTime int32  `json:"proc_time"`
	ServerId int32  `json:"server_id"`
	RoleId   string `json:"role_id"`
	CpAppId  int32  `json:"cp_app_id"`
	Reserve  string `json:"reserve"`
	LogDate  string `json:"log_date"`
	LogTime  string `json:"log_time"`
}

func (*actLog) eventLog() string {
	return "act_log"
}

func NewActLog(actId string, actStat int32, req, res string, procTime, serverId, cpAppId int32, roleId, reserve string, logTime int64) *actLog {
	logTime_ := time.Unix(logTime, 0).In(timeLocation).Format("2006-01-02 15:04:05")
	logDate := strings.Split(logTime_, " ")[0]
	act := actLog{
		ActId:    actId,
		ActStat:  actStat,
		Req:      req,
		Res:      res,
		ProcTime: procTime,
		ServerId: serverId,
		CpAppId:  cpAppId,
		RoleId:   roleId,
		Reserve:  reserve,
		LogTime:  logTime_,
		LogDate:  logDate,
	}
	return &act
}

type loginLog struct {
	UserId     string `json:"user_id"`
	DeviceId   string `json:"device_id"`
	DeviceType string `json:"device_type"`
	DeviceOs   string `json:"device_os"`
	RetailId   int32  `json:"retail_id"`
	Lvl        int32  `json:"lvl"`
	UpTime     int32  `json:"uptime"`
	Ip         string `json:"ip"`
	ServerId   int32  `json:"server_id"`
	RoleId     string `json:"role_id"`
	CpAppId    int32  `json:"cp_app_id"`
	Reserve    string `json:"reserve"`
	LogTime    string `json:"log_time"`
	LogDate    string `json:"log_date"`
}

func (*loginLog) eventLog() string {
	return "login_log"
}

func NewLoginLog(userId, deviceId, deviceType, deviceOs string, retailId, lvl, upTime, serverId, cpAppId int32, ip, roleId, reserve string, logTime int64) *loginLog {
	logTime_ := time.Unix(logTime, 0).In(timeLocation).Format("2006-01-02 15:04:05")
	logDate := strings.Split(logTime_, " ")[0]
	login := loginLog{
		UserId:     userId,
		DeviceId:   deviceId,
		DeviceType: deviceType,
		DeviceOs:   deviceOs,
		RetailId:   retailId,
		Lvl:        lvl,
		UpTime:     upTime,
		ServerId:   serverId,
		Ip:         ip,
		RoleId:     roleId,
		CpAppId:    cpAppId,
		Reserve:    reserve,
		LogTime:    logTime_,
		LogDate:    logDate,
	}
	return &login
}

type regLog struct {
	RoleId     string `json:"role_id"`
	UserId     string `json:"user_id"`
	DeviceId   string `json:"device_id"`
	DeviceType string `json:"device_type"`
	DeviceOs   string `json:"device_os"`
	RetailId   int32  `json:"retail_id"`
	ServerId   int32  `json:"server_id"`
	CpAppId    int32  `json:"cp_app_id"`
	Reserve    string `json:"reserve"`
	LogTime    string `json:"log_time"`
	LogDate    string `json:"log_date"`
	IP         string `json:"ip"`
}

func (*regLog) eventLog() string {
	return "reg_log"
}

func NewRegLog(roleId, userId, deviceId, deviceType, deviceOs string, retailId, serverId, cpAppId int32, reserve string, logTime int64, ip string) *regLog {
	logTime_ := time.Unix(logTime, 0).In(timeLocation).Format("2006-01-02 15:04:05")
	logDate := strings.Split(logTime_, " ")[0]
	reg := regLog{
		RoleId:     roleId,
		UserId:     userId,
		DeviceId:   deviceId,
		DeviceType: deviceType,
		DeviceOs:   deviceOs,
		RetailId:   retailId,
		ServerId:   serverId,
		CpAppId:    cpAppId,
		Reserve:    reserve,
		LogTime:    logTime_,
		LogDate:    logDate,
		IP:         ip,
	}
	return &reg
}

type varLog struct {
	FieldName string `json:"field_name"`
	OldVal    int64  `json:"old_val"`
	NewVal    int64  `json:"new_val"`
	Scene     string `json:"scene"`
	Remark    string `json:"remark"`
	ServerId  int32  `json:"server_id"`
	RoleID    string `json:"role_id"`
	CpAppId   int32  `json:"cp_app_id"`
	Reserve   string `json:"reserve"`
	LogTime   string `json:"log_time"`
	LogDate   string `json:"log_date"`
}

func (*varLog) eventLog() string {
	return "var_log"
}

func NewVarLog(fieldName, scene, roleId string, oldVal, newVal int64, serverID, cpAppId int32, remark, reserve string, logTime int64) *varLog {
	logTime_ := time.Unix(logTime, 0).In(timeLocation).Format("2006-01-02 15:04:05")
	logDate := strings.Split(logTime_, " ")[0]
	Var := varLog{
		FieldName: fieldName,
		OldVal:    oldVal,
		NewVal:    newVal,
		Scene:     scene,
		RoleID:    roleId,
		ServerId:  serverID,
		CpAppId:   cpAppId,
		Remark:    remark,
		Reserve:   reserve,
		LogTime:   logTime_,
		LogDate:   logDate,
	}
	return &Var
}

type chatLog struct {
	Id          string `json:"id"`
	GameId      int32  `json:"game_id"`
	CpAppId     int32  `json:"cp_app_id"`
	ServerId    int32  `json:"server_id"`
	Channel     string `json:"channel"`
	UserId      string `json:"user_id"`
	RoleID      string `json:"role_id"`
	Alliance    string `json:"alliance"`
	DeviceId    string `json:"device_id"`
	Ext         string `json:"ext"`
	MsgType     int32  `json:"msg_type"`
	Text        string `json:"text"`
	RawImage    string `json:"raw_image"`
	Thumbnail   string `json:"thumbnail"`
	Voice       string `json:"voice"`
	VoiceLength int32  `json:"voice_length"`
	LogTime     string `json:"log_time"`
	LogDate     string `json:"log_date"`
	Reserve     string `json:"reserve"`
}

func (*chatLog) eventLog() string {
	return "chat_log"
}

func NewChatLog(id string, gameId, cpAppId, serverId int32, channel, userId, roleID, alliance, deviceId, ext string, msgType int32, text, rawImage, thumbnail, voice string, voiceLength int32, logTime int64, reserve string) *chatLog {
	logTime_ := time.Unix(logTime, 0).In(timeLocation).Format("2006-01-02 15:04:05")
	logDate := strings.Split(logTime_, " ")[0]
	Chat := chatLog{
		Id:          id,
		GameId:      gameId,
		CpAppId:     cpAppId,
		ServerId:    serverId,
		Channel:     channel,
		UserId:      userId,
		RoleID:      roleID,
		Alliance:    alliance,
		DeviceId:    deviceId,
		Ext:         ext,
		MsgType:     msgType,
		Text:        text,
		RawImage:    rawImage,
		Thumbnail:   thumbnail,
		Voice:       voice,
		VoiceLength: voiceLength,
		LogTime:     logTime_,
		LogDate:     logDate,
		Reserve:     reserve,
	}
	return &Chat
}
