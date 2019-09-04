// Time        : 2019/07/11
// Description :

package main

import (
	"fmt"
	"golearn/something-sundry/crawler/parseurl"
)

var url = "https://movie.douban.com/top250?start="

func main() {
	var result []interface{}
	parseurl.ParseBuiltIn(url+"0", &result)
	parseurl.ParseGoQuery(url+"25", &result)
	for i := range result {
		fmt.Println(i, result[i])
	}
}
