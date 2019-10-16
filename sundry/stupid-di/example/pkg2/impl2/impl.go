// Time        : 2019/09/26
// Description :

package impl2

import (
	"fmt"
	"golearn/sundry/stupid-di/example/pkg2"
)

type impl struct {
}

func (i impl) Pkg2() {
	fmt.Println("implement Pkg2Interface")
}

func NewImpl() pkg2.Pkg2Interface {
	return impl{}
}
