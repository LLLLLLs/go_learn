//@author: lls
//@date: 2020/7/1
//@desc:

package _struct

import "testing"

func TestAlias(t *testing.T) {
	var (
		ea EnumA
		eb EnumB
		ec EnumC
	)
	ea = 1
	eb = 1
	ec = 1
	ec = EnumC(eb)
	ec = EnumC(ea)
	eb = ea
	eb = EnumB(ec)
}
