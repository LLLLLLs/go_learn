//@time:2020/04/07
//@desc:

package calc_business_days

import "time"

func CalcBusinessDays(start, end time.Time) int {
	start = DayStart(start)
	end = DayStart(end)
	if start.After(end) || start.Equal(end) {
		return 0
	}
	offDays := int(end.Sub(start).Hours()) / 24
	weeks := offDays / 7
	daysRemain := offDays % 7
	offSaturday := convertWeekday(time.Saturday) - convertWeekday(start.Weekday())
	if offSaturday > 0 { // monday to friday
		if daysRemain-offSaturday == 1 || daysRemain-offSaturday == 2 {
			daysRemain = offSaturday
		}
		if daysRemain-offSaturday > 2 {
			daysRemain -= 2
		}
	} else { // saturday or sunday
		daysRemain -= 2 + offSaturday
		if daysRemain < 0 {
			daysRemain = 0
		}
		if daysRemain > 5 {
			daysRemain = 5
		}
	}
	return weeks*5 + daysRemain
}

func convertWeekday(wd time.Weekday) int {
	if wd == time.Sunday {
		return 7
	}
	return int(wd)
}

func DayStart(t time.Time) time.Time {
	return time.Unix(t.Unix()-(t.Unix()%int64(24*time.Hour/time.Second)), 0)
}
