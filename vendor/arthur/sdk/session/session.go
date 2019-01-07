/*
Created on 2018/4/19 15:26

author: ChenJinLong

Content:
*/
package session

import (
	pb "arthur/sdk/session/grpc_project/chaossession/protos"
	"arthur/utils/errors"
	"arthur/utils/http"
	"github.com/bitly/go-simplejson"
	"github.com/tidwall/gjson"
	"golang.org/x/net/context"
	ggrpc "google.golang.org/grpc"
	"strings"
	"time"
)

var (
	sess *Manager
	//TODO：连接池配置参数后续需要配置在ZK或其他地方
	client = http.NewNetClient(2000, 100, 20*time.Second, 5*time.Second)
)

type Info map[string]string

type Manager struct {
	url     string
	method  string
	version string
	timeout int
}

func Init(timeout int, method, url string) {
	s := &Manager{
		url:     url,
		method:  method,
		version: "1.6",
		timeout: timeout,
	}
	sess = s
}

//添加一个用户到会话管理
func AddRole(appId, serverId int, jwt, uid, deviceId, rid, retailId, traceId string) error {
	temp := simplejson.New()
	temp.Set("act_id", 303)
	temp.Set("version", sess.version)
	temp.Set("server_id", serverId)
	temp.Set("app_id", appId)

	dataTmp := map[string]string{}
	dataTmp["jwt"] = jwt
	dataTmp["uid"] = uid
	dataTmp["device_id"] = deviceId
	dataTmp["rid"] = rid
	dataTmp["retail_id"] = retailId

	temp.Set("data", dataTmp)
	temp.Set("trace_id", traceId)

	data, err := temp.MarshalJSON()
	if err != nil {
		return err
	}
	res, err := sendReq(data)
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

//删除一个用户
func DelRole(appId, serverId int, jwt, traceId string) error {

	temp := simplejson.New()
	temp.Set("act_id", 304)
	temp.Set("version", sess.version)
	temp.Set("app_id", appId)
	temp.Set("server_id", serverId)
	dataTmp := map[string]string{}
	dataTmp["jwt"] = jwt
	temp.Set("data", dataTmp)
	temp.Set("trace_id", traceId)

	data, err := temp.MarshalJSON()
	if err != nil {
		return err
	}
	res, err := sendReq(data)
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

//获取在线全区服在线用户数已经总数
func GetOnlineUser(appId, serverId int, traceId string) (map[string]int64, int64, error) {
	temp := simplejson.New()
	temp.Set("act_id", 313)
	temp.Set("version", sess.version)
	temp.Set("app_id", appId)
	temp.Set("trace_id", traceId)

	data, err := temp.MarshalJSON()
	if err != nil {
		return nil, 0, err
	}
	res, err := sendReq(data)
	if err != nil {
		return nil, 0, err
	}
	code := gjson.GetBytes(res, "stat").Int()
	if code != 1 {
		info := gjson.GetBytes(res, "info").String()
		return nil, 0, errors.New(info)
	}
	userSumMap := map[string]int64{}
	userSumMapTmp := gjson.GetBytes(res, "data.user_sum_map").Map()
	for k, v := range userSumMapTmp {
		userSumMap[k] = v.Int()
	}
	sumMapTmp := gjson.GetBytes(res, "data.user_sum").Int()
	return userSumMap, sumMapTmp, nil

}

//获取在线用户Rid
func GetAllOnlineRoleId(appId, serverId int, traceId string) ([]string, error) {
	temp := simplejson.New()
	temp.Set("act_id", 310)
	temp.Set("version", sess.version)
	temp.Set("app_id", appId)
	temp.Set("server_id", serverId)
	temp.Set("trace_id", traceId)

	data, err := temp.MarshalJSON()
	if err != nil {
		return nil, err
	}
	res, err := sendReq(data)
	if err != nil {
		return nil, err
	}
	code := gjson.GetBytes(res, "stat").Int()
	if code != 1 {
		info := gjson.GetBytes(res, "info").String()
		return nil, errors.New(info)
	}
	onlineRoleIds := make([]string, 0)
	onlineRoleIdsTmp := gjson.GetBytes(res, "data.rids").Array()
	for _, v := range onlineRoleIdsTmp {
		onlineRoleIds = append(onlineRoleIds, v.String())
	}
	return onlineRoleIds, nil
}

//获取单个用户的会话信息
func GetRoleSession(appId, serverId int, uid, jwt, rid, traceId string) (Info, error) {
	temp := simplejson.New()
	temp.Set("act_id", 306)
	temp.Set("version", sess.version)
	temp.Set("app_id", appId)
	temp.Set("server_id", serverId)
	temp.Set("trace_id", traceId)

	dataTmp := map[string]string{}
	if uid != "" {
		dataTmp["uid"] = uid
		dataTmp["operation"] = "1"
	}
	if jwt != "" {
		dataTmp["jwt"] = jwt
		dataTmp["operation"] = "2"
	}
	if rid != "" {
		dataTmp["rid"] = rid
		dataTmp["operation"] = "3"
	}
	temp.Set("data", dataTmp)
	data, err := temp.MarshalJSON()
	if err != nil {
		return nil, err
	}
	res, err := sendReq(data)
	if err != nil {
		return nil, err
	}
	code := gjson.GetBytes(res, "stat").Int()
	if code != 1 {
		info := gjson.GetBytes(res, "info").String()
		return nil, errors.New(info)
	}
	sessionInfo := Info{}
	sessionInfoTmp := gjson.GetBytes(res, "data").Map()
	for k, v := range sessionInfoTmp {
		sessionInfo[k] = v.String()
	}
	return sessionInfo, nil
}

//获取多个用户的会话信息
func GetRoleSessions(appId, serverId int, uids, jwts, rids []string, traceId string) (map[string]Info, error) {
	temp := simplejson.New()
	temp.Set("act_id", 308)
	temp.Set("version", sess.version)
	temp.Set("app_id", appId)
	temp.Set("server_id", serverId)
	temp.Set("trace_id", traceId)

	dataTmp := map[string]interface{}{}
	if len(uids) != 0 {
		dataTmp["uid"] = uids
		dataTmp["operation"] = "1"
	}
	if len(jwts) != 0 {
		dataTmp["jwts"] = jwts
		dataTmp["operation"] = "2"
	}
	if len(rids) != 0 {
		dataTmp["rid"] = rids
		dataTmp["operation"] = "3"
	}
	temp.Set("data", dataTmp)
	data, err := temp.MarshalJSON()
	if err != nil {
		return nil, err
	}
	res, err := sendReq(data)
	if err != nil {
		return nil, err
	}
	code := gjson.GetBytes(res, "stat").Int()
	if code != 1 {
		info := gjson.GetBytes(res, "info").String()
		return nil, errors.New(info)
	}

	sessionInfos := map[string]Info{}
	sessionInfoTmps := gjson.GetBytes(res, "data").Map()
	for k, v := range sessionInfoTmps {
		sessionInfoMini := Info{}
		for i, j := range v.Map() {
			sessionInfoMini[i] = j.String()
		}
		sessionInfos[k] = sessionInfoMini
	}
	return sessionInfos, nil
}

//剔除某区服所有玩家
func KickAllOnlineRole(appId, serverId int, traceId string) (int64, error) {
	temp := simplejson.New()
	temp.Set("act_id", 311)
	temp.Set("version", sess.version)
	temp.Set("app_id", appId)
	temp.Set("server_id", serverId)
	temp.Set("trace_id", traceId)

	data, err := temp.MarshalJSON()
	if err != nil {
		return 0, err
	}
	res, err := sendReq(data)
	if err != nil {
		return 0, err
	}
	code := gjson.GetBytes(res, "stat").Int()
	if code != 1 {
		info := gjson.GetBytes(res, "info").String()
		return 0, errors.New(info)
	}
	kickNum := gjson.GetBytes(res, "data.count").Int()
	return kickNum, nil
}

func sendReq(data []byte) (b []byte, err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "session request error: ")
		}
	}()
	if strings.ToLower(sess.method) == "http" {
		code, resp, err := client.PostJson(sess.url, data)
		return http.ToResp(code, resp, err)
	} else {
		timeout := time.Duration(sess.timeout) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		client, err := ggrpc.DialContext(ctx, sess.url, ggrpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		conn := pb.NewSessionServiceClient(client)
		res, err := conn.SessionAction(ctx, &pb.SessionRequest{ActionData: string(data)})
		if err != nil {
			return nil, err
		} else {
			return []byte(res.ReplyData), nil
		}

	}

}
