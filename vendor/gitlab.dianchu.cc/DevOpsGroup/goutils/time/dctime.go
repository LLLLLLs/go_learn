package time

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/json-iterator/go"
)

const (
	//FMT 通用时间格式
	FMT = "2006-01-02 15:04:05"
	//MSFMT 游戏内带毫秒时间格式
	MSFMT   = "2006-01-02 15:04:05.000"
	DTFMT   = "2006-01-02"
	TIMEFMT = "15:04:05"
)

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

type Month = time.Month

const (
	January   Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday = time.Weekday

const (
	Sunday    Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var (
	TZ  = "Asia/Shanghai"
	Loc *Location
)

type Time time.Time
type OldTime = time.Time
type Duration = time.Duration
type Location = time.Location

func InitTimeutils(tz string) {
	TZ = tz
	loc, err := time.LoadLocation(TZ)
	if err != nil {
		panic(err)
	}
	Loc = loc
}

func Now(millLen ...int) Time {
	t := time.Now()
	nLen := 3
	if len(millLen) > 0 {
		if millLen[0] <= 0 || millLen[0] > 9 {
			panic(fmt.Sprintf("millLen:%d out of range", millLen))
		}
		nLen = millLen[0]
	}
	nSec := 0
	if nLen != 0 {
		divN := 10
		for i := 0; i < 8-nLen; i++ {
			divN *= 10
		}
		nSec = t.Nanosecond() / divN * divN
	}
	year, month, day := t.Date()
	h, m, s := t.Clock()
	dt := time.Date(year, month, day, h, m, s, nSec, Loc)
	return Time(dt)
}

// UTC returns t with the location set to UTC.
func (t Time) UTC() Time {
	return Time(OldTime(t).UTC())
}

// Local returns t with the location set to local time.
func (t Time) Local() Time {
	return Time(OldTime(t).Local())
}

// In returns t with the location information set to loc.
//
// In panics if loc is nil.
func (t Time) In(loc *Location) Time {
	return Time(OldTime(t).In(loc))
}

// Location returns the time zone information associated with t.
func (t Time) Location() *Location {
	return OldTime(t).Location()
}

// Zone computes the time zone in effect at time t, returning the abbreviated
// name of the zone (such as "CET") and its offset in seconds east of UTC.
func (t Time) Zone() (name string, offset int) {
	return OldTime(t).Zone()
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC.
func (t Time) Unix() int64 {
	return OldTime(t).Unix()
}

// UnixNano returns t as a Unix time, the number of nanoseconds elapsed
// since January 1, 1970 UTC. The result is undefined if the Unix time
// in nanoseconds cannot be represented by an int64 (a date before the year
// 1678 or after 2262). Note that this means the result of calling UnixNano
// on the zero Time is undefined.
func (t Time) UnixNano() int64 {
	return OldTime(t).UnixNano()
}

func (t Time) Add(d Duration) Time {
	return Time(OldTime(t).Add(d))
}

func (t Time) Sub(u Time) Duration {
	return OldTime(t).Sub(OldTime(u))
}

func (t Time) AddDate(years, months, days int) Time {
	return Time(OldTime(t).AddDate(years, months, days))
}

func (t Time) Date() (year int, month Month, day int) {
	return OldTime(t).Date()
}
func (t Time) Year() int {
	return OldTime(t).Year()
}

// Month returns the month of the year specified by t.
func (t Time) Month() Month {
	return OldTime(t).Month()
}

// Day returns the day of the month specified by t.
func (t Time) Day() int {
	return OldTime(t).Day()
}

// Weekday returns the day of the week specified by t.
func (t Time) Weekday() Weekday {
	return OldTime(t).Weekday()
}

func (t Time) ISOWeek() (year, week int) {
	return OldTime(t).ISOWeek()
}

// Clock returns the hour, minute, and second within the day specified by t.
func (t Time) Clock() (hour, min, sec int) {
	return OldTime(t).Clock()
}

// Hour returns the hour within the day specified by t, in the range [0, 23].
func (t Time) Hour() int {
	return OldTime(t).Hour()
}

// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
func (t Time) Minute() int {
	return OldTime(t).Minute()
}

// Second returns the second offset within the minute specified by t, in the range [0, 59].
func (t Time) Second() int {
	return OldTime(t).Second()
}

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (t Time) Nanosecond() int {
	return OldTime(t).Nanosecond()
}

// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
func (t Time) YearDay() int {
	return OldTime(t).YearDay()
}

func (t Time) Truncate(d Duration) Time {
	return Time(OldTime(t).Truncate(d))
}

func (t Time) Round(d Duration) Time {
	return Time(OldTime(t).Round(d))
}

func (t Time) Format(layout string) string {
	return OldTime(t).Format(layout)
}

func (t Time) Equal(u Time) bool {
	return OldTime(t).Equal(OldTime(u))
}

// Before reports whether the time instant t is before u.
func (t Time) Before(u Time) bool {
	return OldTime(t).Before(OldTime(u))
}

func (t Time) After(u Time) bool {
	return OldTime(t).After(OldTime(u))
}

// Since returns the time elapsed since t.
// It is shorthand for time.Now().Sub(t).
func Since(t Time) Duration {
	return Now().Sub(t)
}

// Until returns the duration until t.
// It is shorthand for t.Sub(time.Now()).
func Until(t Time) Duration {
	return t.Sub(Now())
}

func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time {
	return Time(time.Date(year, month, day, hour, min, sec, nsec, loc))
}

func Unix(sec int64, nsec int64) Time {
	return Time(time.Unix(sec, nsec))
}

func ParseInLocation(layout, value string, loc *Location) (Time, error) {
	dt, err := time.ParseInLocation(layout, value, loc)
	return Time(dt), err
}

//自定义参数

func (t Time) String() string {
	return Time2Str(t)
}

func (t Time) MSString() string {
	return Time2StrWithMS(t)
}

func (t Time) Seconds() float64 {
	return Time2Secs(t)
}

func (t Time) SecondStr() string {
	return fmt.Sprintf("%.3f", Time2Secs(t))
}

//从字符串获取时间
func ParseDCTimeStr(str string) Time {
	t := Str2Time(str)
	return Time(t)
}

//从时间戳/秒获取时间
func ParseDCTimeSec(sec float64) Time {
	t := Sec2Time(sec)
	return Time(t)
}

func init() {
	//自定义类型解析
	jsoniter.RegisterTypeEncoderFunc("time.Time", func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
		t := (*Time)(ptr)
		stream.WriteString(t.String())
	}, nil)

	jsoniter.RegisterTypeDecoderFunc("time.Time", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		tt := iter.ReadString()
		if tt == "" {
			*((*Time)(ptr)) = Time{}
			return
		}
		*((*Time)(ptr)) = ParseDCTimeStr(tt)
	})

}
