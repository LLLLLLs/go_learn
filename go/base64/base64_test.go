// Time        : 2019/09/23
// Description :

package base64

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	bs := []byte{0b00000100, 0b00100000, 0b11000100, 0b00011101}
	// encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	b64 := base64.StdEncoding.EncodeToString(bs)
	fmt.Println(b64)
	// 原二进制 = 00000100 00100000 11000100 00000000
	// base64每个字符6bit = 000001 000010 000011 000100 000111 01
	// 最后只占2bit,所以需0补齐
	// 结果 = 000001 000010 000011 000100 000111 010000 = 1 2 3 4 7 16 (BCDEHQ)
	// 此时并未结束,b64长度一定是4的倍数,因为byte长度=8,b64长度=6,最小公倍数=24=8*3=6*4
	// 所以结果应再补两位偏移量占位符"="得到最终结果"BCDEHQ=="
	// PS：常见的base64字符串最后的"="就是用来补位的,但值得注意的是不可能出现"===",因为第一个byte将对应两个b64字符

	// 使用自定义encoding来编码b64
	// encoder表示从0~63分别用什么字符来编码(偏移量占位符始终是"=")
	myEncoding := base64.NewEncoding("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	myB64 := myEncoding.EncodeToString(bs)
	fmt.Println(myB64)
}
