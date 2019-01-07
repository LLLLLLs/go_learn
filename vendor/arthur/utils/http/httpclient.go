/*
@Time : 2018/12/13 16:56
@Author : linfeng
@File : httpclient
@Desc:
*/

package http

import (
	"bytes"
	"fmt"
	"gitlab.dianchu.cc/DevOpsGroup/goutils/encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	TEXT = "text/plain;charset=utf-8"
	JSON = "application/json;charset=utf-8"
)

type HttpClient struct {
	*http.Client
}

func (h *HttpClient) Get(url string) (statusCode int, body []byte, err error) {
	log := newHttpLog().StartRequest()
	resp, err := h.Client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.EndRequest().SetData(url, nil, nil, -1, err).RecordLog()
		return
	}
	statusCode = resp.StatusCode
	body, err = ReadResp(resp.Body)
	log.EndRequest().SetData(url, nil, body, resp.StatusCode, err).RecordLog()
	return
}

func (h *HttpClient) PostForm(url string, data url.Values) (statusCode int, body []byte, err error) {
	params := json.Marshal(data)
	log := newHttpLog().StartRequest()
	resp, err := h.Client.PostForm(url, data)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.EndRequest().SetData(url, params, nil, -1, err).RecordLog()
		return
	}
	statusCode = resp.StatusCode
	body, err = ReadResp(resp.Body)
	log.EndRequest().SetData(url, params, body, resp.StatusCode, err).RecordLog()
	return
}

func (h *HttpClient) post(url string, data []byte, isJson bool) (statusCode int, body []byte, err error) {
	var resp *http.Response
	reader := bytes.NewReader(data)
	log := newHttpLog().StartRequest()
	if isJson {
		resp, err = h.Client.Post(url, JSON, reader)
	} else {
		resp, err = h.Client.Post(url, TEXT, reader)
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.EndRequest().SetData(url, data, nil, -1, err).RecordLog()
		return
	}
	statusCode = resp.StatusCode
	body, err = ReadResp(resp.Body)
	log.EndRequest().SetData(url, data, body, resp.StatusCode, err).RecordLog()
	return
}

func (h *HttpClient) Post(url string, data []byte) (statusCode int, body []byte, err error) {
	return h.post(url, data, false)
}

func (h *HttpClient) PostJson(url string, data []byte) (statusCode int, body []byte, err error) {
	return h.post(url, data, true)
}

func (h *HttpClient) PostJsonObj(url string, req, reply interface{}) (err error) {
	var response *http.Response
	if req != nil {
		b, err := json.MarshalWithErr(req)
		if err != nil {
			return err
		}
		response, err = h.Client.Post(url, JSON, bytes.NewReader(b))
	} else {
		response, err = h.Client.Post(url, JSON, nil)
	}

	if err != nil {
		return
	}

	if response != nil {
		defer response.Body.Close()
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("http response status code !=200,code:%d", response.StatusCode)
	}
	return json.Config.NewDecoder(response.Body).Decode(reply)

}

func ReadResp(resp io.Reader) ([]byte, error) {
	respBody, err := ioutil.ReadAll(resp)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
