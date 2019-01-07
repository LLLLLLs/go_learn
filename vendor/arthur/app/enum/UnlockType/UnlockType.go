
package UnlockType

type Type int16

func (t Type) ToInt16() int16 {
	return int16(t)
}

const (
	//关卡章数
	PhaseChapter Type = 1
	//角色等级
	RoleLv Type = 2
	//英雄数量
	HeroNum Type = 3
	//红颜数量
	BeautyNum Type = 4
	//学员数量
	ChildNum Type = 5
)
