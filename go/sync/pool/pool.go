// Time        : 2019/11/01
// Description :

package pool

import "sync"

var intPool *sync.Pool
var slicePool *sync.Pool

func init() {
	intPool = &sync.Pool{
		New: func() interface{} {
			return 123
		},
	}
	slicePool = &sync.Pool{
		New: func() interface{} {
			return make([]int, 10)
		},
	}
}
