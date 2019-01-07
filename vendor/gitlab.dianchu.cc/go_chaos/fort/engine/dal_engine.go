package engine

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"gitlab.dianchu.cc/go_chaos/fort/engine/cache"
	"gitlab.dianchu.cc/go_chaos/fort/engine/nodeListen"
	"gitlab.dianchu.cc/go_chaos/fort/tools/syslog"
	"gitlab.dianchu.cc/go_chaos/fort/utils"
)

const (
	GZipEncoder  = "gzip"
	CloseEncoder = ""
)

type DALEngine struct {
	dbName       string                       //在zk里配置的数据库别名
	hostList     []string                     //数据访问服务地址列表
	roundRobin   RoundRobinData               //轮询获取主机
	nodeListen   *nodeListen.NodeListenClient //监听zk地址更新的客户端
	httpClient   *http.Client                 //httpClient
	mysqlCacher  *cache.MySQLCacher           //mysql缓存
	encodingType string                       //编码类型
	triggerSize  int                          //编码(压缩)触发大小
}

type RoundRobinData struct {
	HostWeightList map[int]int
	LastHost       int
	CW             int //表示当前调度的权值
	GCD            int //当前所有权重的最大公约数
}

var (
	dalStorer = &dalEngineStore{
		Store: make(map[string]*DALEngine),
	}
)

//并发安全
type dalEngineStore struct {
	sync.RWMutex
	Store map[string]*DALEngine
}

func (store *dalEngineStore) getEngine(sourceName string) *DALEngine {
	store.RLock()
	dalEngine, ok := store.Store[sourceName]
	store.RUnlock()
	if ok {
		return dalEngine
	}
	return nil
}

func (store *dalEngineStore) addEngine(sourceName string, dalEngine *DALEngine) {
	store.Lock()
	store.Store[sourceName] = dalEngine
	store.Unlock()
}

func NewDALEngine(dbName, zkPath string, zkHost, zkAuth []string) (*DALEngine, error) {
	var (
		err    error
		source bytes.Buffer
		client *DALEngine
	)
	source.WriteString(dbName)
	source.WriteString(zkPath)
	if client = dalStorer.getEngine(source.String()); client != nil {
		return client, nil
	}

	transport := http.Transport{
		//DisableKeepAlives: true, //是否禁用重用连接，测试时使用
		IdleConnTimeout: 30 * time.Second, //socket在该时间内没有交互则自动关闭连接
	}
	httpClient := &http.Client{
		Timeout:   time.Second * 60,
		Transport: &transport,
	}

	client = &DALEngine{
		dbName: dbName,
		//ctx:          &ctx,
		nodeListen:   nodeListen.NewNodeListenClient(zkPath, zkHost, zkAuth),
		httpClient:   httpClient,
		encodingType: GZipEncoder, //使用gzip压缩
		triggerSize:  16,          //>=16KB时,启用压缩
	}
	dalStorer.addEngine(source.String(), client)
	return client, err
}

func (dal *DALEngine) SetMySQLCacher(cacher *cache.MySQLCacher) {
	dal.mysqlCacher = cacher
}

func (dal *DALEngine) MySQLCacher() *cache.MySQLCacher {
	return dal.mysqlCacher
}

//func (dal *DALEngine) SetDBName(dbName string) {
//	dal.dbName = dbName
//}

func (dal *DALEngine) GetDBName() string {
	return dal.dbName
}

type DALResp struct {
	Info    []byte //gob序列化后的查询结果
	ErrInfo string //业务错误信息
	Count   int64  //查询到的条数、执行后的影响数
	Code    int64  //业务代码
}

type DALQuery struct {
	DB      string
	IsArray bool
	Data    CmdData
}

type DALTransaction struct {
	DB   string
	Data []CmdData
}

type CmdData struct {
	Statement string
	Args      []interface{}
}

