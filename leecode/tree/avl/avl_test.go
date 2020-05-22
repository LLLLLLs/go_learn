//@author: lls
//@time: 2020/05/21
//@desc:

package avl

import (
	"fmt"
	"testing"
)

func TestAvl(t *testing.T) {
	avl := NewNode(10, 10)
	avl.print()
	fmt.Println("插入8")
	avl = avl.Insert(NewNode(8, 8))
	avl.print()
	fmt.Println("插入9")
	avl = avl.Insert(NewNode(9, 9))
	avl.print()
	fmt.Println("插入11")
	avl = avl.Insert(NewNode(11, 11))
	avl.print()
	fmt.Println("插入12")
	avl = avl.Insert(NewNode(12, 12))
	avl.print()
	fmt.Println("插入6")
	avl = avl.Insert(NewNode(6, 6))
	avl.print()
	fmt.Println("插入7")
	avl = avl.Insert(NewNode(7, 7))
	avl.print()

}
