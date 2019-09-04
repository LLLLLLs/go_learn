// Time        : 2019/09/03
// Description :

package viper

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"golearn/utils"
	"testing"
)

func TestGetConfig(t *testing.T) {
	af := initConfig()
	stat := map[string]interface{}{"stat": af}
	viper.SetConfigType("json")
	err := viper.ReadConfig(bytes.NewBuffer(utils.Marshal(stat)))
	utils.OkOrPanic(err)
	students := viper.Get("stat.role.role.students")
	fmt.Println(students)
	allConf := viper.Get("stat")
	fmt.Println(allConf)
}
