//@time:2020/04/02
//@desc:

package _interface

import (
	"fmt"
	"strings"
)

type WriteF func(p []byte) (n int, err error)

func (fw WriteF) Write(p []byte) (n int, err error) {
	return fw(p)
}

func MyWriteFunction(p []byte) (n int, err error) {
	// this function implements the Writer interface but is not named "Write"
	fmt.Printf("%v", p)
	return len(p), nil
}

type ProxyWrite interface {
	Write(p []byte) (n int, err error)
	SpecialWrite(p []byte) (n int, err error)
	Result() string
}

type Implementer struct {
	buff strings.Builder
}

func (i *Implementer) Write(p []byte) (n int, err error) {
	fmt.Printf("Normal write: %v\n", p)
	return i.buff.Write(p)
}

func (i *Implementer) SpecialWrite(p []byte) (n int, err error) {
	fmt.Printf("Special write: %v\n", p)
	return i.buff.Write(p)
}

func (i *Implementer) Result() string {
	return i.buff.String()
}

type WriteFunc func(p []byte) (n int, err error)

func (wf WriteFunc) Write(p []byte) (n int, err error) {
	return wf(p)
}
