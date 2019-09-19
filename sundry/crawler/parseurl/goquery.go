// Time        : 2019/07/11
// Description :

package parseurl

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"golearn/utils"
	"strings"
)

// <ol class="grid_view">
//   <li>
//     <div class="item">
//       <div class="info">
//         <div class="hd">
//           <a href="https://movie.douban.com/subject/1292052/" class="">
//             <span class="title">肖申克的救赎</span>
//             <span class="title">&nbsp;/&nbsp;The Shawshank Redemption</span>
//             <span class="other">&nbsp;/&nbsp;月黑高飞(港)  /  刺激1995(台)</span>
//           </a>
//           <span class="playable">[可播放]</span>
//         </div>
//       </div>
//     </div>
//   </li>
//   ....
// </ol>

func ParseGoQuery(url string, result *[]interface{}) {
	body := fetch(url)
	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(body))
	utils.OkOrPanic(err)
	doc.Find("ol.grid_view li").Find(".hd").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Find("a").Attr("href")
		*result = append(*result, []interface{}{strings.Split(url, "/")[4], selection.Find(".title").Eq(0).Text()})
	})
}
