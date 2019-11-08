// Time        : 2019/07/08
// Description :

package util

import (
	"bytes"
	"encoding/gob"
)

func DeepCopy(dst, src interface{}) {
	io := new(bytes.Buffer)
	MustNil(gob.NewEncoder(io).Encode(src))
	MustNil(gob.NewDecoder(io).Decode(dst))
}
