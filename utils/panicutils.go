// Time        : 2019/06/25
// Description :

package utils

func OkOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}
