// Time        : 2019/07/11
// Description :

package parseurl

import (
	"regexp"
	"strings"
)

func ParseBuiltIn(url string, result *[]interface{}) {
	body := fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	rp := regexp.MustCompile(`<div class="hd">(.*?)</div>`)
	titleRe := regexp.MustCompile(`<span class="title">(.*?)</span>`)
	idRe := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(\d+)/"`)
	items := rp.FindAllStringSubmatch(body, -1)
	for _, item := range items {
		*result = append(*result,
			[]interface{}{
				idRe.FindStringSubmatch(item[1])[1],
				titleRe.FindStringSubmatch(item[1])[1],
			},
		)
	}
}
