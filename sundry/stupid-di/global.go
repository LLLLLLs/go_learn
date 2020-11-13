//@author: lls
//@time: 2020/08/03
//@desc:

package stupiddi

var globalDI = NewDI()

func Provide(constructor ...interface{}) {
	globalDI.Provide(constructor...)
}

func Get(model interface{}) interface{} {
	return globalDI.Get(model)
}
