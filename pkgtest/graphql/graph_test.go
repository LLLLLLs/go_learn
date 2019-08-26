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
		student{
			Id
			Name
			Prof
			Status
		}
	}
`
	fmt.Println(execQuery(query))
}
