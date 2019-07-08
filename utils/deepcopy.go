// Time        : 2019/07/08
// Description :

package utils

import (
	"bytes"
	"encoding/gob"
)

func DeepCopy(dst, src interface{}) {
	io := new(bytes.Buffer)
	OkOrPanic(gob.NewEncoder(io).Encode(src))
	OkOrPanic(gob.NewDecoder(io).Decode(dst))
}
