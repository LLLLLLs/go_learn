package cache

import (
	"bytes"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"

	"gitlab.dianchu.cc/go_chaos/fort/utils"
)

const (
	versionKey = "version"
	dataKey    = "data"
	idsqlKey   = "sql"
	lock       = -1
)

//CacheAsidePattern:预留缓存模式
//缓存内容:gob struct
//写:直接写数据库,并且淘汰缓存（更新缓存版本号）,更新期间缓存被锁定,直到数据写入数据才解锁缓存
//读:缓存版本号与缓存数据版本号一致,并且缓存未被锁定，则读缓存。否则读数据库更新缓存(缓存被锁定不更新缓存)。
type CacheAsidePattern struct {
	expiration time.Duration //缓存自动淘汰时间,0则永不淘汰
	rc         *utils.RedisManager
}





//获取数据版本号的Key
func (c *CacheAsidePattern) genKey(database, table, pkVal string, tag string) string {
	var key bytes.Buffer
	key.WriteString(tag)
	key.WriteString(":")
	key.WriteString(utils.MD5Hash(database))
	key.WriteString(":")
	key.WriteString(table)
	key.WriteString(":")
	key.WriteString(pkVal)
	return key.String()
}

func (c *CacheAsidePattern) genIdSQLKey(idSQL string) string {
	var key bytes.Buffer
	key.WriteString(idsqlKey)
	key.WriteString(":")
	key.WriteString(utils.MD5Hash(idSQL))
	return key.String()
}

func (c *CacheAsidePattern) SetExpiration(exp time.Duration) {
	c.expiration = exp
}

func (c *CacheAsidePattern) SetRedisClient(client *utils.RedisManager) {
	c.rc = client
}

//更新版本号
//todo 如果有cacher,所有写操作都要去更新数据库.是否要把cacher改为全局?
//更新期间缓存被锁定,直到数据写入数据才解锁缓存。缓存锁定期间,直接读数据库。锁会在3秒内过期,3秒后没解锁,删除缓存记录。
func (c *CacheAsidePattern) UpdateCacheVersion(database, table string, pkVal interface{}, lock bool) error {
	var (
		pkStr string
		err   error
	)
	if pkStr, err = utils.InterfaceToString(pkVal); err != nil {
		return err
	}
	key := c.genKey(database, table, pkStr, versionKey)
	if lock {
		return c.rc.Set(key, lock, 3*time.Second) //防止高并发情况下生成脏数据
	}
	return c.rc.Set(key, time.Now().UnixNano(), c.expiration)
}

func (c *CacheAsidePattern) GetCache(database, table string, pkVals []interface{}, model interface{}) (bool, int64, error) {
	var (
		exist bool
		err   error
	)
	switch pkLen := len(pkVals); pkLen {
	case 0: // 不支持的语句
		return false, 0, nil
	case 1: // 单结果
		exist, err = c.getLineCache(database, table, pkVals[0], model)
		return exist, int64(pkLen), err
	default: // 多结果
		slice := reflect.MakeSlice(reflect.TypeOf(model).Elem(), pkLen, pkLen)
		modelValue := reflect.ValueOf(model)
		for i := 0; i < pkLen; i++ {
			mValue := reflect.New(slice.Index(i).Type())
			if exist, err = c.getLineCache(database, table, pkVals[i], mValue); err != nil || !exist {
				return exist, int64(pkLen), err
			}
			slice.Index(i).Set(mValue.Elem())
		}
		modelValue.Elem().Set(reflect.AppendSlice(modelValue.Elem(), slice))
		return true, int64(pkLen), nil
	}
}

func (c *CacheAsidePattern) UpdateCache(database, table string, pkVals []interface{}, data []byte) error {
	var err error
	for _, p := range pkVals {
		if e := c.updateLineCache(database, table, p, data); e != nil {
			err = e
		}
	}
	return err

}

