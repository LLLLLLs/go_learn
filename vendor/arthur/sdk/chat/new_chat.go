//chat
//created: 2018/11/13
//author: wdj

package chat

import (
	"fmt"
	"strings"
)

//频道
const ChannelGlobal = "g"   //跨服
const ChannelSystem = "s"   //系统
const ChannelAlliance = "a" //联盟
const ChannelPrivate = "p"  //私聊
const ChannelCommon = "c"   //普通
const ChannelShield = "b"   //屏蔽
//1文本2图片3音频
const TextMsg = 1
const PictureMsg = 2
const VoiceMsg = 3

//shield type
//1客服禁言 2消息分类禁言 3信用评级禁言 4频发广告聊天禁言 5黑名单文本直接禁言 6举报调用 7遗漏调用
const BannedKF = 1
const BannedFilter = 2
const BannedCredit = 3
const BannedAD = 4
const BannedBlack = 5
const BannedReport = 6
const BannedLoss = 7

var chatConf *ChatsConfManager

type ChatsConfManager struct {
	url       string
	method    string
	pageLimit int
	version   string
	timeout   int
	isPush    bool
}

func Init(timeout, pageLimit int, method, grpcAddr, httpAddr, version string, isPush bool) {
	var url string
	if strings.ToLower(method) == "http" {
		url = httpAddr
	} else {
		url = grpcAddr
	}
	c := &ChatsConfManager{
		url:       url,
		method:    method,
		pageLimit: pageLimit,
		version:   version,
		timeout:   timeout,
		isPush:    isPush,
	}
	chatConf = c
}

type QueryChat struct {
	conf *ChatsConfManager
	q    queryChat
}

//文本聊天信息
type TextInfo struct {
	ID         string
	AllianceID string
	Ext        map[string]interface{}
	Text       string
	Channel    string
	ToRole     string //私聊
	Time       int64
	RoleID     string //发言人
	UserID     string
	DeviceID   string
	MsgType    int
	Source     string //翻译来源
}

func NewQueryChat(gameID int, appID int, serverID int, logFlag int) QueryChat {
	qc := new(QueryChat)
	q := newQueryChat(gameID, appID, serverID, logFlag)
	qc.conf = chatConf
	qc.q = q
	return *qc
}

//添加非联盟非私聊文本信息, 返回新增文本id
func (qc QueryChat) AddTextMsg(traceID string, channel string, roleID string, userID string, deviceID string, text string, ext map[string]interface{}, wordFilter bool, garbFilter bool) (TextInfo, error, int) {
	if channel == ChannelAlliance {
		return TextInfo{}, fmt.Errorf("channel can not be alliance"), NOT_CHAT_CODE
	}
	condition := map[string]interface{}{}
	condition["channel"] = channel
	condition["roleID"] = roleID
	condition["deviceID"] = deviceID
	condition["text"] = text
	condition["wordFilter"] = wordFilter
	condition["garbFilter"] = garbFilter
	condition["userID"] = userID
	condition["ext"] = ext
	resp, err, code := qc.q.addOneTextMsg(traceID, condition)
	if err != nil {
		return TextInfo{}, err, code
	}
	return qc.getTextInfo(resp), nil, code
}

//添加联盟非私聊文本信息, 返回新增文本id
func (qc QueryChat) AddAllianceTextMsg(traceID string, allianceID string, roleID string, deviceID string, text string, wordFilter bool, garbFilter bool) (TextInfo, error, int) {
	condition := map[string]interface{}{}
	condition["channel"] = ChannelAlliance
	condition["roleID"] = roleID
	condition["deviceID"] = deviceID
	condition["text"] = text
	condition["wordFilter"] = wordFilter
	condition["garbFilter"] = garbFilter
	condition["allianceID"] = allianceID
	resp, err, code := qc.q.addOneTextMsg(traceID, condition)
	if err != nil {
		return TextInfo{}, err, code
	}
	return qc.getTextInfo(resp), nil, code
}

//查询最新pageLimit条消息
func (qc QueryChat) GetLastTextMsg(traceID string, channel string, size int) ([]TextInfo, error, int) {
	condition := map[string]interface{}{}
	condition["channel"] = channel
	condition["size"] = size
	resp, err, code := qc.q.queryMsgWithCondition(traceID, condition)
	if err != nil {
		return nil, err, code
	}
	return qc.getTextInfoList(resp), nil, code
}

