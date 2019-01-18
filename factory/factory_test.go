// Time        : 2019/01/11
// Description :

package factory

import "testing"

func TestFactory(t *testing.T) {
	tp := new(aFactory).create()
	tp.add()
	tp.print()
	tp.sub()
	tp.print()
	tp = new(bFactory).create()
	tp.add()
	tp.print()
	tp.sub()
	tp.print()
}
