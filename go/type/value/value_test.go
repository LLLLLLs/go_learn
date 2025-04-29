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

func TestValueEqual(t *testing.T) {
	type v struct{ a int }
	type vv struct {
		v v
		b string
	}
	fmt.Println(v{a: 1} == v{a: 1})
	fmt.Println(vv{v: v{a: 1}, b: "hello"} == vv{v: v{a: 1}, b: "hello"})

	mm := map[v]bool{}
	mm[v{a: 1}] = true
	fmt.Println(mm[v{a: 1}])
	type vvv struct{ v *v }
	mmm := map[*vvv]bool{}
	mmm[&vvv{v: &v{a: 1}}] = true
	fmt.Println(mmm[&vvv{v: &v{a: 1}}])
	fmt.Println(vvv{v: &v{a: 1}} == vvv{v: &v{a: 1}})
}
