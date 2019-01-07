/*
Author : 洪坤安
Time   : 2018/4/25
Desc   : 静态数据相关的变量
*/

package model

import (
	"arthur/app/info/errors"
	"arthur/app/infra"
	"arthur/conf"
	"arthur/env"
	"arthur/utils/log"
	"arthur/utils/redlock"
	"arthur/utils/struitls"
	"arthur/utils/timeutils"
	"encoding/gob"
	"github.com/bsm/redis-lock"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"io"
	"os"
	"path/filepath"
	"reflect"
)

const cacheFileExpreTime = 1800

func LoadDynTable() {
	infra.RegisterFactory(CharacterSet...)
}

func InitStatKey(m ProfileCenter) {
	e := m.KeyValue
	for _, r := range e {
		KeyValue[r.Key] = r.Value
	}
}
func LoadStatTableInLocal(d conf.DSN, center *ProfileCenter) error {
	projectAddr := env.ProjectRoot()
	fileAddr := filepath.Join(projectAddr, env.StatCache_Name)
	machineTag := GetMachineTag()
	redLock := redlock.NewLock(machineTag, nil)
	ok, err := redLock.Lock()
	if ok && err == nil { //只有获取到锁的程序才可以释放锁
		defer redLock.Unlock()
	}
	//如果不是等待超时错误
	if err != nil && err != lock.ErrLockNotObtained {
		log.Errorf("redLock err:%s", err.Error())
		return err
	}
	if checkFileIsExist(fileAddr) && !isFileExpired(fileAddr) {
		log.Info("Load Stat Table From Local")
		f, err := os.Open(fileAddr)
		if err != nil {
			log.Errorf("Open err:%s", err.Error())
			return err
		}
		defer f.Close()
		dec := gob.NewDecoder(f)
		err = dec.Decode(center)
		if err != nil && err == io.EOF { //如果无法解析文件，则加载静态表
			if !LoadStatTable(d, center) {
				return errors.New("load stat table failed")
			}
		} else if err != nil && err != io.EOF {
			log.Errorf("Decode err:%s", err.Error())
			return err
		}
	} else {
		log.Info("Load Stat Table From SQL")
		if !LoadStatTable(d, center) {
			return errors.New("load stat table failed")
		}
		return saveProfCenter(*center, fileAddr)
	}
	return nil
}

/* 执行LoadTable后，会从数据库中取出表内所有记录，并根据表名写入到center内对应的field中,
之后可使用LINQ来查询[]Table01Struct中的记录

示例：

	type Book struct{
		Title string
		Contents string
		PageCount	int
	}

	type Center struct{
		Book		[]Book
		Pen			[]Pen
	}

	center = new(Center)
	LoadTable(db, center)
	record := linq.From(center.Book).FirstWith(func(x interface{}) bool{
			row := x.(*Book)
			return row.Name == "Norwegian Wood"
		}
	)
*/
func LoadStatTable(d conf.DSN, center interface{}) bool {
	conn, err := xorm.NewEngine(infra.DriverName, string(d))
	if err != nil {
		log.Errorf("NewEngine err:%s", err.Error())
		return false
	}
	defer conn.Close()
	//TODO：最大连接、可空闲连接后期可能用ZK配置
	maxIdleConns := 10
	maxOpenConns := 10

	conn.DB().SetMaxIdleConns(maxIdleConns)
	conn.DB().SetMaxOpenConns(maxOpenConns)
	conn.SetTableMapper(core.SnakeMapper{})
	conn.SetColumnMapper(core.SnakeMapper{})

	centerValue := reflect.ValueOf(center).Elem()
	centerType := centerValue.Type()

	ok := true

	// 初始化每个成员变量，并加载静态数据
	for i := 0; i < centerValue.NumField(); i++ {
		field := centerValue.Field(i)

		// 只关注切片类型的成员变量
		if field.Kind() != reflect.Slice {
			continue
		}

		structField := centerType.Field(i)

		// 检查：成员变量首字母必须大写（否则无法被动态赋值）
		if !field.CanSet() {
			log.Error("error! field[", structField.Name, "] the first alphabet must be uppercase")
			return false
		}

		// 初始化成员变量（相当于这样初始化：xxxx = make([]TcXxxxx, 0)）
		sliceValue := reflect.MakeSlice(field.Type(), 0, 1)

		x := reflect.New(sliceValue.Type())
		x.Elem().Set(sliceValue)

		slice := x.Interface()

		// xorm方式加载静态数据到成员变量中
		err := conn.Find(slice)
		if err != nil {
			log.Error("error! load xorm table [", field.Type().String(), "] failed!")
			log.Error("-----> ", err)

			ok = false
		}
		// 赋值给成员变量
		field.Set(reflect.ValueOf(slice).Elem())
	}
	return ok
}

func saveProfCenter(center ProfileCenter, fileAddr string) error {
	f, err := os.Create(fileAddr)
	if err != nil {
		log.Errorf("Create err:%s", err.Error())
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	err = enc.Encode(center)
	if err != nil {
		log.Errorf("Encode err:%s", err.Error())
		return err
	}
	return nil
}

//判断文件是否存在
func checkFileIsExist(fileAddr string) bool {
	var exist = true
	if _, err := os.Stat(fileAddr); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func isFileExpired(fileAddr string) bool {
	fileInfo, err := os.Stat(fileAddr)
	if err != nil {
		return true
	}
	modTime := fileInfo.ModTime().Unix()
	now := timeutils.Now()
	return now-modTime > cacheFileExpreTime
}

func GetMachineTag() string {
	machineTag, err := struitls.MachineTag()
	if err != nil {
		//TODO:如果获取本机TAG失败，则写死TAG，后期TAG可能配置在其他地方
		machineTag = "arthur_stat_table_tag"
	}
	return machineTag
}
