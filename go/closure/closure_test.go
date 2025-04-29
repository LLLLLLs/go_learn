package closure

import (
	"fmt"
	"testing"
)

func TestExcludeArtifactContract(t *testing.T) {
	list := make([]int, 0)
	for i := 0; i <= 10000; i++ {
		list = append(list, i)
	}
	ex := ExcludeArtifactContract(1, 2)
	res := make([]int, 0)
	for _, x := range list {
		if ex(res, x) {
			res = append(res, x)
		}
	}
	fmt.Println(res)
}
