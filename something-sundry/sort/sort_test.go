// Time        : 2019/07/08
// Description :

package sort

import (
	"fmt"
	"testing"
)

// 看起来是一层循环 实则和普通冒泡没啥区别
func TestBubble(t *testing.T) {
	data := []byte{75, 76, 2, 3, 12, 77, 4, 7, 1, 6}
	length := len(data)
	i := 0
	j := 0
	for i < length {
		if data[j] < data[j+1] {
			data[j], data[j+1] = data[j+1], data[j]
		}
		if length-i-2 == 0 {
			break
		} else if j == length-i-2 {
			i++
			j = 0
			continue
		} else {
			j++
		}
	}
	fmt.Println(data)
}
