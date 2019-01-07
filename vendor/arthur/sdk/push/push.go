/*
Created on 2018/8/13 16:37

author: ChenJinLong

Content:
*/
package push

import (
	pb "arthur/sdk/push/pushservice"
	"arthur/utils/errors"
	"arthur/utils/jsonutils"
	"arthur/utils/log"
	"strings"
	"time"

	"arthur/utils/http"
	"github.com/tidwall/gjson"
	"golang.org/x/net/context"
	ggrpc "google.golang.org/grpc"
)

var pushConf *ConfManager

type ConfManager struct {
	url     string
	method  string
	timeout int
}

var (
	//TODO：连接池配置参数后续需要配置在ZK或其他地方
	client = http.NewNetClient(2000, 100, 20*time.Second, 5*time.Second)
)

func Init(timeout int, method, httpAddr, grpcAddr string) {
	var url string
	if strings.ToLower(method) == "http" {
		url = httpAddr
	} else {
		url = grpcAddr
	}
	p := &ConfManager{
		url:     url,
		method:  method,
		timeout: timeout,
	}
	pushConf = p
}

//推送消息请求结构体
type pushMsgReq struct {
	ActId      int                          //接口标识
	GlobalMsg  int                          //0表示非全局消息，推送给RidList字段中指定的用户，1表示全局消息推送给对应AppId和ServerId里的所有在线用户
	AppId      int                          //应用标号
	ServerId   int                          //区服编号
	ServerList []int    `json:",omitempty"` //区服编号列表
	RidList    []string `json:",omitempty"` //由多个角色idRid组成的列表，可以推送消息给指定的角色（全局消息时为空列表）
	MsgType    int                          //推送消息类型编号，转发给客户端处理
	MsgData    string                       //要推送的消息
	ConnLabel  []string `json:",omitempty"` //客户端连接标志，发送给有标志的所有客户端
}

//发送消息至推送服务
func PushsMsg(appId, serverId int, serverList []int, ridList, connLabel []string, globalMsg, msgType int, msgData string) error {
	newReq := new(pushMsgReq)
	newReq.ActId = 30001
	newReq.AppId = appId
	newReq.ServerId = serverId
	newReq.ServerList = serverList
	newReq.RidList = ridList
	newReq.MsgType = msgType
	newReq.MsgData = msgData
	newReq.GlobalMsg = globalMsg
	newReq.ConnLabel = connLabel
	log.Debug("push request: ", newReq)
	data := json.Marshal(newReq)
	res, err := sendReq(data)
	if err != nil {
		return err
	}
	code := gjson.GetBytes(res, "code").Int()
	if code != 0 {
		info := gjson.GetBytes(res, "info").String()
		return errors.New(info)
	}
	return nil
}

func sendReq(data []byte) (b []byte, err error) {
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "request push error:")
		}
	}()
	if strings.ToLower(pushConf.method) == "http" {
		code, resp, err := client.PostJson(pushConf.url, data)
		return http.ToResp(code, resp, err)
	} else {
		timeout := time.Duration(pushConf.timeout) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		client, err := ggrpc.DialContext(ctx, pushConf.url, ggrpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		conn := pb.NewPushServiceClient(client)
		res, err := conn.PushAction(ctx, &pb.PushRequest{Data: string(data)})
		if err != nil {
			return nil, err
		} else {
			return []byte(res.Info), nil
		}

	}

}
