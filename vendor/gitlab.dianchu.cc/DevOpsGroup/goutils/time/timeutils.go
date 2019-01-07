package time

import "time"

func init() {
	InitTimeutils(TZ)
}

func Str2Time(strtime string) Time {
	tLen := len(strtime)
	timefmt := FMT
	if tLen == 23 {
		timefmt = MSFMT
	} else if tLen == 10 {
		timefmt = DTFMT
	} else if tLen == 8 {
		timefmt = TIMEFMT
	}

	dt, ok := ParseInLocation(timefmt, strtime, Loc)
	if ok != nil {
		panic(ok)
	}
	return dt
}

func NowSecs() float64 {
	dt := Now()
	return Time2Secs(dt)
}

func Sleep(d Duration) {
	time.Sleep(d)
}

//func Today() string {
//	//timeUnix := time.Now().Unix() //已知的时间戳
//	formatTimeStr := Now().Format(FMT)
//	return formatTimeStr
//}
//
//func Tomorrow() string {
//	//oneday, _ := time.ParseDuration("24h")
//	//timeUnix := time.Now().Add(oneday).Unix() //已知的时间戳
//	t := Now()
//	t.AddDate(0, 0, 1)
//	formatTimeStr := t.Format(FMT)
//	return formatTimeStr
//}

func Sec2Time(secs float64) Time {
	secInt := int64(secs)
	secFloat := secs*1000 - float64(secInt*1000)
	t := Unix(secInt, int64(secFloat*1000000))
	return ConvTimeZone(t)
}

func ConvTimeZone(t Time) Time {
	year, month, day := t.Date()
	h, m, s := t.Clock()
	dt := Date(year, month, day, h, m, s, t.Nanosecond(), Loc)
	return dt
}

func Time2Secs(dt Time) float64 {
	//默认保留3位精度
	return float64(dt.UnixNano()/1000000) / 1000
}

func Time2Str(dt Time) string {
	return dt.Format(FMT)
}

func Time2StrWithMS(dt Time) string {
	return dt.Format(MSFMT)
}

func Str2sec(strtime string) float64 {
	dt := Str2Time(strtime)
	return Time2Secs(dt)
}

func Sec2Str(secs float64) string {
	dt := Sec2Time(secs)
	return Time2Str(dt)
}

func Sec2StrWithMS(secs float64) string {
	dt := Sec2Time(secs)
	return Time2StrWithMS(dt)
}
