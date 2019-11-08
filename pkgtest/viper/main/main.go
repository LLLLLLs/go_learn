// Time        : 2019/09/04
// Description :

package main

import (
	"fmt"
	"github.com/spf13/viper"
	myviper "golearn/pkgtest/viper"
	"golearn/util"
)

func main() {
	var isContinue = 1
	for isContinue == 1 {
		myviper.InitConfigFromFile("/go_learn/pkgtest/viper/config.json")
		fmt.Println(viper.Get("Cd"))
		fmt.Println(viper.Get("Length"))
		fmt.Println("0 = exit 1 = continue")
		_, err := fmt.Scanln(&isContinue)
		util.MustNil(err)
	}
}