//encodingType(压缩编码类型):默认为"",表示不压缩 否则会判断 是否压缩请求数据.  目前支持类型:"gzip"
//triggerSize(触发编码大小,KB):默认为0,表示返回响应不需要被压缩。只在查询API上有效
func (dal *DALEngine) SetEncoding(encodingType string, triggerSize int) {
	dal.encodingType = encodingType
	dal.triggerSize = triggerSize
}

func (dal *DALEngine) DALDisTransaction(ctx context.Context, data []DALTransaction) error {
	ok, resp, err := dal.dalApiHandle(ctx, "810", data)
	if err != nil {
		return err
	}
	if !ok {
		var e bytes.Buffer
		e.WriteString("code:")
		e.WriteString(strconv.FormatInt(resp.Code, 10))
		e.WriteString(" - ")
		e.WriteString(resp.ErrInfo)
		syslog.FortLog.ShowLog(syslog.ERROR, e.String())
		return errors.New(e.String())
	}
	return err
}

func (dal *DALEngine) DALTransaction(ctx context.Context, data []CmdData) error {
	var err error
	//Post Body
	body := DALTransaction{
		DB:   dal.dbName,
		Data: data,
	}
	ok, resp, err := dal.dalApiHandle(ctx, "808", &body)
	if err != nil {
		return err
	}
	if !ok {
		var e bytes.Buffer
		e.WriteString("code:")
		e.WriteString(strconv.FormatInt(resp.Code, 10))
		e.WriteString(" - ")
		e.WriteString(resp.ErrInfo)
		syslog.FortLog.ShowLog(syslog.ERROR, e.String())
		return errors.New(e.String())
	}
	return err
}

func (dal *DALEngine) DALQuery(ctx context.Context, data CmdData) (int64, []map[string]interface{}, error) {
	//Post Body
	body := DALQuery{
		DB:   dal.dbName,
		Data: data,
	}
	ok, resp, err := dal.dalApiHandle(ctx, "809", &body)
	if err != nil {
		return 0, nil, err
	}

	// 检查ORM服务业务是否成功
	if !ok {
		var e bytes.Buffer
		e.WriteString("code:")
		e.WriteString(strconv.FormatInt(resp.Code, 10))
		e.WriteString(" - ")
		e.WriteString(resp.ErrInfo)
		syslog.FortLog.ShowLog(syslog.ERROR, e.String())
		return 0, nil, errors.New(e.String())
	}

	var res []map[string]interface{}
	if err = utils.Unmarshal(resp.Info, &res); err != nil {
		return 0, nil, err
	}
	return resp.Count, res, err
}

func (dal *DALEngine) dalApiHandle(ctx context.Context, api string, body interface{}) (ok bool, resp DALResp, err error) {
	bodyRawByte, err := utils.Marshal(body)
	if err != nil {
		return
	}
	//Get Trace ID
	traceID, idExist := ctx.Value(utils.TRACE_ID).(string)
	if !idExist {
		err = errors.New("The transaction has not trace id! ")
		syslog.FortLog.ShowLog(syslog.ERROR, err.Error())
		return
	}
	//捕获HttpPost ReadResp 时可能出现的异常
	defer func() {
		if panicErr := recover(); panicErr != nil {
			syslog.FortLog.ShowLog(syslog.ERROR, fmt.Sprint(panicErr))
			err = errors.New(fmt.Sprint(panicErr))
		}
	}()
	url, err := dal.getUrl(traceID, api)
	if err != nil {
		return
	}
	var needEncoding, isEncoded = false, false
	if dal.encodingType != "" && len(bodyRawByte) >= dal.triggerSize*1024 { //压缩请求
		switch dal.encodingType {
		case GZipEncoder:
			if g, err := utils.GzipEncode(bodyRawByte); err == nil {
				bodyRawByte = g
				isEncoded = true
				if api == "809" {
					needEncoding = true
				}
			}
		}
	}
	//只有查询请求的响应才需要判断要不要压缩
	ok, err = dal.post(url, bodyRawByte, &resp, needEncoding, isEncoded)
	return
}

