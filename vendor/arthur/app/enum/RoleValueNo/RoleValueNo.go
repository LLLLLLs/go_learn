package RoleValueNo

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	//空
	Emtpy Type = 0
	//角色等级
	Lv Type = 1
	//VIP等级
	VIPLv Type = 2
	//当前穿戴时装编号
	FashionNo Type = 3

	//充值获得的元宝
	GoldPay Type = 5
	//最后登录时间
	LastLoginTime Type = 6
	//注册时间
	RegisterTime Type = 7
	//红颜活力时间
	InviteTime Type = 8
	//单独邀约次数
	InviteTimes Type = 12
	// 征收事件进度
	LevyEventProgress Type = 13

	// 许愿池许愿时间
	WishTime Type = 14

	//当前穿戴称号编号
	TitleNo Type = 15

	//竞技币
	ArenaCoin Type = 16
	//非第一次解锁的剑
	UnlockedSword Type = 17
	//关卡是否需要提示解锁
	HintSword Type = 19

	// 竞技场1v1今日次数
	ArenaSoloRandomTimes = 18
	//聊天每日领取奖励次数
	DailyNum Type = 20
	//当天领取奖励的时间
	DailyTime Type = 22

	// 签到轮次
	WishPeriodRound Type = 21

	// 单人副本当前关卡数
	DungeonSingleLevel Type = 23
	// 单人副本当前状态 0=探索 1=战斗
	DungeonSingleStat Type = 24
	// 龙币
	DungeonCoin Type = 25

	//总元宝数
	Gold Type = 100
	//经验
	Exp Type = 101

	//银两
	Silver Type = 201
	//食物
	Food Type = 202
	//士兵
	Solider Type = 203
	//食物时间
	FoodTime Type = 211
	//银两时间
	SilverTime Type = 212
	//士兵时间
	SoliderTime Type = 213
	// 关卡-章数
	Chapter Type = 301
	// 关卡-
	Section Type = 302
	// 关卡-关数
	Phase Type = 303
	// 活跃度
	Activity Type = 304
	// 当前主线任务ID
	CurMainTask Type = 311
	// 当前主线任务进度
	CurMainTaskProgress Type = 312
	// 总许愿（签到）天数
	TotalWishDays Type = 500
)

//总计类
const (
	ArenaScore   Type = 2003 //竞技场分数
	DungeonScore Type = 2004 //副本分数

	Wise     Type = 3001 //智慧
	Diligent Type = 3002 //勤勉
	Loyalty  Type = 3003 //忠诚
	Heroic   Type = 3004 //英勇
)
