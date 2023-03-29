package main

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	. "golearn/pkgtest/tablestore"
	"time"
)

const (
	tbColPid = int64(1181947970)
)

var (
	client *tablestore.TableStoreClient
)

func main() {
	client = tablestore.NewClient(EndPoint, InstanceName, AccessKeyId, AccessKeySecret)
	resp, err := client.DescribeTable(&tablestore.DescribeTableRequest{TableName: TbNameTest})
	if err != nil {
		fmt.Println("describe table error:", err)
	}
	fmt.Printf("%+v\n", resp)
	if tsErr, ok := err.(*tablestore.OtsError); ok {
		if tsErr.Code == "OTSObjectNotExist" {
			createTestTable()
		} else {
			panic(fmt.Sprintf("unexpect table error:%v", err))
		}
	}
	for {
		var (
			id   int64
			nick string
		)
		_, _ = fmt.Scanln(&id)
		fmt.Println("id", id)
		_, _ = fmt.Scanln(&nick)
		fmt.Println("nick", nick)
		if nick == "exit" {
			break
		}
		insertRow(id, nick)
		getRow()
	}
}

func createTestTable() {
	tm := &tablestore.TableMeta{
		TableName: TbNameTest,
	}
	tm.AddPrimaryKeyColumn("pid", tablestore.PrimaryKeyType_INTEGER)
	tm.AddDefinedColumn("server", tablestore.DefinedColumn_STRING)
	tm.AddDefinedColumn("login_time", tablestore.DefinedColumn_INTEGER)
	tm.AddDefinedColumn("nick", tablestore.DefinedColumn_STRING)
	tm.AddDefinedColumn("last_sync_sequence", tablestore.DefinedColumn_INTEGER)
	tm.AddDefinedColumn("last_group_sync_sequence", tablestore.DefinedColumn_INTEGER)

	_, err := client.CreateTable(&tablestore.CreateTableRequest{
		TableMeta: tm,
		TableOption: &tablestore.TableOption{
			TimeToAlive: -1,
			MaxVersion:  1,
		},
		ReservedThroughput: &tablestore.ReservedThroughput{
			Readcap:  0,
			Writecap: 0,
		},
	})

	fmt.Println("create table error:", err)
}

func insertRow(id int64, nick string) {
	updateRowRequest := new(tablestore.UpdateRowRequest)
	updateRowChange := new(tablestore.UpdateRowChange)
	updateRowChange.TableName = TbNameTest
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("pid", id)

	updateRowChange.PrimaryKey = putPk
	updateRowChange.PutColumn("server", "123")
	updateRowChange.PutColumn("login_time", time.Now().Unix())
	updateRowChange.PutColumn("nick", nick)
	updateRowChange.PutColumn("last_sync_sequence", time.Now().UnixNano())
	updateRowChange.PutColumn("last_group_sync_sequence", time.Now().UnixNano())
	updateRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	updateRowRequest.UpdateRowChange = updateRowChange
	_, err := client.UpdateRow(updateRowRequest)
	fmt.Println("insert row error:", err)
}

func getRow() {
	resp, err := client.GetRow(&tablestore.GetRowRequest{
		SingleRowQueryCriteria: &tablestore.SingleRowQueryCriteria{
			TableName: TbNameTest,
			PrimaryKey: &tablestore.PrimaryKey{PrimaryKeys: []*tablestore.PrimaryKeyColumn{{
				ColumnName: "pid",
				Value:      tbColPid,
			}}},
			MaxVersion: 1,
		},
	})
	fmt.Printf("%+v,%v\n", resp, err)
}
