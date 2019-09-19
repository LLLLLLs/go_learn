// Time        : 2019/09/09
// Description :

package model

type Version struct {
	Version string
}

func init() {
	RegisterModel(Version{})
}
