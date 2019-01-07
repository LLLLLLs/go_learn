package RankType

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	//角色势力
	Power Type = 1
	//远征关卡数
	Phase Type = 2
	//角色总友好度
	TotalFriendly Type = 3
	//名媛友好度
	BeautyFriendly Type = 5
	//英雄属性
	HeroAttr Type = 4
	//竞技场
	ArenaScore Type = 6
	//副本分数
	DungeonScore Type = 7
	//联盟等级
	AllianceLv Type = 8
)

var (
	List = []Type{
		Power,
		Phase,
		TotalFriendly,
		BeautyFriendly,
		HeroAttr,
		ArenaScore,
		DungeonScore,
		AllianceLv,
	}

	//永久数据库的Id
	PermanentRankId = map[Type]string{
		Power:          "DefaultPowerRank",
		Phase:          "DefaultPhaseRank",
		TotalFriendly:  "DefaultTotalFriendlyRank",
		BeautyFriendly: "DefaultBeautyFriendlyRank",
		HeroAttr:       "DefaultHeroAttrRank",
		ArenaScore:     "DefaultArenaScorerRank",
		DungeonScore:     "DefaultDungeonScorerRank",
	}

	//永久跨服排行类型
	PermanentRankInterType = []Type{Power, AllianceLv, TotalFriendly}
)
