package HeroFaithType

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	None    Type = 0 //无神论者
	Holy    Type = 1 //神圣
	Natural Type = 2 //自然
	Blood   Type = 3 //血族
)

type SkillType int16

func (s SkillType) ToInt16() int16 {
	return int16(s)
}

const (
	IncBySkillLevel    SkillType = 1 // 技能效果根据等级增长
	IncBySkillOwnerNum SkillType = 2 // 技能效果根据同信仰人数增长
)
