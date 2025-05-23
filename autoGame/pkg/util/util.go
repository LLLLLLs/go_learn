package util

import (
	"bytes"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"gocv.io/x/gocv"
	"image"
	"image/png"
	"path/filepath"
)

func MustOK(err error) {
	if err != nil {
		panic(err)
	}
}

type Path string

func (p Path) Join(list ...string) Path {
	return Path(filepath.Join(append([]string{p.String()}, list...)...))
}

func (p Path) String() string {
	return string(p)
}

func (p Path) AbsPath() string {
	res, err := filepath.Abs(string(p))
	MustOK(err)
	return res
}

func Screenshoot() gocv.Mat {
	// 设置截取区域：从 (x=100, y=100)，宽度 300，高度 200
	rect := image.Rect(40, 150, 420, 780)

	img, err := screenshot.CaptureRect(rect)
	if err != nil {
		panic(err)
	}
	return RGBAImageToMat(img)
}

func RGBAImageToMat(img *image.RGBA) gocv.Mat {
	// 将 image.RGBA 编码为 PNG 字节流
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err)
	}

	// 解码为 OpenCV Mat
	mat, err := gocv.IMDecode(buf.Bytes(), gocv.IMReadUnchanged)
	if err != nil {
		panic(err)
	}
	return mat
}

func Click(x, y int) {
	resetX, resetY := robotgo.Location()
	robotgo.MoveClick(x, y)
	robotgo.MoveClick(x, y)
	robotgo.Move(resetX, resetY)
}