// 读取行级缓存 返回:(缓存是否有效,缓存查询结果数量,错误)
// 错误返回:Redis SDK产生的错误 \ 缓存结果反序列化失败错误（model类型错误）\ 主键值类型错误 \ 版本号不存在
// 其他错误(缓存内容错误)不返回，统一视为缓存失效，然后重新生成
func (c *CacheAsidePattern) getLineCache(database, table string, pkVal interface{}, model interface{}) (bool, error) {
	var (
		err      error
		exist    bool
		data     []byte
		dataVer  int64 //数据版本号
		cacheVer int64 //缓存版本号
		pkValStr string
	)

	// 主键值类型不支持
	if pkValStr, err = utils.InterfaceToString(pkVal); err != nil {
		return false, err
	}

	// 根据 数据库、表名、ID 获取[数据版本]
	if exist, data, err = c.rc.Get(c.genKey(database, table, pkValStr, versionKey)); !exist {
		if err == nil {
			return false, errors.New("The version does not exist ")
		}
		return false, err
	}

	// 检查[数据版本是否被锁定]
	if dataVer, err = strconv.ParseInt(string(data), 10, 64); err != nil {
		return false, nil
	}
	if dataVer == lock {
		return false, errors.New("The cache is locked ")
	}

	// 根据 数据库、表名、ID 获取[:8]缓存版本号,[8:]缓存
	if exist, data, err = c.rc.Get(c.genKey(database, table, pkValStr, dataKey)); !exist || len(data) < 8 {
		return false, err
	}

	// [:8]缓存版本号
	if cacheVer, err = utils.BytesToInt64(data[:8]); err != nil || cacheVer < dataVer {
		return false, nil
	}

	// [8:]缓存
	// 缓存结果反序列化失败,返回错误.用于检查模型是否传错
	if err = utils.Unmarshal(data[8:], model); err != nil {
		return false, err
	}
	return true, nil
}

// 更新行级缓存
// 参数:库名,表名,主键[值],Struct gob data ,查询结果条数
// redist储存分区:[:7]版本号,[8:15]查询结果条数,[16:]Struct gob data
func (c *CacheAsidePattern) updateLineCache(database, table string, pkVal interface{}, data []byte) error {
	var (
		putData []byte
		pkStr   string
		err     error
	)
	if pkStr, err = utils.InterfaceToString(pkVal); err != nil {
		return err
	}
	key := c.genKey(database, table, pkStr, dataKey)
	//版本号
	putData = utils.Int64ToBytes(time.Now().UnixNano())
	//Struct gob data
	putData = append(putData, data...)
	return c.rc.Set(key, putData, c.expiration)
}




//更新idSQL语句内容
func (c *CacheAsidePattern) updateSQLInfo(idSQL, database, table string, pkVals []interface{}) error {
	var (
		sqlInfoKey string //用于获取语句设计的表与ID
		pkStr      string
		err        error
		pkStrs     bytes.Buffer
	)
	sqlInfoKey = c.genIdSQLKey(idSQL)
	for i, p := range pkVals {
		if pkStr, err = utils.InterfaceToString(p); err != nil {
			return err
		}
		pkStrs.WriteString(pkStr)
		if i < len(pkVals)+1{
			pkStrs.WriteString(",")
		}
	}
	//todo 错误处理
	c.rc.HSet(sqlInfoKey, "database", database)
	c.rc.HSet(sqlInfoKey, "table", table)
	c.rc.HSet(sqlInfoKey, "ids", pkStrs.String())
	return nil
}

func (c *CacheAsidePattern) getSQLInfo(idSQL string) (exist bool, database, table string, ids []string) {
	var (
		id         string
		sqlInfoKey string //用于获取语句设计的表与ID
	)
	//根据语句hash获取语句的信息:数据库、表名、ID
	sqlInfoKey = c.genIdSQLKey(idsqlKey)
	if exist, database = c.rc.HGet(sqlInfoKey, "database"); !exist {
		return
	}
	if exist, table = c.rc.HGet(sqlInfoKey, "table"); !exist {
		return
	}
	if exist, id = c.rc.HGet(sqlInfoKey, "ids"); !exist {
		return
	}
	if ids = strings.Split(id, ","); len(ids) == 0 {
		return
	}
	return
}
