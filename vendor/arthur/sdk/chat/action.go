//chat
//created: 2018/11/13
//author: wdj

package chat

import (
	pb "arthur/sdk/chat/protos"
	"arthur/utils/jsonutils"
	"arthur/utils/log"
	"fmt"
	"strconv"
	"strings"
	"time"

	"arthur/utils/http"

	"github.com/tidwall/gjson"
	"golang.org/x/net/context"
	ggrpc "google.golang.org/grpc"
)

//请求基础字段
type mustReqFields struct {
	ActID    int         `json:"act_id"`
	GameID   int         `json:"game_id"`
	AppID    int         `json:"app_id"`
	ServerID int         `json:"server_id"`
	LogFlag  int         `json:"log_flag"`
	TraceID  string      `json:"trace_id"`
	Version  string      `json:"version"`
	Data     interface{} `json:"data" `
}

type msgBody struct {
	Id        string                 `json:"_id,omitempty" ps:"消息id"`
	Channel   string                 `json:"channel,omitempty"`
	UserID    string                 `json:"user_id,omitempty"`
	RoleID    string                 `json:"role_id,omitempty"`
	Alliance  string                 `json:"alliance,omitempty" ps:"联盟ID，当频道为a必传"`
	DeviceID  string                 `json:"device_id,omitempty"`
	Ext       map[string]interface{} `json:"ext,omitempty"`
	MsgType   int                    `json:"msg_type,omitempty"`
	Text      string                 `json:"text,omitempty"`
	RawImage  string                 `json:"raw_image,omitempty" ps:"原始图像msg_type为2时"`
	Thumbnail string                 `json:"thumbnail,omitempty" ps:"缩略图msg_type为2时"`
	Voice     string                 `json:"voice,omitempty" ps:"音频msg_type为3时"`
	Length    float64                `json:"length,omitempty" ps:"语音长度msg_type为3时"`
	Time      int64                  `json:"time,omitempty"`
	ToRole    string                 `json:"to_role,omitempty"`
	Source    string                 `json:"source,omitempty"`
	SeverID   int                    `json:"server_id,omitempty"` //禁言
	Status    int                    `json:"status,omitempty"`    //禁言
}

type msgResp struct {
	Info string `json:"info,omitempty"`
	//操作结果状态 1:成功其他失败
	Stat int `json:"stat"`
	//接口代号
	ActID   int    `json:"act_id"`
	TraceID string `json:"trace_id"`
	//数据
	Data msgBody
}

type msgNumResp struct {
	Info string `json:"info,omitempty"`
	//操作结果状态 1:成功其他失败
	Stat int `json:"stat"`
	//接口代号
	ActID   int    `json:"act_id"`
	TraceID string `json:"trace_id"`
	//数据
	Data string
}

type multipleMsgResp struct {
	Info string `json:"info,omitempty"`
	//操作结果状态 1:成功其他失败
	Stat int `json:"stat"`
	//接口代号
	ActID   int    `json:"act_id"`
	TraceID string `json:"trace_id"`
	//数据
	Data []msgBody
}

type multipleMapMsgResp struct {
	Info string `json:"info,omitempty"`
	//操作结果状态 1:成功其他失败
	Stat int `json:"stat"`
	//接口代号
	ActID   int    `json:"act_id"`
	TraceID string `json:"trace_id"`
	//数据
	Data map[string][]msgBody
}

