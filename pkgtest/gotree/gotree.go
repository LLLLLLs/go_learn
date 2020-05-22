//@author: lls
//@time: 2020/05/22
//@desc:

package gotree

import (
	"fmt"
	"github.com/disiqueira/gotree"
)

func gt() {
	t := gotree.New("abc")
	t.Add("cde")
	fmt.Println(t.Print())
}
