// Time        : 2019/01/22
// Description :

package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {
	linkList := GetCircleLinkList(10, 5)
	//code.PrintLinkList(linkList)
	//linkList = code.FirstNodeInCircle(linkList)
	linkList = FirstNodeInCircleImprove(linkList)
	fmt.Println(linkList)
	//code.PrintLinkList(linkList)
}
