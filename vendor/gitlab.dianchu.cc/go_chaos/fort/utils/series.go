package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"time"
)

func init() {
	gob.Register(time.Time{})
}

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func Marshal(object interface{}) ([]byte, error) {
	var (
		buf bytes.Buffer
		err error
	)
	err = gob.NewEncoder(&buf).Encode(object)
	return buf.Bytes(), err
}

func Unmarshal(data []byte, object interface{}) error {
	var (
		err error
	)
	err = gob.NewDecoder(bytes.NewBuffer(data)).Decode(object)
	return err
}
func GzipEncode(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GzipDecode(data []byte) ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer gr.Close()
	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return data, err
}

func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func BytesToInt64(data []byte) (int64, error) {
	if len(data) != 8 {
		return -1, errors.New("bytes len error")
	}
	return int64(binary.LittleEndian.Uint64(data)), nil
}

func Int64ToBytes(data int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(data))
	return b
}
