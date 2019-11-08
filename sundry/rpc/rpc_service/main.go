/*
Author      : lls
Time        : 2018/09/11
Description :
*/

package main

import (
	"github.com/name5566/leaf/log"
	"golearn/util"
	"net"
	"net/rpc"
	"strconv"
)

type HelloService struct{}

func (hs *HelloService) Hello(request int, reply *[]interface{}) error {
	log.Debug("hello service:%d", request)
	*reply = make([]interface{}, request)
	for i := 0; i < request; i++ {
		(*reply)[i] = "hello" + strconv.Itoa(i)
	}
	return nil
}

func (hs *HelloService) Calc(p []int, reply *int) error {
	log.Debug("calc service:a=%d,b=%d", p[0], p[1])
	*reply = p[0] + p[1]
	return nil
}

func main() {
	err := rpc.RegisterName("HelloService", new(HelloService))
	util.MustNil(err)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Tcp listen err:", err)
		return
	}
	log.Debug("service start ok:localhost:1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept err:", err)
		}
		go rpc.ServeConn(conn)
	}
}
