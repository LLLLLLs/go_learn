//stat 玩家升级配置表
//created: 2018/6/27
//author: wdj

package stat

type LvUpgrade struct {
	Lv                 int16
	Name               string
	Exp                int
	GainTimes          int
	Salary             int
	GovEventTimes      int
	GovAchievement     int
	GovAchievementRate float64
}