//增加一条消息
type addOneMsg400 struct {
	Channel    string                 `json:"channel" ps:"频道ID(区服频道：s系统,a联盟,p私聊,c普通,b屏蔽 跨服频道:g)"`
	UserID     string                 `json:"user_id,omitempty" ps:"	用户ID"`
	RoleID     string                 `json:"role_id,omitempty" ps:"角色ID"`
	Alliance   string                 `json:"alliance,omitempty" ps:"联盟ID，当频道为a必传"`
	DeviceID   string                 `json:"device_id,omitempty" ps:"设备ID"`
	Ext        map[string]interface{} `json:"ext,omitempty" ps:"预留字段，无扩展可不传"`
	MsgType    int                    `json:"msg_type,omitempty" ps:"消息类型默认1，1文本2图片3音频"`
	Text       string                 `json:"text,omitempty" ps:"文本msg_type为1时"`
	RawImage   string                 `json:"raw_image,omitempty" ps:"原始图像msg_type为2时"`
	Thumbnail  string                 `json:"thumbnail,omitempty" ps:"缩略图msg_type为2时"`
	Voice      string                 `json:"voice,omitempty" ps:"音频msg_type为3时"`
	Length     float64                `json:"length,omitempty" ps:"语音长度msg_type为3时"`
	Time       int64                  `json:"time,omitempty" ps:"时间统一发送时间戳，可不传自动生成"`
	ToRole     string                 `json:"to_role,omitempty" ps:"接收私聊角色Id"`
	WordFilter int                    `json:"word_filter" ps:"默认True过滤敏感词"`
	GarbFilter int                    `json:"garb_filter" ps:"默认True过滤垃圾分类"`
}

//查询一页消息
type queryOneMsg410 struct {
	Channel string `json:"channel" ps:"频道ID(区服频道：s系统,a联盟,p私聊,c普通,b屏蔽 跨服频道:g)" `
	Text    string `json:"text,omitempty" ps:"模糊查询"`
	//RoleId    string   `json:"role_id,omitempty" ps:"获取这个玩家的消息"`
	RoleIDS []string `json:"role_ids,omitempty" ps:"	匹配角色ID列表"`
	//ToRole    string   `json:"to_role,omitempty" ps:"接收私聊角色ID"`
	ToRoles   string   `json:"to_roles,omitempty" ps:"接收私聊角色ID列表"`
	Alliance  string   `json:"alliance,omitempty" ps:"联盟ID，channel为a匹配该联盟消息"`
	Size      int      `json:"size,omitempty" ps:"查询的返回条数，默认10"`
	NoRoles   []string `json:"no_roles,omitempty" ps:"不获取这些玩家的消息"`
	StartTime int64    `json:"start_time,omitempty" ps:"匹配该时间以后的"`
	EndTime   int64    `json:"end_time,omitempty" ps:"匹配该时间以前的"`
}

//查询多频道消息
type queryMultiple411 struct {
	Channels  []string `json:"channels" ps:"频道ID列表(区服频道：s系统,a联盟,p私聊,c普通,b屏蔽 跨服频道:g)"`
	Text      string   `json:"text,omitempty" ps:"模糊查询"`
	RoleID    string   `json:"role_id,omitempty" ps:"获取这个玩家的消息"`
	ToRole    string   `json:"to_role,omitempty" ps:"接收私聊角色ID"`
	Alliance  string   `json:"alliance,omitempty" ps:"联盟ID，channel为a匹配该联盟消息"`
	Size      int      `json:"size,omitempty" ps:"查询的返回条数，默认10"`
	NoRoles   []string `json:"no_roles,omitempty" ps:"不获取这些玩家的消息"`
	StartTime int64    `json:"start_time,omitempty" ps:"匹配该时间以后的"`
	EndTime   int64    `json:"End_time,omitempty" ps:"匹配该时间以前的"`
}

//查询多频道消息（不合并）
type queryMultiple412 struct {
	Channels  []string `json:"channels" ps:"频道ID(区服频道：s系统,a联盟,p私聊,c普通,b屏蔽 跨服频道:g)" `
	RoleID    string   `json:"role_id,omitempty" ps:"获取这个玩家的消息"`
	Size      int      `json:"size,omitempty" ps:"查询的返回条数，默认10"`
	ToRole    string   `json:"to_role,omitempty" ps:"接收私聊角色ID"`
	ToRoles   []string `json:"to_roles,omitempty" ps:"接收私聊角色ID列表"`
	Alliance  string   `json:"alliance,omitempty" ps:"联盟ID，当频道为a必传"`
	NoRoles   []string `json:"no_roles,omitempty" ps:"不获取这些玩家的消息"`
	StartTime int64    `json:"start_time,omitempty" ps:"匹配该时间以后的"`
	EndTime   int64    `json:"End_time,omitempty" ps:"匹配该时间以前的"`
}

