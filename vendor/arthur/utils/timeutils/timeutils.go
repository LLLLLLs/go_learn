/*
Author : Haoyuan Liu
Doc	   : 时间工具包，整个应用内，不直接使用time包获取时间对象，全部用精确到秒的
		 Unix时间戳来表示时间（int64），不论是接口请求返回的字段、数据库存储的
 		 字段，均用时间戳，这样服务器就不用关心时区问题，前端根据自己的需要，格式化
		 成当地时间。
*/
package timeutils

import (
	"arthur/app/info/errors"
	"arthur/conf"
	"fmt"
	"strconv"
	"time"
)

const (
	OneDaySeconds = 86400
)

//与格林威治相差的秒速
func TimeZoneDelta() int64 {
	return int64(conf.Config.Server.TimeZone) * 60 * 60
}

//获取当前配置的时区
func Location() *time.Location {
	return time.FixedZone("CST", int(TimeZoneDelta()))
}

//获取当前时间戳
func Now() int64 {
	return time.Now().Unix()
}

//昨天的这个时间点
func Yesterday() int64 {
	dt := Now()
	return AddDays(dt, -1)
}

//设置时间戳当天的小时,分钟,秒
//	例：SetTime(Now(),0,0,0)将获得当地今天零点的时间戳
func SetTime(now int64, hour, min, sec int) int64 {
	delta := TimeZoneDelta()
	day := (now + delta) / OneDaySeconds
	return day*OneDaySeconds - delta + int64(hour*60*60+min*60+sec)
}

//获取上次时间距离现在相差秒数
func Since(lastTime int64) int64 {
	return Now() - lastTime
}

//判断某个日期是否在今天之前
func IsPassDay(dt int64) bool {
	dateStart := DateStartTime(Now())
	if dt < dateStart {
		return true
	} else {
		return false
	}
}

//获取相差的天数-与另一个时间比较
func BetweenDays(newer int64, older int64) int64 {
	delta := TimeZoneDelta()
	dayFirst := (newer + delta) / OneDaySeconds
	daySecond := (older + delta) / OneDaySeconds
	return dayFirst - daySecond

}

//增加时间---增加天数
func AddDays(dt int64, days int64) int64 {
	addSeconds := days * 24 * 60 * 60
	return dt + addSeconds
}

//增加时间---增加小时数
func AddHours(dt int64, hours int64) int64 {
	addSeconds := hours * 60 * 60
	return dt + addSeconds
}

//增加时间---增加分钟数
func AddMinutes(dt int64, minutes int64) int64 {
	addSeconds := minutes * 60
	return dt + addSeconds
}

//增加时间---增加秒数
func AddSeconds(dt int64, seconds int64) int64 {
	return dt + seconds
}

//获取当天的最早时间
func DateStartTime(today int64) int64 {
	return today - DateStartTimePass(today)
}

//距离今天的零点过去了多少秒
func DateStartTimePass(t int64) int64 {
	return (t + TimeZoneDelta()) % OneDaySeconds
}

//获取当天的最晚时间
func DateEndTime(today int64) int64 {
	return DateStartTime(today) + OneDaySeconds - 1
}

func TimestampToTime(dt int64) time.Time {
	t := time.Unix(dt, 0)
	return t
}

//获取月和日组成的四位整数, 格式为 "0102"
func GetMonthDay(t int64) int {
	tLocal := time.Unix(t, 0).In(Location())
	_, month, day := tLocal.Date()
	return int(month)*1e2 + day
}

func GetMonthDayStr(t int64) (string, error) {
	return MonthDay2Str(GetMonthDay(t))
}

func MonthDay2Str(t int) (string, error) {
	s := strconv.Itoa(t)
	l := len(s)
	if l < 4 {
		s = fmt.Sprintf("%04s", s)
	} else if l > 4 {
		return s, errors.New("int must small than 10000")
	}
	return s, nil
}
