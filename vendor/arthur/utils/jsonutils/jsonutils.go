/*
Author : Haoyuan Liu
Doc	   : 对json-iterator库的封装
*/
package json

import (
	"arthur/utils/panicutils"
	"github.com/json-iterator/go"
	"time"
	"unsafe"
)

var (
	//Config 默认配置
	Config = jsoniter.ConfigFastest
	//TZ 时区
	TZ = "Asia/Shanghai"
	//Loc 时区对象
	Loc *time.Location
)

const (
	//FMT 通用时间格式
	FMT = "2006-01-02 15:04:05"
	//MSFMT 游戏内带毫秒时间格式
	MSFMT = "2006-01-02 15:04:05.000"
	//DTFMT 短日期格式
	DTFMT = "2006-01-02"
	//TIMEFMT 短时间格式
	TIMEFMT = "15:04:05"
)

type Any = jsoniter.Any

func Get(data []byte, path ...interface{}) Any {
	return Config.Get(data, path...)
}

//Exist 校验Any对象中是否有键
func Exist(any Any, key string) bool {
	for _, v := range any.Keys() {
		if v == key {
			return true
		}
	}
	return false
}

//Format 序列化
func Marshal(v interface{}) []byte {
	b, err := Config.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

//MarshalWithErr
func MarshalWithErr(v interface{}) ([]byte, error) {
	return Config.Marshal(v)
}

//MarshalToString 直接序列化成string
func MarshalToString(v interface{}) string {
	s, err := Config.MarshalToString(v)
	panicutils.OkOrPanic(err)
	return s
}

//Unmarshal 反序列化
func Unmarshal(data []byte, v interface{}) error {
	return Config.Unmarshal(data, v)
}

//UnmarshalFromString 直接从字符串反序列化
func UnmarshalFromString(str string, v interface{}) error {
	return Config.UnmarshalFromString(str, v)
}

func UnmarshalFromInterface(origin, target interface{}) error {
	data, err := MarshalWithErr(origin)
	if err != nil {
		return err
	}
	return Unmarshal(data, target)
}

func init() {
	loc, err := time.LoadLocation(TZ)
	panicutils.OkOrPanic(err)
	Loc = loc

	//自定义类型解析
	jsoniter.RegisterTypeEncoderFunc("time.Time", func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
		t := (*time.Time)(ptr)
		stream.WriteString(t.Format("2006-01-02 15:04:05"))
	}, nil)

	jsoniter.RegisterTypeDecoderFunc("time.Time", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		tt := iter.ReadString()
		if tt == "" {
			*((*time.Time)(ptr)) = time.Time{}.In(Loc)
			return
		}
		tLen := len(tt)
		timefmt := FMT
		if tLen == 23 {
			timefmt = MSFMT
		} else if tLen == 10 {
			timefmt = DTFMT
		} else if tLen == 8 {
			timefmt = TIMEFMT
		}

		dt, err := time.ParseInLocation(timefmt, tt, Loc)
		if err != nil {
			panic(err)
		}
		*((*time.Time)(ptr)) = dt
	})
}