//413
// 获取完整的消息——一般获取音频、图片消息需要调用
type queryByID413 struct {
	Channel  string `json:"channel" ps:"频道ID(区服频道：s系统,a联盟,p私聊,c普通,b屏蔽 跨服频道:g)"`
	ID       string `json:"id,omitempty" ps:"信息Id"`
	Alliance string `json:"alliance,omitempty" ps:"联盟ID，当频道为a必传"`
}

//翻译消息
type transMsg414 struct {
	Channel  string `json:"channel" ps:"频道ID(区服频道：s系统,a联盟,p私聊,c普通,b屏蔽 跨服频道:g)"`
	Language string `json:"language" ps:"翻译目标语言"`
	ID       string `json:"id,omitempty" ps:"信息Id, 如果无text，channel和id是必填的"`
	Alliance string `json:"alliance,omitempty" ps:"联盟id（如过channel是g或者a ， 则必填）"`
	Text     string `json:"text,omitempty" ps:"翻译文本（如果有text，channel和id不是必填的，优先翻译text内容返回）"`
}

//删除匹配消息
type delMsg420 struct {
	Channel   string   `json:"channel" ps:"频道ID(区服频道：s系统,a联盟,p私聊,c普通,b屏蔽 跨服频道:g)" `
	Alliance  string   `json:"alliance,omitempty" ps:"联盟ID，当频道为a必传"`
	RoleID    string   `json:"role_id,omitempty" ps:"获取这个玩家的消息"`
	ID        string   `json:"id,omitempty" ps:"消息ID"`
	IDS       []string `json:"ids,omitempty" ps:"消息id列表"`
	StartTime int64    `json:"start_time,omitempty" ps:"匹配该时间以后的"`
	EndTime   int64    `json:"End_time,omitempty" ps:"匹配该时间以前的"`
}

var (
	//TODO：连接池配置参数后续需要配置在ZK或其他地方
	client = http.NewNetClient(2000, 100, 20*time.Second, 5*time.Second)
)

//添加禁言
type addBannedRole430 struct {
	RoleID     string `json:"role_id" ps:"角色ID"`
	Status     int    `json:"status,omitempty" ps:"禁言类型"`
	ExpiryTime int64  `json:"expiry_time,omitempty" ps:"解禁时间"`
	Text       string `json:"text,omitempty" ps:"禁言记录文本"`
}

//查询一页禁言，返回用户和被标记状态、时间、次数（不是频发广告被禁的无次数）
type queryBannedRole431 struct {
	ServerID  int    `json:"server_id" ps:"区服ID"`
	Page      int    `json:"page,omitempty" ps:"页数"`
	Size      int    `json:"size,omitempty" ps:"返回一页的数量"`
	Status    int    `json:"status,omitempty" ps:"匹配状态"`
	RoleID    string `json:"role_id,omitempty" ps:"角色ID"`
	StartTime int64  `json:"start_time,omitempty" ps:"匹配改时间以后"`
	EndTime   int64  `json:"end_time,omitempty" ps:"匹配改时间以前"`
}

//取消禁言
type delBannedRole432 struct {
	RoleID string `json:"role_id" ps:"角色ID"`
}

