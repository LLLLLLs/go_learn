//@time:2020/04/02
//@desc:

package io

import (
	"fmt"
	"io"
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
	WriteTimes() int
}

type Implementer struct {
	buff  strings.Builder
	times int
}

func (i *Implementer) Write(p []byte) (n int, err error) {
	i.times++
	fmt.Printf("Normal write: %v\n", p)
	return i.buff.Write(p)
}

func (i *Implementer) SpecialWrite(p []byte) (n int, err error) {
	i.times++
	fmt.Printf("Special write: %v\n", p)
	return i.buff.Write(p)
}

func (i *Implementer) Result() string {
	return i.buff.String()
}

func (i *Implementer) WriteTimes() int {
	return i.times
}

type WriteFunc func(p []byte) (n int, err error)

func (wf WriteFunc) Write(p []byte) (n int, err error) {
	return wf(p)
}

type MyReader struct {
	data  []byte
	index int
}

func NewMyReader(d []byte) *MyReader {
	return &MyReader{data: d}
}

func (mr *MyReader) Read(p []byte) (n int, err error) {
	var writeToBuf = func(begin, end int) {
		for i := 0; i < len(p) && begin+i < end; i++ {
			p[i] = mr.data[begin+i]
		}
	}

	initIndex := mr.index
	mr.index += len(p)
	fin := mr.index
	if fin >= len(mr.data) {
		err = io.EOF
		fin = len(mr.data)
		mr.index = 0
	}
	writeToBuf(initIndex, fin)
	n = fin - initIndex
	return
}
