package main

import (
	"context"
	"fmt"
	"golearn/pkgtest/grpc/proto/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"time"
)

func cost(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Println("cost:", time.Now().Sub(start).Microseconds(), "μs")
	return err
}

func main() {
	conn, err := grpc.Dial(":7777", grpc.WithInsecure(), grpc.WithChainUnaryInterceptor(cost))
	if err != nil {
		log.Fatalf("did not connect： %v", err)
	}
	defer conn.Close()
	cli := protos.NewGreeterClient(conn)

	md := metadata.New(map[string]string{
		"trace_id": "trace_id",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	errResp, err := cli.TryError(ctx, &protos.TryErrorReq{Hello: "hello"})
	fmt.Println(errResp, err)
	resp, err := cli.SayHello(ctx, &protos.HelloRequest{Name: "123"})
	fmt.Println(resp.Message)
	fmt.Println(err)
	data := &protos.Data{
		Id:   222,
		Name: "hello",
	}
	anyData, err := anypb.New(data)
	if err != nil {
		panic(err)
	}
	anyResp, err := cli.Any(ctx, &protos.Request{
		Id:   123,
		Data: anyData,
	})
	fmt.Println(anyResp)
}
