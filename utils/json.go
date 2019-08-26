// Time        : 2019/08/19
// Description :

package utils

import "encoding/json"

func MarshalToString(v interface{}) string {
	b, err := json.Marshal(v)
	OkOrPanic(err)
	return string(b)
}