//排除角色查询消息
func (qc QueryChat) GetMsgExceptRole(traceID string, channel string, noRoles []string, size int, beforeTime int64) ([]TextInfo, error, int) {
	condition := map[string]interface{}{}
	condition["channel"] = channel
	condition["noRoles"] = noRoles
	condition["size"] = size
	condition["endTime"] = beforeTime
	resp, err, code := qc.q.queryMsgWithCondition(traceID, condition)
	if err != nil {
		return nil, err, code
	}
	return qc.getTextInfoList(resp), nil, code
}

//查询多频道消息（不合并）
func (qc QueryChat) GetLastMultipleChannelMsg(traceID string, channelList []string, noRoles []string, startTime int64) (map[string][]TextInfo, error, int) {
	condition := map[string]interface{}{}
	condition["channels"] = channelList
	condition["noRoles"] = noRoles
	condition["startTime"] = startTime
	resp, err, code := qc.q.queryMulChannelIndependent(traceID, condition)
	if err != nil {
		return nil, err, code
	}
	return qc.getMapTextInfo(resp), nil, code
}

func (qc QueryChat) GetMsgByID(traceID string, channel string, msgID string) (TextInfo, error, int) {
	condition := map[string]interface{}{}
	condition["channel"] = channel
	condition["id"] = msgID
	resp, err, code := qc.q.queryByID(traceID, condition)
	if err != nil {
		return TextInfo{}, err, code
	}
	info := qc.getTextInfo(resp)
	return info, nil, code

}

func (qc QueryChat) getMapTextInfo(resp map[string][]msgBody) map[string][]TextInfo {
	textInfo := map[string][]TextInfo{}
	for k, v := range resp {
		textInfoList := make([]TextInfo, len(v))
		for i := range v {
			textInfoList[i] = qc.getTextInfo(v[i])
		}
		textInfo[k] = textInfoList
	}
	return textInfo
}

func (qc QueryChat) getTextInfoList(resp []msgBody) []TextInfo {
	textInfo := make([]TextInfo, len(resp))
	for i := range resp {
		textInfo[i] = qc.getTextInfo(resp[i])
	}
	return textInfo
}

func (qc QueryChat) getTextInfo(resp msgBody) TextInfo {
	info := new(TextInfo)
	info.Text = resp.Text
	info.ID = resp.Id
	info.Channel = resp.Channel
	info.Ext = resp.Ext
	info.ToRole = resp.ToRole
	info.AllianceID = resp.Alliance
	info.Time = resp.Time
	info.RoleID = resp.RoleID
	info.UserID = resp.UserID
	info.DeviceID = resp.DeviceID
	info.MsgType = resp.MsgType
	info.Source = resp.Source
	return *info
}

//添加禁言
func (qc QueryChat) AddBannedRole(traceID string, shieldType int, expire int64, text string, roleID string) (error, int) {
	condition := map[string]interface{}{}
	condition["expire"] = expire
	condition["status"] = shieldType
	condition["text"] = text
	condition["roleID"] = roleID
	_, err, code := qc.q.addBannedMsg(traceID, condition)
	return err, code
}

//取消禁言
func (qc QueryChat) DelBannedRole(traceID string, roleID string) (bool, error, int) {
	return qc.q.delBannedMsg(traceID, roleID)
}

////查询本服禁言记录
//func (qc QueryChat) QuerySelfShieldMsg() {
//	condition := map[string]interface{}{}
//	condition["serverID"] = qc.q.serverID
//}

//翻译消息
func (qc QueryChat) TranslateText(traceID string, channel string, text string, lang string) (string, error, int) {
	condition := map[string]interface{}{}
	condition["channel"] = channel
	condition["text"] = text
	condition["language"] = lang
	msg, err, code := qc.q.translate(traceID, condition)
	if err != nil {
		return "", err, code
	}
	return msg.Text, nil, code
}

func (qc QueryChat) TranslateTextByID(traceID string, channel string, msgID string, lang string) (TextInfo, error, int) {
	condition := map[string]interface{}{}
	condition["channel"] = channel
	condition["id"] = msgID
	condition["language"] = lang
	msg, err, code := qc.q.translate(traceID, condition)
	if err != nil {
		return TextInfo{}, err, code
	}
	return qc.getTextInfo(msg), nil, code
}
