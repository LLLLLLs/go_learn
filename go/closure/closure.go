//@author: lls
//@time: 2020/04/10
//@desc: 闭包

package closure

func A() {

}

func B() {

}

type IncludeFunc func(cur []int, candidate int) bool

func ExcludeArtifactContract(hero ...int) IncludeFunc {
	heroMap := make(map[int]bool)
	for _, h := range hero {
		heroMap[h] = true
	}
	return func(cur []int, candidate int) bool {
		if heroMap[candidate] {
			heroMap[candidate*2] = true
			return true
		}
		return false
	}
}
