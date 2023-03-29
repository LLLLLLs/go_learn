package official

import (
	"fmt"
	es8 "github.com/elastic/go-elasticsearch/v8"
	"golearn/util"
)

func Client() {
	cli, err := es8.NewClient(es8.Config{
		Addresses: []string{"10.30.40.199:19200"},
		Username:  "elastic",
		Password:  "lianleshuo123456",
	})
	util.MustNil(err)
	res, err := cli.Info()
	util.MustNil(err)
	fmt.Println(res)
}
