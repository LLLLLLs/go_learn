package conf

import (
	"arthur/sdk/dclog"
	"arthur/sdk/zookeeper"
	"arthur/utils/jsonutils"
	"path"
)

var (
	Config      ServerConfig
	nodeName    = "config"
	watchedNode zookeeper.INodeWatched
)

func Init(zkRoot string) {
	nodePath := path.Join(zkRoot, nodeName)
	loadFromZk(nodePath)
}

func loadFromZk(zkPath string) {
	var initData []byte
	watchedNode, initData = zookeeper.StartWatch(zookeeper.NewNode(zkPath), updateConfig)
	updateConfig(initData)
}

func updateConfig(data []byte) {
	// 反序列化zk节点值，并载入
	c := new(ServerConfig)
	err := json.Unmarshal(data, c)
	if err == nil {
		Config = *c
	} else {
		dclog.Error("unmarshal server config error", "sys", "", nil, "")
	}
}

// 游戏服务配置
type ServerConfig struct {
	Server struct {
		GameId   string
		HttpAddr string
		GrpcAddr string
		TimeZone int
	}

	Database struct {
		DBConn    dbConn
		DalZkHost []string
		DalZkAuth []string
		DalZkPath string
		Charset   string
		CenterDB  string

		ProfileDB string
	}

	Redis RedisConf

	Push       thirdPartyServer
	Session    thirdPartyServer
	Uuid       thirdPartyServer
	WordFilter thirdPartyServer

	Mail struct {
		thirdPartyServer
		IsPush  bool
		Version string
	}

	Chat struct {
		thirdPartyServer
		IsPush    bool
		PageLimit int
		Version   string
	}

	BehaviorLog struct {
		DBConn struct {
			dbConn
			Database string
		}
		ZKConn struct {
			Servers   []string
			FlumePath string
			ConfPath  string
		}
	}

	DCLog struct {
		Host string
		Port int
	}

	MsgClassifier struct {
		HttpAddr string
		Timeout  int
		Version  string
	}
}

// 第三方服务的连接信息
type thirdPartyServer struct {
	Method   string
	HttpAddr string
	GrpcAddr string
	Timeout  int
	PoolSize int
}

// sql数据库连接信息
type dbConn struct {
	Host     string
	Port     int
	User     string
	Password string
}

type RedisConf struct {
	Host        string
	Port        int
	Password    string
	RWTimeout   int //读写超时（秒）
	ConnTimeout int //连接超时（秒）
	MaxIdle     int //最大待机连接数
	MaxActive   int //最大活跃连接数
}
