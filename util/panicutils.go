// Time        : 2019/06/25
// Description :

package util

func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}

func MustTrue(b bool) {
	if !b {
		panic("must be true")
	}
}
