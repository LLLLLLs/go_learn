package log_output

import (
	"testing"
)

func TestBytesBuffer(t *testing.T) {
	buf := GetBytesBuffer()
	defer PutBytesBuffer(buf)
	buf.Write([]byte("1234567890"))
	//s1 := buf.String()
	//PutBytesBuffer(buf)
	//fmt.Println(s1)
	//buf2 := GetBytesBuffer()
	//buf.Write([]byte("123"))
	//s2 := buf2.String()
	//PutBytesBuffer(buf)
	//fmt.Println(s2)
}
