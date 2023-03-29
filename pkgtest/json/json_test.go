package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Data struct {
	A string
	B int
	C Data2
}

type Data2 struct {
	E string
	F int
}

type Target struct {
	A string
	B int
}

func TestJson(t *testing.T) {
	obj := []Data{
		{A: "1", B: 1, C: Data2{E: "11", F: 11}},
		{A: "2", B: 2, C: Data2{E: "21", F: 21}},
	}
	raw, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	var tar interface{}
	if err = json.Unmarshal(raw, &tar); err != nil {
		panic(err)
	}
	fmt.Println(tar)
	d1 := tar.([]interface{})[0]
	d1d2 := d1.(map[string]interface{})["C"]
	fmt.Println(d1d2.(map[string]interface{})["E"])

	var tar2 []Target
	if err = json.Unmarshal(raw, &tar2); err != nil {
		panic(err)
	}
	fmt.Println(tar2)
}
