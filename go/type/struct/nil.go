package _struct

import "fmt"

type nilType struct {
	A string
}

func (n *nilType) PrintA() {
	fmt.Println(n.A)
}
