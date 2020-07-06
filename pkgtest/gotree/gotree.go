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
	t.Add("aaa")
	sub := gotree.New("bbb")
	sub.Add("ccc")
	sub.Add("eee")
	t.AddTree(sub)
	t.Add("ddd")
	t.Add("fff")
	fmt.Println(t.Print())
}
