package olivere

import (
	"fmt"
	es "github.com/olivere/elastic/v7"
	"golearn/util"
	"testing"
)

func TestESClient(t *testing.T) {
	cli, err := es.NewClient(es.SetSniff(false), es.SetURL(host), es.SetBasicAuth("elastic", "p3Ntz292A0k2P9bBMwA3"))
	util.MustNil(err)
	res, err := cli.ElasticsearchVersion(host)
	util.MustNil(err)
	fmt.Println(res)
}
