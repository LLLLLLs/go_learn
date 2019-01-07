/*
author: ZhongXH
date: 2017/12/25
src:https://github.com/jinzhu/now/blob/master/now.go
*/
package time

const fixEndNanoSecond = 1000000 * Nanosecond

func (t *Time) BeginningOfMinute() Time {
	return t.Truncate(Minute)
}

func (t *Time) BeginningOfHour() Time {
	return t.Truncate(Hour)
}

func (t *Time) BeginningOfDay() Time {
	d := Duration(-t.Hour()) * Hour
	return t.BeginningOfHour().Add(d)
}

func (t Time) BeginningOfWeek() Time {
	dt := t.BeginningOfDay()
	weekday := int(t.Weekday())
	//if FirstDayMonday {
	//	if weekday == 0 {
	//		weekday = 7
	//	}
	//	weekday = weekday - 1
	//}

	d := Duration(-weekday) * 24 * Hour
	return dt.Add(d)
}

func (t Time) BeginningOfMonth() Time {
	dt := t.BeginningOfDay()
	d := Duration(-int(t.Day())+1) * 24 * Hour
	return dt.Add(d)
}

func (t *Time) BeginningOfQuarter() Time {
	month := t.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

func (t *Time) BeginningOfYear() Time {
	dt := t.BeginningOfDay()
	d := Duration(-int(t.YearDay())+1) * 24 * Hour
	return dt.Truncate(Hour).Add(d)
}

func (t *Time) EndOfMinute() Time {
	return t.BeginningOfMinute().Add(Minute - fixEndNanoSecond)
}

func (t *Time) EndOfHour() Time {
	return t.BeginningOfHour().Add(Hour - fixEndNanoSecond)
}

func (t *Time) EndOfDay() Time {
	return t.BeginningOfDay().Add(24*Hour - fixEndNanoSecond)
}

func (t *Time) EndOfWeek() Time {
	return t.BeginningOfWeek().AddDate(0, 0, 7).Add(-fixEndNanoSecond)
}

func (t *Time) EndOfMonth() Time {
	return t.BeginningOfMonth().AddDate(0, 1, 0).Add(-fixEndNanoSecond)
}

func (t *Time) EndOfQuarter() Time {
	return t.BeginningOfQuarter().AddDate(0, 3, 0).Add(-fixEndNanoSecond)
}

func (t *Time) EndOfYear() Time {
	return t.BeginningOfYear().AddDate(1, 0, 0).Add(-fixEndNanoSecond)
}
