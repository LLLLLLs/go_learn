// Time        : 2019/08/27
// Description :

package graphql

import (
	"github.com/graphql-go/graphql"
	"golearn/util"
)

var schema graphql.Schema

func init() {
	var err error
	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	util.MustNil(err)
}

func ExecQuery(query string, variables ...map[string]interface{}) interface{} {
	param := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}
	if len(variables) != 0 {
		param.VariableValues = variables[0]
	}
	return graphql.Do(param)
}
