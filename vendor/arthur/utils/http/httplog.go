/*
@Time : 2018/12/13 17:10
@Author : linfeng
@File : httplog
@Desc:
*/

package http

import (
	"arthur/sdk/dclog"
	"fmt"
	"time"
)

type httpLog interface {
	StartRequest() httpLog
	EndRequest() httpLog
	SetURL(url string) httpLog
	SetParams(params []byte) httpLog
	SetResp(resp []byte) httpLog
	SetStatusCode(statusCode int) httpLog
	SetErr(err error) httpLog
	SetData(url string, params, resp []byte, statusCode int, err error) httpLog
	RecordLog()
}

type httpRecord struct {
	Starttime  int64
	Endtime    int64
	Spend      int64
	URL        string
	Params     []byte
	Resp       []byte
	StatusCode int
	Err        error
}

func (h *httpRecord) SetResp(resp []byte) httpLog {
	h.Resp = resp
	return h
}

func (h *httpRecord) SetData(url string, params, resp []byte, statusCode int, err error) httpLog {
	h.SetURL(url)
	h.SetParams(params)
	h.SetResp(resp)
	h.SetStatusCode(statusCode)
	h.SetErr(err)
	return h
}

func (h *httpRecord) RecordLog() {
	if h.Starttime == 0 {
		panic("ErrorStartTime")
	}
	if h.Endtime == 0 {
		panic("ErrorEndTime")
	}
	if h.URL == "" {
		panic("URL Forbid Empty")
	}
	if h.StatusCode == 0 {
		panic("StatusCode Error")
	}
	h.Spend = h.Endtime - h.Starttime
	msg := fmt.Sprintf("url:%s starttime:%d endtime:%d", h.URL, h.Starttime, h.Endtime)
	m := map[string]interface{}{
		"spend": h.Spend,
		"code":  h.StatusCode,
		"msg":   msg,
	}
	if h.Err != nil {
		m["err"] = h.Err.Error()
	}
	if h.Params != nil {
		m["req"] = string(h.Params)
	}
	if h.Resp != nil {
		m["resp"] = string(h.Resp)
	}
	dclog.Debug(msg, "HTTP_LOG", "", m)
}

func (h *httpRecord) StartRequest() httpLog {
	h.Starttime = time.Now().UnixNano() / int64(time.Millisecond)
	return h
}

func (h *httpRecord) EndRequest() httpLog {
	h.Endtime = time.Now().UnixNano() / int64(time.Millisecond)
	return h
}

func (h *httpRecord) SetURL(url string) httpLog {
	h.URL = url
	return h
}

func (h *httpRecord) SetParams(params []byte) httpLog {
	h.Params = params
	return h
}

func (h *httpRecord) SetStatusCode(statusCode int) httpLog {
	h.StatusCode = statusCode
	return h
}

func (h *httpRecord) SetErr(err error) httpLog {
	if err != nil {
		h.Err = err
	}
	return h
}

func newHttpLog() httpLog {
	return &httpRecord{}
}
