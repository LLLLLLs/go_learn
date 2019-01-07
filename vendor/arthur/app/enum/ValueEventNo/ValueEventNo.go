package ValueEventNo

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

/*===========================
======= TargetType: Role ======
===========================*/

const (
	Lv             Type = 100 //角色等级
	VIPLv          Type = 101 //VIP等级
	StudentSeat    Type = 102 //子嗣席位数量
	GoldPay        Type = 103 //充值获得的元宝
	InviteTimes    Type = 104 //单独邀约次数
	Gold           Type = 105 //总元宝
	Exp            Type = 106 //经验
	Silver         Type = 107 //银两
	Food           Type = 108 //食物
	Solider        Type = 109 //士兵
	TravelTimes    Type = 110 //寻访次数
	MarryTimes     Type = 111 //联姻次数
	CheckinTimes   Type = 113 //签到次数
	SoliderTimes   Type = 114 //征收士兵次数
	FoodTimes      Type = 115 //征收粮食次数
	SilverTimes    Type = 116 //征收银两次数
	PoliticalEvent Type = 117 //处理政务事件次数
	StoreBuyTimes  Type = 118 //商店购买次数
	Power          Type = 119 //角色势力
	StudentNum     Type = 120 //子嗣总数
	AllianceLv     Type = 121 //联盟等级
	ArenaScore     Type = 122 //竞技场分数
	DungeonScore   Type = 123 //副本分数
	TotalWishDays  Type = 125 //签到总天数
	EntrustTime    Type = 126 //委派次数
	Wise           Type = 127 //角色智慧属性
	PhaseSuccess   Type = 132 //关卡胜利次数
	Worship        Type = 133 //膜拜
	RandInvite     Type = 134 //随机邀约
	SpecInvite     Type = 135 //指定邀约
	ArenaFight     Type = 136 //竞技场出使
	InspireStu     Type = 137 //鼓励学员
	GraduateStu    Type = 139 //新增出坑学员
	PhaseProcess   Type = 140 //关卡进度

)

/*===========================
======= TargetType: Hero ======
===========================*/

const (
	HeroPeerage Type = 201 //英雄爵位
	HeroAttr    Type = 202 //英雄总属性
	HeroLv      Type = 203 //英雄等级
	HeroTalent  Type = 204 //英雄天赋增量升级

)

/*===========================
======= TargetType: Beauty ======
===========================*/

const (
	BeautyCharm Type = 300 //名媛魅力
	BeautyAmity Type = 301 //名媛友好度
	BeautyAward Type = 302 //赏赐名媛

)

/*===========================
======= TargetType: Item ======
===========================*/

const (
	ItemAdd      Type = 400 //道具增加个数
	ItemCompound Type = 401 //道具合成

)
