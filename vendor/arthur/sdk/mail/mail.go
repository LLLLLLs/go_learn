/*
Created on 2018/6/21

author: WenHao Shan

Content:
*/

package mail

import (
	pb "arthur/sdk/mail/protos"
	"arthur/utils/errors"
	"arthur/utils/jsonutils"
	"arthur/utils/timeutils"
	"context"
	"strings"
	"time"

	"log"

	"arthur/utils/http"
	"fmt"
	"github.com/tidwall/gjson"
	ggrpc "google.golang.org/grpc"
)

var (
	StatusNormal = 1 // 正常邮件
	StatusShield = 2 // 屏蔽邮件
	//TODO：连接池配置参数后续需要配置在ZK或其他地方
	client = http.NewNetClient(2000, 100, 20*time.Second, 5*time.Second)
)

type MailConfManager struct {
	url     string
	method  string
	version string
	timeout int
	isPush  bool
}

var mailConf *MailConfManager

func InitMail(timeout int, method string, grpcAddr string, httpAddr string, version string, isPush bool) {
	var url string
	if strings.ToLower(method) == "http" {
		url = httpAddr
	} else {
		url = grpcAddr
	}
	conf := &MailConfManager{
		url:     url,
		method:  method,
		version: version,
		timeout: timeout,
		isPush:  isPush,
	}
	mailConf = conf
}

// 请求基础字段
type BaseReqFields struct {
	ActId    int    `json:"act_id"`
	GameId   int    `json:"game_id"`
	AppId    int    `json:"app_id"`
	ServerId int    `json:"server_id"`
	LogFlag  int    `json:"log_flag"`
	TraceId  string `json:"trace_id"`
	Version  string `json:"version"`
}

// 邮件接口对象
type Manager struct {
	BaseReqFields
	Data interface{} `json:"data"`
}

// 请求响应数据字段
type RespDataFields struct {
	Id        string  `json:"_id" ps:"消息id"`
	Time      float64 `json:"time" ps:"发件时间"`
	IsRead    int     `json:"is_read" ps:"是否已读"`
	IsReceive int     `json:"is_receive" ps:"是否已领取"`
	ToRole    string  `json:"to_role" ps:"收件人id"`
	Status    int     `json:"status" ps:"状态id"`
	CelId     string  `json:"cel_id" ps:"奖励id"`
	SenderId  string  `json:"sender_id" ps:"发送者id"`
}

// 请求响应数据(查询未读未领取数量)
type RespNrNoRFields struct {
	NoRead    int `json:"no_read" ps:"未读数量"`
	NoReceive int `json:"no_receive" ps:"未领取数量"`
}

func New(gameId int, appId int, serverId int, logFlag int) *Manager {
	c := new(Manager)
	c.GameId = gameId
	c.AppId = appId
	c.ServerId = serverId
	c.LogFlag = logFlag
	c.Version = mailConf.version
	return c
}

type SendMailReqData struct {
	SenderId   string      `json:"sender_id" ps:"发件人id"`
	DeviceId   string      `json:"sender_device" ps:"发件人设备号(系统邮件时可不填)"`
	Ip         string      `json:"sender_ip" ps:"发件人ip(系统邮件时可不填)"`
	ToRole     string      `json:"to_role" ps:"收件人id"`
	ToRoles    []string    `json:"to_roles" ps:"收件人列表"`
	Title      string      `json:"title" ps:"邮件标题"`
	Content    string      `json:"content" ps:"邮件内容"`
	Award      string      `json:"attachments" ps:"附加奖励, 暂定奖励格式为[[ValNo,No,Count]]"`
	CelId      string      `json:"cel_id" ps:"邮件奖励id"`
	Status     int         `json:"status" ps:"邮件自定义类型"`
	Ext        interface{} `json:"ext" ps:"预留字段"`
	IsRead     int         `json:"is_read" ps:"是否已读, 只用于测试, 默认都是未读"`
	ExpiryTime int64       `json:"expiry_time" ps:"邮件过期时间"`
}

