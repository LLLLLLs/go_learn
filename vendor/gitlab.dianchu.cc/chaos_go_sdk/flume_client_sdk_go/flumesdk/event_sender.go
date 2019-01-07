package flumesdk

import (
	"encoding/json"
	"fmt"
	"gitlab.dianchu.cc/chaos_go_sdk/flume_client_sdk_go/flumesdk/flume"
	"gitlab.dianchu.cc/chaos_go_sdk/flume_client_sdk_go/flumesdk/thrift"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type EventSender struct {
	conf *ConfigManager //配置模块
	//todo 前期计数验证sdk可靠性，后期删除
	debug      bool             //输出详细日志
	count      int              //计数
	lock       sync.Mutex       //计数锁
	cacheQueue chan interface{} //缓存队列
	pool       *connPool        //连接池
	daemon     bool             //后台标志
	stopFlag   chan int         //结束标志
	logger     *log.Logger      //日志组件
	backgroup  *sync.WaitGroup  //back发送group
}

// EventSender为全局单例模式，仅进行一次初始化操作
var myEventSender *EventSender
var lock sync.Mutex

func NewEventSender(zkServer []string, flumePath, confPath string, logger *log.Logger, debug bool) *EventSender {
	lock.Lock()
	defer lock.Unlock()
	if myEventSender != nil {
		return myEventSender
	}

	zk := newConfigManagerParam(zkServer, flumePath, confPath)

	myEventSender = &EventSender{
		conf: zk,
		//todo 计数，后期删除
		debug:      debug,
		count:      0,
		cacheQueue: make(chan interface{}, 100000),
		pool:       newConnPool(zk, logger, debug),
		daemon:     true,
		logger:     logger,
		stopFlag:   make(chan int),
		backgroup:  new(sync.WaitGroup),
	}
	//开启后台
	myEventSender.init()
	go myEventSender.daemonWork()
	return myEventSender
}

func (S *EventSender) init() {
	filePath := S.conf.GetFlumePath()
	info, err := os.Stat(filePath)
	if err != nil || !info.IsDir() {
		os.MkdirAll(filePath, os.ModePerm)
	}
}

func (S *EventSender) ToCoin(fieldName string, oldVal, newVal int64, lifeTime int16, scene, remark, roleId, reserve string, serverId, cpAppId int32, logTime int64) error {
	var (
		err     error
		content []byte
	)
	//todo 参数验证
	coin := NewCoinLog(fieldName, oldVal, newVal, lifeTime, scene, remark, roleId, reserve, serverId, cpAppId, logTime)
	content, err = json.Marshal(coin)
	if err != nil {
		S.logger.Println("json encode error", err)
		return err
	}
	event := flume.NewThriftFlumeEvent()
	event.Body = content
	event.Headers["State"] = "coin_log"
	S.send(event)
	return nil
}

func (S *EventSender) ToAct(actId string, actStat int32, req, res string, procTime, serverId, cpAppId int32, roleId, reserve string, logTime int64) error {
	var (
		err     error
		content []byte
	)
	//todo 参数验证
	act := NewActLog(actId, actStat, req, res, procTime, serverId, cpAppId, roleId, reserve, logTime)
	content, err = json.Marshal(act)
	if err != nil {
		S.logger.Println("json encode error", err)
		return err
	}
	event := flume.NewThriftFlumeEvent()
	event.Body = content
	event.Headers["State"] = "act_log"
	S.send(event)
	return nil
}

func (S *EventSender) ToLogin(userId, deviceId, deviceType, deviceOs string, retailId, lvl, upTime, serverId, cpAppId int32, ip, roleId, reserve string, logTime int64) error {
	var (
		err     error
		content []byte
	)
	//todo 参数验证
	login := NewLoginLog(userId, deviceId, deviceType, deviceOs, retailId, lvl, upTime, serverId, cpAppId, ip, roleId, reserve, logTime)
	content, err = json.Marshal(login)
	if err != nil {
		S.logger.Println("json encode error", err)
		return err
	}
	event := flume.NewThriftFlumeEvent()
	event.Body = content
	event.Headers["State"] = "login_log"
	S.send(event)
	return nil
}

func (S *EventSender) ToReg(roleId, userId, deviceId, deviceType, deviceOs string, retailId, serverId, cpAppId int32, reserve string, logTime int64, ip string) error {
	var (
		err     error
		content []byte
	)
	//todo 参数验证
	reg := NewRegLog(roleId, userId, deviceId, deviceType, deviceOs, retailId, serverId, cpAppId, reserve, logTime, ip)
	content, err = json.Marshal(reg)
	if err != nil {
		S.logger.Println("json encode error", err)
		return err
	}
	event := flume.NewThriftFlumeEvent()
	event.Body = content
	event.Headers["State"] = "reg_log"
	S.send(event)
	return nil
}

func (S *EventSender) ToVar(fieldName, scene, roleId string, oldVal, newVal int64, serverID, cpAppId int32, remark, reserve string, logTime int64) error {
	var (
		err     error
		content []byte
	)
	//todo 参数验证
	varlog := NewVarLog(fieldName, scene, roleId, oldVal, newVal, serverID, cpAppId, remark, reserve, logTime)
	content, err = json.Marshal(varlog)
	if err != nil {
		S.logger.Println("json encode error", err)
		return err
	}
	event := flume.NewThriftFlumeEvent()
	event.Body = content
	event.Headers["State"] = "var_log"
	S.send(event)
	return nil
}

func (S *EventSender) ToChat(id string, gameId, cpAppId, serverId int32, channel, userId, roleID, alliance, deviceId, ext string, msgType int32, text, rawImage, thumbnail, voice string, voiceLength int32, logTime int64, reserve string) error {
	var (
		err     error
		content []byte
	)
	//todo 参数验证
	chatlog := NewChatLog(id, gameId, cpAppId, serverId, channel, userId, roleID, alliance, deviceId, ext, msgType, text, rawImage, thumbnail, voice, voiceLength, logTime, reserve)
	content, err = json.Marshal(chatlog)
	if err != nil {
		S.logger.Println("json encode error", err)
		return err
	}
	event := flume.NewThriftFlumeEvent()
	event.Body = content
	event.Headers["State"] = "chat_log"
	S.send(event)
	return nil
}

func (S *EventSender) send(event *flume.ThriftFlumeEvent) {
	defer func() {
		if err := recover(); err != nil {
			S.logger.Printf("panic error %s \n", err)
		}
	}()
	//S.logger.Println(S.backgroup)
	if strings.Contains(S.conf.GetBatchTable(), event.Headers["State"]) {
		S.cacheQueue <- event
	} else {
		S.backgroup.Add(1)
		go S.sendOne(event)
	}
}

func (S *EventSender) sendOne(event *flume.ThriftFlumeEvent) {
	defer S.backgroup.Add(-1)
	timeOut := S.conf.GetSingle().Timeout
	client, ok := S.pool.GetClient(time.Duration(timeOut)*time.Millisecond, false)
	if !ok {
		S.cacheQueue <- event
		return
	}
	start := time.Now()
	state, err := client.Append(event)
	if err != nil {
		//todo 错误处理连接异常
		S.logger.Println(err)
		S.cacheQueue <- event
		S.pool.KillConnect(client)
		return
	}
	if state != flume.Status_OK {
		//发送失败
		S.logger.Println("state error", state)
		S.cacheQueue <- event
		S.pool.PutClient(client)
		return
	}
	if S.debug {
		S.lock.Lock()
		S.count++
		S.lock.Unlock()
		S.logger.Printf("send success %v coast time %v total send %d \n", client.TransPort.HostPort, time.Since(start), S.count)
	}
	// 回收连接
	S.pool.PutClient(client)
}

func (S *EventSender) sendBatch(events []*flume.ThriftFlumeEvent) {
	defer S.backgroup.Add(-1)
	timeOut := S.conf.GetBatch().Timeout
	client, ok := S.pool.GetClient(time.Duration(timeOut)*time.Millisecond, true)
	if !ok {
		//获取连接失败
		if S.debug {
			S.logger.Println("bad connect")
		}
		S.writeFile(events)
		return
	}
	start := time.Now()
	state, err := client.AppendBatch(events)

	if err != nil {
		S.logger.Println(err)
		if e, ok := err.(thrift.TTransportException); ok && e.TypeId() == 4 {
			flag := len(events) / 2
			S.writeFile(events[:flag])
			S.writeFile(events[flag:])
		} else {
			S.writeFile(events)
		}
		S.pool.KillConnect(client)
		return
	}
	if state != flume.Status_OK {
		//写文件
		S.logger.Println("state error", state)
		S.writeFile(events)
		S.pool.PutClient(client)
		return
	}
	if S.debug {
		S.lock.Lock()
		S.count += len(events)
		S.lock.Unlock()
		S.logger.Printf("batch send success %v coast time %v total send %d \n", client.TransPort.HostPort, time.Since(start), S.count)
	}
	S.pool.PutClient(client)
}

func (S *EventSender) scanQueue() {
	count := 0
	interval := int64(S.conf.GetScan().Interval / 1000)
	batchSize := S.conf.GetBatch().BatchCount
	start := time.Now().Unix()
	var events []*flume.ThriftFlumeEvent
	for count < batchSize && time.Now().Unix()-start < interval {
		//select {
		//case event, ok := <-S.cacheQueue:
		//	if !ok {
		//		break
		//	}
		//	count++
		//	events = append(events, event.(*flume.ThriftFlumeEvent))
		//case <-time.After(time.Second):
		//	continue
		//}
		event, ok := <-S.cacheQueue
		if !ok {
			break
		}
		count++
		events = append(events, event.(*flume.ThriftFlumeEvent))
	}
	if len(events) > 0 {
		S.backgroup.Add(1)
		go S.sendBatch(events)
	}
}

func (S *EventSender) scanFile() {
	filePath := S.conf.GetFlumePath()
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		S.logger.Println(err)
		return
	}
	if S.debug {
		S.logger.Println("scanFile (time ,fileNum) == ", time.Now().Unix(), len(files))
	}
	if len(files) == 0 {
		return
	}
	t := time.Now()
	num := rand.Intn(len(files))
	name := files[num].Name()
	fileName := filePath + "/" + name
	events, err := S.readFile(fileName)
	if S.debug {
		S.logger.Println("read file cost time", time.Since(t), fileName)
	}
	if len(events) > 0 {
		S.backgroup.Add(1)
		go S.sendBatch(events)
	}
	os.Remove(fileName)
}

