package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline"
	. "golearn/pkgtest/tablestore"
	"math"
	"time"
)

func main() {
	storeBuilder := timeline.StoreOption{
		Endpoint:         EndPoint,
		Instance:         InstanceName,
		TableName:        TlNameTest,
		AkId:             AccessKeyId,
		AkSecret:         AccessKeySecret,
		TTL:              60 * 24 * 3600, // Data time to alive, eg: almost one year
		TableStoreConfig: tablestore.NewDefaultTableStoreConfig(),
	}
	history, err := timeline.NewDefaultStore(storeBuilder)
	if err != nil {
		panic(err)
	}
	// if table is not exist, sync will create table
	// if table is already exist and StoreOption.TTL is not zero, sync will check and update table TTL if needed
	err = history.Sync()
	if err != nil {
		panic(err)
	}
	get(history)
	for {
		var (
			id   string
			nick string
		)
		_, _ = fmt.Scanln(&id)
		fmt.Println("id", id)
		_, _ = fmt.Scanln(&nick)
		fmt.Println("nick", nick)
		if nick == "exit" {
			break
		}
		msg := &timeline.StreamMessage{
			Id:        id,
			Content:   nick,
			Timestamp: time.Now().UnixNano(),
			Attr: map[string]interface{}{
				"Attr_Type": 1,
				"Attr_From": 2,
				"Attr_To":   3,
			},
		}
		send(history, msg)
	}
}

func get(store timeline.MessageStore) {
	receiver, err := timeline.NewTmLine("test", timeline.DefaultStreamAdapter, store)
	if err != nil {
		panic(err)
	}
	iterator := receiver.Scan(&timeline.ScanParameter{
		From:        math.MaxInt64,
		To:          0,
		MaxCount:    100,
		BufChanSize: 10,
	})
	entries := make([]*timeline.Entry, 0)
	// avoid scanner goroutine leak
	defer iterator.Close()
	for {
		entry, err := iterator.Next()
		if err != nil {
			if err == timeline.ErrorDone {
				break
			} else {
				panic(err)
			}
		}
		entries = append(entries, entry)
	}
	return
}

func send(store timeline.MessageStore, msg timeline.Message) {
	sender, err := timeline.NewTmLine("test", timeline.DefaultStreamAdapter, store)
	if err != nil {
		panic(err)
	}
	_, err = sender.BatchStore(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("send msg", msg)
}
