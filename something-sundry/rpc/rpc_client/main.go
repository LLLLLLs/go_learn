/*
Author      : lls
Time        : 2018/09/11
Description :
*/

package main

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"go_learn/utils"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Tcp connect err:", err)
	}
	var reply []interface{}
	err = client.Call("HelloService.Hello", 6, &reply)
	utils.OkOrPanic(err)
	fmt.Println(reply)
	var reply2 int
	err = client.Call("HelloService.Calc", []int{1, 2}, &reply2)
	utils.OkOrPanic(err)
	fmt.Println(reply2)
}
