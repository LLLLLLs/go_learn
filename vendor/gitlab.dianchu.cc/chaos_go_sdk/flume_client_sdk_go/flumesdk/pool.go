package flumesdk

import (
	"gitlab.dianchu.cc/chaos_go_sdk/flume_client_sdk_go/flumesdk/flume"
	"gitlab.dianchu.cc/chaos_go_sdk/flume_client_sdk_go/flumesdk/thrift"
	"log"
	"sync/atomic"
	"time"
)

type conn struct {
	conn           *thrift.TFramedTransport // transport
	sock           *thrift.TSocket
	maxLifeTime    int64   //连接的最大生存时间
	createTime     int64   //创建的时间
	lastActiveTime int64   //最后活跃的时间
	retreatLevel   int     // 退避等级
	retreat        []int64 // 退避时间数组
	HostPort       string  //ip端口
}

func newConn(hostPort string, maxLifeTime int64, timeout time.Duration) (*conn, error) {
	trans, err := thrift.NewTSocketTimeout(hostPort, timeout)
	if err != nil {
		return nil, err
	}

	conn := conn{
		conn:           thrift.NewTFramedTransport(trans),
		sock:           trans,
		maxLifeTime:    maxLifeTime,
		retreatLevel:   0,
		retreat:        []int64{0, 2, 6, 10, 20, 40, 80, 100, 200},
		createTime:     time.Now().Unix(),
		lastActiveTime: time.Now().Unix(),
		HostPort:       hostPort,
	}
	return &conn, nil
}

func (C *conn) IsOld() bool {
	return time.Now().Unix()-C.createTime > C.maxLifeTime
}

func (C *conn) IsActive() bool {
	if C.retreatLevel >= len(C.retreat) {
		C.retreatLevel = len(C.retreat) - 1
	}
	return C.lastActiveTime+C.retreat[C.retreatLevel] <= time.Now().Unix()
}

func (C *conn) Open() error {
	var err error

	if C.conn.IsOpen() {
		return nil
	}
	err = C.conn.Open()
	if err != nil {
		return err
	}
	return nil
}

func (C *conn) IsOpen() bool {
	return C.conn.IsOpen()
}

func (C *conn) SetTimeout(timeout time.Duration) {
	C.sock.SetTimeout(timeout)
}

func (C *conn) Close() {
	err := C.conn.Close()
	if err != nil {
		//todo 连接断开错误处理，记录日志
	}
}

type flumeClient struct {
	TransPort *conn
	Client    *flume.ThriftSourceProtocolClient
}

func newFlumeClient(conn *conn) *flumeClient {
	client := flumeClient{
		TransPort: conn,
		Client:    flume.NewThriftSourceProtocolClientFactory(conn.conn, thrift.NewTCompactProtocolFactory()),
	}
	return &client
}

func (FC *flumeClient) Append(event *flume.ThriftFlumeEvent) (r flume.Status, err error) {
	return FC.Client.Append(event)
}

func (FC *flumeClient) AppendBatch(events []*flume.ThriftFlumeEvent) (r flume.Status, err error) {
	return FC.Client.AppendBatch(events)
}

type connPool struct {
	pool      *queue
	hostIndex int
	//hostCount map[string]int //对连接进行计数需要加同步锁，降低大约20%的性能
	config *ConfigManager
	logger *log.Logger
	debug  bool
	count  int64
	patch  bool
}

func newConnPool(config *ConfigManager, log *log.Logger, debug bool) *connPool {
	timeOut := time.Millisecond * time.Duration(config.Conf.Control.getMaxWait())
	maxSize := config.Conf.Control.getPoolSize()
	pool := connPool{
		pool:      newQueue(maxSize, timeOut),
		hostIndex: 0,
		//hostCount: make(map[string]int),
		config: config,
		logger: log,
		debug:  debug,
		count:  0,
		patch:  true,
	}
	pool.patchConn()
	return &pool
}

//使用轮询的策略获取flume的节点
func (P *connPool) GetHost() string {
	hosts := P.config.GetHosts().GetValidNode()
	index := P.hostIndex % len(hosts)
	P.hostIndex++
	if P.hostIndex > 100 {
		P.hostIndex = 0
	}
	return hosts[index]
}

func (P *connPool) patchConn() {
	//P.logger.Println("run patch conn number is",P.count)
	for i := P.count; i < int64(P.config.Conf.Control.getConnMin()); i++ {
		conn := P.createConn()
		if conn == nil {
			continue
		}
		P.pool.Put(conn)
	}
}

func (P *connPool) createConn() *conn {
	host := P.GetHost()
	atomic.AddInt64(&P.count, 1)
	connect, err := newConn(host, 600, time.Millisecond*200)
	if err != nil {
		P.logger.Println(err)
		return nil
	}
	if P.debug {
		P.logger.Println("create a new connect ", host)
	}
	return connect
}

func (P *connPool) GetClient(timeOut time.Duration, create bool) (*flumeClient, bool) {
	var (
		connect *conn
		ok      bool
		err     error
		temp    interface{}
	)
	//P.logger.Println("send conn is ", P.count)
	if create {
		if P.count < int64(P.config.Conf.Control.getConnMin()) && P.patch {
			P.patch = false
			go func() {
				conn := P.createConn()
				if conn != nil {
					P.pool.Put(conn)
				}
				P.patch = true
			}()
		}
		temp, ok = P.pool.Get()
	} else {
		temp, ok = P.pool.GetNoWait()
	}

	if !ok && !create {
		return nil, false
	}

	if !ok {
		if P.count > int64(P.pool.maxSize) {
			return nil, false
		}
		connect = P.createConn()
		P.logger.Println("batch send create connect")
		if connect == nil {
			return nil, false
		}
	} else {
		connect = temp.(*conn)
		connect.SetTimeout(timeOut)
	}

	//连接超时，关闭连接
	if connect.IsOld() {
		atomic.AddInt64(&P.count, -1)
		connect.Close()
		return nil, false
	}
	//连接不活跃放回连接池
	if !connect.IsActive() {
		P.pool.Put(connect)
		return nil, false
	}
	//正常连接
	client := newFlumeClient(connect)
	err = connect.Open()
	//正常连接,打开失败
	if err != nil {
		P.logger.Println(err)
		connect.retreatLevel++
		P.pool.Put(connect)
		return nil, false
	}
	return client, true
}

func (P *connPool) PutClient(client *flumeClient) {
	ok := P.pool.Put(client.TransPort)
	if !ok {
		atomic.AddInt64(&P.count, -1)
		client.TransPort.Close()
	}
	client = nil
}

func (P *connPool) KillConnect(client *flumeClient) {
	atomic.AddInt64(&P.count, -1)
	client.TransPort.Close()
	client = nil
}

func (P *connPool) Close() {
	var (
		connect *conn
		ok      bool
		temp    interface{}
	)
	for !P.pool.Empty() {
		temp, ok = P.pool.Get()
		if !ok {
			continue
		}
		connect = temp.(*conn)
		atomic.AddInt64(&P.count, -1)
		connect.Close()
	}
	//释放指针
	P.logger = nil
	P.config = nil
	P.pool.Close()
}
