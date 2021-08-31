// @author: lls
// @date: 2021/7/14
// @desc:

package gozstd

import (
	"bytes"
	"github.com/valyala/gozstd"
	"golearn/util"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriteData(t *testing.T) {
	bd := bytes.Buffer{}
	for i := 0; i < 10000000; i++ {
		bd.WriteByte('a')
	}
	util.MustNil(ioutil.WriteFile("data.txt", bd.Bytes(), os.ModePerm))
}

func TestCompress(t *testing.T) {
	data, err := ioutil.ReadFile("data.json")
	util.MustNil(err)
	util.MustNil(ioutil.WriteFile("compress3.txt", gozstd.CompressLevel(nil, data, 1), os.ModePerm))
}

func TestDecompress(t *testing.T) {
	data, err := ioutil.ReadFile("compress3.txt")
	util.MustNil(err)
	res, err := gozstd.Decompress(nil, data)
	util.MustNil(err)
	util.MustNil(ioutil.WriteFile("decompress3.txt", res, os.ModePerm))
}
