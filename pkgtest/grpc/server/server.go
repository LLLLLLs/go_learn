package main

import (
	"context"
	"fmt"
	"golearn/pkgtest/grpc/proto/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"time"
)

type server struct{}

func (s server) TryError(ctx context.Context, req *protos.TryErrorReq) (*protos.TryErrorResp, error) {
	return nil, fmt.Errorf(req.Hello)
}

func (s server) Any(ctx context.Context, request *protos.Request) (*protos.Request, error) {
	fmt.Println(request)
	return &protos.Request{
		Id:   321,
		Data: request.Data,
	}, nil
}

func (s server) SayHello(ctx context.Context, request *protos.HelloRequest) (*protos.HelloResponse, error) {
	fmt.Println(request.Name)
	return &protos.HelloResponse{Message: "world"}, nil
}

func logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("service %s, trace_id %s ,server %v\n", info.FullMethod, md["trace_id"], info.Server)
	return handler(ctx, req)
}

func cost(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	resp, err = handler(ctx, req)
	fmt.Println("cost:", time.Now().Sub(start).Microseconds(), "μs")
	return
}

func main() {
	lsn, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("failed to listen： %v", err)
	}
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(logger, cost))
	protos.RegisterGreeterServer(s, &server{})
	_ = s.Serve(lsn)
}