// 发送系统邮件
// toRole指个人, toRoles指多人id列表, 二者只可选一个
// isNormal 为True时表示是正常邮件, 为False时表示屏蔽邮件
func (req *Manager) SysSendMail(celId string, expiryTime int64, traceId string,
	toRole string, toRoles []string, award string, isNormal bool) error {
	sendMailData := SendMailReqData{
		SenderId:   "system",
		CelId:      celId,
		ExpiryTime: expiryTime,
	}
	if toRole != "" {
		sendMailData.ToRole = toRole
	} else {
		sendMailData.ToRoles = toRoles
	}

	// TODO 确定奖励格式
	if len(award) != 0 {
		sendMailData.Award = award
	}

	if isNormal {
		sendMailData.Status = StatusNormal
	} else {
		sendMailData.Status = StatusShield
	}
	err := req.sendMail(traceId, sendMailData)
	return err
}

// 玩家发送邮件, 只能p2p发送
func (req *Manager) PSendMail(sendId string, deviceId string, ip string, celId string, expiryTime int64,
	traceId string, title string, content string, toRole string, award string) error {
	sendMailData := SendMailReqData{
		DeviceId:   deviceId,
		Ip:         ip,
		SenderId:   sendId,
		ToRole:     toRole,
		Title:      title,
		Content:    content,
		CelId:      celId,
		Status:     StatusNormal,
		ExpiryTime: expiryTime,
	}

	// TODO 确定奖励格式
	if len(award) != 0 {
		sendMailData.Award = award
	}

	err := req.sendMail(traceId, sendMailData)
	return err
}

// 测试所需的发送邮件
func (req *Manager) MockSendMail(celId string, toRole string, isRead bool) error {
	sendMailData := SendMailReqData{
		SenderId:   "system",
		CelId:      celId,
		ToRole:     toRole,
		Status:     StatusNormal,
		ExpiryTime: timeutils.Now() + 3600,
	}
	if isRead {
		sendMailData.IsRead = 1
	}
	err := req.sendMail("test111", sendMailData)
	return err
}

func (req *Manager) sendMail(traceId string, data SendMailReqData) error {
	req.ActId = 500
	req.TraceId = traceId
	req.Data = data

	err := req.reqFunction()
	return err
}

type QryMailBsReqData struct {
	Id         string   `json:"id" ps:"邮件id"`
	Ids        []string `json:"ids" ps:"邮件id列表"`
	SenderId   string   `json:"sender_id" ps:"发件人"`
	ToRole     string   `json:"to_role" ps:"收件人"`
	NeedDelete int      `json:"need_delete" ps:"是否需要被玩家删除的"`
	CelId      string   `json:"cel_id" ps:"邮件奖励id"`
	CelIds     []string `json:"cel_ids" ps:"邮件奖励id列表"`
	Size       int      `json:"size" ps:"查询条数"`
	// Page       int      `json:"page" ps:"页数, 查询可翻页查询, page默认为0的问题"`
}

// 拆成两个结构体是因为is_read 和is_receive一定会匹配到, 但有的查询不需要
type QryMailReqData struct {
	QryMailBsReqData
	Status    int   `json:"status" ps:"邮件状态"`
	IsRead    int   `json:"is_read" ps:"是否已读"`
	IsReceive int   `json:"is_receive" ps:"是否已领取"`
	StartTime int64 `json:"start_time" ps:"匹配该时间以后的"`
	EndTime   int64 `json:"end_time" ps:"匹配该时间以前的"`
}

type QryMailReqStatus struct {
	CelId      string `json:"cel_id" ps:"邮件奖励id"`
	Status     int    `json:"status" ps:"邮件状态"`
	NeedDelete int    `json:"need_delete" ps:"是否需要被玩家删除的"`
	IsReceive  int    `json:"is_receive" ps:"是否已领取"`
}

type QryMailSTUSReqData struct {
	QryMailBsReqData
	Status int `json:"status" ps:"邮件状态"`
}

