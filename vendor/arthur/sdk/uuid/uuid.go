package uuid

import (
	"arthur/utils/errors"
	"fmt"
	"gitlab.dianchu.cc/DevOpsGroup/goutils/time"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetUuidBySystem(count int) []string {
	var uuids []string
	for count > 0 {
		uuids = append(uuids, newUuid())
		count -= 1
	}
	return uuids
}

func GetOneUuidBySystem() string {
	return newUuid()
}

func GetUuids(count int) []string {
	uuidAddrFinal := fmt.Sprintf("%s/%s/%d", uuidAddr, "guid", count)

	request, err := http.NewRequest("GET", uuidAddrFinal, nil)
	if err != nil {
		uuidPanic(err)
	}
	var client = &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode != 200 {
		uuidPanic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		uuidPanic(err)
	}
	bodyString := string(body)[1 : len(string(body))-1]
	bodyString = strings.Replace(bodyString, "\"", "", -1)
	ids := strings.Split(bodyString, ", ")
	return ids
}

func uuidPanic(err error) {
	err = errors.Wrap(err, "uuid error:")
	panic(err)
}
