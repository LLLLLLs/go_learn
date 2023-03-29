package test

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"golearn/pkgtest/grpc/proto/protos"
	"google.golang.org/protobuf/types/known/anypb"
	"testing"
)

func TestAny(t *testing.T) {
	data := &protos.Data{
		Id:   1,
		Name: "hello",
	}
	any, err := anypb.New(data)
	fmt.Println(1, err)
	req := &protos.Request{
		Id:   123,
		Data: any,
	}
	raw, err := proto.Marshal(req)
	fmt.Println(2, err)
	res := new(protos.Request)
	fmt.Println(3, proto.Unmarshal(raw, res))
	fmt.Println(res)
	fin := new(protos.Data)
	fmt.Println(4, res.Data.UnmarshalTo(fin))
	fmt.Println(fin)
}
