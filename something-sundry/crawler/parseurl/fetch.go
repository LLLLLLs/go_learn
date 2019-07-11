// Time        : 2019/07/11
// Description :

package parseurl

import (
	"go_learn/utils"
	"io/ioutil"
	"net/http"
)

func fetch(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	utils.OkOrPanic(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	resp, err := (&http.Client{}).Do(req)
	utils.OkOrPanic(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	utils.OkOrPanic(err)
	return string(body)
}
