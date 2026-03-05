package main

import (
	"encoding/json"
	"fmt"
	"golearn/script/json1/foo_conf/param"
	"os"
)

type GameConf struct {
	Area             string
	IP               string
	Port             int32
	DataPath         string
	LogLevel         string
	DefaultRedisAddr string
	DefaultRedisPass string
	CreateTime       string
	Season           int32
	SeasonTime       string
	RegionList       []Item
}

type Item struct {
	ID      int          `json:"ID"`
	Name    string       `json:"Name"`
	MergeID int          `json:"MergeID"`
	DB      param.DBInfo `json:"DB"`
}

var MyItem = Item{
	ID:      param.MergeID,
	MergeID: param.MergeID,
	DB:      param.GameDB,
}

func main() {
	items := []Item{MyItem}
	for i := param.ServerStart; i <= param.ServerEnd; i++ {
		items = append(items, Item{
			ID:      i,
			Name:    "",
			MergeID: param.MergeID,
			DB: param.DBInfo{
				Addr: "mongodb://rogue121:KLJdvNeUbt7Rwl!P@dds-uf67240cbe6532841516-pub.mongodb.rds.aliyuncs.com:3717,dds-uf67240cbe6532842120-pub.mongodb.rds.aliyuncs.com:3717/?replicaSet=mgset-92442973&authSource=admin",
				Name: "foo_game_" + itoa(i),
			},
		})
	}

	file, err := os.Create("game_conf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	gc := GameConf{
		Area:             param.Area,
		IP:               param.IP,
		Port:             int32(param.Port),
		DataPath:         param.DataPath,
		LogLevel:         param.LogLevel,
		DefaultRedisAddr: param.DefaultRedisAddr,
		DefaultRedisPass: param.DefaultRedisPass,
		CreateTime:       param.CreateTime,
		Season:           param.Season,
		SeasonTime:       param.SeasonTime,
		RegionList:       items,
	}
	if err := encoder.Encode(gc); err != nil {
		panic(err)
	}
}

// itoa converts int to string (you can also use strconv.Itoa)
func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