// 玩家查询已读邮件
func (req *Manager) QryRMail(toRole string, traceId string, size int,
	timeBefore int64) ([]*RespDataFields, error) {
	queryMailData := QryMailReqData{
		QryMailBsReqData: QryMailBsReqData{
			ToRole: toRole,
			Size:   size,
		},
		EndTime:   timeBefore,
		Status:    StatusNormal, // 只获取正常邮件(屏蔽邮件不获取), 也可不设置, 研发默认只获取正常邮件
		IsRead:    1,
		IsReceive: 1,
	}
	result, err := req.queryMail(traceId, queryMailData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 根据邮件id查询邮件(只会返回StatusNormal状态的邮件)
func (req *Manager) QryMailById(toRole string, traceId string, mailId string) ([]*RespDataFields, error) {
	queryMailData := QryMailSTUSReqData{
		QryMailBsReqData: QryMailBsReqData{
			ToRole: toRole,
			Id:     mailId,
		},
		Status: StatusNormal,
	}
	result, err := req.queryMail(traceId, queryMailData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 根据邮件id列表查询邮件
func (req *Manager) QryMailByIds(toRole string, traceId string, mailId []string) ([]*RespDataFields, error) {
	queryMailData := QryMailBsReqData{
		ToRole: toRole,
		Ids:    mailId,
		Size:   50, // -1 表示不做size限定, 查询所有 (已失效)
	}
	result, err := req.queryMail(traceId, queryMailData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 根据邮件奖励id列表查询邮件
func (req *Manager) QryMailByCelId(toRole string, traceId string, celId []string, needDelete bool) ([]*RespDataFields, error) {
	queryMailData := QryMailBsReqData{}
	if needDelete {
		queryMailData = QryMailBsReqData{
			ToRole:     toRole,
			CelIds:     celId,
			NeedDelete: 1,
			Size:       50,
		}
	} else {
		queryMailData = QryMailBsReqData{
			ToRole: toRole,
			CelIds: celId,
			Size:   50,
		}
	}

	result, err := req.queryMail(traceId, queryMailData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 根据邮件奖励id和屏蔽状态查询邮件详情
func (req *Manager) QryMailByCelIdStatus(traceId string, celId string, status, isReceive int,
	needDelete bool) ([]*RespDataFields, error) {
	queryMailData := QryMailReqStatus{}
	if needDelete {
		queryMailData = QryMailReqStatus{
			CelId:      celId,
			Status:     status,
			IsReceive:  isReceive,
			NeedDelete: 1,
		}
	} else {
		queryMailData = QryMailReqStatus{
			CelId:     celId,
			Status:    status,
			IsReceive: isReceive,
		}
	}
	result, err := req.queryMail(traceId, queryMailData)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (req *Manager) queryMail(traceId string, data interface{}) ([]*RespDataFields, error) {
	req.ActId = 510
	req.TraceId = traceId
	req.Data = data
	reqDataByte, err := json.MarshalWithErr(req)
	if err != nil {
		return nil, err
	}
	res, err := sendReq(reqDataByte)
	if err != nil {
		return nil, err
	}
	code := gjson.GetBytes(res, "stat").Int()
	if code != 1 {
		info := gjson.GetBytes(res, "info").String()
		return nil, errors.New(info)
	}
	resData := gjson.GetBytes(res, "info").Array()
	var infoArray []*RespDataFields
	for _, vl := range resData {
		info := new(RespDataFields)
		err = json.Unmarshal([]byte(vl.String()), info)
		infoArray = append(infoArray, info)
	}
	return infoArray, nil
}

type QryNoRMailReqData struct {
	ToRole string `json:"to_role" ps:"收件人id"`
	Status int    `json:"status" ps:"邮件状态"`
	Size   int    `json:"size" ps:"查询条数, 保底查询多少封"`
	IsFill int    `json:"is_fill" ps:"是否用已读补满size"`
}

// 查询所有未读未领取, 如果超过size, size不生效(传回所有未读未领取),
// 如果不足且is_fill为1则会返回size条未读和已读邮件的混合
func (req *Manager) QryANoRMail(toRole string, traceId string, size int) ([]*RespDataFields, error) {
	req.ActId = 511
	req.TraceId = traceId
	queryMailData := QryNoRMailReqData{
		ToRole: toRole,
		Status: StatusNormal,
		Size:   size, // 一次性获取所有的未读未领取邮件, 保底最少为size 邮件
		IsFill: 1,    // 如果不够size封邮件, 则用未读的补全
	}
	req.Data = queryMailData
	reqDataByte, err := json.MarshalWithErr(req)
	if err != nil {
		return nil, err
	}
	res, err := sendReq(reqDataByte)
	if err != nil {
		return nil, err
	}
	code := gjson.GetBytes(res, "stat").Int()
	if code != 1 {
		info := gjson.GetBytes(res, "info").String()
		return nil, errors.New(info)
	}
	resData := gjson.GetBytes(res, "info").Array()
	var infoArray []*RespDataFields
	nows := time.Now().UnixNano()
	for _, vl := range resData {
		info := new(RespDataFields)
		json.Unmarshal([]byte(vl.String()), info)
		infoArray = append(infoArray, info)
	}
	log.Println(time.Now().UnixNano() - nows)
	//  未读未领取是否排好序
	var result []*RespDataFields
	var unRead []*RespDataFields
	var read []*RespDataFields
	var readReceive []*RespDataFields
	for _, item := range infoArray {
		if item.IsRead == 0 && item.IsReceive == 0 {
			unRead = append(unRead, item)
		} else if item.IsRead == 1 && item.IsReceive == 0 {
			read = append(read, item)
		} else {
			readReceive = append(readReceive, item)
		}
	}
	result = append(result, unRead...)
	result = append(result, read...)
	result = append(result, readReceive...)
	return result, nil
}

type QryNrNoRMailReqData struct {
	ToRole    string `json:"to_role" ps:"收件人id"`
	Status    int    `json:"status" ps:"邮件状态"`
	StartTime int64  `json:"start_time" ps:"匹配该时间以后的"`
	EndTime   int64  `json:"end_time" ps:"匹配该时间之前的"`
}

// 查询所有未领取数量
func (req *Manager) QryNrNoRMail(toRole string, traceId string) (int, error) {
	req.ActId = 512
	req.TraceId = traceId
	queryMailData := QryNrNoRMailReqData{
		ToRole:  toRole,
		Status:  StatusNormal,
		EndTime: timeutils.Now() + 10,
	}
	req.Data = queryMailData

	reqDataByte, err := json.MarshalWithErr(req)
	if err != nil {
		return 0, err
	}
	res, err := sendReq(reqDataByte)
	if err != nil {
		return 0, err
	}
	code := gjson.GetBytes(res, "stat").Int()
	if code != 1 {
		info := gjson.GetBytes(res, "info").String()
		return 0, errors.New(info)
	}

	resData := gjson.GetBytes(res, "info").String()

	info := new(RespNrNoRFields)
	json.Unmarshal([]byte(resData), info)
	// TODO 未领取是否排好序
	return info.NoReceive, nil
}

type TagRMailReqData struct {
	Id     string   `json:"id" ps:"邮件id"`
	Ids    []string `json:"ids" ps:"邮件id列表"`
	ToRole string   `json:"to_role" ps:"收件人id"`
	All    int      `json:"all" ps:"为1时标记所有该收件人的邮件"`
}

// 标记邮件已读(单封)
func (req *Manager) TagRMailSGL(traceId string, id string) error {
	tagMailData := TagRMailReqData{
		Id: id,
	}
	err := req.tagRMail(traceId, tagMailData)
	return err
}

// 标记邮件已读(多封)
func (req *Manager) TagRMailMul(traceId string, ids []string) error {
	tagMailData := TagRMailReqData{
		Ids: ids,
	}
	err := req.tagRMail(traceId, tagMailData)
	return err
}

// 标记邮件已读(根据收件人标记, 标记该玩家所有未读邮件为已读)
func (req *Manager) TagRMailRole(traceId string, toRole string) error {
	tagMailData := TagRMailReqData{
		ToRole: toRole,
		All:    1,
	}
	err := req.tagRMail(traceId, tagMailData)
	return err
}

func (req *Manager) tagRMail(traceId string, data TagRMailReqData) error {
	req.ActId = 520
	req.TraceId = traceId
	req.Data = data

	err := req.reqFunction()
	return err
}

type TagGMailReqData struct {
	Id     string   `json:"id" ps:"邮件id"`
	Ids    []string `json:"ids" ps:"邮件id列表"`
	ToRole string   `json:"to_role" ps:"收件人id"`
	All    int      `json:"all" ps:"为1时该收件的未领取邮件都会标记为已领取"`
}

// 标记邮件已领取(单封)
func (req *Manager) TagGMailSGL(traceId string, id string) error {
	tagMailData := TagGMailReqData{
		Id: id,
	}
	err := req.tagGMail(traceId, tagMailData)
	return err
}

// 标记邮件已领取(多封)
func (req *Manager) TagGMailMul(traceId string, ids []string) error {
	tagMailData := TagGMailReqData{
		Ids: ids,
	}
	err := req.tagGMail(traceId, tagMailData)
	return err
}

// 标记邮件已领取(根据收件人标记, 标记该玩家所有未领取邮件为已领取)
func (req *Manager) TagGMailRole(traceId string, toRole string) error {
	tagMailData := TagGMailReqData{
		ToRole: toRole,
		All:    1,
	}
	err := req.tagGMail(traceId, tagMailData)
	return err
}

func (req *Manager) tagGMail(traceId string, data TagGMailReqData) error {
	req.ActId = 521
	req.TraceId = traceId
	req.Data = data

	err := req.reqFunction()
	return err
}

type TagDMailReqData struct {
	Id     string   `json:"id" ps:"邮件id"`
	Ids    []string `json:"ids" ps:"邮件id列表"`
	CelId  string   `json:"cel_id" ps:"邮件奖励id"`
	ToRole string   `json:"to_role" ps:"收件id"`
	All    int      `json:"all" ps:"该收件人所有的邮件都会标记为删除"`
}

// 玩家标记删除邮件(单封指定邮件id)
func (req *Manager) TagDMailSGL(traceId string, id string) error {
	tagMailData := TagDMailReqData{
		Id: id,
	}
	err := req.tagDMail(traceId, tagMailData)
	return err
}

// 玩家标记删除邮件(多封指定邮件id)
func (req *Manager) TagDMailMul(traceId string, ids []string) error {
	tagMailData := TagDMailReqData{
		Ids: ids,
	}
	err := req.tagDMail(traceId, tagMailData)
	return err
}

// 玩家标记删除邮件(单封指定cel_id, 需匹配玩家）
func (req *Manager) TagDMailCelId(traceId string, celId string, toRole string) error {
	tagMailData := TagDMailReqData{
		CelId:  celId,
		ToRole: toRole,
	}
	err := req.tagDMail(traceId, tagMailData)
	return err
}

// 标记删除玩家所有邮件(包含未读未领取的邮件)
func (req *Manager) TagDMailRole(traceId string, toRole string) error {
	tagMailData := TagDMailReqData{
		ToRole: toRole,
		All:    1,
	}
	err := req.tagDMail(traceId, tagMailData)
	return err
}

func (req *Manager) tagDMail(traceId string, data TagDMailReqData) error {
	req.ActId = 522
	req.TraceId = traceId
	req.Data = data

	err := req.reqFunction()
	return err
}

type ChgMailReqData struct {
	Id      string   `json:"id" ps:"邮件id"`
	Ids     []string `json:"ids" ps:"邮件id列表"`
	CelId   string   `json:"cel_id" ps:"邮件奖励id"`
	Title   string   `json:"title" ps:"邮件标题"`
	Content string   `json:"content" ps:"邮件内容"`
	Status  int      `json:"status" ps:"邮件状态"`
}

// 修改邮件内容(注: 修改之前需要将屏蔽)
func (req *Manager) ChgMail(traceId string, celId string, title string, content string) error {
	chgMailData := ChgMailReqData{
		CelId: celId,
	}
	if title != "" {
		chgMailData.Title = title
	}
	if content != "" {
		chgMailData.Content = content
	}
	err := req.chgMail(traceId, chgMailData)
	return err
}

// 屏蔽邮件(即修改邮件状态)
func (req *Manager) SHLDMail(traceId string, celId string) error {
	chgMailData := ChgMailReqData{
		CelId:  celId,
		Status: StatusShield,
	}
	err := req.chgMail(traceId, chgMailData)
	return err
}

// 解除屏蔽
func (req *Manager) UnSHLDMail(traceId string, celId string) error {
	chgMailData := ChgMailReqData{
		CelId:  celId,
		Status: StatusNormal,
	}
	err := req.chgMail(traceId, chgMailData)
	return err
}

func (req *Manager) chgMail(traceId string, data ChgMailReqData) error {
	req.ActId = 523
	req.TraceId = traceId
	req.Data = data

	err := req.reqFunction()
	return err
}

type DELMailReqData struct {
	Id     string   `json:"id" ps:"邮件id"`
	Ids    []string `json:"ids" ps:"邮件id列表"`
	CelId  string   `json:"cel_id" ps:"奖励id列表"`
	ToRole string   `json:"to_role" ps:"收件人id"`
	All    int      `json:"all" ps:"删除该收件人所有邮件"`
}

// 物理删除邮件(不可恢复, 保留至数据仓库), 按照邮件id删除
func (req *Manager) DELMailSGL(traceId string, id string) error {
	delMailData := DELMailReqData{
		Id: id,
	}
	err := req.delMail(traceId, delMailData)
	return err
}

// 物理删除邮件(不可恢复, 保留至数据仓库), 按照邮件id列表删除
func (req *Manager) DELMailMul(traceId string, ids []string) error {
	delMailData := DELMailReqData{
		Ids: ids,
	}
	err := req.delMail(traceId, delMailData)
	return err
}

// 物理删除邮件(不可恢复, 保留至数据仓库), 按照角色id删除
func (req *Manager) DELMailMulByRole(traceId string, toRole string) error {
	delMailData := DELMailReqData{
		ToRole: toRole,
	}
	err := req.delMail(traceId, delMailData)
	return err
}

// 物理删除邮件(不可恢复, 保留至数据仓库), 按照奖励id删除所有该将奖励id的邮件
func (req *Manager) DELMailCelId(traceId string, celId string) error {
	delMailData := DELMailReqData{
		CelId: celId,
	}
	err := req.delMail(traceId, delMailData)
	return err
}

// 物理删除个人邮件(不可恢复, 保留至数据仓库), 按照邮件奖励id和玩家id删除
func (req *Manager) DELMailRoleSGL(traceId string, toRole string, celId string) error {
	delMailData := DELMailReqData{
		CelId:  celId,
		ToRole: toRole,
	}
	err := req.delMail(traceId, delMailData)
	return err
}

// 物理删除个人所有邮件(不可恢复, 保留至数据仓库)
func (req *Manager) DELMailRoleAll(traceId string, toRole string) error {
	// TODO check 是否有ALL参数
	delMailData := DELMailReqData{
		ToRole: toRole,
		All:    1,
	}
	err := req.delMail(traceId, delMailData)
	return err
}

func (req *Manager) delMail(traceId string, data DELMailReqData) error {
	req.ActId = 530
	req.TraceId = traceId
	req.Data = data

	err := req.reqFunction()
	return err
}

// req通用请求function, 只适用于只返回是否成功的请求
func (req *Manager) reqFunction() error {
	reqDataByte, err := json.MarshalWithErr(req)
	if err != nil {
		return err
	}
	res, err := sendReq(reqDataByte)
	if err != nil {
		return err
	}
	code := gjson.GetBytes(res, "stat").Int()
	if code != 1 {
		info := gjson.GetBytes(res, "info").String()
		return errors.New(info)
	}
	return nil
}

// 发送请求
func sendReq(data []byte) (b []byte, err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "request mail error:")
		}
	}()
	if strings.ToLower(mailConf.method) == "http" {
		code, resp, err := client.PostJson(mailConf.url, data)
		if err != nil {
			fmt.Println("err:", err.Error())
			fmt.Println("resp:", string(resp))
		}
		return http.ToResp(code, resp, err)
		//return conn.PostJson(mailConf.url, data, mailConf.timeout)
	} else {
		timeout := time.Duration(mailConf.timeout) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		client, err := ggrpc.DialContext(ctx, mailConf.url, ggrpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		conn := pb.NewActionClient(client)
		res, err := conn.ExecAction(ctx, &pb.ActionRequest{Action: string(data)})
		if err != nil {
			return nil, err
		} else {
			return []byte(res.Message), nil
		}
	}
}
