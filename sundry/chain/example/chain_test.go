//@author: lls
//@time: 2020/05/19
//@desc:

package example

import (
	"fmt"
	"golearn/sundry/chain"
	"testing"
)

func TestChain(t *testing.T) {
	c := chain.NewContext()
	c.Append(Generate{}, Add{}, Add{}, Multi{}, Multi{})
	fmt.Println(c.Next())
}
