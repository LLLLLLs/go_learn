/*
Author : Haoyuan Liu
Time   : 2018/4/20
*/
package conf

import (
	"arthur/utils/osutils"
	"fmt"
)

//DataSourceName 数据库连接信息
type DSN string

func CenterUri() DSN {
	dbName := Config.Database.CenterDB
	return GetUri(dbName)
}

func ProfileDbUri() DSN {
	return GetUri(Config.Database.ProfileDB)
}

func GetUri(db string) DSN {
	user := osutils.GetEnvWithDefault("DB_USER", Config.Database.DBConn.User)
	password := osutils.GetEnvWithDefault("DB_PASSWORD", Config.Database.DBConn.Password)
	s := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		user,
		password,
		Config.Database.DBConn.Host,
		Config.Database.DBConn.Port,
		db,
		Config.Database.Charset,
	)
	return DSN(s)
}
