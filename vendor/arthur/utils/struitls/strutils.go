/*
@Time : 2018/11/26 9:27
@Author : linfeng
@File : strutils
@Desc:
*/

package struitls

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"net"
	"sort"
	"strconv"
	"strings"
)

//XxxYyy转xxx_yyy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

//xxx_yyy转XxxYyy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func RemoveBlank(str string) string {
	return strings.Replace(str, " ", "", -1)
}

func DeepReplace(str string, oldstr string, newstr string) string {
	if strings.Contains(str, oldstr) {
		str = strings.Replace(str, oldstr, newstr, -1)
		return DeepReplace(str, oldstr, newstr)
	}
	return str
}

//前缀、后缀、分隔和内容生成
func PrefixPostfixSepText(prefix, postfix, sep, text string, textTimes int) string {
	var content bytes.Buffer
	var textContent bytes.Buffer
	for i := 0; i < textTimes; i++ {
		if textContent.String() != "" {
			textContent.WriteString(sep)
		}
		textContent.WriteString(text)
	}
	content.WriteString(prefix)
	content.WriteString(textContent.String())
	content.WriteString(postfix)
	return content.String()
}

func ToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func ToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

//生成机器编号
func MachineTag() (string, error) {
	//获取本机IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	array := make([]string, 0)
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				array = append(array, ipnet.IP.String())
			}
		}
	}
	//对IP排序，以免每次返回本机IP的顺序不同，导致MD5加密时结果不同
	sort.Strings(array)
	return MD5String(array...), nil
}

func MD5String(args ...string) string {
	var whereStr bytes.Buffer
	for _, arg := range args {
		whereStr.WriteString(arg)
	}
	h := md5.New()
	h.Write([]byte(whereStr.String()))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