//查询禁言用户后序发言
type queryAfterBanned433 struct {
	ServerID  int    `json:"server_id" ps:"区服ID"`
	Page      int    `json:"page,omitempty" ps:"页数"`
	Size      int    `json:"size,omitempty" ps:"返回一页的数量"`
	Status    int    `json:"status,omitempty" ps:"匹配状态"`
	RoleID    string `json:"role_id,omitempty" ps:"角色ID"`
	StartTime int64  `json:"start_time,omitempty" ps:"匹配改时间以后"`
	EndTime   int64  `json:"end_time,omitempty" ps:"匹配改时间以前"`
}

type queryChat struct {
	conf     *ChatsConfManager
	gameID   int //游戏id
	appID    int
	serverID int
	logFlag  int
}

func newQueryChat(gameID int, appID int, serverID int, logFlag int) queryChat {
	qc := new(queryChat)
	qc.gameID = gameID
	qc.appID = appID
	qc.serverID = serverID
	qc.logFlag = logFlag
	qc.conf = chatConf
	return *qc
}

//添加文本信息
func (qc queryChat) addOneTextMsg(traceID string, condition map[string]interface{}) (msgBody, error, int) {
	data := new(addOneMsg400)
	data.MsgType = TextMsg
	data.Time = time.Now().UnixNano()
	//默认过滤
	data.WordFilter = 0
	data.GarbFilter = 0
	channel, ok := condition["channel"]
	if !ok {
		return msgBody{}, fmt.Errorf("channel required"), NOT_CHAT_CODE
	}
	data.Channel = channel.(string)
	roleID, ok := condition["roleID"]
	if ok {
		data.RoleID = roleID.(string)
	}
	userID, ok := condition["userID"]
	if ok {
		data.UserID = userID.(string)
	}
	deviceID, ok := condition["deviceID"]
	if ok {
		data.DeviceID = deviceID.(string)
	}
	allianceID, ok := condition["alliance"]
	if ok {
		data.Alliance = allianceID.(string)
	}
	text, ok := condition["text"]
	if ok {
		data.Text = text.(string)
	}
	wordFilter, ok := condition["wordFilter"]
	if ok {
		wf := wordFilter.(bool)
		if !wf {
			//不过滤
			data.WordFilter = 1
		}
	}
	garbFilter, ok := condition["garbFilter"]
	if ok {
		gf := garbFilter.(bool)
		if !gf {
			//不过滤
			data.GarbFilter = 1
		}
	}
	ext, ok := condition["ext"]
	if ok {
		data.Ext = ext.(map[string]interface{})
	}
	toRole, ok := condition["toRole"]
	if ok {
		data.ToRole = toRole.(string)
	}
	return qc.getMsgBodyResp(400, traceID, data)
}

//查询一页消息, 只有roleIDS可以用于筛选,roleID不用
func (qc queryChat) queryMsgWithCondition(traceID string, condition map[string]interface{}) ([]msgBody, error, int) {
	data := new(queryOneMsg410)
	channel, ok := condition["channel"]
	if !ok {
		return nil, fmt.Errorf("channel required"), NOT_CHAT_CODE
	}
	data.Channel = channel.(string)
	data.Size = qc.conf.pageLimit
	//msgID, ok := condition["id"]
	//if ok {
	//	data.RoleIDS
	//}
	text, ok := condition["text"]
	if ok {
		data.Text = text.(string)
	}
	roleIDS, ok := condition["roleIDS"]
	if ok {
		data.RoleIDS = roleIDS.([]string)
	}
	allianceID, ok := condition["alliance"]
	if ok {
		data.Alliance = allianceID.(string)
	}
	size, ok := condition["size"]
	if ok {
		data.Size = size.(int)
	}
	noRoles, ok := condition["noRoles"]
	if ok {
		data.NoRoles = noRoles.([]string)
	}
	startTime, ok := condition["startTime"]
	if ok {
		data.StartTime = startTime.(int64)
	}
	endTime, ok := condition["endTime"]
	if ok {
		data.EndTime = endTime.(int64)
	}
	return qc.getMulMsgBodyResp(410, traceID, data)
}

