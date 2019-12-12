// Time        : 2019/07/09
// Description :

package util

import (
	"fmt"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 0)
	DeepCopy(&dst, &src)
	dst[0] = 10
	fmt.Println(dst)
	fmt.Println(src)
}

var jsonStr = []byte(`{
	"code": 0,
	"info": "123",
	"data": {
		"code": 0,
		"info": "123",
		"data": {
			"captcha_id": "123456",
			"number": "654321"
		}
	}
}`)

//获取验证码响应参数
type YfGetCaptchaResponse struct {
	Code uint   `json:"code"`
	Info string `json:"info"`
	Data struct {
		Code uint   `json:"code"`
		Info string `json:"info"`
		Data struct {
			CaptchaId string `json:"captcha_id"`
			Captcha   string `json:"number"`
		}
	}
}

func TestUnmarshal(t *testing.T) {
	var resp YfGetCaptchaResponse
	Unmarshal(jsonStr, &resp)
	fmt.Println(resp)
}
