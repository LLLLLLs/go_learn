/*
Author      : lls
Time        : 2018/09/11
Description :
*/

package main

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Tcp connect err:", err)
	}
	var reply []string
	err = client.Call("HelloService.Hello", 5, &reply)
	if err != nil {
		log.Fatal("Rpc call err:", err)
	}
	fmt.Println(reply)
}