//查询多频道消息,获取各个频道的数据一页，合并所有数据排序再得到一页返回
func (qc queryChat) queryMulChannel(traceID string, condition map[string]interface{}) ([]msgBody, error, int) {
	data := new(queryMultiple411)
	channels, ok := condition["channels"]
	if !ok {
		return nil, fmt.Errorf("channels required"), NOT_CHAT_CODE
	}
	data.Channels = channels.([]string)
	data.Size = qc.conf.pageLimit
	text, ok := condition["text"]
	if ok {
		data.Text = text.(string)
	}
	roleID, ok := condition["roleID"]
	if ok {
		data.RoleID = roleID.(string)
	}
	allianceID, ok := condition["alliance"]
	if ok {
		data.Alliance = allianceID.(string)
	}
	size, ok := condition["size"]
	if ok {
		data.Size = size.(int)
	}
	noRoles, ok := condition["noRoles"]
	if ok {
		data.NoRoles = noRoles.([]string)
	}
	startTime, ok := condition["startTime"]
	if ok {
		data.StartTime = startTime.(int64)
	}
	endTime, ok := condition["endTime"]
	if ok {
		data.EndTime = endTime.(int64)
	}
	return qc.getMulMsgBodyResp(411, traceID, data)
}

//查询多频道消息（不合并）
func (qc queryChat) queryMulChannelIndependent(traceID string, condition map[string]interface{}) (map[string][]msgBody, error, int) {
	data := new(queryMultiple412)
	channels, ok := condition["channels"]
	if !ok {
		return nil, fmt.Errorf("channels required"), NOT_CHAT_CODE
	}
	data.Channels = channels.([]string)
	roleID, ok := condition["roleID"]
	if ok {
		data.RoleID = roleID.(string)
	}
	allianceID, ok := condition["alliance"]
	if ok {
		data.Alliance = allianceID.(string)
	}
	noRoles, ok := condition["noRoles"]
	if ok {
		data.NoRoles = noRoles.([]string)
	}
	startTime, ok := condition["startTime"]
	if ok {
		data.StartTime = startTime.(int64)
	}
	return qc.getMulMapMsgBodyResp(412, traceID, data)
}

func (qc queryChat) queryByID(traceID string, condition map[string]interface{}) (msgBody, error, int) {
	data := new(queryByID413)
	id, ok := condition["id"]
	if !ok {
		return msgBody{}, fmt.Errorf("id required"), NOT_CHAT_CODE
	}
	data.ID = id.(string)
	channel, ok := condition["channel"]
	if !ok {
		return msgBody{}, fmt.Errorf("channel required"), NOT_CHAT_CODE
	}
	data.Channel = channel.(string)
	allianceID, ok := condition["alliance"]
	if ok {
		data.Alliance = allianceID.(string)
	}
	return qc.getMsgBodyResp(413, traceID, data)
}

//翻译消息
func (qc queryChat) translate(traceID string, condition map[string]interface{}) (msgBody, error, int) {
	data := new(transMsg414)
	channel, ok := condition["channel"]
	if !ok {
		return msgBody{}, fmt.Errorf("channel required"), NOT_CHAT_CODE
	}
	data.Channel = channel.(string)
	lang, ok := condition["language"]
	if !ok {
		return msgBody{}, fmt.Errorf("language required"), NOT_CHAT_CODE
	}
	data.Language = lang.(string)
	text, ok := condition["text"]
	if ok {
		data.Text = text.(string)
	}
	allianceID, ok := condition["alliance"]
	if ok {
		data.Alliance = allianceID.(string)
	}
	id, ok := condition["id"]
	if ok {
		data.ID = id.(string)
	}
	return qc.getMsgBodyResp(414, traceID, data)
}

