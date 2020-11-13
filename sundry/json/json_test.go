//@author: lls
//@time: 2020/07/10
//@desc:

package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJS(t *testing.T) {
	data, _ := json.Marshal(JS{
		A: "",
	})
	fmt.Println(string(data))
}
