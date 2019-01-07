/*
@Time : 2018/12/7 10:36
@Author : linfeng
@File : http
@Desc:
*/

package http

import (
	"arthur/utils/errors"
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

type IClient interface {
	Get(url string) (statusCode int, body []byte, err error)
	Post(url string, data []byte) (statusCode int, body []byte, err error)
	PostForm(url string, data url.Values) (statusCode int, body []byte, err error)
	PostJson(url string, data []byte) (statusCode int, body []byte, err error)
	PostJsonObj(url string, req, reply interface{}) (err error)
}

func ToResp(statusCode int, resp []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, errors.New(string(resp))
	}
	return resp, nil
}

func NewNetClient(maxIdleConns, maxIdleConnsPerHost int, idleConnTimeout, timeout time.Duration) IClient {
	customTransport := &http.Transport{
		DisableKeepAlives:   false,               // 短连接影响性能
		MaxIdleConns:        maxIdleConns,        // 本HTTP client最大的闲置连接数量，实际上相当于设置并发连接数
		MaxIdleConnsPerHost: maxIdleConnsPerHost, // 每个host最大的闲置连接数量
		IdleConnTimeout:     idleConnTimeout,     // 闲置连接超时时间
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 关闭https证书验证
		},
	}
	return &HttpClient{Client: &http.Client{Timeout: timeout, Transport: customTransport}}
}
