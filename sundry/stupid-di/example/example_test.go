// Time        : 2019/09/26
// Description :

package example

import (
	"golearn/sundry/stupid-di"
	"golearn/sundry/stupid-di/example/pkg1"
	"golearn/sundry/stupid-di/example/pkg1/impl1"
	"golearn/sundry/stupid-di/example/pkg2"
	"golearn/sundry/stupid-di/example/pkg2/impl2"
	"testing"
)

type TestStruct struct {
	pkg1 pkg1.Pkg1Interface
	pkg2 pkg2.Pkg2Interface
}

func NewTs(p1 pkg1.Pkg1Interface, p2 pkg2.Pkg2Interface) TestStruct {
	return TestStruct{
		pkg1: p1,
		pkg2: p2,
	}
}

func TestExample(t *testing.T) {
	stupiddi.Provide(impl1.NewImpl, impl2.NewImpl)
	stupiddi.Provide(NewTs)
	ts := stupiddi.Get(TestStruct{}).(TestStruct)
	ts.pkg1.Pkg1()
	ts.pkg2.Pkg2()
}
