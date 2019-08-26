// Time        : 2019/08/26
// Description :

package graphql

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"go_learn/utils"
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

var mockStudent = StudentValue{
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

var studentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "student",
	Fields: graphql.Fields{
		"Id": &graphql.Field{
			Type: graphql.String,
		},
		"Name": &graphql.Field{
			Type: graphql.String,
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
	},
	Description: "",
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"student": &graphql.Field{
			Type: studentType,
			Args: nil,
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				return map[string]interface{}{
					"Id": "666666",
					"Name": func() interface{} {
						return "李四"
					},
					"BeautyNo":      1,
					"Sex":           2,
					"Talent":        3,
					"Power":         9999999,
					"Prof":          1,
					"Status":        2,
					"Exp":           99,
					"RecoverRemain": 1800,
				}, nil
				//return mockStudent, nil
			},
		},
	},
})

var schema graphql.Schema

func init() {
	var err error
	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	utils.OkOrPanic(err)
}

func execQuery(query string) interface{} {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	data, err := json.Marshal(result)
	utils.OkOrPanic(err)
	return string(data)
}