func (S *EventSender) daemonWork() {
	for S.daemon {
		scanQueue := S.conf.GetScan().ScanQueue
		scanFile := S.conf.GetScan().ScanFile
		// 大于等于0表示开启该扫描
		if scanQueue >= 0 {
			S.scanQueue()
		}
		if scanFile >= 0 {
			if scanQueue < 0 {
				interval := S.conf.GetScan().Interval / 1000
				time.Sleep(time.Duration(interval) * time.Second)
			}
			S.scanFile()
		}
	}
	S.stopFlag <- 1
}

func (S *EventSender) flashQueue() {
	for len(S.cacheQueue) > 0 {
		S.scanQueue()
	}
}

func (S *EventSender) Close() {
	S.logger.Println("程序退出")
	S.backgroup.Wait()
	S.daemon = false
	close(S.cacheQueue) //关闭缓存
	<-S.stopFlag        //等待扫描结束
	S.flashQueue()      //清空队列
	S.backgroup.Wait()
	S.pool.Close()
	println(S.pool.count)
	S.logger.Println("退出成功")
	S.logger = nil
	S.conf = nil
	myEventSender = nil
}

func (S *EventSender) writeFile(events []*flume.ThriftFlumeEvent) error {
	var err error
	filePath := S.conf.GetFlumePath()
	name := time.Now().UnixNano() + rand.Int63n(99999)
	fileName := fmt.Sprintf("%s/%d", filePath, name)
	content, err := json.Marshal(events)
	if err != nil {
		S.logger.Println(err)
		return err
	}
	t := time.Now()
	err = ioutil.WriteFile(fileName, content, os.ModePerm)
	if err != nil {
		S.logger.Println(err)
		return err
	}
	if S.debug {
		S.logger.Println("write file coast time ", time.Since(t))
	}
	return nil
}

func (S *EventSender) readFile(fileName string) ([]*flume.ThriftFlumeEvent, error) {
	var events []*flume.ThriftFlumeEvent
	var err error
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		S.logger.Println(err)
		return nil, err
	}
	err = json.Unmarshal(content, &events)
	if err != nil {
		S.logger.Println(err)
	}
	return events, nil
}

type TransAction struct {
	sender *EventSender
	buff   []*flume.ThriftFlumeEvent
}

func (T *TransAction) Add(log EventLog) error {
	content, err := json.Marshal(log)
	if err != nil {
		return err
	}
	event := flume.NewThriftFlumeEvent()
	event.Body = content
	event.Headers["State"] = log.eventLog()
	T.buff = append(T.buff, event)
	return nil
}

func (T *TransAction) Commit() {
	T.sender.backgroup.Add(1)
	T.sender.sendBatch(T.buff)
}

func (S *EventSender) TransAction() *TransAction {
	return &TransAction{
		sender: S,
		buff:   []*flume.ThriftFlumeEvent{},
	}
}
