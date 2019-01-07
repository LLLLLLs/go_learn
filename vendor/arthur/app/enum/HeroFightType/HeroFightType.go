package HeroFightType

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	//远征
	Anabasis Type = 1
	//单人副本
	DungeonSingle Type = 2
	//全局副本
	DungeonGlobal Type = 3
	// 竞技场1v1随机次数
	ArenaSoloRandom Type = 4
	// 竞技场指定派遣次数
	ArenaAppoint Type = 5
)
