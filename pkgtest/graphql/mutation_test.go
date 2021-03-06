// Time        : 2019/08/27
// Description :

package graphql

import (
	"fmt"
	"testing"
	"time"
)

var allField = `
	Id
	Name
	BeautyNo
	Sex      
	Talent
	Power
	Prof
	Status
	Exp
	RecoverRemain
`

func TestCreate(t *testing.T) {
	mutation := `
	mutation {
		create(Name:"sss1",BeautyNo:12){` + allField + `}
	}`
	fmt.Println(ExecQuery(mutation))
}

func TestNow(t *testing.T) {
	fmt.Println(time.Now().Unix())
}
