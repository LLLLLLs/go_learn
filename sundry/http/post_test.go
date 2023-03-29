package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
)

// http post 请求
func HPost(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("new request is fail: %v", err))
	}
	req.Header.Set("Content-type", "application/json")

	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	client := &http.Client{}
	return client.Do(req)
}

type MissionMatchedRoomInfoReq struct {
	RoomId int64   `json:"room_id"`
	Pids   []int64 `json:"pids"`
}

func TestPost(t *testing.T) {
	req := &MissionMatchedRoomInfoReq{
		RoomId: 123456,
		Pids:   []int64{123465, 22222},
	}
	_, err := HPost(fmt.Sprintf("http://%s/mission/matchedRoomInfo", "172.16.22.19:30633"), req, nil, nil)
	fmt.Println(err)
}
