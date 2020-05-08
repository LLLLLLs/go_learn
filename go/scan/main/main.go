//@time:2020/04/02
//@desc:

package main

import (
	"fmt"
)

func main() {
	var input string
	e, err := fmt.Scanln(&input)
	fmt.Println(err)
	fmt.Println(e)
	fmt.Println(input)
}
