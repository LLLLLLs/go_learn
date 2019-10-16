// Time        : 2019/09/26
// Description :

package impl1

import (
	"fmt"
	"golearn/sundry/stupid-di/example/pkg1"
)

type impl struct {
}

func (i impl) Pkg1() {
	fmt.Println("implement Pkg1Interface")
}

func NewImpl() pkg1.Pkg1Interface {
	return impl{}
}
