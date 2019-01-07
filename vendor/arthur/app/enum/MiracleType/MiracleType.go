package MiracleType

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	LevySilver       Type = 1 // 征税10倍神迹
	HeroStudy        Type = 2 // 英雄学习三倍收益
	HeroUpgrade      Type = 3 // 英雄升级连升三级
	Worship          Type = 4 // 膜拜排行榜
	BeautyInvitation Type = 5 // 红颜邀约
	StudentInspire   Type = 6 // 学员鼓励
)
