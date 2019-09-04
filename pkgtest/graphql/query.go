// Time        : 2019/08/26
// Description :

package graphql

import (
	"errors"
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"student": &graphql.Field{
			Type: studentType,
			Args: graphql.FieldConfigArgument{
				"Id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				id := p.Args["Id"]
				i, ok := students[id.(string)]
				if !ok {
					e = errors.New("no student")
				}
				return
			},
		},
		"list": &graphql.Field{
			Type: graphql.NewList(studentType),
			Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
				list := make([]StudentValue, 0, len(students))
				for _, stu := range students {
					list = append(list, *stu)
				}
				return list, nil
			},
		},
	},
})
