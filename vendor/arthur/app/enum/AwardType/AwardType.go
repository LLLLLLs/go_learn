package AwardType

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	Item   Type = 1 //道具奖励
	Value  Type = 2 //值奖励
	Hero   Type = 3 //英雄奖励
	Beauty Type = 4 //红颜奖励
)
