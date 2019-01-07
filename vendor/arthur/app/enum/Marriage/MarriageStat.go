package Marriage

// 学员婚姻状态
type Stat int16

func (t Stat) ToInt16() int16 {
	return int16(t)
}

const (
	None             Stat = 0 // 无状态
	Waiting          Stat = 1 // 寻缘中
	Proposing        Stat = 2 // 求婚中
	Rejected         Stat = 3 // 求婚被拒
	SeekingOutDated  Stat = 4 // 寻缘超时
	ProposalOutDated Stat = 5 // 求婚超时
)

// 联姻类型
type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	Seeking  Type = 1 // 寻缘
	Proposal Type = 2 // 求婚
)

type Consume int16

func (c Consume) ToInt16() int16 {
	return int16(c)
}

const (
	Item    Consume = 1 // 消耗道具(优先)
	Diamond Consume = 2 // 消耗钻石(道具不足)
)
