package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tunnel"
	"go.uber.org/zap"
	. "golearn/pkgtest/tablestore"
	"log"
	"os"
	"strings"
	"time"
)

const (
	tnNameTest        = "testTunnelTimeline"
	errTunnelNotExist = "tunnel not exist"
)

func tunnelName() string {
	return tnNameTest + "-" + os.Getenv("ID")
}

var tunnelClient tunnel.TunnelClient

type nop struct{}

func (n2 nop) Write(p []byte) (n int, err error) {
	return
}

func (n2 nop) Sync() error {
	return nil
}

func main() {
	tunnelClient = tunnel.NewTunnelClient(EndPoint, InstanceName, AccessKeyId, AccessKeySecret)

	// 配置callback到SimpleProcessFactory，配置消费端TunnelWorkerConfig。
	workConfig := &tunnel.TunnelWorkerConfig{
		HeartbeatInterval: time.Millisecond * 500,
		LogWriteSyncer:    nop{},
		ProcessorFactory: &tunnel.SimpleProcessFactory{
			Logger:      zap.NewNop(),
			CustomValue: "user custom interface{} value",
			ProcessFunc: exampleConsumeFunction,
		},
	}

	tunnelId := getOrCreateTunnel()
	fmt.Println("tunnel id", tunnelId)

	// 使用TunnelDaemon持续消费指定tunnel。
	daemon := tunnel.NewTunnelDaemon(tunnelClient, tunnelId, workConfig)
	log.Fatal(daemon.Run())
}

func getOrCreateTunnel() string {
	_, err := tunnelClient.DescribeTunnel(&tunnel.DescribeTunnelRequest{
		TableName:  TlNameTest,
		TunnelName: tunnelName(),
	})
	if err == nil {
		_, err = tunnelClient.DeleteTunnel(&tunnel.DeleteTunnelRequest{
			TableName:  TlNameTest,
			TunnelName: tunnelName(),
		})
		fmt.Println("delete tunnel error", err)
		return createTunnel()
	} else {
		if tErr, ok := err.(*tunnel.TunnelError); ok && strings.Contains(tErr.Message, errTunnelNotExist) {
			return createTunnel()
		}
		log.Fatal("describe tunnel error", err)
		return ""
	}
}

func createTunnel() string {
	req := &tunnel.CreateTunnelRequest{
		TableName:  TlNameTest,
		TunnelName: tunnelName(),
		Type:       tunnel.TunnelTypeStream, // 创建全量加增量类型的Tunnel。
	}
	resp, err := tunnelClient.CreateTunnel(req)
	if err != nil {
		log.Fatal("create test tunnel failed", err)
	}
	return resp.TunnelId
}

// 根据业务自定义数据消费callback函数。
func exampleConsumeFunction(ctx *tunnel.ChannelContext, records []*tunnel.Record) error {
	fmt.Println("user-defined information", ctx.CustomValue)
	for _, rec := range records {
		fmt.Println("tunnel record detail:", rec.String())
	}
	fmt.Println("a round of records consumption finished")
	return nil
}
