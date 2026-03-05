package main

import (
	"encoding/json"
	"golearn/script/json1/foo_conf/param"
	"os"
)

type ZoneConf struct {
	Area          string
	IP            string
	Port          int32
	DataPath      string
	LogLevel      string
	DefaultDBAddr string
	DefaultDBName string
	Season        int32
	CreateTime    string
	RegionList    []Item
}

type Item struct {
	ID      int    `json:"ID"`
	MergeID int    `json:"MergeID"`
	RPCAddr string `json:"RPCAddr"`
}

func main() {
	items := []Item{
		{
			ID:      param.MergeID,
			MergeID: param.MergeID,
			RPCAddr: param.GameRpcAddr,
		},
	}
	for i := param.ServerStart; i <= param.ServerEnd; i++ {
		items = append(items, Item{
			ID:      i,
			MergeID: param.MergeID,
			RPCAddr: param.GameRpcAddr,
		})
	}

	file, err := os.Create("zone_conf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	zc := ZoneConf{
		Area:          param.Area,
		IP:            "0.0.0.0",
		Port:          int32(param.ZonePort),
		DataPath:      param.DataPath,
		LogLevel:      param.LogLevel,
		DefaultDBAddr: param.ZoneDBAddr,
		DefaultDBName: param.ZoneDBName,
		Season:        param.Season,
		CreateTime:    param.CreateTime,
		RegionList:    items,
	}
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(zc); err != nil {
		panic(err)
	}
}
