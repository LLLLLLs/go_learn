//@author: lls
//@time: 2021/01/04
//@desc:

package value

import (
	"fmt"
	"testing"
)

type test struct {
	ti int
	ts string
	tb bool
	ns string
	nest
}

type nest struct {
	ni int
	ns string
	nb bool
}

func (t test) readOnly() test {
	t.nest.nb = true
	return t
}

func TestValue(t *testing.T) {
	v := test{
		ti: 1,
		ts: "2",
		tb: false,
		ns: "2",
		nest: nest{
			ni: 11,
			ns: "22",
			nb: false,
		},
	}
	fmt.Println(v)
	nv := v.readOnly()
	fmt.Println(v)
	fmt.Println(nv)
	fmt.Println(v.ns, v.nest.ns)
}
