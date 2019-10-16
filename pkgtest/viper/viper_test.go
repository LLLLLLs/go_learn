// Time        : 2019/09/03
// Description :

package viper

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"golearn/util"
	"testing"
)

func TestGetConfig(t *testing.T) {

	af := initConfigFromMongo()
	stat := map[string]interface{}{"stat": af}
	viper.SetConfigType("json")
	err := viper.ReadConfig(bytes.NewBuffer(util.Marshal(stat)))
	util.OkOrPanic(err)
	//conf.GetVersion(1).GetTable("phase").Get("1.2")
	//conf.GetVersion(1).GetTable("").GetAll()
	students := viper.Get("stat.role.role.students")
	fmt.Println(students)
	allConf := viper.Get("stat")
	fmt.Println()
	for k, v := range allConf.(map[string]interface{}) {
		for k2, v2 := range v.(map[string]interface{}) {
			fmt.Printf("%s %v %+v\n", k, k2, v2)
		}
	}
	var conf struct {
		Id      int `bson:"_id"`
		F32     float32
		F64     float64
		B       bool
		UI8     uint8
		UI16    uint16
		UI      uint
		UI32    uint32
		UI32Max uint32
		UI64    uint64
	}
	data := viper.Get("stat.test.36318")
	util.Unmarshal(util.Marshal(data), &conf)
	fmt.Printf("%+v\n", conf)
}

func TestGetConfigFromJson(t *testing.T) {
	InitConfigFromFile("/viper/config.json")
	fmt.Println(viper.Get("Cd"))
	fmt.Println(viper.Get("Length"))
}
