//@author: lls
//@time: 2020/04/30
//@desc:

package md5_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	h := md5.New()
	h.Write([]byte("hello word"))
	res := h.Sum(nil)
	fmt.Println(hex.EncodeToString(res))

	h.Reset()
	h.Write([]byte("http://dldir1.qq.com/WechatWebDev/release/0.7.0/wechat_web_devtools_0.7.0_x64.exe"))
	res = h.Sum(nil)
	fmt.Println(hex.EncodeToString(res))
}
