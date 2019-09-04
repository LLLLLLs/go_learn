// Time        : 2019/08/26
// Description :

package graphql

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	query := `
	{
		student(Id:"11"){
			Id
			Name
			Prof
			Status
			Power
		}
		list{
			Id
			Name
		}
	}
`
	fmt.Println(ExecQuery(query))
}
