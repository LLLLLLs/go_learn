// Time        : 2019/08/27
// Description :

package graphql

import (
	"github.com/graphql-go/graphql"
)

type StudentValue struct {
	Id            string // 唯一ID
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

var mockStudentValue = StudentValue{
	Id:            "666666",
	Name:          "张三",
	BeautyNo:      1,
	Sex:           2,
	Talent:        3,
	Power:         9999999,
	Prof:          1,
	Status:        2,
	Exp:           99,
	RecoverRemain: 1800,
}

var mockStudentMap = map[string]interface{}{
	"Id": "666666",
	"Name": func() interface{} {
		return "李四"
	},
	"BeautyNo":      1,
	"Sex":           2,
	"Talent":        3,
	"Power":         9999999,
	"Prof":          int16(1),
	"Status":        int16(2),
	"Exp":           99,
	"RecoverRemain": 1800,
}

var studentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "student",
	Fields: graphql.Fields{
		"Id": &graphql.Field{
			Type: graphql.String,
		},
		"Name": &graphql.Field{
			Type:    graphql.String,
			Resolve: graphql.DefaultResolveFn,
		},
		"BeautyNo": &graphql.Field{
			Type: graphql.Int,
		},
		"Sex": &graphql.Field{
			Type: graphql.Int,
		},
		"Talent": &graphql.Field{
			Type: graphql.Int,
		},
		"Power": &graphql.Field{
			Type: graphql.Int,
		},
		"Prof": &graphql.Field{
			Type: profEnum,
		},
		"Status": &graphql.Field{
			Type: statusEnum,
		},
		"Exp": &graphql.Field{
			Type: graphql.Int,
		},
		"RecoverRemain": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var profEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "prof",
	Values: graphql.EnumValueConfigMap{
		"Knight": &graphql.EnumValueConfig{
			Value: int16(1),
		},
		"Wizard": &graphql.EnumValueConfig{
			Value: int16(2),
		},
		"Warrior": &graphql.EnumValueConfig{
			Value: int16(3),
		},
		"Ranger": &graphql.EnumValueConfig{
			Value: int16(4),
		},
	},
	Description: "",
})

var statusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "status",
	Values: graphql.EnumValueConfigMap{
		"Baby": &graphql.EnumValueConfig{
			Value: int16(1),
		},
		"Infant": &graphql.EnumValueConfig{
			Value: int16(2),
		},
		"Adult": &graphql.EnumValueConfig{
			Value: int16(3),
		},
		"WaitConfer": &graphql.EnumValueConfig{
			Value: int16(4),
		},
		"Conferred": &graphql.EnumValueConfig{
			Value: int16(5),
		},
	},
	Description: "",
})

var students = make(map[string]*StudentValue)