//删除匹配消息
func (qc queryChat) delMsg(traceID string, condition map[string]interface{}) (bool, error, int) {
	data := new(delMsg420)
	data.Channel = condition["channel"].(string)
	id, ok := condition["ID"]
	if ok {
		data.ID = id.(string)
	}
	idx, ok := condition["IDS"]
	if ok {
		data.IDS = idx.([]string)
	}
	roleID, ok := condition["roleID"]
	if ok {
		data.RoleID = roleID.(string)
	}
	allianceID, ok := condition["alliance"]
	if ok {
		data.Alliance = allianceID.(string)
	}
	startTime, ok := condition["startTime"]
	if ok {
		data.StartTime = startTime.(int64)
	}
	endTime, ok := condition["endTime"]
	if ok {
		data.EndTime = endTime.(int64)
	}
	_, err, code := qc.getMsgNumResp(420, traceID, data)
	if err != nil {
		return false, err, code
	}
	return true, nil, code
}

//查询一页禁言记录
func (qc queryChat) queryBannedMsg(traceID string, condition map[string]interface{}) ([]msgBody, error, int) {
	data := new(queryBannedRole431)
	server, ok := condition["serverID"]
	if !ok {
		return nil, fmt.Errorf("serverID required"), NOT_CHAT_CODE
	}
	data.ServerID = server.(int)
	data.Size = qc.conf.pageLimit
	size, ok := condition["size"]
	if ok {
		data.Size = size.(int)
	}
	status, ok := condition["status"]
	if ok {
		data.Status = status.(int)
	}
	page, ok := condition["page"]
	if ok {
		data.Page = page.(int)
	}
	roleID, ok := condition["roleID"]
	if ok {
		data.RoleID = roleID.(string)
	}
	startTime, ok := condition["startTime"]
	if ok {
		data.StartTime = startTime.(int64)
	}
	endTime, ok := condition["endTime"]
	if ok {
		data.EndTime = endTime.(int64)
	}
	return qc.getMulMsgBodyResp(431, traceID, data)
}

//查询禁言用户后序发言
func (qc queryChat) queryAfterBannedMsg(traceID string, condition map[string]interface{}) ([]msgBody, error, int) {
	data := new(queryAfterBanned433)
	server, ok := condition["serverID"]
	if !ok {
		return nil, fmt.Errorf("serverID required"), NOT_CHAT_CODE
	}
	data.ServerID = server.(int)
	size, ok := condition["size"]
	if ok {
		data.Size = size.(int)
	}
	status, ok := condition["status"]
	if ok {
		data.Status = status.(int)
	}
	page, ok := condition["page"]
	if ok {
		data.Page = page.(int)
	}
	roleID, ok := condition["roleID"]
	if ok {
		data.RoleID = roleID.(string)
	}
	startTime, ok := condition["startTime"]
	if ok {
		data.StartTime = startTime.(int64)
	}
	endTime, ok := condition["endTime"]
	if ok {
		data.EndTime = endTime.(int64)
	}
	return qc.getMulMsgBodyWithEmptyResp(433, traceID, data)
}

//添加禁言
func (qc queryChat) addBannedMsg(traceID string, condition map[string]interface{}) (bool, error, int) {
	data := new(addBannedRole430)
	text, ok := condition["text"]
	if ok {
		//发什么文本导致禁言
		data.Text = text.(string)
	}
	status, ok := condition["status"]
	if ok {
		data.Status = status.(int)
	}
	expire, ok := condition["expire"]
	if ok {
		data.ExpiryTime = expire.(int64)
	}
	roleID, ok := condition["roleID"]
	if ok {
		data.RoleID = roleID.(string)
	}
	_, err, code := qc.getMsgBodyResp(430, traceID, data)
	if err != nil {
		return false, err, code
	}
	return true, nil, code
}

//取消禁言
func (qc queryChat) delBannedMsg(traceID string, roleID string) (bool, error, int) {
	data := new(delBannedRole432)
	data.RoleID = roleID
	_, err, code := qc.getMsgBodyResp(432, traceID, data)
	if err != nil {
		return false, err, code
	}
	return true, nil, code
}

