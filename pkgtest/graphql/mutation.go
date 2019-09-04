// Time        : 2019/08/27
// Description :

package graphql

import (
	"errors"
	"github.com/graphql-go/graphql"
	"golearn/utils"
	"strconv"
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type: studentType,
			Args: graphql.FieldConfigArgument{
				"Name": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "学员名称",
				},
				"BeautyNo": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "名媛编号",
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				stu := StudentValue{
					Id:            strconv.Itoa(utils.RandInt(10, 99)),
					Name:          p.Args["Name"].(string),
					BeautyNo:      int16(p.Args["BeautyNo"].(int)),
					Sex:           int16(utils.RandInt(1, 2)),
					Talent:        int16(utils.RandInt(1, 5)),
					Power:         0,
					Prof:          int16(utils.RandInt(1, 4)),
					Status:        1,
					Exp:           0,
					RecoverRemain: 1800,
				}
				students[stu.Id] = &stu
				return stu, nil
			},
		},
		"train": &graphql.Field{
			Type: studentType,
			Args: graphql.FieldConfigArgument{
				"Id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "学员ID",
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				id := p.Args["Id"].(string)
				stu, ok := students[id]
				if !ok {
					return nil, errors.New("找不到学员")
				}
				if stu.Status == 5 {
					return nil, errors.New("已授勋学员无法培养")
				}
				stu.Power += 100
				stu.Exp += 10
				if stu.Exp >= 100 {
					stu.Status += 1
					stu.Exp = 0
				}
				return stu, nil
			},
		},
		"reject": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"Id": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "学员ID",
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				id := p.Args["Id"].(string)
				stu, ok := students[id]
				if !ok {
					return false, errors.New("找不到学员")
				}
				if stu.Status != 1 {
					return false, errors.New("学员已成长无法遗弃")
				}
				delete(students, id)
				return true, nil
			},
		},
	},
})
