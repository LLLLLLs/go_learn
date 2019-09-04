// Time        : 2019/09/02
// Description :

package mongotest

type StudentValue struct {
	Id            string `bson:"_id"` // 唯一ID
	Name          string // 名字
	BeautyNo      int16  // 名媛
	Sex           int16  // 性别
	Talent        int16  // 资质
	Power         int64  // 属性
	Prof          int16  // 职业
	Status        int16  // 状态 1=婴儿 2=幼年 3=成年 4=待授勋
	Exp           int    // 经验
	RecoverRemain int64  // 活力回满剩余时间
}

type Role struct {
	RoleId   string `bson:"_id"`
	Students []StudentValue
}