//get response

//返回信息为int
func (qc queryChat) getMsgNumResp(actID int, traceID string, data interface{}) (msg int, err error, code int) {
	code = NOT_CHAT_CODE
	req, err := qc.buildReq(actID, traceID, data)
	if err != nil {
		return 0, err, code
	}

	rawData, err := qc.post(req)
	if err != nil {
		return 0, err, code
	}
	err, code = qc.checkRespError(rawData)
	if err != nil {
		return 0, err, code
	}
	resp := new(msgNumResp)
	log.Debugf("%q\n", rawData)
	err = json.Unmarshal(rawData, resp)
	if err != nil {
		return 0, err, code
	}
	msg, err = strconv.Atoi(resp.Data)
	return
}

func (qc queryChat) getMsgBodyResp(actID int, traceID string, data interface{}) (msgBody, error, int) {
	req, err := qc.buildReq(actID, traceID, data)
	if err != nil {
		return msgBody{}, err, NOT_CHAT_CODE
	}

	rawData, err := qc.post(req)
	if err != nil {
		return msgBody{}, err, NOT_CHAT_CODE
	}
	err, code := qc.checkRespError(rawData)
	if err != nil {
		return msgBody{}, err, code
	}
	resp := new(msgResp)
	log.Debugf("%q\n", rawData)
	err = json.Unmarshal(rawData, resp)
	return resp.Data, err, code
}

func (qc queryChat) getMulMsgBodyResp(actID int, traceID string, data interface{}) ([]msgBody, error, int) {

	req, err := qc.buildReq(actID, traceID, data)
	if err != nil {
		return nil, err, NOT_CHAT_CODE
	}
	rawData, err := qc.post(req)
	if err != nil {
		return nil, err, NOT_CHAT_CODE
	}
	err, code := qc.checkRespError(rawData)
	if err != nil {
		return nil, err, code
	}
	log.Debugf("%q\n", rawData)
	resp := new(multipleMsgResp)
	err = json.Unmarshal(rawData, resp)
	return resp.Data, err, code
}

//存在空记录
func (qc queryChat) getMulMsgBodyWithEmptyResp(actID int, traceID string, data interface{}) ([]msgBody, error, int) {
	req, err := qc.buildReq(actID, traceID, data)
	if err != nil {
		return nil, err, NOT_CHAT_CODE
	}
	rawData, err := qc.post(req)
	if err != nil {
		return nil, err, NOT_CHAT_CODE
	}
	err, code := qc.checkRespError(rawData)
	if err != nil {
		return nil, err, code
	}
	log.Debugf("%q\n", rawData)
	x := make(map[string]interface{})
	err = json.Unmarshal(rawData, &x)
	if err != nil {
		return nil, err, code
	}
	mbs := make([]msgBody, 0)
	for _, msg := range x["data"].([]interface{}) {
		if msg == nil {
			continue
		}
		mapMsg := msg.(map[string]interface{})
		text := mapMsg["text"]
		if text == nil {
			mb := toMsgBody(mapMsg)
			mbs = append(mbs, mb)
			continue
		}
		//多条记录
		mulText, ok := text.([]interface{})
		if !ok {
			//单条记录
			mb := toMsgBody(mapMsg)
			mbs = append(mbs, mb)
			continue
		}
		for i := range mulText {
			mapMsg["text"] = mulText[i]
			mb := toMsgBody(mapMsg)
			mbs = append(mbs, mb)

		}
	}
	return mbs, nil, code
}

