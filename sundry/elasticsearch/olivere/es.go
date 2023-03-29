package olivere

import (
	"fmt"
	es "github.com/olivere/elastic/v7"
	"golearn/util"
)

var host = "http://localhost:9200"

func Client() {
	cli, err := es.NewClient(es.SetURL(host))
	util.MustNil(err)
	res, err := cli.ElasticsearchVersion(host)
	util.MustNil(err)
	fmt.Println(res)
}
