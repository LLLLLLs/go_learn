package log_output

import (
	"bytes"
	"io"
	"sync"
)

var bytesBufPool = sync.Pool{
	// New is called when a new instance is needed
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBytesBuffer() *bytes.Buffer {
	buf := bytesBufPool.Get().(*bytes.Buffer)
	return buf
}

func PutBytesBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bytesBufPool.Put(buf)
}

var logRecoderHandlePool = sync.Pool{
	// New is called when a new instance is needed
	New: func() interface{} {
		return &logRecoderHandle{}
	},
}

type logRecoderHandle struct {
	Buf *bytes.Buffer
	W   io.Writer
}

func GetlogRecoderHandle() *logRecoderHandle {
	lgr := logRecoderHandlePool.Get().(*logRecoderHandle)
	return lgr
}

func PutlogRecoderHandle(lgr *logRecoderHandle) {
	lgr.W = nil
	logRecoderHandlePool.Put(lgr)
}