func (qc queryChat) getMulMapMsgBodyResp(actID int, traceID string, data interface{}) (map[string][]msgBody, error, int) {
	req, err := qc.buildReq(actID, traceID, data)
	if err != nil {
		return nil, err, NOT_CHAT_CODE
	}
	rawData, err := qc.post(req)
	if err != nil {
		return nil, err, NOT_CHAT_CODE
	}
	err, code := qc.checkRespError(rawData)
	if err != nil {
		return nil, err, code
	}
	log.Debugf("%q\n", rawData)
	resp := new(multipleMapMsgResp)
	err = json.Unmarshal(rawData, resp)
	return resp.Data, err, code
}

//请求参数
func (qc queryChat) buildReq(actID int, traceID string, data interface{}) ([]byte, error) {
	req := new(mustReqFields)
	req.ActID = actID
	req.AppID = qc.appID
	req.GameID = qc.gameID
	req.LogFlag = qc.logFlag
	req.ServerID = qc.serverID
	req.TraceID = traceID
	req.Version = qc.conf.version
	req.Data = data
	reqDataByte, err := json.MarshalWithErr(req)
	if err != nil {
		return nil, err
	}
	return reqDataByte, nil
}

//发送请求
func (qc queryChat) post(data []byte) ([]byte, error) {
	log.Debug(string(data))
	if strings.ToLower(qc.conf.method) == "http" {
		code, resp, err := client.PostJson(qc.conf.url, data)
		return http.ToResp(code, resp, err)
	} else {
		timeout := time.Duration(qc.conf.timeout) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		client, err := ggrpc.DialContext(ctx, qc.conf.url, ggrpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		conn := pb.NewActionClient(client)
		res, err := conn.ExecAction(ctx, &pb.ActionRequest{Action: string(data)})
		if err != nil {
			return nil, err
		}
		return []byte(res.Message), nil
	}
}

//检查错误， 状态码
func (qc queryChat) checkRespError(raw []byte) (err error, code int) {
	code = int(gjson.GetBytes(raw, "stat").Int())
	if code != SUCCESS {
		info := gjson.GetBytes(raw, "info").String()
		return fmt.Errorf("chat fail with %s, status == %d", info, code), code
	}

	return nil, code
}

func toMsgBody(info map[string]interface{}) msgBody {
	mb := new(msgBody)
	ID, ok := info["_id"]
	//"消息id"
	if ok {

		mb.Id = ID.(string)
	}
	channel, ok := info["channel"]
	if ok {
		mb.Channel = channel.(string)
	}
	uid, ok := info["user_id"]
	if ok {
		mb.UserID = uid.(string)
	}
	rid, ok := info["role_id"]
	if ok {
		mb.RoleID = rid.(string)
	}
	allianceID, ok := info["alliance"]
	if ok {
		mb.Alliance = allianceID.(string)
	}
	did, ok := info["device_id"]
	if ok {
		mb.DeviceID = did.(string)
	}
	ext, ok := info["ext"]
	if ok {
		mb.Ext = ext.(map[string]interface{})
	}
	mt, ok := info["msg_type"]
	if ok {
		mb.MsgType = int(mt.(float64))
	}
	text, ok := info["text"]
	if ok {
		mb.Text = text.(string)
	}
	rawImg, ok := info["raw_image"]
	if ok {
		mb.RawImage = rawImg.(string)
	}
	tb, ok := info["thumbnail"]
	if ok {
		mb.Thumbnail = tb.(string)
	}
	voice, ok := info["voice"]
	if ok {
		mb.Voice = voice.(string)
	}
	l, ok := info["length"]
	if ok {
		mb.Length = l.(float64)
	}
	t, ok := info["time"]
	if ok {
		mb.Time = int64(t.(float64))
	}
	tr, ok := info["to_role"]
	if ok {
		mb.ToRole = tr.(string)
	}
	src, ok := info["source"]
	if ok {
		mb.Source = src.(string)
	}
	sid, ok := info["server_id"]
	if ok {
		mb.SeverID = int(sid.(float64))
	}
	stat, ok := info["status"]
	if ok {
		mb.Status = int(stat.(float64))
	}
	return *mb
}
