//@time:2020/04/02
//@desc:

package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func TestIOCopy(t *testing.T) {
	w, err := io.Copy(WriteF(MyWriteFunction), os.Stdin)
	fmt.Println(w, err)
}

func TestCopy(t *testing.T) {
	ast := assert.New(t)
	Proxies := make(map[int]ProxyWrite)
	Proxies[1] = new(Implementer)
	Proxies[2] = new(Implementer)
	Proxies[3] = new(Implementer)
	Proxies[4] = new(Implementer)

	// 使用原生实现Writer的结构体作为接收者
	wtn, err := io.Copy(Proxies[1], strings.NewReader("Hello world"))
	fmt.Println(wtn, Proxies[1].Result())
	ast.Nil(err)

	// 使用WriteFunc包裹的Writer作为接收者
	wtn, err = io.Copy(WriteFunc(Proxies[2].SpecialWrite), strings.NewReader("Hello world 123"))
	fmt.Println(wtn, Proxies[2].Result())
	ast.Nil(err)

	// 使用自定义buffer,并且自己实现了Reader(因为strings.Reader实现了WriterTo,不走Buf)
	buf := make([]byte, 5)
	wtn, err = io.CopyBuffer(Proxies[3], NewMyReader([]byte("Hello world")), buf)
	fmt.Println(wtn, Proxies[3].Result(), Proxies[3].WriteTimes())
	ast.Nil(err)

	// 不借助io,手动从Reader copy to Writer
	buf = make([]byte, 1024)
	n, err := strings.NewReader("hello world").Read(buf)
	ast.Nil(err)
	wt, err := Proxies[4].Write(buf[:n])
	ast.Nil(err)
	fmt.Println(wt, Proxies[4].Result())
}
