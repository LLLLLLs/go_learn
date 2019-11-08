// Time        : 2019/08/19
// Description :

package util

import "encoding/json"

func MarshalToString(v interface{}) string {
	b, err := json.Marshal(v)
	MustNil(err)
	return string(b)
}

func Marshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	MustNil(err)
	return b
}

func Unmarshal(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	MustNil(err)
}
