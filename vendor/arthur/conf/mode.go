/*
Author : Haoyuan Liu
Time   : 2018/8/1
*/
package conf

type Mode int8

const (
	RELEASE Mode = iota
	DEBUG
	TEST
)

var (
	mode = RELEASE
)

func IsMode(m Mode) bool {
	if m == mode {
		return true
	}
	return false
}

func SetMode(m Mode) {
	mode = m
}
