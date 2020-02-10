//@time:2020/01/16
//@desc:

package main

import (
	datamock "golearn/sundry/data-mock"
	"time"
)

var begin = int64(8100000)
var end = int64(8100000)

func main() {
	for count := begin; count <= end; count += 100000 {
		datamock.Do(count)
		time.Sleep(time.Second * 5)
	}
}
