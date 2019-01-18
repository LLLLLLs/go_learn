/*
Author      : lls
Time        : 2018/09/11
Description :
*/

package main

import (
	"github.com/name5566/leaf/log"
	"net"
	"net/rpc"
	"strconv"
)

type HelloService struct{}

func (hs *HelloService) Hello(request int, reply *[]string) error {
	(*reply) = make([]string, request)
	for i := 0; i < request; i++ {
		(*reply)[i] = "hello" + strconv.Itoa(i)
	}
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Tcp listen err:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept err:", err)
	}
	rpc.ServeConn(conn)

}
