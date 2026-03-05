package param

const (
	ServerStart      = 1
	ServerEnd        = 401
	Area             = "weixin"
	DataPath         = "../../../foo-data/Template"
	DefaultRedisAddr = "192.168.1.63:6379"
	DefaultRedisPass = "jU766R7zT8MZec8J6MND"
	ZoneDBAddr       = "mongodb://192.168.1.63:27017"
	CreateTime       = "2025-08-19,00:00:00"
	LogLevel         = "info"
	Season           = 8
	SeasonTime       = "2025-08-19,00:00:00"
)

// lls
//var (
//	MergeID     = 38
//	GameRpcAddr = "ws://0.0.0.0:50301"
//	IP          = "10.22.51.64"
//	Port        = 50301
//	ZonePort    = 50401
//	ZoneDBName  = "foo_zone_lls2"
//	GameDB      = DBInfo{
//		Addr: "mongodb://192.168.1.63:27017",
//		Name: "foo_game_lls2",
//	}
//)

//// yc
//var (
//	MergeID     = 13
//	GameRpcAddr = "ws://0.0.0.0:40301"
//	IP          = "10.22.51.97"
//	Port        = 40301
//	ZonePort    = 40401
//	ZoneDBName  = "foo_zone_yc"
//	GameDB      = DBInfo{
//		Addr: "mongodb://192.168.1.63:27017",
//		Name: "foo_game_yc",
//	}
//)

// lh
var (
	MergeID     = 2
	GameRpcAddr = "ws://0.0.0.0:40301"
	IP          = "10.22.51.115"
	Port        = 40301
	ZonePort    = 40401
	ZoneDBName  = "foo_zone_lh"
	GameDB      = DBInfo{
		Addr: "mongodb://192.168.1.63:27017",
		Name: "foo_game_lh",
	}
)

type DBInfo struct {
	Addr string `json:"Addr"`
	Name string `json:"Name"`
}
