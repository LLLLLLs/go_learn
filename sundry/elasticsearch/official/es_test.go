package official

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"golearn/util"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestESClient(t *testing.T) {
	// cli, err := es8.NewClient(es8.Config{
	// 	Addresses: []string{"http://localhost:9200"},
	// 	Username:  "elastic",
	// 	Password:  "p3Ntz292A0k2P9bBMwA3",
	// })
	cli, err := es8.NewClient(es8.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "p3Ntz292A0k2P9bBMwA3",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	})
	util.MustNil(err)

	res, err := cli.Info()
	util.MustNil(err)
	fmt.Println(res)

	var b strings.Builder
	b.WriteString(`{"title" : "`)
	b.WriteString("title")
	b.WriteString(`"}`)

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      "test",
		DocumentID: "1",
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err = req.Do(context.Background(), cli)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status(), 1)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

}