func (dal *DALEngine) getAddr() (string, error) {
	var hostList []string
	if hostList = dal.nodeListen.GetHostList(); hostList == nil {
		return "", errors.New("The address list is empty ")
	}
	dal.setHostList(hostList)
	getMaxWeight := func() int {
		for i := range dal.hostList {
			if weight := dal.roundRobin.HostWeightList[i]; weight >= 0 {
				return weight
			}
		}
		return 0
	}
	for {
		dal.roundRobin.LastHost = (dal.roundRobin.LastHost + 1) % len(dal.hostList)
		if dal.roundRobin.LastHost == 0 {
			dal.roundRobin.CW = dal.roundRobin.CW - dal.roundRobin.GCD
			if dal.roundRobin.CW <= 0 {
				if dal.roundRobin.CW = getMaxWeight(); dal.roundRobin.CW == 0 {
					return dal.hostList[0], nil
				}
			}
		}
		if weight := dal.roundRobin.HostWeightList[dal.roundRobin.LastHost]; weight >= dal.roundRobin.CW {
			return dal.hostList[dal.roundRobin.LastHost], nil
		}
	}
}

func (dal *DALEngine) getUrl(traceID, actID string) (string, error) {
	var (
		url     bytes.Buffer
		address []string
		host    string
		addr    string
		port    int
		err     error
	)
	if addr, err = dal.getAddr(); err != nil {
		return "", err
	}
	address = strings.Split(addr, ":")
	host = address[0]
	port, err = strconv.Atoi(address[1])
	if err != nil {
		return "", err
	}
	url.WriteString("http://")
	url.WriteString(host)
	url.WriteString(":")
	url.WriteString(strconv.Itoa(port - 1))
	url.WriteString("/not_sql/")
	url.WriteString(traceID)
	url.WriteString("/")
	url.WriteString(actID)
	url.WriteString("/")
	return url.String(), nil
}

//请求url,请求数据,返回结果反序列化结构体,请求响应是否需要被压缩,请求是否被编码
//返回业务是否成功+HttpBody
func (dal *DALEngine) post(url string, data []byte, dalResp *DALResp, needEncoding, isEncoded bool) (bool, error) {
	var (
		resp          *http.Response
		respBody      []byte
		businessState string
		execAction    string
		compression   string
		err           error
	)

	body := bytes.NewBuffer(data)
	postRequest, err := http.NewRequest("POST", url, body)

	//设置响应是否需要被压缩
	if needEncoding {
		postRequest.Header.Add("EncodingType", dal.encodingType)
		postRequest.Header.Add("TriggerSize", strconv.Itoa(dal.triggerSize))
	}

	//设置该请求是否需要解码
	if isEncoded {
		postRequest.Header.Add("IsEncoded", dal.encodingType)
	}

	if resp, err = dal.httpClient.Do(postRequest); err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return false, err
	}

	businessState = resp.Header.Get("BusinessState")
	execAction = resp.Header.Get("Action")
	compression = resp.Header.Get("IsEncoded")
	switch compression {
	case GZipEncoder:
		if gzipByte, err := utils.GzipDecode(respBody); err == nil {
			respBody = gzipByte
		}
	}

	if businessState == "Failure" {
		dalResp.Code = int64(binary.LittleEndian.Uint64(respBody[:8]))
		dalResp.ErrInfo = string(respBody[8:])
		return false, nil
	}

	if businessState == "Success" {
		switch execAction {
		case "DisTransaction", "Transaction":
			return true, nil // 此处有返回影响条数，但是后面的逻辑没有去用,所以直接返回
		case "Query":
			if err = utils.Unmarshal(respBody, &dalResp); err != nil {
				return false, err
			}
			return true, nil
		default:
			return false, errors.New("ActionID Error ")
		}
	}
	return false, errors.New("State Error ")
}

func (dal *DALEngine) setHostList(data []string) {
	dal.hostList = data
	weightList := make(map[int]int)
	for i := range dal.hostList {
		weightList[i] = 0
	}
	dal.roundRobin.HostWeightList = weightList
}
