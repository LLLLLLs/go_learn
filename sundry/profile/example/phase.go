// Time        : 2019/09/05
// Description :

package example

import "golearn/sundry/profile/model"

type Index int16

type Phase struct {
	Index1 int   `index:"true"`
	Index2 Index `index:"true"`
	Index3 int   `index:"true"`
	Conf   string
}

func init() {
	model.RegisterModel(Phase{})
}
